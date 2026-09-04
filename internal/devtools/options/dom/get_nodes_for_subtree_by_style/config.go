package get_nodes_for_subtree_by_style

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

// Config Конфигурация.
type Config struct {
	// Идентификатор узла, указывающий на корень поддерева.
	NodeId *dom.NodeId `json:"nodeId"`
	// Стиль фильтрации узлов (включает узлы, если какое-либо
	// из свойств соответствует условию).
	ComputedStyles []*dom.CSSComputedStyleProperty `json:"computedStyles"`
	// Указывает, следует ли обрабатывать iframe и теневые корневые элементы
	// в одном и том же целевом объекте при возврате результатов (по умолчанию — false).
	Pierce bool `json:"pierce,omitempty"`
}
