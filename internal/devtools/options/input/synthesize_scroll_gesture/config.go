package synthesize_scroll_gesture

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Config Конфигурация.
type Config struct {
	// Координата X начала жеста в пикселях CSS.
	X int `json:"x"`
	// Координата Y начала жеста в пикселях CSS.
	Y int `json:"y"`
	// Расстояние прокрутки вдоль оси X
	// (положительное значение соответствует прокрутке влево).
	XDistance int `json:"xDistance,omitempty"`
	// Расстояние прокрутки вдоль оси Y
	// (положительное значение означает прокрутку вверх).
	YDistance int `json:"yDistance,omitempty"`
	// Количество дополнительных пикселей, которые необходимо
	// прокрутить назад вдоль оси X, помимо заданного расстояния.
	XOverscroll int `json:"xOverscroll,omitempty"`
	// Количество дополнительных пикселей, которые необходимо
	// прокрутить назад вдоль оси Y, помимо заданного расстояния.
	YOverscroll int `json:"yOverscroll,omitempty"`
	// Предотвратить запуск (по умолчанию: true).
	PreventFling bool `json:"preventFling,omitempty"`
	// Скорость прокрутки в пикселях в секунду (по умолчанию: 800).
	Speed int `json:"speed,omitempty"`
	// Какой тип входных событий следует генерировать
	// (по умолчанию: 'default', что означает запрос к платформе
	// для определения предпочтительного типа входных данных).
	GestureSourceType *input.GestureSourceType `json:"gestureSourceType,omitempty"`
	// Количество повторений жеста (по умолчанию: 0).
	RepeatCount int `json:"repeatCount,omitempty"`
	// Задержка в миллисекундах между каждым повторением (по умолчанию: 250).
	RepeatDelayMs int `json:"repeatDelayMs,omitempty"`
	// Название генерируемых маркеров взаимодействия,
	// если оно не пустое (по умолчанию: "").
	InteractionMarkerName string `json:"interactionMarkerName,omitempty"`
}
