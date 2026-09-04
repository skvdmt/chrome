package domains

import (
	"github.com/skvdmt/chrome/internal/devtools/options/target/attach_to_target"
	"github.com/skvdmt/chrome/internal/devtools/options/target/create_browser_context"
	"github.com/skvdmt/chrome/internal/devtools/options/target/create_target"
	"github.com/skvdmt/chrome/internal/devtools/options/target/expose_dev_tools_protocol"
	"github.com/skvdmt/chrome/internal/devtools/options/target/open_dev_tools"
	"github.com/skvdmt/chrome/internal/devtools/options/target/set_auto_attach"
	"github.com/skvdmt/chrome/internal/devtools/types/browser"
	"github.com/skvdmt/chrome/internal/devtools/types/target"
	"github.com/skvdmt/chrome/internal/model"
)

// Target Вкладка.
type Target struct {
	client *model.Client
	debug  *model.Debug
}

// NewTarget Конструктор.
func NewTarget(c *model.Client, d *model.Debug) *Target {
	d.Debug("target created")
	return &Target{
		client: c,
		debug:  d,
	}
}

// ActivateTarget Активирует (фокусирует) вкладку.
func (t *Target) ActivateTarget(targetId *target.TargetId) error {
	return t.client.Exec(
		target.ACTIVATE_TARGET,
		model.ForceJSONMarshal(struct {
			TargetId *target.TargetId `json:"targetId"`
		}{
			TargetId: targetId,
		}),
	)
}

// AttachToTarget Прикрепляется к вклдаке с заданным идентификатором..
func (t *Target) AttachToTarget(
	targetId *target.TargetId,
	options ...attach_to_target.Option,
) (*target.SessionId, error) {
	c := &attach_to_target.Config{
		TargetId: targetId,
	}
	for _, o := range options {
		o(c)
	}
	r, err := t.client.Query(target.ATTACH_TO_TARGET, model.ForceJSONMarshal(c))
	if err != nil {
		return nil, err
	}
	s := struct {
		SessionId *target.SessionId `json:"sessionId"`
	}{}
	model.ForceJSONUnmarshal(r, &s)
	return s.SessionId, nil
}

// CloseTarget Закрывает вкладку.
func (t *Target) СloseTarget(targetId *target.TargetId) error {
	return t.client.Exec(
		target.CLOSE_TARGET,
		model.ForceJSONMarshal(struct {
			TargetId *target.TargetId `json:"targetId"`
		}{
			TargetId: targetId,
		}),
	)
}

// CreateBrowserContext Создает новый пустой BrowserContext. Аналогично
// профилю инкогнито, но можно иметь несколько таких профилей.
func (t *Target) CreateBrowserContext(
	options ...create_browser_context.Option,
) (*browser.BrowserContextID, error) {
	c := &create_browser_context.Config{}
	for _, o := range options {
		o(c)
	}
	r, err := t.client.Query(target.CREATE_BROWSER_CONTEXT, model.ForceJSONMarshal(c))
	if err != nil {
		return nil, err
	}
	b := struct {
		// Идентификатор созданного контекста.
		BrowserContextId *browser.BrowserContextID `json:"browserContextId"`
	}{}
	model.ForceJSONUnmarshal(r, &b)
	return b.BrowserContextId, nil
}

// CreateTarget Создание новой вкладки.
func (t *Target) CreateTarget(
	u string,
	options ...create_target.Option,
) (*target.TargetId, error) {
	c := &create_target.Config{
		Url: u,
	}
	for _, o := range options {
		o(c)
	}
	r, err := t.client.Query(target.CREATE_TARGET, model.ForceJSONMarshal(c))
	if err != nil {
		return nil, err
	}
	i := struct {
		// Идентификатор открытой страницы.
		TargetId *target.TargetId `json:"targetId"`
	}{}
	model.ForceJSONUnmarshal(r, &i)
	return i.TargetId, nil
}

// DetachFromTarget Отключает сессию с заданным идентификатором.
func (t *Target) DetachFromTarget(sessionId *target.SessionId) error {
	return t.client.Exec(
		target.DETACH_FROM_TARGET,
		model.ForceJSONMarshal(struct {
			SessionId *target.SessionId `json:"sessionId"`
		}{
			SessionId: sessionId,
		}),
	)
}

