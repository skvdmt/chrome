package set_attributes_as_text

// Option Опция.
type Option func(c *Confug)

// WithName Указать имя атрибута.
func WithName(name string) Option {
	return func(c *Confug) {
		c.Name = name
	}
}
