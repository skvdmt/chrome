package set_dock_tile

// Option Опция.
type Option func(c *Config)

// WithImage Указание изображения.
func WithImage(image string) Option {
	return func(c *Config) {
		c.Image = image
	}
}
