package create_target

import (
	"github.com/skvdmt/chrome/internal/devtools/types/browser"
	"github.com/skvdmt/chrome/internal/devtools/types/target"
)

// Option Опция.
type Option func(c *Config)

// WithLeft Указывает фрейм слева от начала координат в DIP.
func WithLeft(left int) Option {
	return func(c *Config) {
		c.Left = left
	}
}

// WithTop Указывает фрейм сверху от начала координат в DIP.
func WithTop(top int) Option {
	return func(c *Config) {
		c.Top = top
	}
}

// WithWidth Указывает ширину рамки в DIP.
func WithWidth(width int) Option {
	return func(c *Config) {
		c.Width = width
	}
}

// WithHeight Указывает высоту рамки в DIP.
func WithHeight(height int) Option {
	return func(c *Config) {
		c.Height = height
	}
}

// WithWindowState Указывает состояние окна фрейма.
func WithWindowState(windowState *target.WindowState) Option {
	return func(c *Config) {
		c.WindowState = windowState
	}
}

// WithBrowserContextId Указывает контекст браузера, в котором будет создана вкладка.
func WithBrowserContextId(browserContextId *browser.BrowserContextID) Option {
	return func(c *Config) {
		c.BrowserContextId = browserContextId
	}
}

// WithEnableBeginFrameControl Будет ли управление BeginFrames
// для этой вкладки осуществляться через DevTools.
func WithEnableBeginFrameControl() Option {
	return func(c *Config) {
		c.EnableBeginFrameControl = true
	}
}

// WithNewWindow Создавать ли новое окно.
func WithNewWindow() Option {
	return func(c *Config) {
		c.NewWindow = true
	}
}

// WithBackground Cоздавать вкладку в фоновом режиме.
func WithBackground() Option {
	return func(c *Config) {
		c.Background = true
	}
}

// WithForTab Cоздавать целевой объект типа "вкладка".
func WithForTab() Option {
	return func(c *Config) {
		c.ForTab = true
	}
}

// WithHidden Cоздаваемая вкладка будет скрытой.
func WithHidden() Option {
	return func(c *Config) {
		c.Hidden = true
	}
}

// WithFocus Фокусироваться на новой вкладке.
func WithFocus() Option {
	return func(c *Config) {
		c.Focus = true
	}
}
