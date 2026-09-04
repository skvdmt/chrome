package enable

// Config Конфигурация.
type Config struct {
	// Следует ли включать пробелы в массив дочерних узлов, возвращаемых системой Node.
	// Возможные значения: none, all.
	IncludeWhitespace string `json:"includeWhitespace,omitempty"`
}
