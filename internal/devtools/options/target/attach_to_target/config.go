package attach_to_target

import "github.com/skvdmt/chrome/internal/devtools/types/target"

// Config Конфигурация.
type Config struct {
	TargetId *target.TargetId `json:"targetId"`
	// Позволяет получить "плоский" доступ к сессии, указав
	// атрибут sessionId в командах. Мы планируем сделать это режимом
	// по умолчанию, отказаться от неплоского режима и в конечном итоге
	// полностью его вывести из эксплуатации. См. crbug.com/991325.
	Flatten bool `json:"flatten,omitempty"`
}
