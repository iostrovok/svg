package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
)

type RECT struct {
	INode
	*Node
	x, y   int
	width  int
	height int
}

const (
	rectTag    = `<rect x="%d" y="%d" width="%d" height="%d" %s`
	rectEndTag = `</rect>`
)

// Constructor
func Rect(x, y, width, height int, s ...style.STYLE) *RECT {
	return &RECT{
		Node:   NewNode(s...),
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

func (rect *RECT) AppendTo(n INode) *RECT {
	n.AppendIn(rect)
	return rect
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (rect *RECT) Append(nodes ...INode) *RECT {
	rect.Node.Append(nodes...)
	return rect
}

// Setter
func (rect *RECT) Style(st style.STYLE) *RECT {
	rect.Node.Style(st)
	return rect
}

// Source() returns svg implementation of RECT element
func (rect *RECT) Source() string {
	body := fmt.Sprintf(rectTag, rect.x, rect.y, rect.width, rect.height, rect.st.Source())
	return _Source(body, rectEndTag, rect.inner)

}
