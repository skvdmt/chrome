package get_container_for_node

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

// Option Опция
type Option func(c *Config)

// WithContainerName Указать имя контейнера.
func WithContainerName(containerName string) Option {
	return func(c *Config) {
		c.ContainerName = containerName
	}
}

// WithPhysicalAxes Указать физические оси.
func WithPhysicalAxes(physicalAxes *dom.PhysicalAxes) Option {
	return func(c *Config) {
		c.PhysicalAxes = physicalAxes
	}
}

// WithLogicalAxes Указать логические оси.
func WithLogicalAxes(logicalAxes *dom.LogicalAxes) Option {
	return func(c *Config) {
		c.LogicalAxes = logicalAxes
	}
}

// WithQueriesScrollState С указанием запросов состояния прокрутки.
func WithQueriesScrollState() Option {
	return func(c *Config) {
		c.QueriesScrollState = true
	}
}

// WithQueriesAnchored С указанием закрепленных запросов.
func WithQueriesAnchored() Option {
	return func(c *Config) {
		c.QueriesAnchored = true

	}
}
