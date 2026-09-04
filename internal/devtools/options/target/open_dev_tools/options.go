package open_dev_tools

// Option Опция.
type Option func(c *Config)

// WithPanelId Указывает идентификатор панели.
func WithPanelId(panelId string) Option {
	return func(c *Config) {
		c.PanelId = panelId
	}
}
