package get_cookies

// Config Конфигурация.
type Config struct {
	// Список URL-адресов, для которых будут получены соответствующие
	// файлы cookie. Если список не указан, подразумевается, что он
	// включает URL-адрес страницы и всех её вложенных фреймов.
	Urls []string `json:"urls,omitempty"`
}
