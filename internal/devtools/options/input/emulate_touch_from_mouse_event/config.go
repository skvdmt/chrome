package emulate_touch_from_mouse_event

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Config Конфигурация.
type Config struct {
	// Тип события мыши.
	// Допустимые значения: mousePressed, mouseReleased, mouseMoved, mouseWheel
	Type string `json:"type"`
	// X-координата указателя мыши в DIP.
	X int `json:"x"`
	// Y-координата указателя мыши в DIP.
	Y int `json:"y"`
	// Кнопка мыши. Поддерживаются только значения "none", "left", "right".
	Button *input.MouseButton `json:"button"`
	// Время, когда произошло событие (по умолчанию: текущее время).
	Timestamp *input.TimeSinceEpoch `json:"timestamp,omitempty"`
	// Изменение по оси X в DIP для события прокрутки колесика мыши (по умолчанию: 0).
	DeltaX int `json:"deltaX,omitempty"`
	// Изменение по оси Y в DIP для события прокрутки колесика мыши (по умолчанию: 0).
	DeltaY int `json:"deltaY,omitempty"`
	// Битовое поле, представляющее нажатые клавиши-модификаторы.
	// Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (по умолчанию: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Количество нажатий кнопки мыши (по умолчанию: 0).
	ClickCount int `json:"clickCount,omitempty"`
}
