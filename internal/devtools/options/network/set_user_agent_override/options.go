package set_user_agent_override

// Option Опция.
type Option func(c *Config)

// WithAcceptLanguage Язык браузера для эмуляции.
func WithAcceptLanguage(acceptLanguage string) Option {
	return func(c *Config) {
		c.AcceptLanguage = acceptLanguage
	}
}

// WithPlatform Свойство navigator.platform должно возвращать.
func WithPlatform(platform string) Option {
	return func(c *Config) {
		c.Platform = platform
	}
}
