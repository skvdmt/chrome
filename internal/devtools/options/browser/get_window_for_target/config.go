package get_window_for_target

import "github.com/skvdmt/chrome/internal/devtools/types/target"

// GetWindowForTargetConfig Конфигурация.
type Config struct {
	// Идентификатор хоста агента Devtools. Если вызов осуществляется
	// в рамках сессии, используется связанный targetId.
	TargetId target.TargetId `json:"targetId,omitempty"`
}
