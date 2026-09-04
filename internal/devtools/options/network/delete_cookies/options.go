package delete_cookies

import "github.com/skvdmt/chrome/internal/devtools/types/network"

// Option Опция.
type Option func(c *Config)

// WithUrl Если указано, удаляются все файлы cookie с заданным именем,
// где домен и путь соответствуют указанному URL-адресу.
func WithUrl(url string) Option {
	return func(c *Config) {
		c.Url = url
	}
}

// WithDomain Если указано, удаляет только файлы cookie с точно таким же доменом.
func WithDomain(domain string) Option {
	return func(c *Config) {
		c.Domain = domain
	}
}

// WithPath Если указано, удаляет только файлы cookie с точным путем.
func WithPath(path string) Option {
	return func(c *Config) {
		c.Path = path
	}
}

// WithPartitionKey Если указано, удаляет только файлы cookie с заданным
// именем и параметром partitionKey, где все атрибуты ключа
// раздела соответствуют атрибуту ключа раздела файла cookie.
func WithPartitionKey(partitionKey *network.CookiePartitionKey) Option {
	return func(c *Config) {
		c.PartitionKey = partitionKey
	}
}
