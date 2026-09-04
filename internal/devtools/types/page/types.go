package page

import "github.com/skvdmt/chrome/internal/devtools/types/network"

// AppManifestError Ошибка при сопряжении манифеста приложения.
type AppManifestError struct {
	// Сообщение об ошибке.
	Message string `json:"message"`
	// В критическом случае это неисправимая ошибка синтаксического анализа.
	Critical int `json:"critical"`
	// Строка ошибки.
	Line int `json:"line"`
	// Столбец ошибок.
	Column int `json:"column"`
}

// DialogType Диалоговое окно в формате JavaScript.
// Допустимые значения: alert, confirm, prompt, beforeunload
type DialogType string

// Frame Информация о рамке представлена ​​на странице.
type Frame struct {
	// Уникальный идентификатор кадра.
	Id FrameId `json:"id"`
	// Идентификатор родительского фрейма.
	ParentId FrameId `json:"parentId,omitempty"`
	// Идентификатор погрузчика, связанного с данной рамой.
	LoaderId *network.LoaderId `json:"loaderId"`
	// Название кадра, указанное в теге.
	Name string `json:"name,omitempty"`
	// URL документа без фрагмента.
	Url string `json:"url"`
	// Фрагмент URL-адреса документа-фрейма, включающий символ '#'.
	UrlFragment string `json:"urlFragment,omitempty"`
	// Зарегистрированный домен документа-фрейма с учетом списка
	// общедоступных суффиксов. Извлекается из URL-адреса фрейма.
	// Примеры URL: http://www.google.com/file.html -> "google.com"
	// http://a.b.co.uk/file.html -> "b.co.uk"
	DomainAndRegistry string `json:"domainAndRegistry"`
	// Источник безопасности документа-фрейма.
	SecurityOrigin string `json:"securityOrigin"`
	// Дополнительные сведения об источнике безопасности документа-фрейма.
	SecurityOriginDetails *SecurityOriginDetails `json:"securityOriginDetails,omitempty"`
	// MIME-тип документа фрейма определяется браузером.
	MimeType string `json:"mimeType"`
	// Если фрейм не загрузился, здесь содержится URL-адрес, который не
	// удалось загрузить. Обратите внимание, что в отличие от URL-адреса
	// выше, этот URL-адрес может содержать фрагмент.
	UnreachableUrl string `json:"unreachableUrl,omitempty"`
	// Указывает, был ли данный кадр помечен как реклама и почему.
	AdFrameStatus *AdFrameStatus `json:"adFrameStatus,omitempty"`
	// Указывает, является ли основной документ защищенным
	// контекстом, и объясняет, почему это так.
	SecureContextType *SecureContextType `json:"secureContextType"`
	// Указывает, является ли это контекстом, в котором происходит
	// взаимодействие между различными источниками.
	CrossOriginIsolatedContextType *CrossOriginIsolatedContextType `json:"crossOriginIsolatedContextType"`
	// Указано, какие доступные API/функции с ограниченным доступом доступны.
	GatedAPIFeatures []*GatedAPIFeatures `json:"gatedAPIFeatures"`
}

// FrameId Уникальный идентификатор фрейма.
type FrameId string

// FrameTree Информация об иерархии фреймов.
type FrameTree struct {
	// Информация о структуре этого элемента дерева.
	Frame Frame `json:"frame"`
	// Детские рамки.
	ChildFrames []*FrameTree `json:"childFrames,omitempty"`
}

// LayoutViewport Положение и размеры области просмотра компоновки.
type LayoutViewport struct {
	// Горизонтальное смещение относительно документа (в пикселях CSS).
	PageX int `json:"pageX"`
	// Вертикальное смещение относительно документа (в пикселях CSS).
	PageY int `json:"pageY"`
	// Ширина (в пикселях CSS), без учета полосы прокрутки, если она присутствует.
	ClientWidth int `json:"clientWidth"`
	// Высота (в пикселях CSS), без учета полосы прокрутки, если она присутствует.
	ClientHeight int `json:"clientHeight"`
}

// NavigationEntry Запись в историю навигации.
type NavigationEntry struct {
	// Уникальный идентификатор записи в истории навигации.
	Id int `json:"id"`
	// URL записи истории навигации.
	Url string `json:"url"`
	// URL-адрес, который пользователь ввел в адресную строку браузера.
	UserTypedURL string `json:"userTypedURL"`
	// Заголовок записи истории навигации.
	Title string `json:"title"`
	// Тип перехода.
	TransitionType *TransitionType `json:"transitionType"`
}

