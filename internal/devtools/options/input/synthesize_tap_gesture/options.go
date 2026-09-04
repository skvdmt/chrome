package synthesize_tap_gesture

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Option Опция.
type Option func(c *Config)

// WithDuration Длительность между событиями нажать и отпустить в миллисекундах.
func WithDuration(duration int) Option {
	return func(c *Config) {
		c.Duration = duration
	}
}

// WithTapCount Количество нажатий.
func WithTapCount(tapCount int) Option {
	return func(c *Config) {
		c.TapCount = tapCount
	}
}

// WithGestureSourceType Какой тип входных событий следует генерировать.
func WithGestureSourceType(gestureSourceType *input.GestureSourceType) Option {
	return func(c *Config) {
		c.GestureSourceType = gestureSourceType
	}
}
