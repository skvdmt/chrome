package dom

import (
	"github.com/skvdmt/chrome/internal/devtools/types/network"
	"github.com/skvdmt/chrome/internal/devtools/types/page"
)

// Node Взаимодействие с DOM реализуется с помощью зеркальных объектов,
// представляющих собой фактические узлы DOM. DOMNode — это базовый тип
// зеркального отображения узла.
type Node struct {
	// ID узла.
	NodeId NodeId `json:"nodeId"`
	// ID родительского узла, если таковой имеется.
	ParentId NodeId `json:"parentId"`
	// ID бэкенд-узла.
	BackendNodeId BackendNodeId `json:"backendNodeId"`
	// Тип узла.
	NodeType int `json:"nodeType"`
	// Имя узла.
	NodeName string `json:"nodeName"`
	// Локальное имя узла.
	LocalName string `json:"localName"`
	// Значение узла.
	NodeValue string `json:"nodeValue"`
	// Количество дочерних узлов.
	ChildNodeCount int `json:"childNodeCount"`
	// Дочерние узлы.
	Children []*Node `json:"children"`
	// Аттрибуты узла срез в виде: [name1, value1, name2, value2].
	Attributes []string `json:"attributes"`
	// URL документа, на который указывает узел Document или FrameOwner.
	DocumentURL string `json:"documentURL"`
	// Базовый URL-адрес, который узел Document или FrameOwner
	// использует для автодополнения URL-адресов.
	BaseURL string `json:"baseURL"`
	// Публичный id.
	PublicId string `json:"publicId"`
	// Системный id.
	SystemId string `json:"systemId"`
	// Внутреннее подмножество типа документа.
	InternalSubset string `json:"internalSubset"`
	// В случае XML-документов указывается XML-версия документа.
	XmlVersion string `json:"xmlVersion"`
	// Имя аттрибута.
	Name string `json:"name"`
	// Значение аттрибута.
	Value string `json:"value"`
	// Псевдотип элемента для этого узла.
	PseudoType PseudoType `json:"pseudoType"`
	// Псевдоидентификатор элемента для этого узла.
	// Присутствует только при наличии допустимого псевдотипа.
	PseudoIdentifier string `json:"pseudoIdentifier"`
	// Теневой тип корня.
	ShadowRootType ShadowRootType `json:"shadowRootType"`
	// ID фрейма.
	FrameId *page.FrameId `json:"frameId"`
	// Содержимое документа для элементов, являющихся владельцами фрейма.
	ContentDocument *Node `json:"contentDocument"`
	// Теневые корни.
	ShadowRoots []*Node `json:"shadowRoots"`
	// Фрагмент документа с содержимым для элементов шаблона.
	TemplateContent *Node `json:"templateContent"`
	// Псевдоэлементы, связанные с этим узлом.
	PseudoElements []*Node `json:"pseudoElements"`
	// Распределенные узлы для заданной точки вставки.
	DistributedNodes []*BackendNode `json:"distributedNodes"`
	// Является ли узел SVG.
	IsSVG                    bool                  `json:"isSVG"`
	CompatibilityMode        CompatibilityMode     `json:"compatibilityMode"`
	AssignedSlot             *BackendNode          `json:"assignedSlot"`
	IsScrollable             bool                  `json:"isScrollable"`
	AffectedByStartingStyles bool                  `json:"affectedByStartingStyles"`
	AdoptedStyleSheets       []*StyleSheetId       `json:"adoptedStyleSheets"`
	AdProvenance             *network.AdProvenance `json:"adProvenance"`
}
