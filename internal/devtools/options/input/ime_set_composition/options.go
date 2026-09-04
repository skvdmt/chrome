package ime_set_composition

// Option Опция.
type Option func(c *Config)

// WithReplacementStart Начало замены.
func WithReplacementStart(replacementStart int) Option {
	return func(c *Config) {
		c.ReplacementStart = replacementStart
	}
}

// WithReplacementEnd Конец замены.
func WithReplacementEnd(replacementEnd int) Option {
	return func(c *Config) {
		c.ReplacementEnd = replacementEnd
	}
}
