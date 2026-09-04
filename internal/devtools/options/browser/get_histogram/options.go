package get_histogram

// Option Опция.
type Option func(c *Config)

// WithDelta Получить изменение с момента последнего вызова функции изменений.
func WithDelta() Option {
	return func(c *Config) {
		c.Delta = true
	}
}
