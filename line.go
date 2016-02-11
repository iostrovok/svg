package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
)

type LINE struct {
	//iNode
	*node

	id string

	x1, x2 float64
	y1, y2 float64
}

const (
	lineTag    = `<line x1="%s" y1="%s" x2="%s" y2="%s"`
	lineEndTag = `</line>`
)

// Constructor of "line" object
func Line(x1, y1, x2, y2 float64, s ...style.STYLE) *LINE {
	return &LINE{
		node: newNode(s...),

		x1: x1, x2: x2,
		y1: y1, y2: y2,
	}
}

// Source() returns svg implementation of LINE element
func (line *LINE) Source() string {
	body := fmt.Sprintf(lineTag, printNumber(line.x1), printNumber(line.y1), printNumber(line.x2), printNumber(line.y2))
	return _Source(line, body, lineEndTag, line.node.inner)
}
