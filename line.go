package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

type LINE struct {
	iNode
	*node
	x1, x2 int
	y1, y2 int
}

const (
	lineTag    = `<line x1="%d" y1="%d" x2="%d" y2="%d" %s `
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

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (line *LINE) Append(nodes ...iNode) *LINE {
	line.node.Append(nodes...)
	return line
}

// AppendTo is interface function
func (line *LINE) AppendTo(n iNode) *LINE {
	n.appendIn(line)
	return line
}

// Style sets the "style.STYLE" object
func (line *LINE) Style(st style.STYLE) *LINE {
	line.node.Style(st)
	return line
}

// Transform sets the "transform.TRANSFORM" object
func (line *LINE) Transform(tr transform.TRANSFORM) *LINE {
	line.node.Transform(tr)
	return line
}

// Source() returns svg implementation of LINE element
func (line *LINE) Source() string {
	body := fmt.Sprintf(lineTag, line.x1, line.y1, line.x2, line.y2, line.node.mSource())
	return _Source(body, lineEndTag, line.node.inner)
}