// ScriptIdentifier Уникальный идентификатор скрипта.
type ScriptIdentifier string

// TransitionType Тип перехода.
// Допустимые значения: link, typed, address_bar, auto_bookmark,
// auto_subframe, manual_subframe, generated, auto_toplevel,
// form_submit, reload, keyword, keyword_generated, other
type TransitionType string

// Viewport Область просмотра для создания скриншота.
type Viewport struct {
	// Смещение по оси X в пикселях, не зависящих от устройства (провал).
	X int `json:"x"`
	// Смещение по оси Y в пикселях, не зависящих от устройства (провал).
	Y int `json:"y"`
	// Ширина прямоугольника в пикселях, не зависящих от устройства (dip).
	Width int `json:"width"`
	// Высота прямоугольника в пикселях, не зависящих от устройства (dip).
	Height int `json:"height"`
	// Масштабный коэффициент страницы.
	Scale int `json:"scale"`
}

// VisualViewport Визуальное положение, размеры и масштаб области просмотра.
type VisualViewport struct {
	// Горизонтальное смещение относительно области просмотра макета (в пикселях CSS).
	OffsetX int `json:"offsetX"`
	// Вертикальное смещение относительно области просмотра макета (в пикселях CSS).
	OffsetY int `json:"offsetY"`
	// Горизонтальное смещение относительно документа (в пикселях CSS).
	PageX int `json:"pageX"`
	// Вертикальное смещение относительно документа (в пикселях CSS).
	PageY int `json:"pageY"`
	// Ширина (в пикселях CSS), без учета полосы прокрутки, если она присутствует.
	ClientWidth int `json:"clientWidth"`
	// Высота (в пикселях CSS), без учета полосы прокрутки, если она присутствует.
	ClientHeight int `json:"clientHeight"`
	// Масштабирование относительно идеального размера области просмотра
	// (размер при ширине равен ширине устройства).
	Scale int `json:"scale"`
	// Коэффициент масштабирования страницы (соотношение пикселей,
	// не зависящее от CSS и устройства).
	Zoom int `json:"zoom,omitempty"`
}

// AdFrameExplanation Допустимые значения: ParentIsAd,
// CreatedByAdScript, MatchedBlockingRule
type AdFrameExplanation string

// AdFrameStatus Указывает, был ли кадр идентифицирован как реклама и почему.
type AdFrameStatus struct {
	AdFrameType  *AdFrameType          `json:"adFrameType"`
	Explanations []*AdFrameExplanation `json:"explanations,omitempty"`
}

// AdFrameType Указывает, был ли фрейм идентифицирован как реклама.
// Допустимые значения: none, child, root
type AdFrameType string

// AppManifestParsedProperties Проанализированы свойства манифеста приложения.
type AppManifestParsedProperties struct {
	// Вычисленное значение области действия.
	Scope string `json:"scope"`
}

// BackForwardCacheBlockingDetails
type BackForwardCacheBlockingDetails struct {
	// URL файла, где произошла блокировка. Необязательно,
	// так как используется для тестирования.
	Url string `json:"url,omitempty"`
	// Название функции, в которой произошла блокировка.
	// Необязательно, поскольку используются анонимные функции и тесты.
	Function string `json:"function,omitempty"`
	// Номер строки в скрипте (начиная с 0).
	LineNumber int `json:"lineNumber"`
	// Номер столбца в скрипте (начиная с 0).
	ColumnNumber int `json:"columnNumber"`
}

// BackForwardCacheNotRestoredExplanation
type BackForwardCacheNotRestoredExplanation struct {
	// Тип причины.
	Type *BackForwardCacheNotRestoredReasonType `json:"type"`
	// Причина не восстановлена.
	Reason *BackForwardCacheNotRestoredReason `json:"reason"`
	// Контекст, связанный с причиной. Значение этого контекста зависит от причины:
	// EmbedderExtensionSentMessageToCachedFrame: идентификатор расширения.
	Context string                             `json:"context,omitempty"`
	Details []*BackForwardCacheBlockingDetails `json:"details,omitempty"`
}

