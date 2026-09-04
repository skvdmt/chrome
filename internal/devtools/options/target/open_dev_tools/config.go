package open_dev_tools

import "github.com/skvdmt/chrome/internal/devtools/types/target"

// Config Конфигурация.
type Config struct {
	// Идентификатор целевой страницы или вкладки.
	TargetId *target.TargetId `json:"targetId"`
	// Идентификатор панели, которую мы хотим, чтобы DevTools открыли изначально.
	// В настоящее время поддерживаются следующие панели:
	// elements, console, network, sources, resources, timeline,
	// chrome-recorder, heap-profiler, lighthouse и security.
	PanelId string `json:"panelId,omitempty"`
}
