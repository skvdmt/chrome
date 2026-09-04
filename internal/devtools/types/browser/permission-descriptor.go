package browser

// PermissionDescriptor Определение объекта PermissionDescriptor,
// заданного в API разрешений: https://w3c.github.io/permissions/#dom-permissiondescriptor.
type PermissionDescriptor struct {
	// Название разрешения. Смотрите https://cs.chromium.org/chromium/src/third_party/blink/renderer/modules/permissions/permission_descriptor.idl for valid permission names.
	Name string `json:"name"`
	// Для разрешения "midi" можно также указать управление sysex.
	Sysex bool `json:"sysex,omitempty"`
	// Для разрешения "push" можно указать userVisibleOnly. Обратите внимание,
	// что userVisibleOnly = true — единственный поддерживаемый в настоящее время тип.
	UserVisibleOnly bool `json:"userVisibleOnly,omitempty"`
	// Для разрешения доступа к "clipboard" можно указать allowWithoutSanitization.
	AllowWithoutSanitization bool `json:"allowWithoutSanitization,omitempty"`
	// Для разрешения "fullscreen" необходимо указать allowWithoutGesture:true.
	AllowWithoutGesture bool `json:"allowWithoutGesture,omitempty"`
	// Для разрешения доступа к "camera" можно указать panTiltZoom.
	PanTiltZoom bool `json:"panTiltZoom,omitempty"`
}
