package attach_to_target

// Option Опция.
type Option func(c *Config)

// WithFlatten Указывает получение "плоского" доступа к сессии.
func WithFlatten() Option {
	return func(c *Config) {
		c.Flatten = true
	}
}
