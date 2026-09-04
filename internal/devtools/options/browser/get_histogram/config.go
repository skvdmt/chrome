package get_histogram

// Config Конфигурация.
type Config struct {
	// Запрошенное название гистограммы.
	Name string `json:"name,omitempty"`
	// Если true, получить изменение с момента последнего вызова функции изменений.
	Delta bool `json:"delta,omitempty"`
}
