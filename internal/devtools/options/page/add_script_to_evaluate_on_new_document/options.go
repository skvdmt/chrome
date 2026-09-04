package add_script_to_evaluate_on_new_document

// Option Опция.
type Option func(c *Config)

// WithWorldName создает изолированный мир с заданным именем
// и выполняет в нем заданный скрипт
func WithWorldName(worldName string) Option {
	return func(c *Config) {
		c.WorldName = worldName
	}
}

// WithIncludeCommandLineAPI Скрипт должен быть доступен через API командной строки.
func WithIncludeCommandLineAPI() Option {
	return func(c *Config) {
		c.IncludeCommandLineAPI = true
	}
}

// WithRunImmediately Скрипт запускается немедленно.
func WithRunImmediately() Option {
	return func(c *Config) {
		c.RunImmediately = true
	}
}
