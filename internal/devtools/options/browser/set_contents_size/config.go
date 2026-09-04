package set_contents_size

import "github.com/skvdmt/chrome/internal/devtools/types/browser"

// Config Конфигурация.
type Config struct {
	// ID окна браузера.
	WindowId browser.WindowID `json:"windowId"`
	// Ширина содержимого окна в формате DIP. Если параметр
	// опущен, предполагается текущая ширина. Необходимо указать,
	// если параметр 'height' опущен.
	Width int `json:"width,omitempty"`
	// Высота содержимого окна в формате DIP. Если параметр
	// опущен, предполагается текущая высота. Если параметр
	// 'width' опущен, его необходимо указать.
	Height int `json:"height,omitempty"`
}
