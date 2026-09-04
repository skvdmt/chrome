package perform_search

// Option Опция.
type Option func(c *Config)

// WithIncludeUserAgentShadowDOM Указать поиск в теневом DOM-элементе пользовательского агента.
func WithIncludeUserAgentShadowDOM() Option {
	return func(c *Config) {
		c.IncludeUserAgentShadowDOM = true
	}
}
