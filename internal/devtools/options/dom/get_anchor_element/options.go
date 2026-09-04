package get_anchor_element

// Option Опция.
type Option func(c *Config)

// WithAnchorSpecifier Указать необязательный спецификатор привязки.
func WithAnchorSpecifier(anchorSpecifier string) Option {
	return func(c *Config) {
		c.AnchorSpecifier = anchorSpecifier
	}
}
