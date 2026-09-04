package reset_permissions

import "github.com/skvdmt/chrome/internal/devtools/types/browser"

// Config Конфигурация.
type Config struct {
	// Параметр BrowserContext используется для сброса разрешений.
	// Если он опущен, используется контекст браузера по умолчанию.
	// Указывается с помощью метода ResetPermissionsWithBrowserContextID.
	BrowserContextID browser.BrowserContextID `json:"browserContextID,omitempty"`
}
