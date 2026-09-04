package set_user_agent_override

// Config Конфигурация.
type Config struct {
	// Используемый User-Agent.
	UserAgent string `json:"userAgent"`
	// Язык браузера для эмуляции.
	AcceptLanguage string `json:"acceptLanguage,omitempty"`
	// Свойство navigator.platform должно возвращать...
	Platform string `json:"platform,omitempty"`
}
