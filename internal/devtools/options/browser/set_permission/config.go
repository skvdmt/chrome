package set_permission

import "github.com/skvdmt/chrome/internal/devtools/types/browser"

// Config Конфигурация.
type Config struct {
	// Описание разрешения на отмену.
	Permission browser.PermissionDescriptor `json:"permission"`
	// Настройка разрешений.
	Setting browser.PermissionSetting `json:"setting"`
	// Разрешение распространяется на все источники встраивания, если не указано иное.
	Origin string `json:"origin,omitempty"`
	// Встроенный источник, к которому относится разрешение. Он игнорируется,
	// если встроенный источник отсутствует и недействителен. Если встроенный
	// источник указан, но сам встроенный источник отсутствует, в качестве
	// встроенного источника используется встроенный источник.
	EmbeddedOrigin string `json:"embeddedOrigin,omitempty"`
	// Контекст для переопределения. Если он опущен,
	// используется контекст браузера по умолчанию.
	BrowserContextId browser.BrowserContextID `json:"browserContextId,omitempty"`
}
