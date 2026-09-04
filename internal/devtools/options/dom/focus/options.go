package focus

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Option Опция.
type Option func(c *Config)

// FocusNodeId Фокус на id узла.
func FocusNodeId(nodeId dom.NodeId) Option {
	return func(c *Config) {
		c.NodeId = nodeId
	}
}

// FocusBackendNodeId Фокус id бэкэнд-узла.
func FocusBackendNodeId(backendNodeId dom.BackendNodeId) Option {
	return func(c *Config) {
		c.BackendNodeId = backendNodeId
	}
}

// FocusObjectId Фокус id объекта JavaScript для обертки узла.
func FocusObjectId(objectId rnt.RemoteObjectId) Option {
	return func(c *Config) {
		c.ObjectId = objectId
	}
}
