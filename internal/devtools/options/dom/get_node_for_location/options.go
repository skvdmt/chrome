package get_node_for_location

// Option Опция.
type Option func(c *Config)

// WithIncludeUserAgentShadowDOM включить теневой DOM UserAgent
func WithIncludeUserAgentShadowDOM() Option {
	return func(c *Config) {
		c.IncludeUserAgentShadowDOM = true
	}
}

// WithIgnorePointerEventsNone Игнорировать события указателя.
func WithIgnorePointerEventsNone() Option {
	return func(c *Config) {
		c.IgnorePointerEventsNone = true
	}
}
