package get_outer_html

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Option Опция.
type Option func(c *Config)

// WithNodeId Указание ID узла.
func WithNodeId(nodeId dom.NodeId) Option {
	return func(c *Config) {
		c.NodeId = nodeId
	}
}

// WithBackendNodeId Указание ID бэкэнд-узла.
func WithBackendNodeId(backendNodeId dom.BackendNodeId) Option {
	return func(c *Config) {
		c.BackendNodeId = backendNodeId
	}
}

// WithObjectId Указание ID объекта JavaScript для обертки узла.
func WithObjectId(objectId rnt.RemoteObjectId) Option {
	return func(c *Config) {
		c.ObjectId = objectId
	}
}

// WithIncludeShadowDOM Включить все теневые корни.
func WithIncludeShadowDOM() Option {
	return func(c *Config) {
		c.IncludeShadowDOM = true
	}
}
