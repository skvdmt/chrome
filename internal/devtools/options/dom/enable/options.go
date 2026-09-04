package enable

const (
	NONE = "none"
	ALL  = "all"
)

// Option Опиця.
type Option func(c *Config)

// EnableIncludeWhitespaceNone Не включать пробелы в массив дочерних
// узлов, возвращаемых системой Node.
func EnableIncludeWhitespaceNone() Option {
	return func(c *Config) {
		c.IncludeWhitespace = NONE
	}
}

// EnableIncludeWhitespaceAll Включать пробелы в массив дочерних
// узлов, возвращаемых системой Node.
func EnableIncludeWhitespaceAll() Option {
	return func(c *Config) {
		c.IncludeWhitespace = ALL
	}
}
