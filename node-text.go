package chrome

import (
	"regexp"
	"strings"

	"github.com/skvdmt/chrome/internal/devtools/types/dom"
)

// NodeText Тект узла.
func (d *Driver) NodeText(node *dom.Node) string {
	b := strings.Builder{}
	if node.NodeName == "#text" {
		b.WriteString(node.NodeValue)
		b.WriteRune(' ')
		return b.String()
	}
	for _, c := range node.Children {
		b.WriteString(d.NodeText(c))
		b.WriteRune(' ')
	}
	return strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(b.String(), " "))
}
