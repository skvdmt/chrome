package get_app_manifest

// Config Конфигурация.
type Config struct {
	ManifestId string `json:"manifestId"`
}

// Option Опция.
type Option func(c *Config)

// WithManifestId
func WithManifestId(manifestId string) Option {
	return func(c *Config) {
		c.ManifestId = manifestId
	}
}
