package set_cookie

import "github.com/skvdmt/chrome/internal/devtools/types/network"

// Option Опция.
type Option func(c *Config)

// WithUrl Request-URI, который следует связать с установкой cookie.
func WithUrl(url string) Option {
	return func(c *Config) {
		c.Url = url
	}
}

// WithDomain Домен cookie.
func WithDomain(domain string) Option {
	return func(c *Config) {
		c.Domain = domain
	}
}

// WithPath Путь cookie.
func WithPath(path string) Option {
	return func(c *Config) {
		c.Path = path
	}
}

// WithSecure Cookie защищен.
func WithSecure() Option {
	return func(c *Config) {
		c.Secure = true
	}
}

// WithHttpOnly Значение true, если cookie доступен только по протоколу HTTP.
func WithHttpOnly() Option {
	return func(c *Config) {
		c.HttpOnly = true
	}
}

// WithSameSite Тип SameSite для cookie.
func WithSameSite(sameSite *network.CookieSameSite) Option {
	return func(c *Config) {
		c.SameSite = sameSite
	}
}

// WithExpires Дата истечения срока действия cookie.
func WithExpires(expires *network.TimeSinceEpoch) Option {
	return func(c *Config) {
		c.Expires = expires
	}
}

// WithPriority Тип приоритета cookie.
func WithPriority(priority *network.CookiePriority) Option {
	return func(c *Config) {
		c.Priority = priority
	}
}

// WithSourceScheme Тип схемы источника cookie.
func WithSourceScheme(sourceScheme *network.CookieSourceScheme) Option {
	return func(c *Config) {
		c.SourceScheme = sourceScheme
	}
}

// WithSourcePort Исходный порт cookie.
func WithSourcePort(sourcePort int) Option {
	return func(c *Config) {
		c.SourcePort = sourcePort
	}
}

// WithPartitionKey Ключ секционирования cookie.
func WithPartitionKey(partitionKey *network.CookiePartitionKey) Option {
	return func(c *Config) {
		c.PartitionKey = partitionKey
	}
}
