package enable

// Option Опция.
type Option func(c *Config)

// WithEnableFileChooserOpenedEvent событие Page.fileChooserOpened будет
// сгенерировано независимо от состояния, установленного командой
// Page.setInterceptFileChooserDialog
func WithEnableFileChooserOpenedEvent() Option {
	return func(c *Config) {
		c.EnableFileChooserOpenedEvent = true
	}
}
