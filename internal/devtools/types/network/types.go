package network

import (
	"encoding/json"

	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
	"github.com/skvdmt/chrome/internal/devtools/types/security"
)

// BlockedReason Причина блокировки запроса.
// Допустимые значения: other, csp, mixed-content, origin, inspector, integrity,
// subresource-filter, content-type, coep-frame-resource-needs-coep-header,
// coop-sandboxed-iframe-cannot-navigate-to-coop-page, corp-not-same-origin,
// corp-not-same-origin-after-defaulted-to-same-origin-by-coep,
// corp-not-same-origin-after-defaulted-to-same-origin-by-dip,
// corp-not-same-origin-after-defaulted-to-same-origin-by-coep-and-dip,
// corp-not-same-site, sri-message-signature-mismatch
type BlockedReason string

// CachedResource Информация о кэшированном ресурсе.
type CachedResource struct {
	// URL ресурса. Это URL исходного сетевого запроса.
	Url string `json:"url"`
	// Тип данного ресурса.
	Type *ResourceType `json:"type"`
	// Кэшированные данные ответа.
	Response *Response `json:"response"`
	// Размер кэшированного тела ответа.
	BodySize int `json:"bodySize"`
}

// CertificateTransparencyCompliance Соответствует ли запрос
// политике прозрачности сертификатов?
// Допустимые значения: unknown, not-compliant, compliant
type CertificateTransparencyCompliance string

// ConnectionType Технология подключения, которую предположительно использует браузер.
// Допустимые значения: none, cellular2g, cellular3g, cellular4g,
// bluetooth, ethernet, wifi, wimax, other
type ConnectionType string

// Cookie Объект Cookie.
type Cookie struct {
	// Название файла cookie.
	Name string `json:"name"`
	// Значение файла cookie.
	Value string `json:"value"`
	// Домен файлов cookie.
	Domain string `json:"domain"`
	// Путь к файлу cookie.
	Path string `json:"path"`
	// Срок действия cookie-файла указывается в секундах с начала эпохи UNIX.
	// Если срок действия не указан, значение устанавливается равным -1.
	// Для значений, которые не могут быть представлены в формате JSON (±Inf),
	// значение может быть равно null.
	Expires int `json:"expires"`
	// Размер печенья.
	Size int `json:"size"`
	// Возвращает true, если cookie-файл предназначен только для HTTP-запросов.
	HttpOnly bool `json:"httpOnly"`
	// Возвращает true, если cookie-файл защищен.
	Secure bool `json:"secure"`
	// Возвращает true в случае использования сессионного cookie.
	Session bool `json:"session"`
	// Тип cookie SameSite.
	SameSite *CookieSameSite `json:"sameSite,omitempty"`
	// Приоритет файлов cookie.
	Priority *CookiePriority `json:"priority"`
	// Тип схемы источника файлов cookie.
	SourceScheme *CookieSourceScheme `json:"sourceScheme"`
	// Исходный порт cookie. Допустимые значения: {-1, [1, 65535]}, -1 указывает
	// на неуказанный порт. Неуказанное значение порта позволяет клиентам протокола
	// эмулировать область действия устаревших cookie для этого порта.
	// Это временная возможность, которая будет удалена в будущем.
	SourcePort int `json:"sourcePort"`
	// Ключ раздела cookie.
	PartitionKey CookiePartitionKey `json:"partitionKey,omitempty"`
	// Возвращает true, если ключ раздела cookie непрозрачен.
	PartitionKeyOpaque bool `json:"partitionKeyOpaque,omitempty"`
}

// CookieParam Объект параметра cookie.
type CookieParam struct {
	// Название cookie.
	name string
	// Значение cookie.
	value string
	// URI запроса, связываемый с настройками cookie.
	// Это значение может влиять на значения домена, пути, порта
	// источника и схемы источника по умолчанию для создаваемого cookie.
	url string
	// Домен cookie.
	domain string
	// Путь cookie.
	path string
	// Возвращает true, если cookie защищен.
	secure bool
	// Возвращает true, если cookie предназначен только для HTTP-запросов.
	httpOnly bool
	// Тип cookie SameSite.
	sameSite *CookieSameSite
	// Срок действия cookie-файла, сессионный cookie-файл, если он не установлен.
	expires *TimeSinceEpoch
	// Приоритет cookie.
	priority *CookiePriority
	// Тип схемы источника cookie.
	sourceScheme *CookieSourceScheme
	// Исходный порт cookie. Допустимые значения: {-1, [1, 65535]}, -1 указывает
	// на неуказанный порт. Неуказанное значение порта позволяет клиентам протокола
	// эмулировать область действия устаревших cookie для этого порта.
	// Это временная возможность, которая будет удалена в будущем.
	sourcePort int
	// Ключ разделения cookie-файла. Если не задан, cookie
	// будет помечен как не разделенный.
	partitionKey *CookiePartitionKey
}

