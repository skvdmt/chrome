package create_target

import (
	"github.com/skvdmt/chrome/internal/devtools/types/browser"
	"github.com/skvdmt/chrome/internal/devtools/types/target"
)

// Config Конфигурация.
type Config struct {
	// Исходный URL-адрес, на который будет перенаправлена ​​страница.
	// Пустая строка указывает на about:blank.
	Url string `json:"url"`
	// Фрейм слева от начала координат в DIP (требуется, чтобы newWindow
	// было true или использовалась безголовая оболочка).
	Left int `json:"left,omitempty"`
	// Фрейм сверху от начала координат в DIP (требуется, чтобы newWindow
	// было true или использовалась безголовая оболочка).
	Top int `json:"top,omitempty"`
	// Ширина рамки в DIP (требуется, чтобы newWindow
	// было true или использовалась безголовая оболочка).
	Width int `json:"width,omitempty"`
	// Высота рамки в DIP (требуется, чтобы newWindow
	// было true или использовалась безголовая оболочка).
	Height int `json:"height,omitempty"`
	// Состояние окна фрейма (требуется, чтобы newWindow
	// было true или использовалась безголовая оболочка). По умолчанию — нормальное.
	WindowState *target.WindowState `json:"windowState,omitempty"`
	// Контекст браузера, в котором будет создана страница.
	BrowserContextId *browser.BrowserContextID `json:"browserContextId,omitempty"`
	// Будет ли управление BeginFrames для этой цели
	// осуществляться через DevTools (только в безголовой оболочке,
	// пока не поддерживается на MacOS, по умолчанию false).
	EnableBeginFrameControl bool `json:"enableBeginFrameControl,omitempty"`
	// Создавать ли новое окно или вкладку (по умолчанию - false,
	// не поддерживается в безголовой оболочке).
	NewWindow bool `json:"newWindow,omitempty"`
	// Определяет, создавать ли целевой объект в фоновом или активном режиме
	// (по умолчанию false, не поддерживается безголовой оболочкой).
	Background bool `json:"background,omitempty"`
	// Следует ли создавать целевой объект типа "вкладка".
	ForTab bool `json:"forTab,omitempty"`
	// Создавать ли скрытый объект. Скрытый объект доступен для наблюдения
	// по протоколу, но не отображается в панели пользовательского
	// интерфейса вкладок. Не может быть создан с параметрами
	// forTab: true, newWindow: true или background: false.
	// Время жизни вкладки ограничено временем жизни сессии.
	Hidden bool `json:"hidden,omitempty"`
	// Если указано, определяет, следует ли фокусироваться на новом целевом объекте.
	// По умолчанию поведение фокусировки зависит от параметра фона:
	// - Если background равно false (по умолчанию) и focus не задан, фокус
	//   устанавливается на новый целевой объект, и окно браузера выводится
	//   на передний план.
	// - Если background равно false и focus равно false, целевой объект
	//   открывается, но фокус окна браузера остается неизменным (например,
	//   если окно находилось на заднем плане, оно там и останется).
	// - Если background равно true, установка focus в true не поддерживается
	//   и приведет к ошибке.
	Focus bool `json:"focus,omitempty"`
}
