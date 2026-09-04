package get_outer_html

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Config Конфигурация.
type Config struct {
	// ID узла.
	NodeId dom.NodeId `json:"nodeId,omitempty"`
	// ID бэкэнд-узла.
	BackendNodeId dom.BackendNodeId `json:"backendNodeId,omitempty"`
	// ID объекта JavaScript для обертки узла.
	ObjectId rnt.RemoteObjectId `json:"objectId,omitempty"`
	// Включить все теневые корни. Если не указано иное, равно false.
	IncludeShadowDOM bool `json:"includeShadowDOM,omitempty"`
}
