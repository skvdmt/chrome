package get_content_quads

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Config Конфигурация.
type Config struct {
	// Идентификатор узла.
	NodeId *dom.NodeId `json:"nodeId,omitempty"`
	// Идентификатор бэкэнд-узла.
	BackendNodeId *dom.BackendNodeId `json:"backendNodeId,omitempty"`
	// Идентификатор объекта JavaScript, представляющего собой обертку для узла.
	ObjectId *rnt.RemoteObjectId `json:"objectId,omitempty"`
}

// Option Опция.
type Option func(c *Config)

// WithNodeId Указать идентификатор узла.
func WithNodeId(nodeId *dom.NodeId) Option {
	return func(c *Config) {
		c.NodeId = nodeId
	}
}

// WithBackendNodeId Указать идентификатор бэкэнд-узла.
func WithBackendNodeId(backendNodeId *dom.BackendNodeId) Option {
	return func(c *Config) {
		c.BackendNodeId = backendNodeId
	}
}

// WithObjectId Указать идентификатор объекта JavaScript.
func WithObjectId(objectId *rnt.RemoteObjectId) Option {
	return func(c *Config) {
		c.ObjectId = objectId
	}
}
