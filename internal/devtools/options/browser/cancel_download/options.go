package cancel_download

import "github.com/skvdmt/chrome/internal/devtools/types/browser"

// Option Опция.
type Option func(c *Config)

// WithBrowserContextID Указание контекста браузера.
func WithBrowserContextID(browserContextID browser.BrowserContextID) Option {
	return func(c *Config) {
		c.BrowserContextID = browserContextID
	}
}
