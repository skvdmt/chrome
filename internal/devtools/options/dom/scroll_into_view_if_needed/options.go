package scroll_into_view_if_needed

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Option Опция.
type Option func(c *Config)

// WithNodeId Указать id узла
func WithNodeId(nodeId *dom.NodeId) Option {
	return func(c *Config) {
		c.NodeId = nodeId
	}
}

// WithBackendNodeId Указать id бэкэнд-узла.
func WithBackendNodeId(backendNodeId *dom.BackendNodeId) Option {
	return func(c *Config) {
		c.BackendNodeId = backendNodeId
	}
}

// WithObjectId Указать id объекта JavaScript, представляющего собой обертку для узла.
func WithObjectId(objectId *rnt.RemoteObjectId) Option {
	return func(c *Config) {
		c.ObjectId = objectId
	}
}

// WithRect Указать прямоугольник.
func WithRect(rect *dom.Rect) Option {
	return func(c *Config) {
		c.Rect = rect
	}
}
