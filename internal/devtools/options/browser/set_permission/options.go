package set_permission

import "github.com/skvdmt/chrome/internal/devtools/types/browser"

// Option Опция.
type Option func(c *Config)

// WithOrigin Указавает источники встраивания на которые распространяется разрешение.
func WithOrigin(origin string) Option {
	return func(c *Config) {
		c.Origin = origin
	}
}

// WithEmbeddedOrigin Указывает встроенный источник.
func WithEmbeddedOrigin(embeddedOrigin string) Option {
	return func(c *Config) {
		c.EmbeddedOrigin = embeddedOrigin
	}
}

// WithBrowserContextId Указывает id контекста браузера.
func WithBrowserContextId(browserContextId browser.BrowserContextID) Option {
	return func(c *Config) {
		c.BrowserContextId = browserContextId
	}
}
