package move_to

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

// Config Конфигурация.
type Config struct {
	// ID узла для перемещения.
	NodeId dom.NodeId `json:"nodeId"`
	// ID элемента, в который нужно переместить перемещаемый узел.
	TargetNodeId dom.NodeId `json:"targetNodeId"`
	// Удалить узел, предшествующий этому (если он отсутствует, перемещенный
	// узел становится последним дочерним узлом targetNodeId).
	InsertBeforeNodeId dom.NodeId `json:"insertBeforeNodeId,omitempty"`
}
