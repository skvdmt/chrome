package emulate_touch_from_mouse_event

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Option Опция.
type Option func(c *Config)

// WithTimestamp Время, когда произошло событие.
func WithTimestamp(timestamp *input.TimeSinceEpoch) Option {
	return func(c *Config) {
		c.Timestamp = timestamp
	}
}

// WithDeltaX Изменение по оси X в DIP для события прокрутки колесика мыши.
func WithDeltaX(deltaX int) Option {
	return func(c *Config) {
		c.DeltaX = deltaX
	}
}

// WithDeltaY Изменение по оси Y в DIP для события прокрутки колесика мыши.
func WithDeltaY(deltaY int) Option {
	return func(c *Config) {
		c.DeltaY = deltaY
	}
}

// WithModifiers Битовое поле, представляющее нажатые клавиши-модификаторы.
func WithModifiers(modifiers int) Option {
	return func(c *Config) {
		c.Modifiers = modifiers
	}
}

// WithClickCount Количество нажатий кнопки мыши.
func WithClickCount(clickCount int) Option {
	return func(c *Config) {
		c.ClickCount = clickCount
	}
}
