package chrome

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/skvdmt/chrome/internal/devtools/domains"
	"github.com/skvdmt/chrome/internal/devtools/options/target/attach_to_target"
	"github.com/skvdmt/chrome/internal/devtools/types/target"
	"github.com/skvdmt/chrome/internal/model"
)

const (
	// Путь к chrome.
	PATH = "/usr/bin/google-chrome"

	// Префикс для поиска адреса сервера, на котором запущен сервер управления браузером.
	devtoolsPrefix = "DevTools listening on"
)

// Driver Драйвер.
type Driver struct {
	// Отладчик.
	debug *model.Debug
	// Удаление пользовательской дириктории chrome после закрытия.
	removeUserDataDirAfterClose bool
	// Контекст.
	ctx context.Context
	// Конфигурация.
	config *Config
	// Путь к файлу.
	path string
	// Аргументы.
	args *model.Args
	// Вкладки.
	targets []*target.TargetInfo
	// Сессии.
	sessions map[target.TargetId]target.SessionId
	// Chrome browser
	chrome *exec.Cmd
	// Соединение с сервером chrome.
	client *model.Client
	// Браузер.
	Browser *domains.Browser
	// Объектная модель документа.
	Dom *domains.Dom
	// Ввод.
	Input *domains.Input
	// Сеть
	Network *domains.Network
	// Страница.
	Page *domains.Page
	// Вкладка.
	Target *domains.Target
}

// NewDriver Конструктор.
func NewDriver(option ...driverOption) (*Driver, error) {
	return NewDriverWithContext(context.Background(), option...)
}

// NewDriverWithContext Конструктор с контекстом.
func NewDriverWithContext(ctx context.Context, option ...driverOption) (*Driver, error) {
	// Конфигурация по-умолчанию.
	d := &Driver{
		ctx:    ctx,
		path:   PATH,
		config: NewConfig(),
		args:   model.NewArgs(),
		debug:  model.NewDebug(),
	}
	// Выполнение опций.
	for _, o := range option {
		if err := o(d); err != nil {
			return nil, err
		}
	}
	d.debug.Debug(fmt.Sprintf("path: %s", d.path))
	d.debug.Debug(fmt.Sprintf("args: %v", d.args.Join()))

	d.debug.Debug("chrome driver created")
	return d, nil
}

// Open Открытие.
func (d *Driver) Open() error {
	if err := d.run(); err != nil {
		return err
	}
	d.debug.Debug("chrome opened")
	if err := d.createTargetAndAttach(); err != nil {
		return err
	}
	d.debug.Debug("attached to target")
	return nil
}

// run Запуск.
func (d *Driver) run() error {
	d.chrome = exec.CommandContext(d.ctx, d.path, d.args.Join()...)
	if err := d.exec(); err != nil {
		return err
	}
	d.Browser = domains.NewBrowser(d.client, d.debug)
	d.Dom = domains.NewDom(d.client, d.debug)
	d.Input = domains.NewInput(d.client, d.debug)
	d.Network = domains.NewNetwork(d.client, d.debug)
	d.Page = domains.NewPage(d.client, d.debug)
	d.Target = domains.NewTarget(d.client, d.debug)
	return nil
}

// createTargetAndAttach Создание вкладки и присоединение к ней.
func (d *Driver) createTargetAndAttach() error {
	// Создание пустой вкладки.
	tid, err := d.Target.CreateTarget("")
	if err != nil {
		return nil
	}
	if len(*tid) == 0 {
		return model.ERR_NO_TARGETS
	}
	// Присоединение к вкладке.
	sid, err := d.Target.AttachToTarget(tid, attach_to_target.WithFlatten())
	if err != nil {
		return err
	}
	if len(*sid) == 0 {
		return model.ERR_NO_SESSION
	}
	d.updateCurrent(sid)
	return nil
}

// updateCurrent Обновление текущей сейссии.
func (d *Driver) updateCurrent(sid *target.SessionId) {
	d.Dom.CurrentSessionId = sid
	d.Page.CurrentSessionId = sid
	d.Input.CurrentSessionId = sid
}

