package navigate

import "github.com/skvdmt/chrome/internal/devtools/types/page"

// Config Конфигурация.
type Config struct {
	// URL-адрес для перехода на страницу.
	Url string `json:"url"`
	// URL источника перехода.
	Referrer string `json:"referrer,omitempty"`
	// Предполагаемый тип перехода.
	TransitionType *page.TransitionType `json:"transitionType,omitempty"`
	// Идентификатор кадра для навигации; если не указан,
	// осуществляется навигация по верхнему кадру.
	FrameId *page.FrameId `json:"frameId,omitempty"`
	// Для навигации используется Referrer-policy.
	ReferrerPolicy *page.ReferrerPolicy `json:"referrerPolicy,omitempty"`
}

// Option Опция.
type Option func(c *Config)

// WithReferrer URL источника перехода.
func WithReferrer(referrer string) Option {
	return func(c *Config) {
		c.Referrer = referrer
	}
}

// WithTransitionType Предполагаемый тип перехода.
func WithTransitionType(transitionType *page.TransitionType) Option {
	return func(c *Config) {
		c.TransitionType = transitionType
	}
}

// WithFrameId Идентификатор кадра для навигации.
func WithFrameId(frameId *page.FrameId) Option {
	return func(c *Config) {
		c.FrameId = frameId
	}
}

// WithReferrerPolicy Для навигации используется Referrer-policy.
func WithReferrerPolicy(referrerPolicy *page.ReferrerPolicy) Option {
	return func(c *Config) {
		c.ReferrerPolicy = referrerPolicy
	}
}
