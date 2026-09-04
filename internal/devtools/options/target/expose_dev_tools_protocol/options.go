package expose_dev_tools_protocol

// Option Опция.
type Option func(c *Config)

// WithBindingName Указывает имя привязки.
func WithBindingName(bindingName string) Option {
	return func(c *Config) {
		c.BindingName = bindingName
	}
}

// WithInheritPermissions Наследуются права доступа текущей корневой сессии.
func WithInheritPermissions() Option {
	return func(c *Config) {
		c.InheritPermissions = true
	}
}
