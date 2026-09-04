package target

import (
	"github.com/skvdmt/chrome/internal/devtools/types/browser"
	"github.com/skvdmt/chrome/internal/devtools/types/page"
)

// Уникальный идентификатор подключенной отладочной сессии.
type SessionId string

// TargetId Id вкладки.
type TargetId string

// TargetInfo Информация вкладки.
type TargetInfo struct {
	TargetId *TargetId `json:"targetId"`
	// Список типов: https://source.chromium.org/chromium/chromium/src/+/main:content/browser/devtools/devtools_agent_host_impl.cc?ss=chromium&q=f:devtools%20-f:out%20%22::kTypeTab%5B%5D%22
	Type  string `json:"type"`
	Title string `json:"title"`
	Url   string `json:"url"`
	// Указывает, подключен ли к целевому объекту клиент.
	Attached bool `json:"attached,omitempty"`
	// Идентификатор родительского элемента, если таковой имеется.
	// Например, элемент "iframe" может иметь родительский элемент "page".
	ParentId *TargetId `json:"parentId,omitempty"`
	// Идентификатор цели открывателя.
	OpenerId *TargetId `json:"openerId,omitempty"`
	// Имеет ли целевой объект доступ к исходному окну.
	CanAccessOpener bool `json:"canAccessOpener"`
	// Идентификатор кадра исходного окна (устанавливается только
	// в том случае, если у целевого окна есть открывающее окно).
	OpenerFrameId *page.FrameId `json:"openerFrameId,omitempty"`
	// Идентификатор родительского фрейма, присутствующий для целей
	// типа «iframe» и «worker». Для вложенных рабочих процессов
	// это «предок» фрейма, создавший первый рабочий процесс во вложенной цепочке.
	ParentFrameId    *page.FrameId             `json:"parentFrameId,omitempty"`
	BrowserContextId *browser.BrowserContextID `json:"browserContextId,omitempty"`
	Subtype          string                    `json:"subtype,omitempty"`
}

// FilterEntry Фильтр, используемый в операциях
// запроса/обнаружения/автоматического подключения целевых объектов.
type FilterEntry struct {
	// Если задано, это приводит к исключению соответствующих целей из списка.
	Exclude bool `json:"exclude,omitempty"`
	// Если отсутствует, соответствует любому типу.
	Type string `json:"type,omitempty"`
}

// RemoteLocation
type RemoteLocation struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// WindowState Состояние целевого окна.
// Возможные значения: normal, minimized, maximized, fullscreen
type WindowState string
