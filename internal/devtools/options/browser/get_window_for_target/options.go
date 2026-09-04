package get_window_for_target

import "github.com/skvdmt/chrome/internal/devtools/types/target"

// Option Опция.
type Option func(c *Config)

// WithTargetId Указание id вкладки.
func WithTargetId(targetId target.TargetId) Option {
	return func(c *Config) {
		c.TargetId = targetId
	}
}
