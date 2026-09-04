package enable

// Config
type Config struct {
	// Размер буфера (в байтах) для сохранения полезной нагрузки сетевых
	// запросов (XHR и т. д.). Это максимальный объем данных в байтах,
	// который будет собран в рамках данного сеанса DevTools.
	MaxTotalBufferSize int `json:"maxTotalBufferSize,omitempty"`
	// Размер буфера (в байтах) на каждый ресурс, используемый
	// при сохранении сетевых полезных данных (XHR и т. д.).
	MaxResourceBufferSize int `json:"maxResourceBufferSize,omitempty"`
	// Максимальный размер тела запроса (в байтах), включаемый
	// в уведомление requestWillBeSent.
	MaxPostDataSize int `json:"maxPostDataSize,omitempty"`
	// Следует ли сообщать о событиях отправки и получения фрагментов DirectSocket.
	ReportDirectSocketTraffic bool `json:"reportDirectSocketTraffic,omitempty"`
	// Включает сохранение тел ответов вне процесса рендеринга, чтобы они
	// сохранялись при навигации между процессами. Требует установки
	// параметра `maxTotalBufferSize`. По умолчанию в данный момент
	// установлено значение `false`. Использование этого поля признается
	// устаревшим в пользу специальной команды `configureDurableMessages`
	// из-за риска возникновения взаимных блокировок при ожидании
	// события `Network.enable` перед вызовом `Runtime.runIfWaitingForDebugger`.
	EnableDurableMessages bool `json:"enableDurableMessages,omitempty"`
}
