package reload

import "github.com/skvdmt/chrome/internal/devtools/types/network"

// Option Опция.
type Option func(c *Config)

// WithIgnoreCache Кэш браузера игнорируется
func WithIgnoreCache() Option {
	return func(c *Config) {
		c.IgnoreCache = true
	}
}

// WithScriptToEvaluateOnLoad скрипт будет внедрен во все фреймы проверяемой
// страницы после перезагрузки.
func WithScriptToEvaluateOnLoad(scriptToEvaluateOnLoad string) Option {
	return func(c *Config) {
		c.ScriptToEvaluateOnLoad = scriptToEvaluateOnLoad
	}
}

// WithLoaderId будет выдана ошибка, если
// идентификатор загрузчика основного фрейма целевой
// страницы не совпадает с указанным идентификатором.
func WithLoaderId(loaderId *network.LoaderId) Option {
	return func(c *Config) {
		c.LoaderId = loaderId
	}
}
