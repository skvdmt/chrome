package set_contents_size

// Option Опция.
type Option func(c *Config)

// WithWidth Указание ширины.
func WithWidth(width int) Option {
	return func(c *Config) {
		c.Width = width
	}
}

// WithHeight Указание высоты.
func WithHeight(height int) Option {
	return func(c *Config) {
		c.Height = height
	}
}
