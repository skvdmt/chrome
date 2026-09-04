package copy_to

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

// Option Опция.
type Option func(c *Config)

// WithInsertBeforeNodeId Указать копию перед этим узлом.
func WithInsertBeforeNodeId(insertBeforeNodeId *dom.NodeId) Option {
	return func(c *Config) {
		c.InsertBeforeNodeId = insertBeforeNodeId
	}
}
