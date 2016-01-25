package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
)

type LINE struct {
	Node
	x1, x2 int
	y1, y2 int
	st     style.STYLE
	inner  []Node
}

const (
	lineTag    = `<line x1="%d" y1="%d" x2="%d" y2="%d" %s `
	lineEndTag = `</line>`
)

// Constructor
func Line(x1, y1, x2, y2 int, s ...style.STYLE) LINE {
	return LINE{
		x1: x1, x2: x2,
		y1: y1, y2: y2,
		st:    mstyle(s...),
		inner: []Node{},
	}
}

// Setter
func (line LINE) Style(s style.STYLE) Node {
	line.st = s
	return line
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (line LINE) Append(n ...Node) LINE {
	if len(n) > 0 {
		line.inner = append(line.inner, n...)
	}
	return line
}

// Inner() returns inner elements of LINE
func (line LINE) Inner() []Node {
	return line.inner
}

// Source() returns svg implementation of LINE element
func (line LINE) Source() string {
	body := fmt.Sprintf(lineTag, line.x1, line.y1, line.x2, line.y2, line.st.Source())
	return _Source(body, lineEndTag, line.inner)
}
