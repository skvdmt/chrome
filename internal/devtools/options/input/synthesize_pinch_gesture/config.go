package synthesize_pinch_gesture

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Config Конфигурация.
type Config struct {
	// Координата X начала жеста в пикселях CSS.
	X int `json:"x"`
	// Координата Y начала жеста в пикселях CSS.
	Y int `json:"y"`
	// Относительный масштабный коэффициент после увеличения
	// масштаба (>1,0 — увеличение, <1,0 — уменьшение).
	ScaleFactor int `json:"scaleFactor"`
	// Относительная скорость указателя в пикселях в секунду (по умолчанию: 800).
	RelativeSpeed int `json:"relativeSpeed,omitempty"`
	// Какой тип входных событий следует генерировать
	// (по умолчанию: 'default', что означает запрос к платформе
	// для определения предпочтительного типа входных данных).
	GestureSourceType *input.GestureSourceType `json:"gestureSourceType,omitempty"`
}
