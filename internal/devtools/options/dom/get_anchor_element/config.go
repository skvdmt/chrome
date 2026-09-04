package get_anchor_element

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

// Config Конфигурация.
type Config struct {
	// Идентификатор позиционированного элемента, от которого
	// нужно отталкиваться при определении точки привязки.
	NodeId *dom.NodeId `json:"nodeId"`
	// Необязательный спецификатор привязки, как определено
	// в https://www.w3.org/TR/css-anchor-position-1/#anchor-specifier.
	// Если он не указан, он вернет неявный элемент привязки для
	// заданного позиционированного элемента.
	AnchorSpecifier string `json:"anchorSpecifier,omitempty"`
}
