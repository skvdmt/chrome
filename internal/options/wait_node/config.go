package wait_node

import (
	"time"

	"github.com/skvdmt/chrome/internal/devtools/options/dom/query_selector"
)

const (
	WAIT_NODE_TIMEOUT  = time.Second * 3
	WAIT_NODE_INTERVAL = time.Millisecond * 50
)

// Config Конфигурация.
type Config struct {
	Timeout              time.Duration
	Interval             time.Duration
	QuerySelectorOptions []query_selector.Option
}

// NewConfig Конструктор.
func NewConfig() *Config {
	return &Config{
		Timeout:  WAIT_NODE_TIMEOUT,
		Interval: WAIT_NODE_INTERVAL,
	}
}
