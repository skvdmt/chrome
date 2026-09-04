package describe_node

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Option Опиця.
type Option func(c *Config)

// WithNodeId По id узла.
func WithNodeId(nodeId dom.NodeId) Option {
	return func(c *Config) {
		c.NodeId = nodeId
	}
}

// WithBackendNodeId По id бэкэнд-узла.
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

// WithDepth Указание максимальной глубины, дочерних элементов.
func WithDepth(depth int) Option {
	return func(c *Config) {
		c.Depth = depth
	}
}

// WithPierce Обходить iframe и теневые корни.
func WithPierce() Option {
	return func(c *Config) {
		c.Pierce = true
	}
}
