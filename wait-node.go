package chrome

import (
	"fmt"
	"time"

	"github.com/skvdmt/chrome/internal/devtools/options/dom/describe_node"
	"github.com/skvdmt/chrome/internal/devtools/types/dom"
	"github.com/skvdmt/chrome/internal/model"
	"github.com/skvdmt/chrome/internal/options/wait_node"
)

// WaitNodeText Ожидание узел по селектору и получает его полнное текстовое содержимое.
func (d *Driver) WaitNodeText(selector string, options ...wait_node.Option) (string, error) {
	n, err := d.WaitNode(selector, options...)
	if err != nil {
		return "", err
	}
	return d.NodeText(n), nil
}

// WaitNode Ожидание узла.
func (d *Driver) WaitNode(selector string, options ...wait_node.Option) (*dom.Node, error) {
	c := wait_node.NewConfig()
	s := time.Now()
	for {
		if s.Add(c.Timeout).UnixNano() < time.Now().UnixNano() {
			return nil, model.ERR_RESPONSE_TIMEOUT
		}
		d.debug.Debug(fmt.Sprintf("try get node by %s selector", selector))

		id, err := d.Dom.QuerySelector(selector, c.QuerySelectorOptions...)
		if err != nil {
			time.Sleep(c.Interval)
			continue
		}
		n, err := d.Dom.DescribeNode(
			describe_node.WithNodeId(id),
			describe_node.WithDepth(-1),
		)
		if err != nil {
			time.Sleep(c.Interval)
			continue
		}
		d.debug.Debug(fmt.Sprintf("node geted by %s selector", selector))

		return n, nil
	}
}
