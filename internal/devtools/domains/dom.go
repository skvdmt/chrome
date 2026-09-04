package domains

import (
	"github.com/skvdmt/chrome/internal/devtools/options/dom/copy_to"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/describe_node"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/enable"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/focus"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/force_show_popover"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/get_anchor_element"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/get_box_model"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/get_container_for_node"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/get_content_quads"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/get_document"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/get_node_for_location"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/get_nodes_for_subtree_by_style"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/get_outer_html"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/move_to"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/perform_search"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/query_selector"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/resolve_node"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/scroll_into_view_if_needed"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/set_attributes_as_text"
	"github.com/skvdmt/chrome/internal/devtools/options/dom/set_file_input_files"
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	"github.com/skvdmt/chrome/internal/devtools/types/page"
	rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"
	"github.com/skvdmt/chrome/internal/devtools/types/target"
	"github.com/skvdmt/chrome/internal/model"
)

// Dom Объектная модель документа.
type Dom struct {
	client *model.Client
	debug  *model.Debug
	// Текущая сессия.
	CurrentSessionId *target.SessionId
}

// NewDom Конструктор.
func NewDom(c *model.Client, d *model.Debug) *Dom {
	d.Debug("dom created")
	return &Dom{
		client: c,
		debug:  d,
	}
}

// DescribeNode Описание узла.
func (d *Dom) DescribeNode(options ...describe_node.Option) (*dom.Node, error) {
	c := &describe_node.Config{}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.DESCRIBE_NODE,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		Node *dom.Node `json:"node"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.Node, nil
}

// Disable Отключает DOM-агент для данной страницы.
func (d *Dom) Disable() error {
	return d.client.Exec(dom.DISABLE, nil, model.WithSessionId(d.CurrentSessionId))
}

// Enable Включает DOM-агент для данной страницы.
func (d *Dom) Enable(options ...enable.Option) error {
	c := &enable.Config{}
	for _, o := range options {
		o(c)
	}
	return d.client.Exec(
		dom.ENABLE,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// Focus Фокусирует внимание на заданном элементе.
func (d *Dom) Focus(options ...focus.Option) error {
	c := &focus.Config{}
	for _, o := range options {
		o(c)
	}
	return d.client.Exec(
		dom.FOCUS,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// GetAttributes Возвращает атрибуты для указанного узла.
func (d *Dom) GetAttributes(nodeId dom.NodeId) ([]string, error) {
	r, err := d.client.Query(
		dom.GET_ATTRIBUTES,
		model.ForceJSONMarshal(struct {
			NodeId dom.NodeId `json:"nodeId"`
		}{
			NodeId: nodeId,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	a := struct {
		Attributes []string `json:"attributes"`
	}{}
	model.ForceJSONUnmarshal(r, a)
	return a.Attributes, nil
}

// GetBoxModel Блоки узла.
func (d *Dom) GetBoxModel(options ...get_box_model.Option) (*dom.BoxModel, error) {
	c := &get_box_model.Config{}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.GET_BOX_MODEL,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	b := struct {
		Model *dom.BoxModel `json:"model"`
	}{}
	model.ForceJSONUnmarshal(r, b)
	return b.Model, nil
}

// GetDocument Коревой узел.
func (d *Dom) GetDocument(options ...get_document.Option) (*dom.Node, error) {
	c := &get_document.Config{}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.GET_DOCUMENT,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		Root *dom.Node `json:"root"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.Root, nil
}

// GetNodeForLocation Возвращает идентификатор узла в указанном
// местоположении. В зависимости от того, включен ли домен DOM,
// возвращается либо идентификатор узла, либо нет.
func (d *Dom) GetNodeForLocation(x, y int, options ...get_node_for_location.Option) (
	*dom.BackendNodeId, *page.FrameId, *dom.NodeId, error) {
	c := &get_node_for_location.Config{
		X: x,
		Y: y,
	}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.GET_NODE_FOR_LOCATION,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, nil, nil, err
	}
	e := struct {
		// Resulting node.
		BackendNodeId *dom.BackendNodeId `json:"backendNodeId"`
		//Frame this node belongs to.
		FrameId *page.FrameId `json:"frameId"`
		// Id of the node at given coordinates, only when enabled and requested document.
		NodeId *dom.NodeId `json:"nodeId"`
	}{}
	model.ForceJSONUnmarshal(r, &e)
	return e.BackendNodeId, e.FrameId, e.NodeId, nil
}

