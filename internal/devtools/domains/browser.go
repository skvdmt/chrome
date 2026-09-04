package domains

import (
	"github.com/skvdmt/chrome/internal/devtools/options/browser/cancel_download"
	"github.com/skvdmt/chrome/internal/devtools/options/browser/get_histogram"
	"github.com/skvdmt/chrome/internal/devtools/options/browser/get_histograms"
	"github.com/skvdmt/chrome/internal/devtools/options/browser/get_window_for_target"
	"github.com/skvdmt/chrome/internal/devtools/options/browser/reset_permissions"
	"github.com/skvdmt/chrome/internal/devtools/options/browser/set_contents_size"
	"github.com/skvdmt/chrome/internal/devtools/options/browser/set_dock_tile"
	"github.com/skvdmt/chrome/internal/devtools/options/browser/set_download_behavior"
	"github.com/skvdmt/chrome/internal/devtools/options/browser/set_permission"
	"github.com/skvdmt/chrome/internal/devtools/types/browser"
	"github.com/skvdmt/chrome/internal/model"
)

// Browser Браузер.
type Browser struct {
	client *model.Client
	debug  *model.Debug
}

// NewBrowser Конструктор.
func NewBrowser(c *model.Client, d *model.Debug) *Browser {
	d.Debug("domain browser created")
	return &Browser{
		client: c,
		debug:  d,
	}
}

// AddPrivacySandboxEnrollmentOverride Позволяет сайту использовать функции
// песочницы конфиденциальности, требующие регистрации, без фактической
// регистрации самого сайта. Поддерживается только для целевых страниц.
func (b *Browser) AddPrivacySandboxEnrollmentOverride(u string) error {
	return b.client.Exec(
		browser.ADD_PRIVACY_SANDBOX_ENROLLMENT_OVERRIDE,
		model.ForceJSONMarshal(struct {
			Url string `json:"url"`
		}{u}),
	)
}

// Close Корректное закрытие браузера.
func (b *Browser) Close() error {
	return b.client.Exec(browser.CLOSE, nil)
}

// GetVersion Информация о версии.
func (b *Browser) GetVersion() (
	protocolVersion,
	product,
	revision,
	userAgent,
	jsVersion *string,
	err error,
) {
	r, err := b.client.Query(browser.GET_VERSION, nil)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	d := struct {
		ProtocolVersion string `json:"protocolVersion"`
		Product         string `json:"product"`
		Revision        string `json:"revision"`
		UserAgent       string `json:"userAgent"`
		JsVersion       string `json:"jsVersion"`
	}{}
	model.ForceJSONUnmarshal(r, &d)
	return &d.ProtocolVersion, &d.Product, &d.Revision, &d.UserAgent, &d.JsVersion, nil
}

// ResetPermissions Сбросить все права доступа для всех источников.
func (b *Browser) ResetPermissions(option ...reset_permissions.Option) error {
	c := &reset_permissions.Config{}
	for _, o := range option {
		o(c)
	}
	return b.client.Exec(browser.RESET_PERMISSIONS, model.ForceJSONMarshal(c))
}

// CancelDownload Отмените загрузку, если она уже идёт.
func (b *Browser) CancelDownload(guid string, option ...cancel_download.Option) error {
	c := &cancel_download.Config{
		Guid: guid,
	}
	for _, o := range option {
		o(c)
	}
	return b.client.Exec(browser.CANCEL_DOWNLOAD, model.ForceJSONMarshal(c))
}

// Crash Крушение.
func (b *Browser) Crash() error {
	return b.client.Exec(browser.CRASH, nil)
}

// CrashGpuProcess Крушение GPU процесса.
func (b *Browser) CrashGpuProcess() error {
	return b.client.Exec(browser.CRASH_GPU_PROCESS, nil)
}

// ExecuteBrowserCommand Вызов пользовательских команд браузера, используемых телеметрией.
func (b *Browser) ExecuteBrowserCommand(commandId browser.BrowserCommandId) error {
	return b.client.Exec(browser.EXECUTE_BROWSER_COMMAND, model.ForceJSONMarshal(struct {
		CommandId browser.BrowserCommandId `json:"commandId"`
	}{
		CommandId: commandId,
	}))
}

// GetBrowserCommandLine Возвращает параметры командной строки для процесса браузера
// только в том случае, если параметр --enable-automation присутствует в командной строке.
func (b *Browser) GetBrowserCommandLine() ([]string, error) {
	r, err := b.client.Query(browser.GET_BROWSER_COMMAND_LINE, nil)
	if err != nil {
		return nil, err
	}
	c := struct {
		// Параметры командной строки.
		Arguments []string `json:"arguments"`
	}{}
	model.ForceJSONUnmarshal(r, c)
	return c.Arguments, nil
}

