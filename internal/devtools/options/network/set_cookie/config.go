package set_cookie

import "github.com/skvdmt/chrome/internal/devtools/types/network"

// Config Конфигурация.
type Config struct {
	// Имя файла cookie.
	Name string `json:"name"`
	// Значение cookie.
	Value string `json:"value"`
	// Request-URI, который следует связать с установкой cookie.
	// Это значение может влиять на значения домена, пути, исходного порта
	// и исходной схемы по умолчанию для создаваемого cookie.
	Url string `json:"url,omitempty"`
	// Домен cookie.
	Domain string `json:"domain,omitempty"`
	// Путь cookie.
	Path string `json:"path,omitempty"`
	// True, если cookie защищен.
	Secure bool `json:"secure,omitempty"`
	// Значение true, если cookie доступен только по протоколу HTTP.
	HttpOnly bool `json:"httpOnly,omitempty"`
	// Тип SameSite для cookie.
	SameSite *network.CookieSameSite `json:"sameSite,omitempty"`
	// Дата истечения срока действия cookie; если не задана — сеансовый cookie.
	Expires *network.TimeSinceEpoch `json:"expires,omitempty"`
	// Тип приоритета cookie.
	Priority *network.CookiePriority `json:"priority,omitempty"`
	// Тип схемы источника cookie.
	SourceScheme *network.CookieSourceScheme `json:"sourceScheme"`
	// Исходный порт cookie. Допустимые значения: {-1, [1, 65535]}; значение -1
	// означает, что порт не указан. Отсутствие указанного порта позволяет клиентам
	// протокола эмулировать устаревший механизм определения области действия cookie,
	// привязанный к порту. Эта возможность является временной и в будущем будет удалена.
	SourcePort int `json:"sourcePort,omitempty"`
	// Ключ секционирования cookie. Если он не задан, cookie
	// будет установлено как несекционированное.
	PartitionKey *network.CookiePartitionKey `json:"partitionKey,omitempty"`
}
