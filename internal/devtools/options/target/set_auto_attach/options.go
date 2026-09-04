package set_auto_attach

// Option Опция.
type Option func(c *Config)

// WithFlatten позволяет получить «плоский» доступ к сессии.
func WithFlatten() Option {
	return func(c *Config) {
		c.Flatten = true
	}
}
