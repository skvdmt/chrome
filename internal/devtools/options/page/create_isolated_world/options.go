package create_isolated_world

// Option Опция.
type Option func(c *Config)

// WithWorldName Необязательное имя.
func WithWorldName(worldName string) Option {
	return func(c *Config) {
		c.WorldName = worldName
	}
}

// WithGrantUniveralAccess Предоставлять всеобщий доступ изолированному миру.
func WithGrantUniveralAccess() Option {
	return func(c *Config) {
		c.GrantUniveralAccess = true
	}
}

// WithContentSecurityPolicy Необязательная политика безопасности
// контента (CSP) для изолированной среды.
func WithContentSecurityPolicy(contentSecurityPolicy string) Option {
	return func(c *Config) {
		c.ContentSecurityPolicy = contentSecurityPolicy
	}
}