// CookieSameSite Отражает статус cookie-файла «SameSite»:
// https://tools.ietf.org/html/draft-west-first-party-cookies
// Допустимые значения: Strict, Lax, None
type CookieSameSite string

// CorsError Причина блокировки запроса.
// Допустимые значения: DisallowedByMode, InvalidResponse, WildcardOriginNotAllowed,
// MissingAllowOriginHeader, MultipleAllowOriginValues, InvalidAllowOriginValue,
// AllowOriginMismatch, InvalidAllowCredentials, CorsDisabledScheme,
// PreflightInvalidStatus, PreflightDisallowedRedirect,
// PreflightWildcardOriginNotAllowed, PreflightMissingAllowOriginHeader,
// PreflightMultipleAllowOriginValues, PreflightInvalidAllowOriginValue,
// PreflightAllowOriginMismatch, PreflightInvalidAllowCredentials,
// PreflightMissingAllowExternal, PreflightInvalidAllowExternal,
// InvalidAllowMethodsPreflightResponse, InvalidAllowHeadersPreflightResponse,
// MethodDisallowedByPreflightResponse, HeaderDisallowedByPreflightResponse,
// RedirectContainsCredentials, InsecureLocalNetwork, InvalidLocalNetworkAccess,
// NoCorsRedirectModeNotFollow, LocalNetworkAccessPermissionDenied
type CorsError string

// CorsErrorStatus
type CorsErrorStatus struct {
	CorsError       CorsError `json:"corsError"`
	FailedParameter string    `json:"failedParameter"`
}

// ErrorReason Причина ошибки при получении данных на сетевом уровне.
// Допустимые значения: Failed, Aborted, TimedOut, AccessDenied, ConnectionClosed,
// ConnectionReset, ConnectionRefused, ConnectionAborted, ConnectionFailed,
// NameNotResolved, InternetDisconnected, AddressUnreachable, BlockedByClient,
// BlockedByResponse
type ErrorReason string

// Headers Заголовки запроса/ответа представлены в виде ключей/значений JSON-объекта.
type Headers json.RawMessage

// Initiator Информация об инициаторе запроса.
type Initiator struct {
	// Тип инициатора.
	// Допустимые значения: parser, script, preload, SignedExchange, preflight, FedCM, other
	Type string `json:"type"`
	// Трассировка стека JavaScript инициатора, настроена только для скрипта.
	// Требуется включение домена отладчика.
	Stack *rnt.StackTrace `json:"stack,omitempty"`
	// URL-адрес инициатора, задается для типа парсера, для типа
	// скрипта (если скрипт импортирует модуль) или для типа SignedExchange.
	Url string `json:"url,omitempty"`
	// Номер строки инициатора, задаваемый для типа парсера
	// или для типа скрипта (когда скрипт импортирует модуль) (начиная с 0).
	LineNumber int `json:"lineNumber,omitempty"`
	// Номер столбца "Инициатор", задается для типа парсера или
	// для типа скрипта (если скрипт импортирует модуль) (начиная с 0).
	ColumnNumber int `json:"columnNumber,omitempty"`
	// Укажите, если данный запрос был инициирован
	// другим запросом (например, предварительным запросом).
	RequestId *RequestId `json:"requestId,omitempty"`
}

// LoaderId Уникальный идентификатор погрузчика.
type LoaderId string

// MonotonicTime Монотонно возрастающее время в секундах,
// прошедшее с произвольной точки в прошлом.
type MonotonicTime int

// PostDataEntry Отправьте данные для HTTP-запроса.
type PostDataEntry struct {
	Bytes string `json:"bytesm,omitempty"`
}

