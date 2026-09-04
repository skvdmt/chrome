package query_selector

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

// Option Опция.
type Option func(c *Config)

// WithNodeId С указанием узла в котором проводить поиск по селектору.
func WithNodeId(nodeId dom.NodeId) Option {
	return func(c *Config) {
		c.NodeId = nodeId
	}
}
