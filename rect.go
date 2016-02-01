package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
)

type RECT struct {
	iNode
	*node
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
		node:   newNode(s...),
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

// AppendTo is interface function
func (rect *RECT) AppendTo(n iNode) *RECT {
	n.AppendIn(rect)
	return rect
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (rect *RECT) Append(nodes ...iNode) *RECT {
	rect.node.Append(nodes...)
	return rect
}

// Style sets the "style.STYLE" object
func (rect *RECT) Style(st style.STYLE) *RECT {
	rect.node.Style(st)
	return rect
}

// Source() returns svg implementation of RECT element
func (rect *RECT) Source() string {
	body := fmt.Sprintf(rectTag, rect.x, rect.y, rect.width, rect.height, rect.st.Source())
	return _Source(body, rectEndTag, rect.inner)

}
