package force_show_popover

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

// Config Конфигурация.
type Config struct {
	// Идентификатор всплывающего HTML-элемента.
	NodeId *dom.NodeId `json:"nodeId"`
	// Если значение равно true, всплывающее окно открывается и остается
	// открытым. Если значение равно false, всплывающее окно закрывается,
	// если оно было принудительно открыто ранее.
	Enable bool `json:"enable"`
	// Необязательный идентификатор элемента, вызывающего это
	// всплывающее окно, используемый для установления неявной привязки.
	// Если он не указан, будет использоваться первый вызывающий элемент
	// в документе, отдавая предпочтение элементам с атрибутом popovertarget
	// перед элементами с атрибутом commandfor. Обратите внимание, что если
	// вызывающих элементов несколько, это всего лишь приблизительная оценка.
	InvokerNodeId *dom.BackendNodeId `json:"invokerNodeId,omitempty"`
}

// Option Опция.
type Option func(c *Config)

// WithInvokerNodeId Указать необязательный идентификатор элемента
// вызывающего это всплывающее окно.
func WithInvokerNodeId(invokerNodeId *dom.BackendNodeId) Option {
	return func(c *Config) {
		c.InvokerNodeId = invokerNodeId
	}
}
