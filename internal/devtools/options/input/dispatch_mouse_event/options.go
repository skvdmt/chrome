package dispatch_mouse_event

import "github.com/skvdmt/chrome/internal/devtools/types/input"

// Option Опция.
type Option func(c *Config)

// WithModifiers Битовое поле, представляющее нажатые клавиши-модификаторы.
func WithModifiers(modifiers int) Option {
	return func(c *Config) {
		c.Modifiers = modifiers
	}
}

// WithTimestamp Время, когда произошло событие.
func WithTimestamp(timestamp *input.TimeSinceEpoch) Option {
	return func(c *Config) {
		c.Timestamp = timestamp
	}
}

// WithButton Кнопка мыши.
func WithButton(button *input.MouseButton) Option {
	return func(c *Config) {
		c.Button = button
	}
}

// WithButtons Число, указывающее, какие кнопки мыши нажаты при
// срабатывании события мыши.
func WithButtons(buttons int) Option {
	return func(c *Config) {
		c.Buttons = buttons
	}
}

// WithClickCount Количество нажатий кнопки мыши.
func WithClickCount(clickCount int) Option {
	return func(c *Config) {
		c.ClickCount = clickCount
	}
}

// WithForce Нормализованное давление.
func WithForce(force int) Option {
	return func(c *Config) {
		c.Force = force
	}
}

// WithTangentialPressure Нормализованное тангенциальное давление.
func WithTangentialPressure(tangentialPressure int) Option {
	return func(c *Config) {
		c.TangentialPressure = tangentialPressure
	}
}

// WithTiltX Угол между плоскостью Y-Z и плоскостью, содержащей как
// ось стилуса, так и ось Y, в градусах в диапазоне [-90,90],
// положительный наклон X означает наклон вправо.
func WithTiltX(tiltX int) Option {
	return func(c *Config) {
		c.TiltX = tiltX
	}
}

// WithTiltY Угол между плоскостью X-Z и плоскостью, содержащей как
// ось стилуса, так и ось X, в градусах в диапазоне [-90,90],
// положительный наклон Y направлен в сторону пользователя.
func WithTiltY(tiltY int) Option {
	return func(c *Config) {
		c.TiltY = tiltY
	}
}

// WithTwist Вращение стилуса по часовой стрелке вокруг
// своей главной оси, в градусах в диапазоне [0,359].
func WithTwist(twist int) Option {
	return func(c *Config) {
		c.Twist = twist
	}
}

// WithDeltaX Изменение по оси X в пикселях CSS
// для события прокрутки колесика мыши.
func WithDeltaX(deltaX int) Option {
	return func(c *Config) {
		c.DeltaX = deltaX
	}
}

// WithDeltaY Изменение координаты Y в пикселях CSS
// для события прокрутки колесика мыши.
func WithDeltaY(deltaY int) Option {
	return func(c *Config) {
		c.DeltaY = deltaY
	}
}

// WithPointerType Тип указателя.
func WithPointerType(pointerType string) Option {
	return func(c *Config) {
		c.PointerType = pointerType
	}
}
