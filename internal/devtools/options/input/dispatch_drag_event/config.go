package dispatch_drag_event

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Config Конфигурация.
type Config struct {
	// Тип события торможения.
	// Допустимые значения: dragEnter, dragOver, drop, dragCancel
	Type string `json:"type"`
	// Координата X события относительно области
	// просмотра основного фрейма в пикселях CSS.
	X int `json:"X"`
	// Координата Y события относительно области
	// просмотра основного фрейма в пикселях CSS.
	// 0 соответствует верхней границе области просмотра,
	// а значение Y увеличивается по мере приближения к
	// нижней границе области просмотра.
	Y    int             `json:"Y"`
	Data *input.DragData `json:"data"`
	// Битовое поле, представляющее нажатые клавиши-модификаторы.
	// Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (по умолчанию: 0).
	Modifiers int `json:"modifiers,omitempty"`
}
