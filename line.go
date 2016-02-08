package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
)

type LINE struct {
	//iNode
	*node

	id string

	x1, x2 int
	y1, y2 int
}

const (
	lineTag    = `<line x1="%d" y1="%d" x2="%d" y2="%d"`
	lineEndTag = `</line>`
)

// Constructor of "line" object
func Line(x1, y1, x2, y2 int, s ...style.STYLE) *LINE {
	return &LINE{
		node: newNode(s...),

		x1: x1, x2: x2,
		y1: y1, y2: y2,
	}
}

// Source() returns svg implementation of LINE element
func (line *LINE) Source() string {
	body := fmt.Sprintf(lineTag, line.x1, line.y1, line.x2, line.y2)
	return _Source(line, body, lineEndTag, line.node.inner)
}
