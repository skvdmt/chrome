package dispatch_key_event

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Config Конфигурация.
type Config struct {
	// Тип ключевого события.
	// Возможные значения: keyDown, keyUp, rawKeyDown, char
	Type string `json:"type"`
	// Битовое поле, представляющее нажатые клавиши-модификаторы.
	// Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (по умолчанию: 0).
	Modifiers int `json:"modifiers,omitempty"`
	// Время, когда произошло событие.
	Timestamp input.TimeSinceEpoch `json:"timestamp,omitempty"`
	// Текст, сгенерированный в результате обработки кода виртуальной клавиши с помощью
	// раскладки клавиатуры. Не требуется для событий keyUp и rawKeyDown (по умолчанию: "").
	Text string `json:"text,omitempty"`
	// Текст, который был бы сгенерирован клавиатурой, если бы не были
	// нажаты никакие модификаторы (кроме Shift). Полезно для обработки
	// сочетаний клавиш (акселераторов) (по умолчанию: "").
	UnmodifiedText string `json:"unmodifiedText,omitempty"`
	// Уникальный идентификатор клавиши (например, 'U+0041') (по умолчанию: "").
	KeyIdentifier string `json:"keyIdentifier,omitempty"`
	// Для каждого физического ключа (например, 'KeyA') задано
	// уникальное строковое значение, определяемое DOM (по умолчанию: "").
	Code string `json:"code,omitempty"`
	// Уникальное строковое значение, определяемое DOM, описывающее
	// значение клавиши в контексте активных модификаторов, раскладки
	// клавиатуры и т. д. (например, 'AltGr') (по умолчанию: "").
	Key string `json:"key,omitempty"`
	// Код виртуальной клавиши Windows (по умолчанию: 0).
	WindowsVirtualKeyCode int `json:"windowsVirtualKeyCode,omitempty"`
	// Код виртуальной клавиши (по умолчанию: 0).
	NativeVirtualKeyCode int `json:"nativeVirtualKeyCode,omitempty"`
	// Указывает, было ли событие сгенерировано с помощью
	// автоматического повтора (по умолчанию: false).
	AutoRepeat bool `json:"autoRepeat,omitempty"`
	// Указывает, было ли событие сгенерировано с клавиатуры (по умолчанию: false).
	IsKeypad bool `json:"isKeypad,omitempty"`
	// Указывает, является ли событие событием, связанным с
	// системным ключом (по умолчанию: false).
	IsSystemKey bool `json:"isSystemKey,omitempty"`
	// Указывает, произошло ли событие с левой или правой стороны
	// клавиатуры. 1 = слева, 2 = справа (по умолчанию: 0).
	Location int `json:"location,omitempty"`
	// Команды редактирования, отправляемые вместе с событием клавиши (например, 'selectAll')
	// (по умолчанию: []). Они связаны с именами команд, используемыми
	// в document.execCommand и NSStandardKeyBindingResponding, но не идентичны им.
	// См. https://source.chromium.org/chromium/chromium/src/+/main:third_party/blink/renderer/core/editing/commands/editor_command_names.h
	// для получения информации о допустимых именах команд.
	Commands []string `json:"commands,omitempty"`
}
