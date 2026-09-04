package chrome

import (
	"time"
)

const (
	// Таймаут клиента.
	CLIENT_TIMEOUT = time.Second * 10
	// Таймаут поиска узла по селектору.
	NODE_BY_SELECTOR_TIMEOUT = time.Second * 3
	// Интервал поиска узла по селектору.
	NODE_BY_SELECTOR_INTERVAL = time.Millisecond * 50
)

// Config Конфигурация.
type Config struct {
	ClientTimeout          time.Duration `yaml:"client_timeout"`
	NodeBySelectorTimeout  time.Duration `yaml:"node_by_selector_timeout"`
	NodeBySelectorInterval time.Duration `yaml:"node_by_selector_interval"`
}

// NewConfig Конструктор.
func NewConfig() *Config {
	return &Config{
		ClientTimeout:          CLIENT_TIMEOUT,
		NodeBySelectorTimeout:  NODE_BY_SELECTOR_TIMEOUT,
		NodeBySelectorInterval: NODE_BY_SELECTOR_INTERVAL,
	}
}
