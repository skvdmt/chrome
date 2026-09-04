package cancel_download

import "github.com/skvdmt/chrome/internal/devtools/types/browser"

// Config Конфигурация.
type Config struct {
	// Глобальный уникальный идентификатор загрузки.
	Guid string `json:"guid"`
	// Контекст браузера в котором будет выполняться действие.
	// Если этот параметр опущен, используется контекст браузера по умолчанию.
	BrowserContextID browser.BrowserContextID `json:"browserContextId"`
}
