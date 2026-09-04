package resolve_node

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Config Конфигурация.
type Config struct {
	// Идентификатор узла, который необходимо разрешить.
	NodeId *dom.NodeId `json:"nodeId,omitempty"`
	// Идентификатор бэкэнда узла, который необходимо разрешить.
	BackendNodeId *dom.BackendNodeId `json:"backendNodeId,omitempty"`
	// Символическое имя группы, которое можно использовать
	// для освобождения нескольких объектов.
	ObjectGroup string `json:"objectGroup,omitempty"`
	// Контекст выполнения, в котором следует разрешить узел.
	ExecutionContextId rnt.ExecutionContextId `json:"executionContextId,omitempty"`
}