// GetOuterHTML Возвращает HTML-разметку узла.
func (d *Dom) GetOuterHTML(options ...get_outer_html.Option) (string, error) {
	c := &get_outer_html.Config{}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.GET_OUTER_HTML,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return "", err
	}
	o := struct {
		// Outer HTML markup.
		OuterHTML string `json:"outerHTML"`
	}{}
	model.ForceJSONUnmarshal(r, &o)
	return o.OuterHTML, nil
}

// HideHighlight Скрывает все выделенные области.
func (d *Dom) HideHighlight() error {
	return d.client.Exec(dom.HIDE_HIGHLIGH, nil, model.WithSessionId(d.CurrentSessionId))
}

// HighlightNode Выделяет DOM-узел.
func (d *Dom) HighlightNode() error {
	return d.client.Exec(dom.HIGHLIGHT_NODE, nil, model.WithSessionId(d.CurrentSessionId))
}

// HighlightRect Выделяет заданный прямоугольник.
func (d *Dom) HighlightRect() error {
	return d.client.Exec(dom.HISHLIGHT_RECT, nil, model.WithSessionId(d.CurrentSessionId))
}

// MoveTo Перемещает узел в новый контейнер, размещая его перед заданным якорем.
func (d *Dom) MoveTo(nodeId, targetNodeId dom.NodeId, options ...move_to.Option) (*dom.NodeId, error) {
	c := &move_to.Config{
		NodeId:       nodeId,
		TargetNodeId: targetNodeId,
	}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.MOVE_TO,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// New id of the moved node.
		NodeId *dom.NodeId `json:"nodeId"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeId, nil
}

// QuerySelector Поиск узла по селектору.
func (d *Dom) QuerySelector(selector string, options ...query_selector.Option) (dom.NodeId, error) {
	c := &query_selector.Config{
		Selector: selector,
	}
	for _, o := range options {
		o(c)
	}
	if c.NodeId == 0 {
		r, err := d.GetDocument()
		if err != nil {
			return 0, err
		}
		c.NodeId = r.NodeId
	}
	r, err := d.client.Query(
		dom.QUERY_SELECTOR,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return 0, err
	}
	n := struct {
		NodeId dom.NodeId `json:"nodeId"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	if n.NodeId == 0 {
		return 0, model.ERR_NODE_NOT_FOUND
	}
	return n.NodeId, nil
}