// BackForwardCacheNotRestoredExplanationTree
type BackForwardCacheNotRestoredExplanationTree struct {
	// URL каждого фрейма.
	Url string `json:"url"`
	// Причины, по которым не были восстановлены для каждого кадра.
	Explanations []*BackForwardCacheNotRestoredExplanation `json:"explanations"`
	// Массив дочерних кадров.
	Children []*BackForwardCacheNotRestoredExplanationTree `json:"children"`
}

// BackForwardCacheNotRestoredReason Список причин, по которым
// кэш не был восстановлен (перемотка вперед/назад).
// Допустимые значения: NotPrimaryMainFrame, BackForwardCacheDisabled,
// RelatedActiveContentsExist, HTTPStatusNotOK, SchemeNotHTTPOrHTTPS, Loading,
// WasGrantedMediaAccess, DisableForRenderFrameHostCalled, DomainNotAllowed,
// HTTPMethodNotGET, SubframeIsNavigating, Timeout, CacheLimit,
// JavaScriptExecution, RendererProcessKilled, RendererProcessCrashed,
// SchedulerTrackedFeatureUsed, ConflictingBrowsingInstance, CacheFlushed,
// ServiceWorkerVersionActivation, SessionRestored, ServiceWorkerPostMessage,
// EnteredBackForwardCacheBeforeServiceWorkerHostAdded,
// RenderFrameHostReused_SameSite, RenderFrameHostReused_CrossSite,
// ServiceWorkerClaim, IgnoreEventAndEvict, HaveInnerContents, TimeoutPuttingInCache,
// BackForwardCacheDisabledByLowMemory, BackForwardCacheDisabledByCommandLine,
// NetworkRequestDatapipeDrainedAsBytesConsumer, NetworkRequestRedirected,
// NetworkRequestTimeout, NetworkExceedsBufferLimit, NavigationCancelledWhileRestoring,
// NotMostRecentNavigationEntry, BackForwardCacheDisabledForPrerender,
// UserAgentOverrideDiffers, ForegroundCacheLimit, ForwardCacheDisabled,
// BrowsingInstanceNotSwapped, BackForwardCacheDisabledForDelegate,
// UnloadHandlerExistsInMainFrame, UnloadHandlerExistsInSubFrame,
// ServiceWorkerUnregistration, CacheControlNoStore,
// CacheControlNoStoreCookieModified, CacheControlNoStoreHTTPOnlyCookieModified,
// NoResponseHead, Unknown, ActivationNavigationsDisallowedForBug1234857,
// ErrorDocument, FencedFramesEmbedder, CookieDisabled, HTTPAuthRequired,
// CookieFlushed, BroadcastChannelOnMessage, WebViewSettingsChanged,
// WebViewJavaScriptObjectChanged, WebViewMessageListenerInjected,
// WebViewSafeBrowsingAllowlistChanged, WebViewDocumentStartJavascriptChanged,
// WebSocket, WebTransport, WebRTC, MainResourceHasCacheControlNoStore,
// MainResourceHasCacheControlNoCache, SubresourceHasCacheControlNoStore,
// SubresourceHasCacheControlNoCache, ContainsPlugins, DocumentLoaded,
// OutstandingNetworkRequestOthers, RequestedMIDIPermission,
// RequestedAudioCapturePermission, RequestedVideoCapturePermission,
// RequestedBackForwardCacheBlockedSensors, RequestedBackgroundWorkPermission,
// BroadcastChannel, WebXR, SharedWorker, SharedWorkerMessage,
// SharedWorkerWithNoActiveClient, WebLocks, WebLocksContention, WebHID,
// WebBluetooth, WebShare, RequestedStorageAccessGrant, WebNfc,
// OutstandingNetworkRequestFetch, OutstandingNetworkRequestXHR, AppBanner,
// Printing, WebDatabase, PictureInPicture, SpeechRecognizer, IdleManager,
// PaymentManager, SpeechSynthesis, KeyboardLock, WebOTPService,
// OutstandingNetworkRequestDirectSocket, InjectedJavascript,
// InjectedStyleSheet, KeepaliveRequest, IndexedDBEvent, Dummy,
// JsNetworkRequestReceivedCacheControlNoStoreResource, WebRTCUsedWithCCNS,
// WebTransportUsedWithCCNS, WebSocketUsedWithCCNS, SmartCard,
// LiveMediaStreamTrack, UnloadHandler, ParserAborted, ContentSecurityHandler,
// ContentWebAuthenticationAPI, ContentFileChooser, ContentSerial,
// ContentFileSystemAccess, ContentMediaDevicesDispatcherHost, ContentWebBluetooth,
// ContentWebUSB, ContentMediaSessionService, ContentScreenReader, ContentDiscarded,
// EmbedderPopupBlockerTabHelper, EmbedderSafeBrowsingTriggeredPopupBlocker,
// EmbedderSafeBrowsingThreatDetails, EmbedderAppBannerManager,
// EmbedderDomDistillerViewerSource, EmbedderDomDistillerSelfDeletingRequestDelegate,
// EmbedderOomInterventionTabHelper, EmbedderOfflinePage,
// EmbedderChromePasswordManagerClientBindCredentialManager,
// EmbedderPermissionRequestManager, EmbedderModalDialog, EmbedderExtensions,
// EmbedderExtensionMessaging, EmbedderExtensionMessagingForOpenPort,
// EmbedderExtensionSentMessageToCachedFrame, EmbedderExtensionFrame,
// EmbedderPrivilegedWebContents, RequestedByWebViewClient,
// PostMessageByWebViewClient, CacheControlNoStoreDeviceBoundSessionTerminated,
// CacheLimitPrunedOnModerateMemoryPressure, CacheLimitPrunedOnCriticalMemoryPressure
type BackForwardCacheNotRestoredReason string

