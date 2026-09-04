package add_script_to_evaluate_on_new_document

// Config Конфигурация.
type Config struct {
	Source string `json:"source"`
	// Если указано, создает изолированный мир с заданным именем и выполняет
	// в нем заданный скрипт. Это имя мира будет использоваться в качестве
	// ExecutionContextDescription::name при генерации соответствующего события.
	WorldName string `json:"worldName,omitempty"`
	// Указывает, должен ли скрипт быть доступен через API командной строки;
	// по умолчанию — false.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`
	// Если значение равно true, скрипт запускается немедленно
	// в существующих контекстах выполнения или мирах. По умолчанию: false.
	RunImmediately bool `json:"runImmediately,omitempty"`
}
