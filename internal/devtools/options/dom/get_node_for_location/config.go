package get_node_for_location

// Config Конфигурация.
type Config struct {
	// Координата X.
	X int `json:"x"`
	// Координата Y.
	Y int `json:"y"`
	// Значение false означает переход к ближайшему корневому предку,
	// не относящемуся к UA (по умолчанию: false).
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
	// Следует ли игнорировать события указателя: нет событий для
	// элементов и проверять их на срабатывание.
	IgnorePointerEventsNone bool `json:"ignorePointerEventsNone,omitempty"`
}
