package capture_screenshot

import "github.com/skvdmt/chrome/internal/devtools/types/page"

// Config Конфигурация.
type Config struct {
	// Формат сжатия изображения (по умолчанию — png).
	// Допустимые значения: jpeg, png, webp
	Format string `json:"format,omitempty"`
	// Качество сжатия в диапазоне [0..100] (только JPEG).
	Quality int `json:"quality,omitempty"`
	// Сделайте снимок экрана только выбраной области.
	Clip *page.Viewport `json:"clip,omitempty"`
	// Делать снимок экрана с поверхности, а не из области обзора. По умолчанию — true.
	FromSurface bool `json:"fromSurface,omitempty"`
	// Делать снимок экрана за пределами области просмотра. По умолчанию — false.
	CaptureBeyondViewport bool `json:"captureBeyondViewport,omitempty"`
	// Оптимизировать кодирование изображения по скорости,
	// а не по размеру (по умолчанию - false).
	OptimizeForSpeed bool `json:"optimizeForSpeed,omitempty"`
}

// Option Опция.
type Option func(c *Config)

// WithFormat Формат сжатия изображения (по умолчанию — png).
func WithFormat(format string) Option {
	return func(c *Config) {
		c.Format = format
	}
}

// WithQuality Качество сжатия в диапазоне [0..100] (только JPEG).
func WithQuality(quality int) Option {
	return func(c *Config) {
		c.Quality = quality
	}
}

// WithClip Сделайте снимок экрана только выбраной области.
func WithClip(clip *page.Viewport) Option {
	return func(c *Config) {
		c.Clip = clip
	}
}

// WithFromSurface Делать снимок экрана с поверхности, а не из области обзора.
func WithFromSurface(fromSurface bool) Option {
	return func(c *Config) {
		c.FromSurface = fromSurface
	}
}

// WithCaptureBeyondViewport Делать снимок экрана за пределами области просмотра.
func WithCaptureBeyondViewport() Option {
	return func(c *Config) {
		c.CaptureBeyondViewport = true
	}
}

// WithOptimizeForSpeed Оптимизировать кодирование изображения по скорости.
func WithOptimizeForSpeed() Option {
	return func(c *Config) {
		c.OptimizeForSpeed = true
	}
}
