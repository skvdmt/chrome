package synthesize_scroll_gesture

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Option Опция.
type Option func(c *Config)

// WithXDistance Расстояние прокрутки вдоль оси X.
func WithXDistance(xDistance int) Option {
	return func(c *Config) {
		c.XDistance = xDistance
	}
}

// WithYDistance Расстояние прокрутки вдоль оси Y.
func WithYDistance(yDistance int) Option {
	return func(c *Config) {
		c.YDistance = yDistance
	}
}

// WithXOverscroll Количество дополнительных пикселей, которые необходимо
// прокрутить назад вдоль оси X, помимо заданного расстояния.
func WithXOverscroll(xOverscroll int) Option {
	return func(c *Config) {
		c.XOverscroll = xOverscroll
	}
}

// WithYOverscroll Количество дополнительных пикселей, которые необходимо
// прокрутить назад вдоль оси Y, помимо заданного расстояния.
func WithYOverscroll(yOverscroll int) Option {
	return func(c *Config) {
		c.YOverscroll = yOverscroll
	}
}

// WithPreventFling Предотвратить запуск.
func WithPreventFling(preventFling bool) Option {
	return func(c *Config) {
		c.PreventFling = preventFling
	}
}

// WithSpeed Скорость прокрутки в пикселях в секунду.
func WithSpeed(speed int) Option {
	return func(c *Config) {
		c.Speed = speed
	}
}

// WithGestureSourceType Какой тип входных событий следует генерировать.
func WithGestureSourceType(gestureSourceType *input.GestureSourceType) Option {
	return func(c *Config) {
		c.GestureSourceType = gestureSourceType
	}
}

// WithRepeatCount Количество повторений жеста.
func WithRepeatCount(repeatCount int) Option {
	return func(c *Config) {
		c.RepeatCount = repeatCount
	}
}

// WithRepeatDelayMs Задержка в миллисекундах между каждым повторением.
func WithRepeatDelayMs(repeatDelayMs int) Option {
	return func(c *Config) {
		c.RepeatDelayMs = repeatDelayMs
	}
}

// WithInteractionMarkerName Название генерируемых маркеров взаимодействия.
func WithInteractionMarkerName(interactionMarkerName string) Option {
	return func(c *Config) {
		c.InteractionMarkerName = interactionMarkerName
	}
}
