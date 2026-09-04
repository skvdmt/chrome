package synthesize_pinch_gesture

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Option Опция.
type Option func(c *Config)

// WithRelativeSpeed Относительная скорость указателя в пикселях в секунду.
func WithRelativeSpeed(relativeSpeed int) Option {
	return func(c *Config) {
		c.RelativeSpeed = relativeSpeed
	}
}

// WithGestureSourceType Какой тип входных событий следует генерировать.
func WithGestureSourceType(gestureSourceType *input.GestureSourceType) Option {
	return func(c *Config) {
		c.GestureSourceType = gestureSourceType
	}
}
