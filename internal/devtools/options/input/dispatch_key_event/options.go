package dispatch_key_event

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Option Опция.
type Option func(c *Config)

// WithModifiers Указывает битовое поле, представляющее нажатые клавиши-модификаторы.
func WithModifiers(modifiers int) Option {
	return func(c *Config) {
		c.Modifiers = modifiers
	}
}

// WithTimestamp Указывает время, когда произошло событие.
func WithTimestamp(timestamp input.TimeSinceEpoch) Option {
	return func(c *Config) {
		c.Timestamp = timestamp
	}
}

// WithText Указывает Текст, сгенерированный в результате обработки
// кода виртуальной клавиши с помощью раскладки клавиатуры.
func WithText(text string) Option {
	return func(c *Config) {
		c.Text = text
	}
}

// WithUnmodifiedText Указывает текст, который был бы сгенерирован клавиатурой.
func WithUnmodifiedText(unmodifiedText string) Option {
	return func(c *Config) {
		c.UnmodifiedText = unmodifiedText
	}
}

// WithKeyIdentifier Указывает уникальный идентификатор клавиши.
func WithKeyIdentifier(keyIdentifier string) Option {
	return func(c *Config) {
		c.KeyIdentifier = keyIdentifier
	}
}

// WithCode Указывает уникальное строковое значение, определяемое DOM.
func WithCode(code string) Option {
	return func(c *Config) {
		c.Code = code
	}
}

// WithKey Указывает строковое значение клавиши.
func WithKey(key string) Option {
	return func(c *Config) {
		c.Key = key
	}
}

// WithWindowsVirtualKeyCode Указывает код виртуальной клавиши Windows.
func WithWindowsVirtualKeyCode(windowsVirtualKeyCode int) Option {
	return func(c *Config) {
		c.WindowsVirtualKeyCode = windowsVirtualKeyCode
	}
}

// WithNativeVirtualKeyCode Указывает код виртуальной клавиши.
func WithNativeVirtualKeyCode(nativeVirtualKeyCode int) Option {
	return func(c *Config) {
		c.NativeVirtualKeyCode = nativeVirtualKeyCode
	}
}

// WithAutoRepeat Указывает автоматический повтор.
func WithAutoRepeat() Option {
	return func(c *Config) {
		c.AutoRepeat = true
	}
}

// WithIsKeypad Указывает, что событие сгенерировано с клавиатуры.
func WithIsKeypad() Option {
	return func(c *Config) {
		c.IsKeypad = true
	}
}

// WithIsSystemKey Указывает наличие системного ключа.
func WithIsSystemKey() Option {
	return func(c *Config) {
		c.IsSystemKey = true
	}
}

// WithLocation Указывает сторону клавиатуры.
func WithLocation(location int) Option {
	return func(c *Config) {
		c.Location = location
	}
}

// WithCommands Указывает команды редактирования.
func WithCommands(commands []string) Option {
	return func(c *Config) {
		c.Commands = commands
	}
}
