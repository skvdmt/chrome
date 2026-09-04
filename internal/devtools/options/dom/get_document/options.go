package get_document

// Option Опция.
type Option func(c *Config)

// WithDepth Указание максимальной глубины, дочерних элементов.
func WithDepth(depth int) Option {
	return func(c *Config) {
		c.Depth = depth
	}
}

// WithPierce Обходить iframe и теневые корни.
func WithPierce() Option {
	return func(c *Config) {
		c.Pierce = true
	}
}
