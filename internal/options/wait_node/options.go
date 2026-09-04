package wait_node

import (
	"time"

	"github.com/skvdmt/chrome/internal/devtools/options/dom/query_selector"
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
)

// Option Опция.
type Option func(c *Config)

// WithTimeout Установка таймаута ожидания узла.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithInterval Установка интервала ожидания узла.
func WithInterval(interval time.Duration) Option {
	return func(c *Config) {
		c.Interval = interval
	}
}

// WithNodeId Где корневой узел равен аргументу.
func WithNodeId(nodeId dom.NodeId) Option {
	return func(c *Config) {
		c.QuerySelectorOptions = append(c.QuerySelectorOptions, query_selector.WithNodeId(nodeId))
	}
}
