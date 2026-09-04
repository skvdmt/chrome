package dispatch_mouse_event

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Config Конфигурация.
type Config struct {
	// Тип события мыши.
	// Допустимые значения: mousePressed, mouseReleased, mouseMoved, mouseWheel
	Type string `json:"type"`
	// Координата X события относительно области просмотра
	// основного фрейма в пикселях CSS.
	X int `json:"x"`
	// Координата Y события относительно области просмотра
	// основного фрейма в пикселях CSS. 0 соответствует верхней
	// границе области просмотра, а значение Y увеличивается
	// по мере приближения к нижней границе области просмотра.
	Y int `json:"y"`
	// Битовое поле, представляющее нажатые клавиши-модификаторы.
	// Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (по умолчанию: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Время, когда произошло событие.
	Timestamp *input.TimeSinceEpoch `json:"timestamp,omitempty"`
	// Кнопка мыши (по умолчанию: "none").
	Button *input.MouseButton `json:"button,omitempty"`
	// Число, указывающее, какие кнопки мыши нажаты при срабатывании события мыши. Левая=1, Правая=2, Средняя=4, Назад=8, Вперед=16, None=0.
	Buttons int `json:"buttons,omitempty"`
	// Количество нажатий кнопки мыши (по умолчанию: 0).
	ClickCount int `json:"clickCount,omitempty"`
	// Нормализованное давление, диапазон значений
	// которого составляет [0,1] (по умолчанию: 0).
	Force int `json:"force,omitempty"`
	// Нормализованное тангенциальное давление,
	// диапазон значений которого составляет [-1,1] (по умолчанию: 0).
	TangentialPressure int `json:"tangentialPressure,omitempty"`
	// Угол между плоскостью Y-Z и плоскостью, содержащей как
	// ось стилуса, так и ось Y, в градусах в диапазоне [-90,90],
	// положительный наклон X означает наклон вправо (по умолчанию: 0).
	TiltX int `json:"tiltX,omitempty"`
	// Угол между плоскостью X-Z и плоскостью, содержащей как ось стилуса,
	// так и ось X, в градусах в диапазоне [-90,90], положительный наклон Y
	// направлен в сторону пользователя (по умолчанию: 0).
	TiltY int `json:"tiltY,omitempty"`
	// Вращение стилуса по часовой стрелке вокруг своей главной оси,
	// в градусах в диапазоне [0,359] (по умолчанию: 0).
	Twist int `json:"twist,omitempty"`
	// Изменение по оси X в пикселях CSS для события
	// прокрутки колесика мыши (по умолчанию: 0).
	DeltaX int `json:"deltaX,omitempty"`
	// Изменение координаты Y в пикселях CSS для события
	// прокрутки колесика мыши (по умолчанию: 0).
	DeltaY int `json:"deltaY,omitempty"`
	// Тип указателя (по умолчанию: "mouse").
	// Допустимые значения: mouse, pen
	PointerType string `json:"pointerType,omitempty"`
}
