package set_file_input_files

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Option Опция.
type Option func(c *Config)

// WithNodeId Указание id узла.
func WithNodeId(nodeId *dom.NodeId) Option {
	return func(c *Config) {
		c.NodeId = nodeId
	}
}

// WithBackendNodeId Указание id бэкэнд-узла.
func WithBackendNodeId(backendNodeId *dom.BackendNodeId) Option {
	return func(c *Config) {
		c.BackendNodeId = backendNodeId
	}
}

// WithObjectId Указание id объекта JavaScript, представляющего собой обертку для узла.
func With(objectId *rnt.RemoteObjectId) Option {
	return func(c *Config) {
		c.ObjectId = objectId
	}
}
