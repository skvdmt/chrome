package chrome

import "github.com/skvdmt/chrome/internal/devtools/types/dom"

const (
	LEFT_BUTTON   = "left"
	RIGHT_BUTTON  = "right"
	MIDDLE_BUTTON = "middle"
)

// Click Клик.
func (d *Driver) Click(nodeId dom.NodeId, option ...ClickOption) error {
	c := NewClickConfig()
	for _, o := range option {
		if err := o(c); err != nil {
			return err
		}
	}
	return nil
}

// ClickOption Опция.
type ClickOption func(c *ClickConfig) error

// ClickConfig Конфигурация.
type ClickConfig struct {
	button string
}

// NewClickConfig Конструктор.
func NewClickConfig() *ClickConfig {
	return &ClickConfig{
		button: LEFT_BUTTON,
	}
}

// RightButton Клик правой кнопкой мыши.
func RightButton() ClickOption {
	return func(c *ClickConfig) error {
		c.button = RIGHT_BUTTON
		return nil
	}
}
