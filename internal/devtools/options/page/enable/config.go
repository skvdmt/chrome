package enable

// Config Конфигурация.
type Config struct {
	// Если значение равно true, событие Page.fileChooserOpened будет
	// сгенерировано независимо от состояния, установленного командой
	// Page.setInterceptFileChooserDialog (по умолчанию: false).
	EnableFileChooserOpenedEvent bool `json:"enableFileChooserOpenedEvent,omitempty"`
}