// Request Данные HTTP-запроса.
type Request struct {
	// URL запроса (без фрагмента).
	Url string `json:"url"`
	// Фрагмент запрошенного URL-адреса, начинающийся с хеша, если он присутствует.
	UrlFragment string `json:"urlFragment"`
	// Метод HTTP-запроса.
	Method string `json:"method"`
	// Заголовки HTTP-запроса.
	Headers Headers `json:"headers"`
	// Значение true, если запрос содержит POST-данные. Обратите внимание, что
	// параметр postData может быть пропущен, даже если данные слишком длинные,
	// если этот флаг имеет значение true.
	HasPostData bool `json:"hasPostData"`
	// Элементы тела запроса (данные запроса разбиты на отдельные записи).
	PostDataEntries []*PostDataEntry `json:"postDataEntries"`
	// Тип смешанного содержимого запроса.
	MixedContentType *security.MixedContentType `json:"mixedContentType"`
	// Приоритет запроса ресурсов в момент его отправки.
	InitialPriority *ResourcePriority `json:"initialPriority"`
	// Политика рефереров запроса, определенная в https://www.w3.org/TR/referrer-policy/
	// Допустимые значения: unsafe-url, no-referrer-when-downgrade, no-referrer,
	// origin, origin-when-cross-origin, same-origin, strict-origin,
	// strict-origin-when-cross-origin
	ReferrerPolicy string `json:"referrerPolicy"`
	// Загружается ли файл посредством предварительной загрузки ссылки.
	IsLinkPreload bool `json:"isLinkPreload"`
	// Устанавливается для запросов при использовании API TrustToken.
	// Содержит параметры, передаваемые разработчиком (например, через "fetch"),
	// в том виде, в котором они понимаются бэкэндом.
	TrustTokenParams *TrustTokenParams `json:"trustTokenParams"`
	// Значение true, если этот запрос на ресурс считается относящимся
	// к «тому же сайту», что и запрос, соответствующий мэйнфрейму.
	IsSameSite bool `json:"isSameSite"`
	// Верно, если запрос ресурса связан с рекламой.
	IsAdRelated bool `json:"isAdRelated"`
}

// RequestId Уникальный сетевой идентификатор запроса. Обратите внимание, что
// он не идентифицирует отдельные HTTP-запросы, являющиеся частью сетевого запроса.
type RequestId string

// ResourcePriority Приоритет загрузки запрашиваемого ресурса.
// Допустимые значения: VeryLow, Low, Medium, High, VeryHigh
type ResourcePriority string

// ResourceTiming Информация о сроках выполнения запроса.
type ResourceTiming struct {
	// Значение requestTime в параметре Timing — это базовое значение в секундах,
	// а остальные числа — это миллисекунды, отстоящие от этого значения requestTime.
	RequestTime int `json:"requestTime"`
	// Начато разрешение прокси-сервера.
	ProxyStart int `json:"proxyStart"`
	// Завершено разрешение прокси-сервера.
	ProxyEnd int `json:"proxyEnd"`
	// Начато разрешение DNS-адреса.
	DnsStart int `json:"dnsStart"`
	// Завершено разрешение DNS-адреса.
	DnsEnd int `json:"dnsEnd"`
	// Началось подключение к удалённому хосту.
	ConnectStart int `json:"connectStart"`
	// Подключено к удаленному хосту.
	ConnectEnd int `json:"connectEnd"`
	// Начато SSL-рукопожатие.
	SslStart int `json:"sslStart"`
	// Завершено SSL-рукопожатие.
	SslEnd int `json:"sslEnd"`
	// Запустил ServiceWorker.
	WorkerStart int `json:"workerStart"`
	// Завершен запуск ServiceWorker.
	WorkerReady int `json:"workerReady"`
	// Началось получение события.
	WorkerFetchStart int `json:"workerFetchStart"`
	// Завершено получение события respondWith promise.
	WorkerRespondWithSettled int `json:"workerRespondWithSettled"`
	// Начата оценка источника статической маршрутизации ServiceWorker.
	WorkerRouterEvaluationStart int `json:"workerRouterEvaluationStart,omitempty"`
	// Поиск в кэше начался после того, как источник был оценен и найден в кэше.
	WorkerCacheLookupStart int `json:"workerCacheLookupStart"`
	// Начал отправлять запрос.
	SendStart int `json:"sendStart"`
	// Отправка запроса завершена.
	SendEnd int `json:"sendEnd"`
	// Время, когда сервер начал отправлять запрос.
	PushStart int `json:"pushStart"`
	// Засеките время, за которое сервер завершит отправку запроса.
	PushEnd int `json:"pushEnd"`
	// Начали получать заголовки ответа.
	ReceiveHeadersStart int `json:"receiveHeadersStart"`
	// Получение заголовков ответа завершено.
	ReceiveHeadersEnd int `json:"receiveHeadersEnd"`
}

