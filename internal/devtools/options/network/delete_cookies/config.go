package delete_cookies

import "github.com/skvdmt/chrome/internal/devtools/types/network"

// Config Конфигурация.
type Config struct {
	// Названия файлов cookie, которые нужно удалить.
	Name string `json:"name"`
	// Если указано, удаляются все файлы cookie с заданным именем,
	// где домен и путь соответствуют указанному URL-адресу.
	Url string `json:"url,omitempty"`
	// Если указано иное, удаляет только файлы cookie с точно таким же доменом.
	Domain string `json:"domain,omitempty"`
	// Если указано иное, удаляет только файлы cookie с точным путем.
	Path string `json:"path,omitempty"`
	// Если указано иное, удаляет только файлы cookie с заданным
	// именем и параметром partitionKey, где все атрибуты ключа
	// раздела соответствуют атрибуту ключа раздела файла cookie.
	PartitionKey *network.CookiePartitionKey `json:"partitionKey,omitempty"`
}
