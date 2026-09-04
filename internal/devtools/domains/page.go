package domains

import (
	"github.com/skvdmt/chrome/internal/devtools/options/dom/enable"
	"github.com/skvdmt/chrome/internal/devtools/options/page/add_script_to_evaluate_on_new_document"
	"github.com/skvdmt/chrome/internal/devtools/options/page/capture_screenshot"
	"github.com/skvdmt/chrome/internal/devtools/options/page/create_isolated_world"
	"github.com/skvdmt/chrome/internal/devtools/options/page/get_app_manifest"
	"github.com/skvdmt/chrome/internal/devtools/options/page/handle_java_script_dialog"
	"github.com/skvdmt/chrome/internal/devtools/options/page/navigate"
	"github.com/skvdmt/chrome/internal/devtools/options/page/print_to_pdf"
	"github.com/skvdmt/chrome/internal/devtools/options/page/reload"
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	"github.com/skvdmt/chrome/internal/devtools/types/io"
	"github.com/skvdmt/chrome/internal/devtools/types/network"
	"github.com/skvdmt/chrome/internal/devtools/types/page"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
	"github.com/skvdmt/chrome/internal/devtools/types/target"
	"github.com/skvdmt/chrome/internal/model"
)

// Page Страница.
type Page struct {
	client *model.Client
	debug  *model.Debug
	// Текущая сессия.
	CurrentSessionId *target.SessionId
}

// NewPage Конструктор.
func NewPage(c *model.Client, d *model.Debug) *Page {
	d.Debug("page created")
	return &Page{
		client: c,
		debug:  d,
	}
}

