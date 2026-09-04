package focus

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
}