// BackForwardCacheNotRestoredReasonType Типы причин, по которым кэш
// не восстанавливается при переключении между предыдущим и следующим этапами.
// Допустимые значения: SupportPending, PageSupportNeeded, Circumstantial
type BackForwardCacheNotRestoredReasonType string

// ClientNavigationDisposition
// Допустимые значения: currentTab, newTab, newWindow, download
type ClientNavigationDisposition string

// ClientNavigationReason
// Допустимые значения: anchorClick, formSubmissionGet, formSubmissionPost,
// httpHeaderRefresh, initialFrameNavigation, metaTagRefresh, other,
// pageBlockInterstitial, reload, scriptInitiated
type ClientNavigationReason string

// CompilationCacheParams Параметры кэширования компиляции
// для каждого скрипта в Page.ProduceCompilationCache
type CompilationCacheParams struct {
	// URL скрипта, для которого необходимо создать запись в кэше компиляции.
	Url string `json:"url"`
	// Уведомление бэкэнда о том, рекомендуется ли немедленная компиляция
	// (фактический режим компиляции определяется бэкэндом).
	Eager bool `json:"eager,omitempty"`
}

// CrossOriginIsolatedContextType Указывает, является ли кадр
// изолированным от другого источника и почему это так.
// Допустимые значения: Isolated, NotIsolated, NotIsolatedFeatureDisabled
type CrossOriginIsolatedContextType string

// FileFilter
type FileFilter struct {
	Name    string   `json:"name,omitempty"`
	Accepts []string `json:"accepts,omitempty"`
}

// FileHandler
type FileHandler struct {
	Action string `json:"action"`
	Name   string `json:"name"`
	// Имитация карты: имя — это ключ, значение — accepts.
	Accepts []*FileFilter `json:"accepts,omitempty"`
	// Не буду повторять перечисления, для удобства сравнения буду
	// использовать строки. То же самое, что и с другими перечислениями ниже.
	LaunchType string `json:"launchType"`
}

// FontFamilies Коллекция универсальных семейств шрифтов.
type FontFamilies struct {
	// Стандартное семейство шрифтов.
	Standard string `json:"standard,omitempty"`
	// Фиксированное семейство шрифтов.
	Fixed string `json:"fixed,omitempty"`
	// Семейство шрифтов с засечками.
	Serif string `json:"serif,omitempty"`
	// Семейство шрифтов sansSerif.
	SansSerif string `json:"sansSerif,omitempty"`
	// Семейство курсивных шрифтов.
	Cursive string `json:"cursive,omitempty"`
	// Фантазийное семейство шрифтов.
	Fantasy string `json:"fantasy,omitempty"`
	// Семейство математических шрифтов.
	Math string `json:"math,omitempty"`
}

// FontSizes Размеры шрифта по умолчанию.
type FontSizes struct {
	// Стандартный размер шрифта по умолчанию.
	Standard int `json:"standard,omitempty"`
	// Размер шрифта по умолчанию фиксированный.
	Fixed int `json:"fixed,omitempty"`
}

