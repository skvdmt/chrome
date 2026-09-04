package create_browser_context

// Config Конфигурация.
type Config struct {
	// Если указано, освобождает этот контекст при отладке разрывов сеанса.
	DisposeOnDetach bool `json:"disposeOnDetach,omitempty"`
	// Прокси-сервер, аналогичный тому, который передается в параметр --proxy-server
	ProxyServer string `json:"proxyServer,omitempty"`
	// Список обхода прокси, аналогичный тому, который
	// передается в параметр --proxy-bypass-list.
	ProxyBypassList string `json:"proxyBypassList,omitempty"`
	// Необязательный список источников, которым предоставляется
	// неограниченный доступ из других источников. Части URL-адреса,
	// не являющиеся источниками, игнорируются.
	OriginsWithUniversalNetworkAccess []string `json:"originsWithUniversalNetworkAccess,omitempty"`
}

// Option Опция.
type Option func(c *Config)

// WithDisposeOnDetach Освобождить этот контекст при отладке разрывов сеанса.
func WithDisposeOnDetach() Option {
	return func(c *Config) {
		c.DisposeOnDetach = true
	}
}

// WithProxyServer Указать прокси-сервер.
func WithProxyServer(proxyServer string) Option {
	return func(c *Config) {
		c.ProxyServer = proxyServer
	}
}

// WithProxyBypassList Указать список обхода прокси.
func WithProxyBypassList(proxyBypassList string) Option {
	return func(c *Config) {
		c.ProxyBypassList = proxyBypassList
	}
}

// WithOriginsWithUniversalNetworkAccess Указать список источников, которым
// предоставляется неограниченный доступ из других источников.
func WithOriginsWithUniversalNetworkAccess(originsWithUniversalNetworkAccess []string) Option {
	return func(c *Config) {
		c.OriginsWithUniversalNetworkAccess = originsWithUniversalNetworkAccess
	}
}
