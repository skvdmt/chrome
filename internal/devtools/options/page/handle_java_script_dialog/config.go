package handle_java_script_dialog

// Config Конфигурация.
type Config struct {
	// Принять или отклонить диалог.
	Accept bool `json:"accept"`
	// Текст, который нужно ввести в диалоговое окно перед подтверждением.
	// Используется только в том случае, если это диалоговое окно с
	// запросом на подтверждение.
	PromptText string `json:"promptText,omitempty"`
}
