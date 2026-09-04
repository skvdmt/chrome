package synthesize_tap_gesture

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Config Конфигурация.
type Config struct {
	// Координата X начала жеста в пикселях CSS.
	X int `json:"x"`
	// Координата Y начала жеста в пикселях CSS.
	Y int `json:"y"`
	// Длительность между событиями нажать и
	// отпустить в миллисекундах (по умолчанию: 50).
	Duration int `json:"duration,omitempty"`
	// Количество нажатий (например, 2 для двойного нажатия, по умолчанию: 1).
	TapCount int `json:"tapCount,omitempty"`
	// Какой тип входных событий следует генерировать
	// (по умолчанию: 'default', что означает запрос к платформе
	// для определения предпочтительного типа входных данных).
	GestureSourceType *input.GestureSourceType `json:"gestureSourceType,omitempty"`
}
