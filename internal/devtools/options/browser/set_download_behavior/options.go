package set_download_behavior

import "github.com/skvdmt/chrome/internal/devtools/types/browser"

// Option Опция.
type Option func(c *Config)

// WithBrowserContextId Указание id контекста браузера.
func WithBrowserContextId(browserContextId browser.BrowserContextID) Option {
	return func(c *Config) {
		c.BrowserContextId = browserContextId
	}
}

// WithDownloadPath Установка пути для сохранения загруженных файлов.
func WithDownloadPath(downloadPath string) Option {
	return func(c *Config) {
		c.DownloadPath = downloadPath
	}
}

// WithEventsEnabled Отправлять события загрузки.
func WithEventsEnabled() Option {
	return func(c *Config) {
		c.EventsEnabled = true
	}
}
