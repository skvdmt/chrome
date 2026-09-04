package get_box_model

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Config Конфигурация.
type Config struct {
	// Id узла.
	NodeId dom.NodeId `json:"nodeId,omitempty"`
	// Id бэкенд-узла.
	BackendNodeId dom.BackendNodeId `json:"backendNodeId,omitempty"`
	// Id объекта JavaScript для обертки узла.
	ObjectId rnt.RemoteObjectId `json:"objectId,omitempty"`
}