// ResourceType Тип ресурса, как он был воспринят механизмом рендеринга.
// Допустимые значения: Document, Stylesheet, Image, Media, Font, Script,
// TextTrack, XHR, Fetch, Prefetch, EventSource, WebSocket, Manifest,
// SignedExchange, Ping, CSPViolationReport, Preflight, FedCM, Other
type ResourceType string

// Response Данные HTTP-ответа.
type Response struct {
	// URL ответа. Этот URL может отличаться от CachedResource.url
	// в случае перенаправления.
	Url string `json:"url"`
	// Код состояния HTTP-ответа.
	Status int `json:"status"`
	// Текст статуса HTTP-ответа.
	StatusText string `json:"statusText"`
	// Заголовки HTTP-ответа.
	Headers Headers `json:"headers"`
	// MIME-тип ресурса определяется браузером.
	MimeType string `json:"mimeType"`
	// Кодировка ресурса определяется браузером (если применимо).
	Charset string `json:"charset"`
	// Уточнены заголовки HTTP-запросов, фактически переданные по сети.
	RequestHeaders *Headers `json:"requestHeaders"`
	// Указывает, было ли физическое соединение фактически
	// повторно использовано для этого запроса.
	ConnectionReused bool `json:"connectionReused"`
	// Идентификатор физического соединения, фактически использованный для этого запроса.
	ConnectionId int `json:"connectionId"`
	// Удаленный IP-адрес.
	RemoteIPAddress string `json:"remoteIPAddress"`
	// Удаленный порт.
	RemotePort int `json:"remotePort"`
	// Указывает, что запрос был обработан из дискового кэша.
	FromDiskCache bool `json:"fromDiskCache"`
	// Указывает, что запрос был обработан объектом ServiceWorker.
	FromServiceWorker bool `json:"fromServiceWorker"`
	// Указывает, что запрос был обработан из кэша предварительной выборки.
	FromPrefetchCache bool `json:"fromPrefetchCache"`
	// Указывает, что запрос был обработан из кэша предварительной выборки.
	FromEarlyHints bool `json:"fromEarlyHints"`
	// Информация об использовании API статического маршрутизатора ServiceWorker.
	// Если это поле содержит значение поля matchedSourceType, найдено соответствующее
	// правило. Если это поле не содержит значения matchedSource, соответствующее
	// правило не найдено. В противном случае API не используется.
	ServiceWorkerRouterInfo *ServiceWorkerRouterInfo `json:"serviceWorkerRouterInfo"`
	// Общее количество байтов, полученных по этому запросу на данный момент.
	EncodedDataLength int `json:"encodedDataLength"`
	// Информация о времени выполнения данного запроса.
	Timing *ResourceTiming `json:"timing"`
	// Источник ответа от ServiceWorker.
	ServiceWorkerResponseSource *ServiceWorkerResponseSource `json:"serviceWorkerResponseSource"`
	// Время, когда был сгенерирован полученный ответ.
	ResponseTime *TimeSinceEpoch `json:"responseTime"`
	// Хранилище кэша. Имя кэша.
	CacheStorageCacheName string `json:"cacheStorageCacheName"`
	// Протокол, использованный для получения этого запроса.
	Protocol string `json:"protocol"`
	// Причина, по которой Chrome использует определенный транспортный протокол для семантики HTTP.
	AlternateProtocolUsage *AlternateProtocolUsage `json:"alternateProtocolUsage"`
	// Состояние безопасности запрашиваемого ресурса.
	SecurityState *security.SecurityState `json:"securityState"`
	// Информация о безопасности запроса.
	SecurityDetails *SecurityDetails `json:"securityDetails"`
}

