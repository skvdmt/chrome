package reload

import "github.com/skvdmt/chrome/internal/devtools/types/network"

// Config Конфигурация.
type Config struct {
	// Если true, кэш браузера игнорируется (как если бы
	// пользователь нажал Shift + обновить страницу).
	IgnoreCache bool `json:"ignoreCache,omitempty"`
	// Если задано, скрипт будет внедрен во все фреймы проверяемой
	// страницы после перезагрузки. Этот аргумент будет проигнорирован
	// при перезагрузке источника dataURL.
	ScriptToEvaluateOnLoad string `json:"scriptToEvaluateOnLoad,omitempty"`
	// Если этот параметр задан, будет выдана ошибка, если
	// идентификатор загрузчика основного фрейма целевой
	// страницы не совпадает с указанным идентификатором.
	// Это предотвращает случайную перезагрузку непреднамеренной
	// цели в случае возникновения конфликтов навигации.
	LoaderId *network.LoaderId `json:"loaderId,omitempty"`
}