// GetHistogram Гистограмма Chrome по имени.
func (b *Browser) GetHistogram(name string, options ...get_histogram.Option) (*browser.Histogram, error) {
	c := &get_histogram.Config{
		Name: name,
	}
	for _, o := range options {
		o(c)
	}
	r, err := b.client.Query(browser.GET_HISTOGRAM, model.ForceJSONMarshal(c), nil)
	if err != nil {
		return nil, err
	}
	h := struct {
		Histogram *browser.Histogram `json:"histogram"`
	}{}
	model.ForceJSONUnmarshal(r, &h)
	return h.Histogram, nil
}

// GetHistograms Гистограммы Chrome.
func (b *Browser) GetHistograms(name string, options ...get_histograms.Option) ([]*browser.Histogram, error) {
	c := &get_histograms.Config{}
	for _, o := range options {
		o(c)
	}
	r, err := b.client.Query(browser.GET_HISTOGRAMS, model.ForceJSONMarshal(c))
	if err != nil {
		return nil, err
	}
	h := struct {
		Histograms []*browser.Histogram `json:"histograms"`
	}{}
	model.ForceJSONUnmarshal(r, &h)
	return h.Histograms, nil
}

// GetWindowBounds Положение и размер окна браузера.
func (b *Browser) GetWindowBounds(windowID browser.WindowID) (*browser.Bounds, error) {
	r, err := b.client.Query(
		browser.GET_WINDOW_BOUNDS,
		model.ForceJSONMarshal(struct {
			// ID окна браузера.
			WindowId browser.WindowID `json:"windowId"`
		}{
			WindowId: windowID,
		}),
	)
	if err != nil {
		return nil, err
	}
	bs := struct {
		Bounds *browser.Bounds
	}{}
	model.ForceJSONUnmarshal(r, &b)
	return bs.Bounds, nil
}

// OLD

// GetWindowForTarget Id окна и информация о границах.
// Эксперементально. Это может быть изменино, перемещено или удалено.
func (b *Browser) GetWindowForTarget(option ...get_window_for_target.Option) (
	*browser.WindowID, *browser.Bounds, error) {
	c := &get_window_for_target.Config{}
	for _, o := range option {
		o(c)
	}
	r, err := b.client.Query(
		browser.GET_WINDOW_FOR_TARGET,
		model.ForceJSONMarshal(c),
	)
	if err != nil {
		return nil, nil, err
	}
	d := struct {
		WindowId *browser.WindowID `json:"windowId"`
		Bounds   *browser.Bounds   `json:"bounds"`
	}{}
	model.ForceJSONUnmarshal(r, &d)
	return d.WindowId, d.Bounds, nil
}

// SetContentsSize Установите размер содержимого браузера,
// изменяя размер окна браузера по мере необходимости.
func (b *Browser) SetContentsSize(options ...set_contents_size.Option) error {
	c := &set_contents_size.Config{}
	for _, o := range options {
		o(c)
	}
	return b.client.Exec(browser.SET_CONTENTS_SIZE, model.ForceJSONMarshal(c))
}

// SetDockTile Задайте подробные характеристики плитки
// панели задач, специфичные для каждой платформы.
func (b *Browser) SetDockTile(badgeLabel string, options ...set_dock_tile.Option) error {
	c := &set_dock_tile.Config{
		BadgeLabel: badgeLabel,
	}
	for _, o := range options {
		o(c)
	}
	return b.client.Exec(browser.SET_DOCK_TILE, model.ForceJSONMarshal(c))
}

// SetDownloadBehavior Настройка поведения при загрузке файла.
func (b *Browser) SetDownloadBehavior(behavior string, options ...set_download_behavior.Option) error {
	c := &set_download_behavior.Config{
		Behavior: behavior,
	}
	for _, o := range options {
		o(c)
	}
	return b.client.Exec(browser.SET_DOWNLOAD_BEHAVIOR, model.ForceJSONMarshal(c))
}

// SetPermission Задание параметров разрешений для указанного
// встраивания и источников встраивания.
func (b *Browser) SetPermission(
	permission browser.PermissionDescriptor,
	setting browser.PermissionSetting,
	options ...set_permission.Option) error {
	c := &set_permission.Config{
		Permission: permission,
		Setting:    setting,
	}
	for _, o := range options {
		o(c)
	}
	return b.client.Exec(browser.SET_PERMISSION, model.ForceJSONMarshal(c))
}

// SetWindowBounds Указывает положение и/или размер окна браузера.
func (b *Browser) SetWindowBounds(windowId browser.WindowID, bounds browser.Bounds) error {
	return b.client.Exec(
		browser.SET_WINDOW_BOUNDS,
		model.ForceJSONMarshal(struct {
			// ID окна браузера.
			WindowId browser.WindowID `json:"windowId"`
			// Новые границы окна. Состояния «свернуто», «развернуто» и «полноэкранный режим»
			// нельзя комбинировать с «слева», «сверху», «шириной» или «высотой».
			// Неуказанные поля остаются без изменений.
			Bounds browser.Bounds `json:"bounds"`
		}{
			WindowId: windowId,
			Bounds:   bounds,
		}),
	)
}
