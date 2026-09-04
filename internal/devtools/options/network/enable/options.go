package enable

// Option Опция.
type Option func(c *Config)

// WithMaxTotalBufferSize Размер буфера (в байтах) для
// сохранения полезной нагрузки сетевых запросов.
func WithMaxTotalBufferSize(maxTotalBufferSize int) Option {
	return func(c *Config) {
		c.MaxTotalBufferSize = maxTotalBufferSize
	}
}

// WithMaxResourceBufferSize Размер буфера (в байтах) на каждый ресурс,
// используемый при сохранении сетевых полезных данных.
func WithMaxResourceBufferSize(maxResourceBufferSize int) Option {
	return func(c *Config) {
		c.MaxResourceBufferSize = maxResourceBufferSize
	}
}

// WithMaxPostDataSize Максимальный размер тела запроса в байтах.
func WithMaxPostDataSize(maxPostDataSize int) Option {
	return func(c *Config) {
		c.MaxPostDataSize = maxPostDataSize
	}
}

// WithReportDirectSocketTraffic Сообщать о событиях отправки
// и получения фрагментов DirectSocket.
func WithReportDirectSocketTraffic() Option {
	return func(c *Config) {
		c.ReportDirectSocketTraffic = true
	}
}

// WithEnableDurableMessages Включает сохранение тел ответов вне процесса рендеринга.
func WithEnableDurableMessages() Option {
	return func(c *Config) {
		c.EnableDurableMessages = true
	}
}
