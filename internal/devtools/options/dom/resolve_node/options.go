package resolve_node

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Option Опция.
type Option func(c *Config)

// WithNodeId Указать id узла.
func WithNodeId(nodeId *dom.NodeId) Option {
	return func(c *Config) {
		c.NodeId = nodeId
	}
}

// WithBackendNodeId Указать id бэкэнда узла.
func WithBackendNodeId(BackendNodeId *dom.BackendNodeId) Option {
	return func(c *Config) {
		c.BackendNodeId = BackendNodeId
	}
}

// WithObjectGroup Указать имя группы.
func WithObjectGroup(objectGroup string) Option {
	return func(c *Config) {
		c.ObjectGroup = objectGroup
	}
}

// WithExecutionContextId Указать id контекста выполнения.
func WithExecutionContextId(executionContextId *rnt.ExecutionContextId) Option {
	return func(c *Config) {
		c.ExecutionContextId = *executionContextId
	}
}
