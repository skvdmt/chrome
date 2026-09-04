package browser

// Bounds Информация о границах окна браузера.
type Bounds struct {
	// Смещение от левого края экрана до окна в пикселях.
	Left int `json:"left,omitempty"`
	// Смещение от верхнего края экрана до окна в пикселях.
	Top int `json:"top,omitempty"`
	// Ширина окна в пикселях.
	Width int `json:"width,omitempty"`
	// Высота окна в пикселях.
	Height int `json:"height,omitempty"`
	// Состояние окна. По умолчанию — normal.
	WindowState WindowState `json:"windowState,omitempty"`
}