// SecurityDetails Информация о безопасности запроса.
type SecurityDetails struct {
	// Название протокола (например, "TLS 1.2" или "QUIC").
	Protocol string `json:"protocol"`
	// Ключ обмена, используемый в соединении, или пустая строка, если он неприменим.
	KeyExchange string `json:"keyExchange"`
	// (EC)DH группа, используемая соединением, если применимо.
	KeyExchangeGroup string `json:"keyExchangeGroup,omitempty"`
	// Название шифра.
	Cipher string `json:"cipher"`
	// MAC-адреса TLS. Обратите внимание, что шифры AEAD не имеют отдельных MAC-адресов.
	Mac string `json:"mac,omitempty"`
	// Значение идентификатора сертификата.
	CertificateId *security.CertificateId `json:"certificateId"`
	// Название предмета сертификата.
	SubjectName string `json:"subjectName"`
	// Альтернативное имя субъекта (SAN) — DNS-имена и IP-адреса.
	SanList []string `json:"sanList"`
	// Название центра сертификации, выпустившего сертификат.
	Issuer string `json:"issuer"`
	// Сертификат действителен с указанной даты.
	ValidFrom *TimeSinceEpoch `json:"validFrom"`
	// Сертификат действителен до даты истечения срока действия.
	ValidTo *TimeSinceEpoch `json:"validTo"`
	// Список временных меток подписанных сертификатов (SCT).
	SignedCertificateTimestampList []*SignedCertificateTimestamp `json:"signedCertificateTimestampList"`
	// Соответствовал ли запрос политике прозрачности сертификатов.
	CertificateTransparencyCompliance *CertificateTransparencyCompliance `json:"certificateTransparencyCompliance"`
	// Алгоритм подписи, используемый сервером в TLS-подписи,
	// представленный в виде кодовой точки TLS SignatureScheme.
	// Опускается, если неприменим или неизвестен.
	ServerSignatureAlgorithm int `json:"serverSignatureAlgorithm,omitempty"`
	// Использовалось ли в соединении зашифрованное приглашение ClientHello.
	EncryptedClientHello bool `json:"encryptedClientHello"`
}

// ServiceWorkerResponseSource Источник ответа сервис-воркера.
// Допустимые значения: cache-storage, http-cache, fallback-code, network
type ServiceWorkerResponseSource string

// ServiceWorkerRouterSource Источник маршрутизатора сервис-воркера.
// Допустимые значения: network, cache, fetch-event,
// race-network-and-fetch-handler, race-network-and-cache
type ServiceWorkerRouterSource string

// SignedCertificateTimestamp Подробная информация о временной метке
// подписанного сертификата (SCT).
type SignedCertificateTimestamp struct {
	// Статус проверки.
	Status string `json:"status"`
	// Источник.
	Origin string `json:"origin"`
	// Название/описание журнала.
	LogDescription string `json:"logDescription"`
	// Идентификатор журнала.
	LogId string `json:"logId"`
	// Дата выпуска. В отличие от TimeSinceEpoch, здесь указывается количество
	// миллисекунд с 1 января 1970 года (UTC), а не количество секунд.
	Timestamp int `json:"timestamp"`
	// Хэш-алгоритм.
	HashAlgorithm string `json:"hashAlgorithm"`
	// Алгоритм подписи.
	SignatureAlgorithm string `json:"signatureAlgorithm"`
	// Данные подписи.
	SignatureData string `json:"signatureData"`
}

// TimeSinceEpoch Время UTC в секундах, отсчитываемое с 1 января 1970 года.
type TimeSinceEpoch int

// WebSocketFrame Данные сообщения WebSocket. Это представляет собой целое
// сообщение WebSocket, а не просто фрагментированный кадр, как следует из названия.
type WebSocketFrame struct {
	// Код операции сообщения WebSocket.
	Opcode int `json:"opcode"`
	// Маска сообщения WebSocket.
	Mask bool `json:"mask"`
	// Данные полезной нагрузки сообщения WebSocket. Если код операции равен 1,
	// это текстовое сообщение, а payloadData — строка в кодировке UTF-8. Если
	// код операции не равен 1, то payloadData — это строка в кодировке Base64,
	// представляющая двоичные данные.
	PayloadData string `json:"payloadData"`
}