// Close Закрытие.
func (d *Driver) Close() error {
	if err := d.Browser.Close(); err != nil {
		return err
	}
	d.debug.Debug("chrome browser closed")

	if err := d.client.Close(); err != nil {
		return err
	}
	d.debug.Debug("chrome server connection closed")

	if _, err := d.chrome.Process.Wait(); err != nil {
		return err
	}
	d.debug.Debug("chrome server closed")

	if d.removeUserDataDirAfterClose {
		if err := d.removeUserDataDir(); err != nil {
			return err
		}
	}
	d.debug.Debug("chrome driver closed")
	return nil
}

// CurrentTarget Получить id текущей вклдаки.
func (d *Driver) CurrentTargetId() (*target.TargetId, error) {
	for tid, sid := range d.sessions {
		if sid == *d.Dom.CurrentSessionId {
			return &tid, nil
		}
	}
	return nil, model.ERR_TARGET_NOT_FOUND
}

// exec Выполнение команды запуска chrome.
// Обработка вывода.
// Получение адреса сервера управления.
// Установка соединение с сервером по WebSocket.
// Запуск обработки входящих сообщений.
func (d *Driver) exec() error {
	// Получить источник вывода ошибок команды.
	s, err := d.chrome.StderrPipe()
	if err != nil {
		return err
	}
	// Создание отдельной группы процессов, чтобы не передавать
	// сигнал прерывания приложения внутрь процесса.
	d.chrome.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	// Выполнить команду запуска браузера.
	if err := d.chrome.Start(); err != nil {
		return err
	}
	u, err := d.devtoolsURL(s)
	if err != nil {
		return err
	}

	d.client = model.NewClient(d.debug, d.incomingHandler)
	if err := d.client.Open(d.ctx, u); err != nil {
		return err
	}

	return nil
}

// IncomingHandler Обработчик входящих сообщений.
func (d *Driver) incomingHandler(message []byte) error {
	d.debug.Debug(fmt.Sprintf("<- %+v", string(message)))
	_, o := d.event(message)
	if o {
		// обработка события.
		return nil
	}
	r, o := d.response(message)
	if o {
		d.client.Mu.Lock()
		d.client.Responses[r.Id] = r
		d.client.Mu.Unlock()
		return nil
	}
	return model.ERR_UNKNOWN_DATA
}

// event Событие.
func (d *Driver) event(message []byte) (*model.Event, bool) {
	e := &model.Event{}
	if err := json.Unmarshal(message, e); err != nil {
		return nil, false
	}
	if len(e.Method) == 0 {
		return nil, false
	}
	return e, true
}

// response Ответ.
func (d *Driver) response(message []byte) (*model.Response, bool) {
	r := &model.Response{}
	if err := json.Unmarshal(message, r); err != nil {
		return nil, false
	}
	if r.Id == 0 {
		return nil, false
	}
	return r, true
}

// devtoolsURL Адрес сервера управления Chrome.
func (d *Driver) devtoolsURL(source io.ReadCloser) (string, error) {
	buf := bufio.NewReader(source)
	p := []byte(devtoolsPrefix)
	var u string
	for {
		l, err := buf.ReadBytes('\n')
		if err != nil {
			return "", err
		}
		if bytes.HasPrefix(l, p) {
			l = l[len(p):]
			u = string(bytes.TrimSpace(l))
			if err = source.Close(); err != nil {
				return "", err
			}
			break
		}
	}
	return u, nil
}

// removeUserDataDir Удаление пользовательской дириктории chrome.
func (d *Driver) removeUserDataDir() error {
	p, o := d.args.Get(model.ARG_NAME_USER_DATA_DIR)
	if !o {
		return model.ERR_USER_DATA_DIR_NOT_SET
	}
	// Костыль. 5 попыток удалить дирикторию.
	for i := 1; i <= 5; i++ {
		err := os.RemoveAll(p)
		if err != nil && i <= 5 {
			time.Sleep(time.Millisecond * 20)
			continue
		}
		if err != nil {
			return err
		}
		break
	}
	d.debug.Debug("chrome user data dir removed")
	return nil
}
