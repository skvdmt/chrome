package get_histograms

// Option Опция.
type Option func(c *Config)

// WithQuery С указанием части имени.
func WithQuery(query string) Option {
	return func(c *Config) {
		c.Query = query
	}
}

// WithDelta Получить изменение с момента последнего вызова функции изменений.
func WithDelta() Option {
	return func(c *Config) {
		c.Delta = true
	}
}
