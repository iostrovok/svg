package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
)

type LINE struct {
	INode
	*Node
	x1, x2 int
	y1, y2 int
}

const (
	lineTag    = `<line x1="%d" y1="%d" x2="%d" y2="%d" %s `
	lineEndTag = `</line>`
)

// Constructor
func Line(x1, y1, x2, y2 int, s ...style.STYLE) *LINE {
	return &LINE{
		Node: NewNode(s...),
		x1:   x1, x2: x2,
		y1: y1, y2: y2,
	}
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (line *LINE) Append(nodes ...INode) *LINE {
	line.Node.Append(nodes...)
	return line
}

func (line *LINE) AppendTo(n INode) *LINE {
	n.AppendIn(line)
	return line
}

// Setter
func (line *LINE) Style(st style.STYLE) *LINE {
	line.Node.Style(st)
	return line
}

// Source() returns svg implementation of LINE element
func (line *LINE) Source() string {
	body := fmt.Sprintf(lineTag, line.x1, line.y1, line.x2, line.y2, line.Node.st.Source())
	return _Source(body, lineEndTag, line.Node.inner)
}
