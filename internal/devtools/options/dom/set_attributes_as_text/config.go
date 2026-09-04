package set_attributes_as_text

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

// Config Конфигурация.
type Confug struct {
	// Идентификатор элемента, для которого нужно задать атрибуты.
	NodeId *dom.NodeId `json:"nodeId"`
	// Текст, содержащий ряд атрибутов. Будет проанализирован с помощью HTML-парсера.
	Text string `json:"text"`
	// Имя атрибута, которое будет заменено новыми атрибутами,
	// полученными из текста, в случае успешного анализа текста.
	Name string `json:"name,omitempty"`
}
