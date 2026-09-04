package dispatch_touch_event

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Config Конфигурация.
type Config struct {
	// Тип события касания. События TouchEnd и TouchCancel не должны содержать
	// точек касания, в то время как события TouchStart и TouchMove должны
	// содержать хотя бы одну точку касания.
	// Допустимые значения: touchStart, touchEnd, touchMove, touchCancel
	Type string `json:"type"`
	// Активные точки касания на сенсорном устройстве. Генерируется одно
	// событие на каждую измененную точку (по сравнению с предыдущим
	// событием касания в последовательности), имитирующее
	// нажатие/перемещение/отпускание точек по очереди.
	TouchPoints []*input.TouchPoint `json:"touchPoints"`
	// Битовое поле, представляющее нажатые клавиши-модификаторы.
	// Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (по умолчанию: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Время, когда произошло событие.
	Timestamp *input.TimeSinceEpoch `json:"timestamp,omitempty"`
}
