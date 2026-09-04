package input

// MouseButton Кнопка мыши
// Возможные значения: none, left, middle, right, back, forward
type MouseButton string

// TimeSinceEpoch Время UTC в секундах, отсчитываемое с 1 января 1970 года.
type TimeSinceEpoch int

// TouchPoint
type TouchPoint struct {
	// Координата X события относительно области просмотра основного фрейма в пикселях CSS.
	X int `json:"x"`
	// Координата Y события относительно области просмотра основного фрейма в пикселях CSS.
	// 0 соответствует верхней границе области просмотра, а значение Y увеличивается
	// по мере приближения к нижней границе области просмотра.
	Y int `json:"y"`
	// Радиус сенсорной области по оси X (по умолчанию: 1,0).
	RadiusX int `json:"radiusX,omitempty"`
	// Радиус по оси Y области касания (по умолчанию: 1,0).
	RadiusY int `json:"radiusY,omitempty"`
	// Угол поворота (по умолчанию: 0,0).
	RotationAngle int `json:"rotationAngle,omitempty"`
	// Сила (по умолчанию: 1.0).
	Force int `json:"force,omitempty"`
	// Нормализованное тангенциальное давление, диапазон
	// значений которого составляет [-1,1] (по умолчанию: 0).
	TangentialPressure int `json:"tangentialPressure,omitempty"`
	// Угол между плоскостью Y-Z и плоскостью, содержащей как ось стилуса,
	// так и ось Y, в градусах в диапазоне [-90,90], положительный наклон X
	// означает наклон вправо (по умолчанию: 0).
	TiltX int `json:"tiltX,omitempty"`
	// Угол между плоскостью X-Z и плоскостью, содержащей как ось стилуса,
	// так и ось X, в градусах в диапазоне [-90,90], положительный наклон Y
	// направлен в сторону пользователя (по умолчанию: 0).
	TiltY int `json:"tiltY,omitempty"`
	// Вращение стилуса по часовой стрелке вокруг своей главной оси,
	// в градусах в диапазоне [0,359] (по умолчанию: 0).
	Twist int `json:"twist,omitempty"`
	// Идентификатор, используемый для отслеживания источников касания между
	// событиями, должен быть уникальным в рамках одного события.
	Id int `json:"id,omitempty"`
}

// DragData
type DragData struct {
	Items []*DragDataItem `json:"items"`
	// List of filenames that should be included when dropping
	Files []*string `json:"files,omitempty"`
	// Bit field representing allowed drag operations. Copy = 1, Link = 2, Move = 16
	DragOperationsMask int `json:"dragOperationsMask"`
}

// DragDataItem
type DragDataItem struct {
	// MIME-тип перетаскиваемых данных.
	MimeType string `json:"mimeType"`
	// В зависимости от значения параметра mimeType, он может содержать
	// перетаскиваемую ссылку, текст, HTML-разметку или любые другие данные.
	Data string `json:"data"`
	// Заголовок, связанный со ссылкой.
	// Действительно только при mimeType == "text/uri-list".
	Title string `json:"title,omitempty"`
	// Сохраняет базовый URL-адрес для содержащейся разметки.
	// Действителен только при mimeType == "text/html".
	BaseURL string `json:"baseURL,omitempty"`
}

// GestureSourceType
// Возможные значения: default, touch, mouse
type GestureSourceType string
