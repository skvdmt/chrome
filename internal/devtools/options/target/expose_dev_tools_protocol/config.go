package expose_dev_tools_protocol

import "github.com/skvdmt/chrome/internal/devtools/types/target"

// Config Конфигурация.
type Config struct {
	TargetId *target.TargetId `json:"targetId"`
	// Имя привязки, 'cdp', если не указано.
	BindingName string `json:"bindingName,omitempty"`
	// Если значение равно true, наследуются права доступа
	// текущей корневой сессии (по умолчанию: false).
	InheritPermissions bool `json:"inheritPermissions,omitempty"`
}
