package dom

// BackendNode Узел бэкэнда с понятным именем.
type BackendNode struct {
	NodeType      int            `json:"nodeType"`
	NodeName      string         `json:"nodeName"`
	BackendNodeId *BackendNodeId `json:"backendNodeId"`
}
