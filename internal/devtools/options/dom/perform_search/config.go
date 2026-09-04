package perform_search

// Config Конфигурация.
type Config struct {
	// Простой текст, селектор запроса или поисковый запрос XPath.
	Query string `json:"query"`
	// Значение true означает поиск в теневом DOM-элементе пользовательского агента.
	IncludeUserAgentShadowDOM bool `json:"includeUserAgentShadowDOM,omitempty"`
}