// DisposeBrowserContext Удаляет объект BrowserContext. Все связанные с ним
// страницы будут закрыты без вызова их обработчиков beforeunload.
func (t *Target) DisposeBrowserContext(browserContextId *browser.BrowserContextID) error {
	return t.client.Exec(
		target.DISPOSE_BROWSER_CONTEXT, model.ForceJSONMarshal(struct {
			BrowserContextId *browser.BrowserContextID `json:"browserContextId"`
		}{
			BrowserContextId: browserContextId,
		}),
	)
}

// GetBrowserContexts
func (t *Target) GetBrowserContexts() (
	[]*browser.BrowserContextID,
	*browser.BrowserContextID,
	error,
) {
	r, err := t.client.Query(target.GET_BROWSER_CONTEXTS, nil)
	if err != nil {
		return nil, nil, err
	}
	c := struct {
		// Массив идентификаторов контекста браузера.
		BrowserContextIds []*browser.BrowserContextID `json:"browserContextIds"`
		// Идентификатор контекста браузера по умолчанию, если он доступен.
		DefaultBrowserContextId *browser.BrowserContextID `json:"defaultBrowserContextId,omitempty"`
	}{}
	model.ForceJSONUnmarshal(r, &c)
	return c.BrowserContextIds, c.DefaultBrowserContextId, nil
}

// GetTargets Получает список доступных вкладок.
func (t *Target) GetTargets() ([]*target.TargetInfo, error) {
	r, err := t.client.Query(target.GET_TARGETS, nil)
	if err != nil {
		return nil, err
	}
	i := struct {
		TargetInfos []*target.TargetInfo `json:"targetInfos"`
	}{}
	model.ForceJSONUnmarshal(r, &i)
	return i.TargetInfos, nil
}

// SetAutoAttach Управляет автоматическим подключением к новым целям, которые
// считаются непосредственно связанными с данной (например, iframe или worker).
// При включении подключается также ко всем существующим связанным целям.
// При выключении автоматически отключается от всех подключенных целей.
// Это также удаляет все цели, добавленные с помощью autoAttachRelated,
// из списка целей, за которыми следует следить на предмет создания
// связанных целей. Возможно, вам потребуется вызывать эту функцию рекурсивно,
// чтобы автоматически подключаемые цели подключались ко всем доступным целям.
func (t *Target) SetAutoAttach(autoAttach bool, waitForDebuggerOnStart bool, options ...set_auto_attach.Option) error {
	c := &set_auto_attach.Config{
		AutoAttach:             autoAttach,
		WaitForDebuggerOnStart: waitForDebuggerOnStart,
	}
	for _, o := range options {
		o(c)
	}
	return t.client.Exec(target.SET_AUTHO_ATTACH, model.ForceJSONMarshal(c))
}

// SetDiscoverTargets Управляет процессом обнаружения доступных целей и уведомления
// о них посредством событий targetCreated/targetInfoChanged/targetDestroyed.
func (t *Target) SetDiscoverTargets(discover bool) error {
	return t.client.Exec(
		target.SET_DISCOVER_TARGETS,
		model.ForceJSONMarshal(struct {
			// Стоит ли искать доступные цели.
			Discover bool `josn:"discover"`
		}{
			Discover: discover,
		}),
	)
}

// AttachToBrowserTarget Привязывается к целевому объекту браузера,
// использует только режим с плоским идентификатором сессии.
func (t *Target) AttachToBrowserTarget() (*target.SessionId, error) {
	r, err := t.client.Query(target.ATTACH_TO_BROWSER_TARGET, nil)
	if err != nil {
		return nil, err
	}
	s := struct {
		// Идентификатор, присвоенный сессии.
		SessionId *target.SessionId `json:"sessionId"`
	}{}
	model.ForceJSONUnmarshal(r, &s)
	return s.SessionId, nil
}

