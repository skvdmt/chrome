package dispatch_drag_event

// Option Опция.
type Option func(c *Config)

// WithModifiers Битовое поле, представляющее нажатые клавиши-модификаторы.
func WithModifiers(modifiers int) Option {
	return func(c *Config) {
		c.Modifiers = modifiers
	}
}
