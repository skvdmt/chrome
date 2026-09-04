package scroll_into_view_if_needed

import (
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
)

// Config Конфигурация.
type Config struct {
	// Идентификатор узла.
	NodeId *dom.NodeId `json:"nodeId,omitempty"`
	// Идентификатор бэкэнд-узла.
	BackendNodeId *dom.BackendNodeId `json:"backendNodeId,omitempty"`
	// Идентификатор объекта JavaScript, представляющего собой обертку для узла.
	ObjectId *rnt.RemoteObjectId `json:"objectId,omitempty"`
	// Прямоугольник, который будет прокручиваться в поле зрения, относительно
	// границы узла, в пикселях CSS. Если этот параметр опущен, будет
	// использоваться центр узла, аналогично Element.scrollIntoView.
	Rect *dom.Rect `json:"rect,omitempty"`
}