// QuerySelectorAll Поиск узлов по селектору.
func (d *Dom) QuerySelectorAll(selector string, options ...query_selector.Option) ([]dom.NodeId, error) {
	c := &query_selector.Config{
		Selector: selector,
	}
	for _, o := range options {
		o(c)
	}
	if c.NodeId == 0 {
		r, err := d.GetDocument()
		if err != nil {
			return nil, err
		}
		c.NodeId = r.NodeId
	}
	r, err := d.client.Query(
		dom.QUERY_SELECTOR_ALL,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		NodeIds []dom.NodeId `json:"nodeIds"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeIds, nil
}

// RemoveAttribute Удаляет атрибут с заданным именем из элемента с заданным id.
func (d *Dom) RemoveAttribute(nodeId dom.NodeId, name string) error {
	return d.client.Exec(
		dom.REMOVE_ATTRIBUTE,
		model.ForceJSONMarshal(struct {
			// ID элемента, из которого нужно удалить атрибут.
			NodeId dom.NodeId `json:"nodeId"`
			// Имя атрибута, который нужно удалить.
			Name string `json:"name"`
		}{
			NodeId: nodeId,
			Name:   name,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// RemoveNode Удаляет узел с заданным идентификатором.
func (d *Dom) RemoveNode(nodeId *dom.NodeId) error {
	return d.client.Exec(
		dom.REMOVE_NODE,
		model.ForceJSONMarshal(struct {
			// Идентификатор узла, который нужно удалить.
			NodeId *dom.NodeId `json:"nodeId"`
		}{
			NodeId: nodeId,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// RequestChildNodes Не реализован из-за возврата ответа в виде события.
func (d *Dom) RequestChildNodes() {
}

// RequestNode Запрос на отправку узла вызывающей стороне на основании ссылки
// на объект узла JavaScript. Все узлы, образующие путь от узла к корню, также
// отправляются клиенту в виде серии уведомлений setChildNodes.
func (d *Dom) RequestNode(objectId *rnt.RemoteObjectId) (*dom.NodeId, error) {
	r, err := d.client.Query(
		dom.REQUEST_NODE,
		model.ForceJSONMarshal(struct {
			// Идентификатор объекта JavaScript для преобразования в Node.js.
			ObjectId *rnt.RemoteObjectId `json:"objectId"`
		}{
			ObjectId: objectId,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Node id for given object.
		NodeId *dom.NodeId `json:"nodeId"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeId, nil
}

// ResolveNode Выполняет разрешение объекта узла JavaScript
// для заданного NodeId или BackendNodeId.
func (d *Dom) ResolveNode(options ...resolve_node.Option) (*rnt.RemoteObject, error) {
	c := &resolve_node.Config{}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.RESOLVE_NODE,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	o := struct {
		// Объект JavaScript, обертывающий заданный узел.
		Object *rnt.RemoteObject `json:"object"`
	}{}
	model.ForceJSONUnmarshal(r, &o)
	return o.Object, nil
}

// ScrollIntoViewIfNeeded Прокручивает указанный прямоугольник заданного узла,
// чтобы он стал видимым, если он еще не виден. Примечание: для идентификации
// узла необходимо передать ровно одно значение из диапазона
// nodeId, backendNodeId и objectId.
func (d *Dom) ScrollIntoViewIfNeeded(options ...scroll_into_view_if_needed.Option) error {
	c := &scroll_into_view_if_needed.Config{}
	for _, o := range options {
		o(c)
	}
	return d.client.Exec(
		dom.SCROLL_INTO_VIEW_IF_NEEDED,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// SetAttributesAsText Устанавливает атрибуты для элемента с заданным идентификатором.
// Этот метод полезен, когда пользователь редактирует существующее значение
// атрибута и вводит несколько пар «имя/значение атрибута».
func (d *Dom) SetAttributesAsText(
	nodeId *dom.NodeId,
	text string,
	options ...set_attributes_as_text.Option,
) error {
	c := &set_attributes_as_text.Confug{
		NodeId: nodeId,
		Text:   text,
	}
	for _, o := range options {
		o(c)
	}
	return d.client.Exec(
		dom.SET_ATTRIBUTES_AS_TEXT,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// SetAttributeValue Устанавливает атрибут для элемента с заданным идентификатором.
func (d *Dom) SetAttributeValue(nodeId *dom.NodeId, name, value string) error {
	return d.client.Exec(
		dom.SET_ATTRIBUTE_VALUE,
		model.ForceJSONMarshal(struct {
			// Идентификатор элемента, для которого нужно установить атрибут.
			NodeId *dom.NodeId `json:"nodeId"`
			// Название атрибута.
			Name string `json:"name"`
			// Значение атрибута.
			Value string `json:"value"`
		}{
			NodeId: nodeId,
			Name:   name,
			Value:  value,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// SetFileInputFiles Задает файлы для указанного элемента ввода файла.
func (d *Dom) SetFileInputFiles(
	files []string,
	options ...set_file_input_files.Option,
) error {
	c := &set_file_input_files.Config{
		Files: files,
	}
	for _, o := range options {
		o(c)
	}
	return d.client.Exec(
		dom.SET_FILE_INPUT_FILES,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// SetNodeName Задает имя узла для узла с заданным идентификатором.
func (d *Dom) SetNodeName(nodeId *dom.NodeId, name string) (*dom.NodeId, error) {
	r, err := d.client.Query(
		dom.SET_NODE_NAME,
		model.ForceJSONMarshal(struct {
			// Идентификатор узла, для которого нужно задать имя.
			NodeId *dom.NodeId `json:"nodeId"`
			// Новое название узла.
			Name string `json:"name"`
		}{
			NodeId: nodeId,
			Name:   name,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Идентификатор нового узла.
		NodeId *dom.NodeId `json:"nodeId"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeId, nil
}

// SetNodeValue Устанавливает значение для узла с заданным идентификатором.
func (d *Dom) SetNodeValue(nodeId *dom.NodeId, value string) error {
	return d.client.Exec(
		dom.SET_NODE_VALUE,
		model.ForceJSONMarshal(struct {
			// Идентификатор узла, для которого нужно установить значение.
			NodeId *dom.NodeId `json:"nodeId"`
			// Новое значение узла.
			Value string `json:"value"`
		}{
			NodeId: nodeId,
			Value:  value,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// SetOuterHTML Устанавливает HTML-разметку узла.
func (d *Dom) SetOuterHTML(nodeId *dom.NodeId, outerHTML string) error {
	return d.client.Exec(
		dom.SET_OUTHER_HTML,
		model.ForceJSONMarshal(struct {
			// Идентификатор узла, для которого нужно задать разметку.
			NodeId *dom.NodeId `json:"nodeId"`
			// Внешняя HTML-разметка для установки.
			OuterHTML string `json:"outerHTML"`
		}{
			NodeId:    nodeId,
			OuterHTML: outerHTML,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// CollectClassNamesFromSubtree Собирает имена классов для узла с
// заданным идентификатором и всех его дочерних узлов.
func (d *Dom) CollectClassNamesFromSubtree(nodeId *dom.NodeId) ([]string, error) {
	r, err := d.client.Query(
		dom.COLLECT_CLASS_NAMES_FROM_SUBTREE,
		model.ForceJSONMarshal(struct {
			// Идентификатор узла, для которого нужно собрать имена классов.
			NodeId *dom.NodeId `json:"nodeId"`
		}{
			NodeId: nodeId,
		}),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Список названий классов.
		ClassNames []string `json:"classNames"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.ClassNames, nil
}

// CopyTo Создает глубокую копию указанного узла и помещает
// ее в целевой контейнер перед заданным якорем.
func (d *Dom) CopyTo(
	nodeId, targetNodeId *dom.NodeId,
	options ...copy_to.Option) (*dom.NodeId, error) {
	c := &copy_to.Config{
		NodeId:       nodeId,
		TargetNodeId: targetNodeId,
	}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.COPY_TO,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Идентификатор клонированного узла.
		NodeId *dom.NodeId `json:"nodeId"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeId, nil
}

// DiscardSearchResults Удаляет результаты поиска из сессии с
// заданным идентификатором. Метод getSearchResults больше
// не следует вызывать для этого поиска.
func (d *Dom) DiscardSearchResults() error {
	return d.client.Exec(
		dom.DISCARD_SEARCH_RESULTS,
		model.ForceJSONMarshal(struct {
			// Уникальный идентификатор поисковой сессии.
			SearchId string `json:"searchId"`
		}{}),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// ForceShowInterest При включении этот API заставляет элемент проявлять
// интерес к целевому элементу, поддерживая интерес активным до тех пор,
// пока он не будет отключен.
func (d *Dom) ForceShowInterest(nodeId *dom.NodeId, enable bool) error {
	return d.client.Exec(
		dom.FORCE_SHOW_INTEREST,
		model.ForceJSONMarshal(struct {
			// Идентификатор HTMLElement, вызывающего интерес.
			NodeId *dom.NodeId `json:"nodeId"`
			// Если true, то открывается и удерживается процентная ставка.
			// Если false, то принудительно выплачиваются проценты.
			Enable bool `json:"enable"`
		}{
			NodeId: nodeId,
			Enable: enable,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// ForceShowPopover При включении этот API принудительно открывает
// всплывающее окно, идентифицированное по nodeId, и держит его
// открытым до тех пор, пока не будет отключен.
func (d *Dom) ForceShowPopover(
	nodeId *dom.NodeId,
	enable bool,
	options ...force_show_popover.Option,
) ([]*dom.NodeId, error) {
	c := &force_show_popover.Config{
		NodeId: nodeId,
		Enable: enable,
	}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.FORCE_SHOW_POPOVER,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Список всплывающих окон, которые были закрыты для
		// соблюдения порядка их наложения друг на друга.
		NodeIds []*dom.NodeId `json:"nodeIds"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeIds, nil
}

// GetAnchorElement Возвращает целевой элемент привязки для заданного запроса привязки
// в соответствии с https://www.w3.org/TR/css-anchor-position-1/#target.
func (d *Dom) GetAnchorElement(
	nodeId *dom.NodeId,
	options ...get_anchor_element.Option,
) (*dom.NodeId, error) {
	c := &get_anchor_element.Config{
		NodeId: nodeId,
	}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.GET_ANCHOR_ELEMENT,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Якорный элемент заданного якорного запроса.
		NodeId *dom.NodeId `json:"nodeId"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeId, nil
}

// GetContainerForNode Возвращает контейнер запроса для заданного узла на основе
// условий запроса контейнера: containerName, физические и логические оси, а также
// то, запрашивает ли он элементы с состоянием прокрутки или закрепленные элементы.
// Если оси не указаны и queriesScrollState имеет значение false, возвращается
// контейнер стиля, который является непосредственным родительским элементом или
// ближайшим элементом с совпадающим именем контейнера.
func (d *Dom) GetContainerForNode(
	nodeId *dom.NodeId,
	options ...get_container_for_node.Option,
) (*dom.NodeId, error) {
	c := &get_container_for_node.Config{
		NodeId: nodeId,
	}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.GET_CONTAINER_FOR_NODE,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Контейнерный узел для заданного узла или null, если узел не найден.
		NodeId *dom.NodeId `json:"nodeId,omitempty"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeId, nil
}

// GetContentQuads Возвращает квадраты, описывающие положение узла на странице.
// Для строчных узлов этот метод может возвращать несколько квадратов.
func (d *Dom) GetContentQuads(options ...get_content_quads.Option) ([]*dom.Quad, error) {
	c := &get_content_quads.Config{}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.GET_CONTENT_QUADS,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Квадраты, описывающие расположение узлов относительно области просмотра.
		Quads []*dom.Quad `json:"quads"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.Quads, nil
}

// GetDetachedDomNodes Возвращает список отсоединенных узлов.
func (d *Dom) GetDetachedDomNodes() ([]*dom.DetachedElementInfo, error) {
	r, err := d.client.Query(
		dom.GET_DETACHED_DOM_NODES,
		nil,
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	e := struct {
		// Список отсоединенных узлов.
		DetachedNodes []*dom.DetachedElementInfo `json:"detachedNodes"`
	}{}
	model.ForceJSONUnmarshal(r, &e)
	return e.DetachedNodes, nil
}

// GetElementByRelation Возвращает NodeId соответствующего элемента
// в соответствии с определенными отношениями.
func (d *Dom) GetElementByRelation(nodeId *dom.NodeId, relation string) (*dom.NodeId, error) {
	r, err := d.client.Query(
		dom.GET_ELEMENT_BY_RELATION,
		model.ForceJSONMarshal(struct {
			// Идентификатор узла, из которого следует запрашивать связь.
			NodeId *dom.NodeId `json:"nodeId"`
			// Тип отношения, которое нужно получить.
			// Допустимые значения: PopoverTarget, InterestTarget, CommandFor
			Relation string `json:"relation"`
		}{
			NodeId:   nodeId,
			Relation: relation,
		}),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// NodeId of the element matching the queried relation.
		NodeId *dom.NodeId `json:"nodeId"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeId, nil
}

// GetFileInfo Возвращает информацию о файле для заданной оболочки File.
func (d *Dom) GetFileInfo() (path *string, err error) {
	r, err := d.client.Query(
		dom.GET_FILE_INFO,
		model.ForceJSONMarshal(struct {
			// Идентификатор объекта JavaScript, представляющего собой обертку для узла.
			ObjectId *rnt.RemoteObjectId `json:"objectId"`
		}{}),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	p := struct {
		Path *string `json:"path"`
	}{}
	model.ForceJSONUnmarshal(r, &p)
	return p.Path, nil
}

// GetFrameOwner Возвращает узел iframe, которому принадлежит iframe с заданным доменом.
func (d *Dom) GetFrameOwner(frameId *page.FrameId) (*dom.BackendNodeId, *dom.NodeId, error) {
	r, err := d.client.Query(
		dom.GET_FRAME_OWNER,
		model.ForceJSONMarshal(struct {
			FrameId *page.FrameId `json:"frameId"`
		}{
			FrameId: frameId,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, nil, err
	}
	n := struct {
		// Результирующий узел.
		BackendNodeId *dom.BackendNodeId `json:"backendNodeId"`
		// Идентификатор узла в заданных координатах, только
		// если он включен и запрошен документ.
		NodeId *dom.NodeId `json:"nodeId"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.BackendNodeId, n.NodeId, nil
}

// GetNodesForSubtreeByStyle Находит узлы с заданным вычисленным стилем в поддереве.
func (d *Dom) GetNodesForSubtreeByStyle(
	nodeId *dom.NodeId,
	computedStyles []*dom.CSSComputedStyleProperty,
	options ...get_nodes_for_subtree_by_style.Option,
) ([]*dom.NodeId, error) {
	c := &get_nodes_for_subtree_by_style.Config{
		NodeId:         nodeId,
		ComputedStyles: computedStyles,
	}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.GET_NODES_FOR_SUBTREE_BY_STYLE,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Resulting nodes.
		NodeIds []*dom.NodeId `json:"nodeIds"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeIds, nil
}

// GetNodeStackTraces Получает трассировку стека, связанную с узлом. На данный
// момент предоставляет трассировку стека только для создания узла.
func (d *Dom) GetNodeStackTraces(nodeId *dom.NodeId) (*rnt.StackTrace, error) {
	r, err := d.client.Query(
		dom.GET_NODES_STACK_TRACES,
		model.ForceJSONMarshal(struct {
			// Идентификатор узла, для которого необходимо получить трассировку стека.
			NodeId *dom.NodeId `json:"nodeId"`
		}{
			NodeId: nodeId,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	c := struct {
		// Creation stack trace, if available.
		Creation *rnt.StackTrace `json:"creation"`
	}{}
	model.ForceJSONUnmarshal(r, &c)
	return c.Creation, nil
}

// GetQueryingDescendantsForContainer Возвращает потомков контейнера,
// выполняющего запросы к этому контейнеру.
func (d *Dom) GetQueryingDescendantsForContainer(nodeId *dom.NodeId) ([]*dom.NodeId, error) {
	r, err := d.client.Query(
		dom.GET_QUERYING_DESCENDANTS_FOR_CONTAINER,
		model.ForceJSONMarshal(struct {
			// Идентификатор узла контейнера, от которого нужно найти потомков для запроса.
			NodeId *dom.NodeId `json:"nodeId"`
		}{
			NodeId: nodeId,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Дочерние узлы, содержащие запросы к контейнеру, обращаются к заданному контейнеру.
		NodeIds []*dom.NodeId `json:"nodeIds"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeIds, nil
}

// GetRelayoutBoundary Возвращает идентификатор ближайшего предка,
// являющегося границей точки переадресации.
func (d *Dom) GetRelayoutBoundary(nodeId *dom.NodeId) (*dom.NodeId, error) {
	r, err := d.client.Query(
		dom.GET_RELAYOUT_BOUNDARY,
		model.ForceJSONMarshal(struct {
			// Идентификатор узла.
			NodeId *dom.NodeId `json:"nodeId"`
		}{
			NodeId: nodeId,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Идентификатор граничного узла Relayout для заданного узла.
		NodeId *dom.NodeId `json:"nodeId"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeId, nil
}

// GetSearchResults Возвращает результаты поиска из заданного fromIndex в
// заданный toIndex, полученные в результате поиска по заданному идентификатору.
func (d *Dom) GetSearchResults(searchId string, fromIndex, toIndex int) ([]*dom.NodeId, error) {
	r, err := d.client.Query(
		dom.GET_SEARCH_RESULTS,
		model.ForceJSONMarshal(struct {
			// Уникальный идентификатор поисковой сессии.
			SearchId string `json:"searchId"`
			// Начальный индекс результатов поиска, которые должны быть возвращены.
			FromIndex int `json:"fromIndex"`
			// Конечный индекс возвращаемого результата поиска.
			ToIndex int `json:"toIndex"`
		}{
			SearchId:  searchId,
			FromIndex: fromIndex,
			ToIndex:   toIndex,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Идентификаторы узлов результатов поиска.
		NodeIds []*dom.NodeId `json:"nodeIds"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeIds, nil
}

// GetTopLayerElements Возвращает идентификаторы узлов (NodeIds) элементов верхнего слоя.
// Верхний слой отображается ближе всего к пользователю в пределах области просмотра,
// поэтому его элементы всегда отображаются поверх всего остального контента.
func (d *Dom) GetTopLayerElements() ([]*dom.NodeId, error) {
	r, err := d.client.Query(
		dom.GET_TOP_LAYER_ELEMENTS,
		nil,
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Идентификаторы узлов элементов верхнего слоя.
		NodeIds []*dom.NodeId `json:"nodeIds"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeIds, err
}

// MarkUndoableState Отмечает последнее невыполнимое состояние.
func (d *Dom) MarkUndoableState() error {
	return d.client.Exec(
		dom.MARK_UNDOABLE_STATE,
		nil,
		model.WithSessionId(d.CurrentSessionId),
	)
}

// PerformSearch Выполняет поиск заданной строки в DOM-дереве.
// Используйте GetSearchResults для доступа к результатам поиска или
// CancelSearch для завершения текущей сессии поиска.
func (d *Dom) PerformSearch(query string, options ...perform_search.Option) (SearchId *string, ResultCount *int, err error) {
	c := &perform_search.Config{
		Query: query,
	}
	for _, o := range options {
		o(c)
	}
	r, err := d.client.Query(
		dom.PERFORM_SEARCH,
		model.ForceJSONMarshal(c),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, nil, err
	}
	s := struct {
		// Уникальный идентификатор поисковой сессии.
		SearchId *string `json:"searchId"`
		// Количество результатов поиска.
		ResultCount *int `json:"resultCount"`
	}{}
	model.ForceJSONUnmarshal(r, &s)
	return s.SearchId, s.ResultCount, nil
}

// PushNodeByPathToFrontend Запрашивает отправку узла вызывающей стороне
// по его пути. // FIXME, используйте XPath
func (d *Dom) PushNodeByPathToFrontend(path string) (*dom.NodeId, error) {
	r, err := d.client.Query(
		dom.PUSH_NODE_BY_PATH_TO_FRONTEND,
		model.ForceJSONMarshal(struct {
			// Путь к узлу в проприетарном формате.
			Path string `json:"path"`
		}{}),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Идентификатор узла для заданного пути.
		NodeId *dom.NodeId `json:"nodeId"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeId, nil
}

// PushNodesByBackendIdsToFrontend Запрашивает отправку пакета узлов вызывающей
// стороне на основе идентификаторов их бэкэнд-узлов.
func (d *Dom) PushNodesByBackendIdsToFrontend(backendNodeIds []*dom.BackendNodeId) ([]*dom.NodeId, error) {
	r, err := d.client.Query(
		dom.PUSH_NODES_BY_BACKEND_IDS_TO_FRONTEND,
		model.ForceJSONMarshal(struct {
			// Массив идентификаторов бэкэнд-узлов.
			BackendNodeIds []*dom.BackendNodeId `json:"backendNodeIds"`
		}{
			BackendNodeIds: backendNodeIds,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
	if err != nil {
		return nil, err
	}
	n := struct {
		// Массив идентификаторов добавленных узлов, соответствующих
		// идентификаторам бэкэнда, указанным в backendNodeIds.
		NodeIds []*dom.NodeId `json:"nodeIds"`
	}{}
	model.ForceJSONUnmarshal(r, &n)
	return n.NodeIds, nil
}

// Redo Повторяет последнее отмененное действие.
func (d *Dom) Redo() error {
	return d.client.Exec(
		dom.REDO,
		nil,
		model.WithSessionId(d.CurrentSessionId),
	)
}

// SetInspectedNode Позволяет консоли обращаться к узлу с заданным идентификатором
// через $x (подробнее о функциях $x см. в разделе «API командной строки»).
func (d *Dom) SetInspectedNode(nodeId *dom.NodeId) error {
	return d.client.Exec(
		dom.SET_INSPECTED_NODE,
		model.ForceJSONMarshal(struct {
			// Идентификатор узла DOM должен быть доступен через API командной строки $x.
			NodeId *dom.NodeId `json:"nodeId"`
		}{
			NodeId: nodeId,
		}),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// SetNodeStackTracesEnabled Определяет, следует ли захватывать трассировку стека для узлов.
func (d *Dom) SetNodeStackTracesEnabled(enable bool) error {
	return d.client.Exec(
		dom.SET_NODE_STACK_TRACES_ENABLED,
		model.ForceJSONMarshal(struct {
			// Включить или отключить.
			Enable bool `json:"enable"`
		}{}),
		model.WithSessionId(d.CurrentSessionId),
	)
}

// Undo Отменяет последнее выполненное действие.
func (d *Dom) Undo() error {
	return d.client.Exec(
		dom.UNDO,
		nil,
		model.WithSessionId(d.CurrentSessionId),
	)
}
