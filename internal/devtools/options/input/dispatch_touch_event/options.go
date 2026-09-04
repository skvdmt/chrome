package dispatch_touch_event

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Option Опция.
type Option func(c *Config)

// WithModifiers Битовое поле, представляющее нажатые клавиши-модификаторы.
func WithModifiers(modifiers int) Option {
	return func(c *Config) {
		c.Modifiers = modifiers
	}
}

// WithTimestamp Время, когда произошло событие.
func WithTimestamp(timestamp *input.TimeSinceEpoch) Option {
	return func(c *Config) {
		c.Timestamp = timestamp
	}
}
