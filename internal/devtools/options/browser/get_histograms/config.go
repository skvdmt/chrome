package get_histograms

// Config Конфигурация.
type Config struct {
	// Имя запрашиваемой гистограммы.
	Name string `json:"name,omitempty"`
	// Если true, получить изменение с момента последнего вызова функции изменений.
	Delta bool `json:"delta,omitempty"`
}
