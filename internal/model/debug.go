package model

import "fmt"

// Debug Отладчик.
type Debug struct {
	enable bool
}

// NewDebug Конструктор.
func NewDebug() *Debug {
	return &Debug{}
}

// Enable Включить.
func (d *Debug) Enable() {
	d.enable = true
}

// Disable Выключить.
func (d *Debug) Disable() {
	d.enable = false
}

// Debug Отладка.
func (d *Debug) Debug(message string) {
	if d.enable {
		fmt.Println(message)
	}
}