// AutoAttachRelated Добавляет указанную цель в список целей, которые будут
// отслеживаться на предмет создания связанных целей (таких как дочерние
// фреймы, дочерние рабочие процессы и новые версии сервис-воркеров) и
// сообщаться через attachedToTarget. Указанная цель также автоматически
// подключается. Это отменяет действие любого предыдущего SetAutoAttach
// и отменяется последующим SetAutoAttach. Доступно только для цели Browser.
func (t *Target) AutoAttachRelated(targetId *target.TargetId, waitForDebuggerOnStart bool) error {
	return t.client.Exec(
		target.AUTO_ATTACH_RELATED,
		model.ForceJSONMarshal(struct {
			TargetId *target.TargetId `json:"targetId"`
			// Указывается, следует ли приостанавливать выполнение новых целей
			// при подключении к ним. Используйте runtime.RunIfWaitingForDebugger
			// для запуска приостановленных целей.
			WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
		}{
			TargetId:               targetId,
			WaitForDebuggerOnStart: waitForDebuggerOnStart,
		}),
	)
}

// ExposeDevToolsProtocol Внедрите объект в главный фрейм целевого объекта,
// который обеспечивает канал связи с целевым браузером. Внедренный объект
// будет доступен как window[bindingName]. Объект имеет следующий API:
//   - binding.send(json) - метод для отправки сообщений
//     по протоколу удаленной отладки;
//   - binding.onmessage = json => handleMessage(json) - функция
//     обратного вызова, которая будет вызываться для уведомлений
//     протокола и ответов на команды.
func (t *Target) ExposeDevToolsProtocol(
	targetId *target.TargetId,
	options ...expose_dev_tools_protocol.Option,
) error {
	c := &expose_dev_tools_protocol.Config{
		TargetId: targetId,
	}
	for _, o := range options {
		o(c)
	}
	return t.client.Exec(target.EXPOSE_DEV_TOOLS_PROTOCOL, model.ForceJSONMarshal(c))
}

// GetDevToolsTarget Получает targetId целевого объекта страницы DevTools,
// открытого для указанного целевого объекта (если таковой имеется).
func (t *Target) GetDevToolsTarget(targetId *target.TargetId) (*target.TargetId, error) {
	r, err := t.client.Query(
		target.GET_DEV_TOOLS_TARGET,
		model.ForceJSONMarshal(struct {
			// Page or tab target ID.
			TargetId *target.TargetId `json:"targetId"`
		}{
			TargetId: targetId,
		}),
	)
	if err != nil {
		return nil, err
	}
	i := struct {
		// TargetId целевого объекта страницы DevTools, если он существует.
		TargetId *target.TargetId `json:"targetId,omitempty"`
	}{}
	model.ForceJSONUnmarshal(r, &i)
	return i.TargetId, nil
}

// GetTargetInfo Возвращает информацию о вкладке.
func (t *Target) GetTargetInfo(targetId *target.TargetId) (*target.TargetInfo, error) {
	r, err := t.client.Query(
		target.GET_TARGET_INFO,
		model.ForceJSONMarshal(struct {
			TargetId *target.TargetId `json:"targetId"`
		}{
			TargetId: targetId,
		}),
	)
	if err != nil {
		return nil, err
	}
	i := struct {
		TargetInfo *target.TargetInfo `json:"targetInfo"`
	}{}
	model.ForceJSONUnmarshal(r, &i)
	return i.TargetInfo, nil
}

// OpenDevTools Открывает окно инструментов разработчика для целевого объекта.
func (t *Target) OpenDevTools(
	targetId *target.TargetId,
	options ...open_dev_tools.Option,
) error {
	c := &open_dev_tools.Config{
		TargetId: targetId,
	}
	for _, o := range options {
		o(c)
	}
	return t.client.Exec(target.OPEN_DEV_TOOLS, model.ForceJSONMarshal(c))
}

// SetRemoteLocations Включает обнаружение целевых объектов
// для указанных местоположений, если для параметра
// SetDiscoverTargets установлено значение true.
func (t *Target) SetRemoteLocations(locations []*target.RemoteLocation) error {
	return t.client.Exec(
		target.SET_REMOTE_LOCATIONS,
		model.ForceJSONMarshal(struct {
			// Список удаленных мест.
			Locations []*target.RemoteLocation `json:"locations"`
		}{
			Locations: locations,
		}),
	)
}
