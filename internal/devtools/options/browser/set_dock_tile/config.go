package set_dock_tile

// Config Конфигурация.
type Config struct {
	BadgeLabel string `json:"badgeLabel"`
	// Изображение в формате PNG. (При передаче в
	// формате JSON закодировано как строка base64)
	Image string `json:"image,omitempty"`
}