// WebSocketRequest Данные запроса WebSocket.
type WebSocketRequest struct {
	// Заголовки HTTP-запроса.
	Headers *Headers `json:"headers"`
}

// WebSocketResponse Данные ответа WebSocket.
type WebSocketResponse struct {
	// Код состояния HTTP-ответа.
	Status int `json:"status"`
	// Текст статуса HTTP-ответа.
	StatusText string `json:"statusText"`
	// Заголовки HTTP-ответа.
	Headers *Headers `json:"headers"`
	// Текст заголовков HTTP-ответа.
	HeadersText string `json:"headersText,omitempty"`
	// Заголовки HTTP-запроса.
	RequestHeaders *Headers `json:"requestHeaders,omitempty"`
	// Текст заголовков HTTP-запроса.
	RequestHeadersText string `json:"requestHeadersText,omitempty"`
}

// CookiePriority Отображает статус «Приоритет» файла cookie:
// https://tools.ietf.org/html/draft-west-cookie-priority-00
// Допустимые значения: Low, Medium, High
type CookiePriority string

// CookieSourceScheme Представляет собой исходную схему источника, который
// первоначально установил cookie. Значение "Unset" позволяет клиентам
// протокола эмулировать область действия устаревших cookie для данной схемы.
// Это временная возможность, которая будет удалена в будущем.
// Допустимые значения: Unset, NonSecure, Secure
type CookieSourceScheme string

// CookiePartitionKey Объект cookiePartitionKey. Представление компонентов ключа,
// создаваемых классом cookiePartitionKey, содержащимся в файле
// net/cookies/cookie_partition_key.h.
type CookiePartitionKey struct {
	// Адрес верхнего уровня URL-адреса, который браузер посещал
	// в начале запроса к конечной точке, установившей cookie-файл.
	TopLevelSite string `json:"topLevelSite"`
	// Указывает, есть ли у файла cookie какие-либо предки, находящиеся
	// на других сайтах, помимо сайта верхнего уровня.
	HasCrossSiteAncestor bool `json:"hasCrossSiteAncestor"`
}

// TrustTokenOperationType
// Допустимые значения: Issuance, Redemption, Signing
type TrustTokenOperationType string

// TrustTokenParams Определяет тип выполняемой операции с токеном доверия и,
// в зависимости от типа, некоторые дополнительные параметры. Значения указываются
// в файле third_party/blink/renderer/core/fetch/trust_token.idl.
type TrustTokenParams struct {
	Operation *TrustTokenOperationType `json:"operation"`
	// Настраивается только для операции «погашение токенов» и определяет, следует ли
	// запрашивать новый SRR или использовать все еще действительный кэшированный SRR.
	// Допустимые значения: UseCached, Refresh
	RefreshPolicy string `json:"refreshPolicy"`
	// Источники эмитентов, у которых следует запрашивать
	// токены или записи об их погашении.
	Issuers []string `json:"issuers,omitempty"`
}

// ServiceWorkerRouterInfo
type ServiceWorkerRouterInfo struct {
	// Идентификатор совпавшего правила. Если найдено совпавшее правило,
	// это поле будет заполнено, в противном случае значение не будет установлено.
	RuleIdMatched int `json:"ruleIdMatched,omitempty"`
	// Источник маршрутизатора для соответствующего правила.
	// Если соответствующее правило существует, это поле будет заполнено,
	// в противном случае значение не будет установлено.
	MatchedSourceType *ServiceWorkerRouterSource `json:"matchedSourceType,omitempty"`
	// Фактически использованный источник маршрутизатора.
	ActualSourceType *ServiceWorkerRouterSource `json:"actualSourceType,omitempty"`
}

// AlternateProtocolUsage Причина, по которой Chrome использует определенный
// транспортный протокол для семантики HTTP.
// Допустимые значения: alternativeJobWonWithoutRace, alternativeJobWonRace,
// mainJobWonRace, mappingMissing, broken, dnsAlpnH3JobWonWithoutRace,
// dnsAlpnH3JobWonRace, unspecifiedReason
type AlternateProtocolUsage string
