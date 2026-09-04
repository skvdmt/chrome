package dom

// DetachedElementInfo Структура для хранения узла верхнего уровня
// отсоединенного дерева и массива его сохраненных потомков.
type DetachedElementInfo struct {
	TreeNode        *Node    `json:"treeNode"`
	RetainedNodeIds []NodeId `json:"retainedNodeIds"`
}
