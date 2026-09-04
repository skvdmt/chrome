package get_cookies

// Option Опция.
type Option func(c *Config)

// WithUrls Список URL-адресов, для которых будут получены соответствующие файлы cookie.
func WithUrls(urls []string) Option {
	return func(c *Config) {
		c.Urls = urls
	}
}