// AddScriptToEvaluateOnNewDocument Выполняет проверку заданного скрипта
// в каждом кадре при его создании (перед загрузкой скриптов кадра).
func (p *Page) AddScriptToEvaluateOnNewDocument(
	source string,
	options ...add_script_to_evaluate_on_new_document.Option,
) (*page.ScriptIdentifier, error) {
	c := &add_script_to_evaluate_on_new_document.Config{
		Source: source,
	}
	for _, o := range options {
		o(c)
	}
	r, err := p.client.Query(
		page.ADD_SCRIPT_TO_EVALUATE_ON_NEW_DOCUMENT,
		model.ForceJSONMarshal(c),
		model.WithSessionId(p.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	i := struct {
		// Идентификатор добавленного скрипта.
		Identifier *page.ScriptIdentifier `json:"identifier"`
	}{}
	model.ForceJSONUnmarshal(r, &i)
	return i.Identifier, nil
}

// BringToFront Выводит страницу на передний план (активирует вкладку).
func (p *Page) BringToFront() error {
	return p.client.Exec(
		page.BRING_TO_FRONT,
		nil,
		model.WithSessionId(p.CurrentSessionId),
	)
}

// CaptureScreenshot Сделать снимок экрана страницы.
func (p *Page) CaptureScreenshot(options ...capture_screenshot.Option) error {
	c := &capture_screenshot.Config{}
	for _, o := range options {
		o(c)
	}
	return p.client.Exec(
		page.CAPTURE_SCREENSHOT,
		model.ForceJSONMarshal(c),
		model.WithSessionId(p.CurrentSessionId),
	)
}

// Close Пытается закрыть страницу, запуская соответствующие
// обработчики события beforeunload, если таковые имеются.
func (p *Page) Close() error {
	return p.client.Exec(
		page.CLOSE,
		nil,
		model.WithSessionId(p.CurrentSessionId),
	)
}

// CreateIsolatedWorld Создаёт изолированный мир для заданного кадра.
func (p *Page) CreateIsolatedWorld(
	frameId *page.FrameId,
	options ...create_isolated_world.Option,
) (*rnt.ExecutionContextId, error) {
	c := &create_isolated_world.Config{
		FrameId: frameId,
	}
	for _, o := range options {
		o(c)
	}
	r, err := p.client.Query(
		page.CREATE_ISOLATED_WORLD,
		model.ForceJSONMarshal(c),
		model.WithSessionId(p.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	e := struct {
		// Контекст выполнения в изолированном мире.
		ExecutionContextId *rnt.ExecutionContextId `json:"executionContextId"`
	}{}
	model.ForceJSONUnmarshal(r, &e)
	return e.ExecutionContextId, nil
}

// Disable Отключает уведомления о домене страницы.
func (p *Page) Disable() error {
	return p.client.Exec(
		page.DISABLE,
		nil,
		model.WithSessionId(p.CurrentSessionId),
	)
}

// Enable Включает уведомления о домене страницы.
func (p *Page) Enable(options ...enable.Option) error {
	c := &enable.Config{}
	for _, o := range options {
		o(c)
	}
	return p.client.Exec(
		page.ENABLE,
		model.ForceJSONMarshal(c),
		model.WithSessionId(p.CurrentSessionId),
	)
}

// GetAppManifest Получает обработанный манифест для текущего документа.
// Этот API всегда ожидает загрузки манифеста. Если указан manifestId,
// и он не совпадает с манифестом текущего документа, API выдает ошибку.
// Если страница не загружена, API немедленно выдает ошибку.
func (p *Page) GetAppManifest(options ...get_app_manifest.Option) (
	u string,
	errs []*page.AppManifestError,
	data string,
	parsed *page.AppManifestParsedProperties,
	manifest *page.WebAppManifest,
	err error,
) {
	c := &get_app_manifest.Config{}
	for _, o := range options {
		o(c)
	}
	r, err := p.client.Query(
		page.GET_APP_MANIFEST,
		model.ForceJSONMarshal(c),
		model.WithSessionId(p.CurrentSessionId),
	)
	if err != nil {
		return "", nil, "", nil, nil, err
	}
	m := struct {
		// Местоположение, указанное в списке.
		Url    string                   `json:"url"`
		Errors []*page.AppManifestError `json:"errors"`
		// Проявить контент.
		Data string `json:"data"`
		// Анализируются свойства манифеста.
		// Устарело, используйте вместо этого manifest.
		Parsed   *page.AppManifestParsedProperties `json:"parsed"`
		Manifest *page.WebAppManifest              `json:"manifest"`
	}{}
	model.ForceJSONUnmarshal(r, &m)
	return m.Url, m.Errors, m.Data, m.Parsed, m.Manifest, nil
}

// GetFrameTree Возвращает структуру дерева текущего кадра.
func (p *Page) GetFrameTree() (*page.FrameTree, error) {
	r, err := p.client.Query(
		page.GET_FRAME_TREE,
		nil,
		model.WithSessionId(p.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	f := struct {
		// Представлена ​​древовидная структура каркаса.
		FrameTree *page.FrameTree `json:"frameTree"`
	}{}
	model.ForceJSONUnmarshal(r, &f)
	return f.FrameTree, nil
}

// GetLayoutMetrics Возвращает метрики, относящиеся к компоновке страницы,
// такие как границы/масштаб области просмотра.
func (p *Page) GetLayoutMetrics() (
	*page.LayoutViewport,
	*page.VisualViewport,
	*dom.Rect,
	*page.LayoutViewport,
	*page.VisualViewport,
	*dom.Rect,
	error,
) {
	r, err := p.client.Query(
		page.GET_LAYOUT_METRICS,
		nil,
		model.WithSessionId(p.CurrentSessionId),
	)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}
	l := struct {
		// Устаревшие метрики, относящиеся к области просмотра макета.
		// Используются пиксели устройства. Вместо них используйте cssLayoutViewport.
		LayoutViewport *page.LayoutViewport `json:"layoutViewport"`
		// Устаревшие метрики, относящиеся к визуальной области просмотра.
		// Используются пиксели устройства. Вместо них используйте cssVisualViewport.
		VisualViewport *page.VisualViewport `json:"visualViewport"`
		// Устаревший размер прокручиваемой области. Используется в DP.
		// Вместо него используйте cssContentSize.
		ContentSize *dom.Rect `json:"contentSize"`
		// Метрики, относящиеся к области просмотра макета, в пикселях CSS.
		CssLayoutViewport *page.LayoutViewport `json:"cssLayoutViewport"`
		// Метрики, относящиеся к визуальной области просмотра в пикселях CSS.
		CssVisualViewport *page.VisualViewport `json:"cssVisualViewport"`
		// Размер прокручиваемой области в пикселях CSS.
		CssContentSize *dom.Rect `json:"cssContentSize"`
	}{}
	model.ForceJSONUnmarshal(r, &l)
	return l.LayoutViewport, l.VisualViewport, l.ContentSize,
		l.CssLayoutViewport, l.CssVisualViewport, l.ContentSize, nil
}

// GetNavigationHistory Возвращает историю навигации по текущей странице.
func (p *Page) GetNavigationHistory() (
	currentIndex int,
	entries []*page.NavigationEntry,
	err error,
) {
	r, err := p.client.Query(
		page.GET_NAVIGATION_HISTORY,
		nil,
		model.WithSessionId(p.CurrentSessionId),
	)
	if err != nil {
		return 0, nil, err
	}
	h := struct {
		// Index of the current navigation history entry.
		CurrentIndex int `json:"currentIndex"`
		// Array of navigation history entries.
		Entries []*page.NavigationEntry `json:"entries"`
	}{}
	model.ForceJSONUnmarshal(r, &h)
	return h.CurrentIndex, h.Entries, nil
}

// HandleJavaScriptDialog Принимает или отклоняет диалоговое окно, инициированное
// JavaScript (alert, confirm, prompt или onbeforeunload).
func (p *Page) HandleJavaScriptDialog(
	accept bool,
	options ...handle_java_script_dialog.Option,
) error {
	c := &handle_java_script_dialog.Config{
		Accept: accept,
	}
	return p.client.Exec(
		page.HANDLE_JAVA_SCRIPT_DIALOG,
		model.ForceJSONMarshal(c),
		model.WithSessionId(p.CurrentSessionId),
	)
}

// Navigate Перенаправляет на текущую страницу по указанному URL-адресу.
func (p *Page) Navigate(u string, options ...navigate.Option) (
	FrameId *page.FrameId,
	LoaderId *network.LoaderId,
	ErrorText *string,
	IsDownload *bool,
	err error,
) {
	c := &navigate.Config{
		Url: u,
	}
	r, err := p.client.Query(
		page.NAVIGATE,
		model.ForceJSONMarshal(c),
		model.WithSessionId(p.CurrentSessionId),
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	f := struct {
		// Идентификатор кадра, который был просмотрен (или просмотр не удался).
		FrameId *page.FrameId `json:"frameId"`
		// Идентификатор загрузчика. Он опускается в случае навигации по одному
		// и тому же документу, поскольку ранее сохраненный loaderId не изменится.
		LoaderId *network.LoaderId `json:"loaderId"`
		// Удобное для пользователя сообщение об ошибке, которое
		// появляется только в случае сбоя навигации.
		ErrorText *string `json:"errorText"`
		// Привела ли навигация к загрузке.
		IsDownload *bool `json:"isDownload"`
	}{}
	model.ForceJSONUnmarshal(r, &f)
	return f.FrameId, f.LoaderId, f.ErrorText, f.IsDownload, nil
}

// NavigateToHistoryEntry Перенаправляет на текущую страницу
// к указанной записи в истории.
func (p *Page) NavigateToHistoryEntry(entryId int) error {
	return p.client.Exec(
		page.NAVIGATE_TO_HISTORY_ENTRY,
		model.ForceJSONMarshal(struct {
			// Уникальный идентификатор записи, к которой нужно перейти.
			EntryId int `json:"entryId"`
		}{
			EntryId: entryId,
		}),
		model.WithSessionId(p.CurrentSessionId),
	)
}

// PrintToPDF Распечатать страницу в формате PDF.
func (p *Page) PrintToPDF(options ...print_to_pdf.Option) (
	Data *string,
	Stream *io.StreamHandle,
	err error,
) {
	c := &print_to_pdf.Config{}
	for _, o := range options {
		o(c)
	}
	r, err := p.client.Query(
		page.PRINT_TO_PDF,
		model.ForceJSONMarshal(c),
		model.WithSessionId(p.CurrentSessionId),
	)
	if err != nil {
		return nil, nil, err
	}
	d := struct {
		// Данные PDF, закодированные в Base64. Пусто,
		// если указан параметр |returnAsStream|. (Закодировано
		// как строка Base64 при передаче в формате JSON)
		Data *string `json:"data"`
		// Дескриптор потока, содержащего результирующие данные PDF-файла.
		Stream *io.StreamHandle `json:"stream,omitempty"`
	}{}
	model.ForceJSONUnmarshal(r, &d)
	return d.Data, d.Stream, nil
}

// Reload Перезагружает указанную страницу, при желании игнорируя кэш.
func (p *Page) Reload(options ...reload.Option) error {
	c := &reload.Config{}
	for _, o := range options {
		o(c)
	}
	return p.client.Exec(
		page.RELOAD,
		model.ForceJSONMarshal(c),
		model.WithSessionId(p.CurrentSessionId),
	)
}

// RemoveScriptToEvaluateOnNewDocument Удаляет указанный скрипт из списка.
func (p *Page) RemoveScriptToEvaluateOnNewDocument(identifier *page.ScriptIdentifier) error {
	return p.client.Exec(
		page.REMOVE_SCRIPT_TO_EVALUATE_ON_DOCUMENT,
		model.ForceJSONMarshal(struct {
			Identifier *page.ScriptIdentifier `json:"identifier"`
		}{
			Identifier: identifier,
		}),
		model.WithSessionId(p.CurrentSessionId),
	)
}

// ResetNavigationHistory Сбрасывает историю навигации для текущей страницы.
func (p *Page) ResetNavigationHistory() error {
	return p.client.Exec(
		page.RESET_NAVIGATION_HISTORY,
		nil,
		model.WithSessionId(p.CurrentSessionId),
	)
}

// SetBypassCSP Включите возможность обхода политики безопасности содержимого страницы.
func (p *Page) SetBypassCSP(enabled bool) error {
	// Следует ли обходить CSP страницы.
	return p.client.Exec(
		page.SET_BYPASS_CSP,
		model.ForceJSONMarshal(struct {
			Enabled bool `json:"enabled"`
		}{
			Enabled: enabled,
		}),
		model.WithSessionId(p.CurrentSessionId),
	)
}

// SetDocumentContent Устанавливает заданную разметку в качестве HTML-кода документа.
func (p *Page) SetDocumentContent(frameId *page.FrameId, html string) error {
	return p.client.Exec(
		page.SET_DOCUMENT_CONTENT,
		model.ForceJSONMarshal(struct {
			// Идентификатор фрейма, для которого нужно задать HTML-код.
			FrameId *page.FrameId `json:"frameId"`
			// HTML-контент для настройки.
			Html string `json:"html"`
		}{
			FrameId: frameId,
			Html:    html,
		}),
		model.WithSessionId(p.CurrentSessionId),
	)
}

// SetInterceptFileChooserDialog Перехват запросов выбора файла и передача
// управления клиентам протокола. При включении перехвата запросов выбора
// файла диалоговое окно выбора файла не отображается. Вместо этого
// генерируется событие протокола Page.fileChooserOpened.
func (p *Page) SetInterceptFileChooserDialog(enabled bool, cancel bool) error {
	return p.client.Exec(
		page.SET_INTERCEPT_FILE_CHOISER_DIALOG,
		model.ForceJSONMarshal(struct {
			Enabled bool `json:"enabled"`
			// Если значение равно true, диалог отменяется путем генерации
			// соответствующих событий (если таковые имеются), а также не
			// отображается, если перехват включен (по умолчанию: false).
			Cancel bool `json:"cancel"`
		}{
			Enabled: enabled,
			Cancel:  cancel,
		}),
		model.WithSessionId(p.CurrentSessionId),
	)
}

// SetLifecycleEventsEnabled Определяет, будет ли страница
// генерировать события жизненного цикла.
func (p *Page) SetLifecycleEventsEnabled(enabled bool) error {
	return p.client.Exec(
		page.SET_LIFECYCLE_EVENTS_ENABLED,
		model.ForceJSONMarshal(struct {
			// Если это так, то начинает генерировать события жизненного цикла.
			Enabled bool `json:"enabled"`
		}{
			Enabled: enabled,
		}),
		model.WithSessionId(p.CurrentSessionId),
	)
}

// StopLoading Принудительно остановить все навигационные
// действия и ожидающие запросы к ресурсам на этой странице.
func (p *Page) StopLoading() error {
	return p.client.Exec(
		page.STOP_LOADING,
		nil,
		model.WithSessionId(p.CurrentSessionId),
	)
}
