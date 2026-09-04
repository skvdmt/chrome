package query_selector

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

// Config Конфигурация.
type Config struct {
	NodeId   dom.NodeId `json:"nodeId"`
	Selector string     `json:"selector"`
}
