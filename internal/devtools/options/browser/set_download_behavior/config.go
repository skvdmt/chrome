package set_download_behavior

import "github.com/skvdmt/chrome/internal/devtools/types/browser"

// Config Конфигурация.
type Config struct {
	// Разрешить все или запретить все запросы на загрузку, или
	// использовать стандартное поведение Chrome, если оно доступно
	// (в противном случае запретить). |allowAndName| разрешает загрузку
	// и присваивает имена файлам в соответствии с их GUID для загрузки.
	// Допустимые значения: deny, allow, allowAndName, default
	Behavior string `json:"behavior"`
	// Параметр задает поведение загрузки. Если он
	// опущен, используется контекст браузера по умолчанию.
	BrowserContextId browser.BrowserContextID `json:"browserContextId,omitempty"`
	// Путь по умолчанию для сохранения загруженных файлов. Это необходимо, если
	// для параметра поведения установлено значение 'allow' или 'allowAndName'.
	DownloadPath string `json:"downloadPath,omitempty"`
	// Следует ли отправлять события загрузки (по умолчанию — false).
	EventsEnabled bool `json:"eventsEnabled,omitempty"`
}