// FrameResource Информация о ресурсе представлена ​​на странице.
type FrameResource struct {
	// URL ресурса.
	Url string `json:"url"`
	// Тип данного ресурса.
	Type *network.ResourceType `json:"type"`
	// MIME-тип ресурса определяется браузером.
	MimeType string `json:"mimeType"`
	// Отметка времени последнего изменения, предоставленная сервером.
	LastModified *network.TimeSinceEpoch `json:"lastModified,omitempty"`
	// Размер содержимого ресурса.
	ContentSize int `json:"contentSize,omitempty"`
	// Возвращает true, если ресурс не загрузился.
	Failed bool `json:"failed,omitempty"`
	// Возвращает true, если ресурс был отменен во время загрузки.
	Canceled bool `json:"canceled,omitempty"`
}

// FrameResourceTree Информация об иерархии фреймов,
// а также об их кэшированных ресурсах.
type FrameResourceTree struct {
	// Информация о структуре этого элемента дерева.
	Frame Frame `json:"frame"`
	// Фреймы-потомки.
	ChildFrames []*FrameResourceTree `json:"childFrames,omitempty"`
	// Информация о ресурсах фрейма.
	Resources []*FrameResource `json:"resources"`
}

// GatedAPIFeatures
// Допустимые значения: SharedArrayBuffers, SharedArrayBuffersTransferAllowed,
// PerformanceMeasureMemory, PerformanceProfile
type GatedAPIFeatures string

// ImageResource Разрешение изображения, использованное
// как в значке, так и на скриншоте.
type ImageResource struct {
	// Поле src в определении, но для единообразия заменено на url.
	Url   string `json:"url"`
	Sizes string `json:"sizes,omitempty"`
	Type  string `json:"type,omitempty"`
}

// InstallabilityError Ошибка установки.
type InstallabilityError struct {
	// Идентификатор ошибки (например, 'manifest-missing-suitable-icon').
	ErrorId string `json:"errorId"`
	// Список аргументов для обработки ошибок
	// (например, {name:'minimum-icon-size-in-pixels', value:'64'}).
	ErrorArguments []*InstallabilityErrorArgument `json:"errorArguments"`
}

// InstallabilityErrorArgument
type InstallabilityErrorArgument struct {
	// Имя аргумента (например, name:'minimum-icon-size-in-pixels').
	Name string `json:"name"`
	// Значение аргумента (например, значение: '64').
	Value string `json:"value"`
}

// LaunchHandler
type LaunchHandler struct {
	ClientMode string `json:"clientMode"`
}

// NavigationType Тип события frameNavigated.
// Допустимые значения: Navigation, BackForwardCacheRestore
type NavigationType string

// OriginTrial
type OriginTrial struct {
	TrialName        string                        `json:"trialName"`
	Status           *OriginTrialStatus            `json:"status"`
	TokensWithStatus []*OriginTrialTokenWithStatus `json:"tokensWithStatus"`
}

// OriginTrialStatus Статус пробной версии Origin.
// Допустимые значения: Enabled, ValidTokenNotProvided,
// OSNotSupported, TrialNotAllowed
type OriginTrialStatus string

// OriginTrialToken
type OriginTrialToken struct {
	Origin           string                      `json:"origin"`
	MatchSubDomains  bool                        `json:"matchSubDomains"`
	TrialName        string                      `json:"trialName"`
	ExpiryTime       *network.TimeSinceEpoch     `json:"expiryTime"`
	IsThirdParty     bool                        `json:"isThirdParty"`
	UsageRestriction OriginTrialUsageRestriction `json:"usageRestriction"`
}

// OriginTrialTokenStatus Поддержка пробной версии Origin
// (https://www.chromium.org/blink/origin-trials).
// Статус токена пробной версии Origin.
// Допустимые значения: Success, NotSupported, Insecure, Expired, WrongOrigin,
// InvalidSignature, Malformed, WrongVersion, FeatureDisabled, TokenDisabled,
// FeatureDisabledForUser, UnknownTrial
type OriginTrialTokenStatus string

// OriginTrialTokenWithStatus
type OriginTrialTokenWithStatus struct {
	RawTokenText string `json:"rawTokenText"`
	// Параметр parsedToken присутствует только тогда, когда
	// токен может быть извлечен и разобран.
	ParsedToken OriginTrialToken       `json:"parsedToken,omitempty"`
	Status      OriginTrialTokenStatus `json:"status"`
}

// OriginTrialUsageRestriction
// Допустимые значения: None, Subset
type OriginTrialUsageRestriction string

