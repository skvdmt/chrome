package set_auto_attach

// Config Конфигурация.
type Config struct {
	// Автоматическое ли прикрепление к связанным целям.
	AutoAttach bool `json:"autoAttach"`
	// Указывается, следует ли приостанавливать выполнение новых целей при
	// подключении к ним. Используйте Runtime.runIfWaitingForDebugger для
	// запуска приостановленных целей.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
	// Позволяет получить «плоский» доступ к сессии, указав атрибут sessionId
	// в командах. Мы планируем сделать это режимом по умолчанию, отказаться
	// от неплоского режима и в конечном итоге полностью его отключить.
	// См. crbug.com/991325
	Flatten bool `json:"flatten,omitempty"`
}
