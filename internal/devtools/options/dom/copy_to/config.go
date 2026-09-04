package copy_to

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

// Config Конфигурация.
type Config struct {
	// Идентификатор узла, который нужно скопировать.
	NodeId *dom.NodeId `json:"nodeId"`
	// Идентификатор элемента, в который нужно вставить копию.
	TargetNodeId *dom.NodeId `json:"targetNodeId"`
	// Удалите копию перед этим узлом (если она отсутствует,
	// копия становится последним дочерним узлом targetNodeId).
	InsertBeforeNodeId *dom.NodeId `json:"insertBeforeNodeId,omitempty"`
}
