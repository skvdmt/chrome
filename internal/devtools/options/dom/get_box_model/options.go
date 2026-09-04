package get_box_model

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Option Опция.
type Option func(c *Config)

// WithNodeId По id узла
func WithNodeId(nodeId dom.NodeId) Option {
	return func(c *Config) {
		c.NodeId = nodeId
	}
}

// WithBackendNodeId По id бэкенд-узла
func WithBackendNodeId(backendNodeId dom.BackendNodeId) Option {
	return func(c *Config) {
		c.BackendNodeId = backendNodeId
	}
}

// WithObjectId По id объекта JavaScript для обертки узла.
func WithObjectId(objectId rnt.RemoteObjectId) Option {
	return func(c *Config) {
		c.ObjectId = objectId
	}
}
