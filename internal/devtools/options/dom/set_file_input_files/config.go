package set_file_input_files

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Config Конфигурация.
type Config struct {
	// Массив путей к файлам для установки.
	Files []string `json:"files"`
	// Идентификатор узла.
	NodeId *dom.NodeId `json:"nodeId,omitempty"`
	// Идентификатор бэкэнд-узла.
	BackendNodeId *dom.BackendNodeId `json:"backendNodeId,omitempty"`
	// Идентификатор объекта JavaScript, представляющего собой обертку для узла.
	ObjectId *rnt.RemoteObjectId `json:"objectId,omitempty"`
}
