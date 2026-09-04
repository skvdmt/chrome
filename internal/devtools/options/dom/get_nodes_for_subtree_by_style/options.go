package get_nodes_for_subtree_by_style

// Option Опция.
type Option func(c *Config)

// WithPierce Указать обработку iframe.
func WithPierce() Option {
	return func(c *Config) {
		c.Pierce = true
	}
}