// PermissionsPolicyBlockLocator
type PermissionsPolicyBlockLocator struct {
	FrameId     *FrameId                      `json:"frameId"`
	BlockReason *PermissionsPolicyBlockReason `json:"blockReason"`
}

// PermissionsPolicyBlockReason Причина отключения функции политики разрешений.
// Допустимые значения: Header, IframeAttribute, InFencedFrameTree, InIsolatedApp
type PermissionsPolicyBlockReason string

// PermissionsPolicyFeature  Все функции политики разрешений. Этот перечисление
// должно соответствовать перечислению, определенному в файле
// services/network/public/cpp/permissions_policy/permissions_policy_features.json5.
// LINT.IfChange(PermissionsPolicyFeature)
// Допустимые значения: accelerometer, all-screens-capture, ambient-light-sensor,
// aria-notify, autofill, autoplay, bluetooth, browsing-topics, camera,
// captured-surface-control, ch-dpr, ch-device-memory, ch-downlink, ch-ect,
// ch-prefers-color-scheme, ch-prefers-reduced-motion,
// ch-prefers-reduced-transparency, ch-rtt, ch-save-data, ch-ua, ch-ua-arch,
// ch-ua-bitness, ch-ua-high-entropy-values, ch-ua-platform, ch-ua-model,
// ch-ua-mobile, ch-ua-form-factors, ch-ua-full-version, ch-ua-full-version-list,
// ch-ua-platform-version, ch-ua-wow64, ch-viewport-height, ch-viewport-width,
// ch-width, clipboard-read, clipboard-write, compute-pressure, controlled-frame,
// cross-origin-isolated, deferred-fetch, deferred-fetch-minimal, device-attributes,
// digital-credentials-create, digital-credentials-get, direct-sockets,
// direct-sockets-multicast, display-capture, document-domain, encrypted-media,
// execution-while-out-of-viewport, execution-while-not-rendered,
// focus-without-user-activation, fullscreen, frobulate, gamepad, geolocation,
// gyroscope, hid, identity-credentials-get, idle-detection, interest-cohort,
// keyboard-map, language-detector, language-model, local-fonts, local-network,
// local-network-access, loopback-network, magnetometer, manual-text,
// media-playback-while-not-visible, microphone, midi, on-device-speech-recognition,
// otp-credentials, payment, picture-in-picture, private-state-token-issuance,
// private-state-token-redemption, publickey-credentials-create,
// publickey-credentials-get, rewriter, screen-wake-lock, serial, shared-storage,
// shared-storage-select-url, smart-card, speaker-selection, storage-access,
// sub-apps, summarizer, sync-xhr, tools, translator, unload, usb, usb-unrestricted,
// vertical-scroll, web-app-installation, webnn, web-printing, web-share,
// window-management, writer, xr-spatial-tracking
type PermissionsPolicyFeature string

// PermissionsPolicyFeatureState
type PermissionsPolicyFeatureState struct {
	Feature *PermissionsPolicyFeature      `json:"feature"`
	Allowed bool                           `json:"allowed"`
	Locator *PermissionsPolicyBlockLocator `json:"locator,omitempty"`
}

// ProtocolHandler
type ProtocolHandler struct {
	Protocol string `json:"protocol"`
	Url      string `json:"url"`
}

// ReferrerPolicy Политика рефереров, используемая для навигации.
// Допустимые значения: noReferrer, noReferrerWhenDowngrade, origin,
// originWhenCrossOrigin, sameOrigin, strictOrigin,
// strictOriginWhenCrossOrigin, unsafeUrl
type ReferrerPolicy string

// RelatedApplication
type RelatedApplication struct {
	Id  string `josn:"id,pmitempty"`
	Url string `josn:"url"`
}

// ScopeExtension
type ScopeExtension struct {
	// Вместо кортежа это поле всегда возвращает сериализованную
	// строку для удобства понимания и сравнения.
	Origin            string `json:"origin"`
	HasOriginWildcard bool   `json:"hasOriginWildcard"`
}

