package describe_node

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Config Конфигурация.
type Config struct {
	// Id узла.
	NodeId dom.NodeId `json:"nodeId,omitempty"`
	// Id бэкэнд-узла.
	BackendNodeId dom.BackendNodeId `json:"backendNodeId,omitempty"`
	// Id объекта JavaScript для обертки узла.
	ObjectId rnt.RemoteObjectId `json:"objectId,omitempty"`
	// Максимальная глубина, на которой следует извлекать
	// дочерние элементы, по умолчанию равна 1. Используйте -1 для
	// всего поддерева или укажите целое число больше 0.
	Depth int `json:"depth,omitempty"`
	// Указывает, следует ли обходить iframe и теневые корни
	// при возврате поддерева (по умолчанию — false).
	Pierce bool `json:"pierce,omitempty"`
}
