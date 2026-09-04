package ime_set_composition

// Config Конфигурация.
type Config struct {
	// Текст для вставки.
	Text string `json:"text"`
	// Начало выбора.
	SelectionStart int `json:"selectionStart"`
	// Конец выбора.
	SelectionEnd int `json:"selectionEnd"`
	// Замена старта.
	ReplacementStart int `json:"replacementStart,omitempty"`
	// Замена конца.
	ReplacementEnd int `json:"replacementEnd,omitempty"`
}