// ScreencastFrameMetadata Метаданные кадра видеозаписи.
type ScreencastFrameMetadata struct {
	// Верхнее смещение в DIP.
	OffsetTop int `json:"offsetTop"`
	// Масштабный коэффициент страницы.
	PageScaleFactor int `json:"pageScaleFactor"`
	// Ширина экрана устройства в DIP-дисплее.
	DeviceWidth int `json:"deviceWidth"`
	// Высота экрана устройства в DIP-корпусе.
	DeviceHeight int `json:"deviceHeight"`
	// Положение горизонтальной прокрутки в пикселях CSS.
	ScrollOffsetX int `json:"scrollOffsetX"`
	// Положение вертикальной прокрутки в пикселях CSS.
	ScrollOffsetY int `json:"scrollOffsetY"`
	// Отметка времени замены кадра.
	Timestamp *network.TimeSinceEpoch `json:"timestamp,omitempty"`
}

// Screenshot
type Screenshot struct {
	Image      *ImageResource `json:"image"`
	FormFactor string         `json:"formFactor"`
	Label      string         `json:"label,omitempty"`
}

// ScriptFontFamilies Коллекция шрифтовых семейств для рукописного стиля.
type ScriptFontFamilies struct {
	// Название скрипта, для которого определены эти семейства шрифтов.
	Script string `json:"script"`
	// Универсальная коллекция шрифтов для рукописного стиля.
	FontFamilies *FontFamilies `json:"fontFamilies"`
}

// SecureContextType Указывает, является ли фрейм защищенным
// контекстом и почему это так.
// Допустимые значения: Secure, SecureLocalhost, InsecureScheme, InsecureAncestor
type SecureContextType string

// SecurityOriginDetails Дополнительная информация об
// источнике безопасности документа-фрейма.
type SecurityOriginDetails struct {
	// Указывает, является ли источником безопасности документа-фрейма одно
	// из локальных имен хостов (например, "localhost") или IP-адресов
	// (IPv4 127.0.0.0/8 или IPv6 ::1).
	IsLocalhost bool `json:"isLocalhost"`
}

// ShareTarget
type ShareTarget struct {
	Action  string `json:"action"`
	Method  string `json:"method"`
	Enctype string `json:"enctype"`
	// Вставьте ShareTargetParams.
	Title string        `json:"title,omitempty"`
	Text  string        `json:"text,omitempty"`
	Url   string        `json:"url,omitempty"`
	Files []*FileFilter `json:"files,omitempty"`
}

// Shortcut
type Shortcut struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// WebAppManifest
type WebAppManifest struct {
	BackgroundColor string `json:"backgroundColor,omitempty"`
	// Дополнительное описание, предоставленное манифестом.
	Description string `json:"description,omitempty"`
	Dir         string `json:"dir,omitempty"`
	Display     string `json:"display,omitempty"`
	// Режим отображения, определяемый пользователем.
	DisplayOverrides []string `json:"displayOverrides,omitempty"`
	// Обработчики для открытия файлов.
	FileHandlers []*FileHandler   `json:"fileHandlers,omitempty"`
	Icons        []*ImageResource `json:"icons,omitempty"`
	Id           string           `json:"id,omitempty"`
	Lang         string           `json:"lang,omitempty"`
	// TODO(crbug.com/1231886): Это поле нестандартное и
	// является частью эксперимента Chrome.
	// См.: https://github.com/WICG/web-app-launch/blob/main/launch_handler.md
	LaunchHandler             *LaunchHandler `json:"launchHandler,omitempty"`
	Name                      string         `json:"name,omitempty"`
	Orientation               string         `json:"orientation,omitempty"`
	PreferRelatedApplications bool           `json:"preferRelatedApplications,omitempty"`
	// Обработчики для открытия протоколов.
	ProtocolHandlers    []*ProtocolHandler    `json:"protocolHandlers,omitempty"`
	RelatedApplications []*RelatedApplication `json:"relatedApplications,omitempty"`
	Scope               string                `json:"scope,omitempty"`
	// Нестандартный вариант,
	// см. https://github.com/WICG/manifest-incubations/blob/gh-pages/scope_extensions-explainer.md
	ScopeExtensions []*ScopeExtension `json:"scopeExtensions,omitempty"`
	// Скриншоты, используемые Chromium.
	Screenshots []*Screenshot `json:"screenshots,omitempty"`
	ShareTarget *ShareTarget  `json:"shareTarget,omitempty"`
	ShortName   string        `json:"shortName,omitempty"`
	Shortcuts   []*Shortcut   `json:"shortcuts,omitempty"`
	StartUrl    string        `json:"startUrl,omitempty"`
	ThemeColor  string        `json:"themeColor,omitempty"`
}
