package domains

import (
	"github.com/skvdmt/chrome/internal/devtools/options/network/delete_cookies"
	"github.com/skvdmt/chrome/internal/devtools/options/network/enable"
	"github.com/skvdmt/chrome/internal/devtools/options/network/get_cookies"
	"github.com/skvdmt/chrome/internal/devtools/options/network/set_cookie"
	"github.com/skvdmt/chrome/internal/devtools/options/network/set_user_agent_override"
	"github.com/skvdmt/chrome/internal/devtools/types/network"
	"github.com/skvdmt/chrome/internal/model"
)

// Network Сеть.
type Network struct {
	client *model.Client
	debug  *model.Debug
}

// NewNetwork Конструктор.
func NewNetwork(c *model.Client, d *model.Debug) *Network {
	d.Debug("network created")
	return &Network{
		client: c,
		debug:  d,
	}
}

// ClearBrowserCache Очищает кэш браузера.
func (n *Network) ClearBrowserCache() error {
	return n.client.Exec(network.CLEAR_BROUSER_CACHE, nil)
}

// ClearBrowserCookies Очищает файлы cookie браузера.
func (n *Network) ClearBrowserCookies() error {
	return n.client.Exec(network.CLEAR_BROWSER_COOKIES, nil)
}

// DeleteCookies Удаляет файлы cookie браузера с совпадающим
// именем и URL-адресом или парой домен/путь/ключ раздела.
func (n *Network) DeleteCookies(name string, options ...delete_cookies.Option) error {
	c := &delete_cookies.Config{
		Name: name,
	}
	for _, o := range options {
		o(c)
	}
	return n.client.Exec(
		network.DELETE_COOKIES,
		model.ForceJSONMarshal(c),
	)
}

// Disable Отключает отслеживание сети и предотвращает отправку сетевых событий клиенту.
func (n *Network) Disable() error {
	return n.client.Exec(network.DISABLE, nil)
}

// Enable Включает отслеживание сети; теперь сетевые события будут передаваться клиенту.
func (n *Network) Enable(options ...enable.Option) error {
	c := &enable.Config{}
	for _, o := range options {
		o(c)
	}
	return n.client.Exec(
		network.ENEBLE,
		model.ForceJSONMarshal(c),
	)
}

// GetCookies Возвращает все файлы cookie браузера для текущего URL.
// В зависимости от поддержки на стороне бэкенда, в поле `cookies`
// будет возвращена подробная информация о файлах cookie.
func (n *Network) GetCookies(options ...get_cookies.Option) ([]*network.Cookie, error) {
	c := &get_cookies.Config{}
	for _, o := range options {
		o(c)
	}
	r, err := n.client.Query(
		network.GET_COOKIE,
		model.ForceJSONMarshal(c),
	)
	if err != nil {
		return nil, err
	}
	s := struct {
		// Срез объектов cookie.
		Cookies []*network.Cookie
	}{}
	model.ForceJSONUnmarshal(r, s)
	return s.Cookies, nil
}

// GetRequestPostData Возвращает данные POST, переданные в запросе.
// Возвращает ошибку, если данные в запросе не были переданы.
func (n *Network) GetRequestPostData(requestId *network.RequestId) (
	PostData *string,
	Base64Encoded *bool,
	err error,
) {
	r, err := n.client.Query(
		network.GET_REQUEST_POST_DATA,
		model.ForceJSONMarshal(struct {
			// Идентификатор сетевого запроса для получения контента.
			RequestId *network.RequestId `json:"requestId"`
		}{
			RequestId: requestId,
		}),
	)
	if err != nil {
		return nil, nil, err
	}
	p := struct {
		// Строка тела запроса, за исключением файлов из multipart-запросов.
		PostData *string `json:"postData"`
		// Верно, если контент был отправлен в формате base64.
		Base64Encoded *bool `json:"base64Encoded"`
	}{}
	model.ForceJSONUnmarshal(r, p)
	return p.PostData, p.Base64Encoded, nil
}

// GetResponseBody Возвращает контент, выданный в ответ на данный запрос.
func (n *Network) GetResponseBody(requestId *network.RequestId) (
	body *string,
	base64Encoded *bool,
	err error,
) {
	r, err := n.client.Query(
		network.GET_RESPONSE_BODY,
		model.ForceJSONMarshal(struct {
			// Идентификатор сетевого запроса для получения контента.
			RequestId *network.RequestId `json:"requestId"`
		}{
			RequestId: requestId,
		}),
	)
	if err != nil {
		return nil, nil, err
	}
	b := struct {
		// Тело ответа.
		Body *string `json:"body"`
		// True, если контент был отправлен в формате base64.
		Base64Encoded *bool `json:"base64Encoded"`
	}{}
	model.ForceJSONUnmarshal(r, b)
	return b.Body, b.Base64Encoded, nil
}

// SetBypassServiceWorker Переключает игнорирование сервис-воркера для каждого запроса.
func (n *Network) SetBypassServiceWorker(bypass bool) error {
	return n.client.Exec(
		network.SET_BYPASS_SERVICE_WORKER,
		model.ForceJSONMarshal(struct {
			// Обойти service worker и загрузить из сети.
			Bypass bool `json:"bypass"`
		}{
			Bypass: bypass,
		}),
	)
}

// SetCacheDisabled Переключает режим игнорирования кэша для каждого запроса.
// Если установлено значение true, кэш использоваться не будет.
func (n *Network) SetCacheDisabled(cacheDisabled bool) error {
	return n.client.Exec(
		network.SET_CACHE_DISABLED,
		model.ForceJSONMarshal(struct {
			// Состояние отключенного кэша.
			CacheDisabled bool `json:"cacheDisabled"`
		}{
			CacheDisabled: cacheDisabled,
		}),
	)
}

// SetCookie Устанавливает cookie с заданными данными; может
// перезаписать аналогичные cookie, если они уже существуют.
func (n *Network) SetCookie(name, value string, options ...set_cookie.Option) error {
	c := &set_cookie.Config{
		Name:  name,
		Value: value,
	}
	for _, o := range options {
		o(c)
	}
	return n.client.Exec(
		network.SET_COOKIE,
		model.ForceJSONMarshal(c),
	)
}

// SetCookies Устанавливает указанные файлы cookie.
func (n *Network) SetCookies(cookies []*network.CookieParam) error {
	return n.client.Exec(
		network.SET_COOKIES,
		model.ForceJSONMarshal(struct {
			// Файлы cookie, которые будут установлены.
			Cookies []*network.CookieParam `json:"cookies"`
		}{
			Cookies: cookies,
		}),
	)
}

// SetExtraHTTPHeaders Определяет, следует ли всегда отправлять дополнительные
// HTTP-заголовки вместе с запросами с этой страницы.
func (n *Network) SetExtraHTTPHeaders(headers *network.Headers) error {
	return n.client.Exec(
		network.SET_EXTRA_HTTP_HEADERS,
		model.ForceJSONMarshal(struct {
			// Карта с дополнительными HTTP-заголовками.
			Headers *network.Headers `json:"headers"`
		}{
			Headers: headers,
		}),
	)
}

// SetUserAgentOverride Позволяет переопределить User-Agent с помощью указанной строки.
func (n *Network) SetUserAgentOverride(userAgent string, options ...set_user_agent_override.Option) error {
	c := &set_user_agent_override.Config{
		UserAgent: userAgent,
	}
	for _, o := range options {
		o(c)
	}
	return n.client.Exec(
		network.SET_USER_AGENT_OVERRIDE,
		model.ForceJSONMarshal(c),
	)
}
