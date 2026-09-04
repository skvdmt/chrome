package handle_java_script_dialog

// Option Опция.
type Option func(c *Config)

// WithPromptText Текст, который нужно ввести в диалоговое окно перед подтверждением.
func WithPromptText(promptText string) Option {
	return func(c *Config) {
		c.PromptText = promptText
	}
}
