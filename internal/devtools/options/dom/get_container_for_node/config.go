package get_container_for_node

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

// Config Конфигурация.
type Config struct {
	NodeId             *dom.NodeId       `json:"nodeId"`
	ContainerName      string            `json:"containerName,omitempty"`
	PhysicalAxes       *dom.PhysicalAxes `json:"physicalAxes,omitempty"`
	LogicalAxes        *dom.LogicalAxes  `json:"logicalAxes,omitempty"`
	QueriesScrollState bool              `json:"queriesScrollState,omitempty"`
	QueriesAnchored    bool              `json:"queriesAnchored,omitempty"`
}
