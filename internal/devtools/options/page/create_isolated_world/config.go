package create_isolated_world

import "github.com/skvdmt/chrome/internal/devtools/types/page"

// Config Конфигурация.
type Config struct {
	// Идентификатор фрейма, в котором должен быть создан изолированный мир.
	FrameId *page.FrameId `json:"frameId"`
	// Необязательное имя, которое указывается в контексте выполнения.
	WorldName string `json:"worldName,omitempty"`
	// Вопрос о том, следует ли предоставлять всеобщий доступ изолированному миру.
	// Это мощный инструмент, но его следует использовать с осторожностью.
	GrantUniveralAccess bool `json:"grantUniveralAccess,omitempty"`
	// Необязательная политика безопасности контента (CSP) для изолированной среды.
	// Если она опущена, все существующие политики CSP для этой среды будут удалены.
	// Обратите внимание, что удаление или обновление CSP не сразу влияет на активный
	// контекст в том же документе, поскольку LocalDOMWindow кэширует объект
	// ContentSecurityPolicy. Изменение вступает в силу при последующих переходах
	// при создании нового контекста окна.
	ContentSecurityPolicy string `json:"contentSecurityPolicy,omitempty"`
}
