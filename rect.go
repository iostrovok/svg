package svg

import (
	"github.com/iostrovok/svg/style"
)

type RECT struct {
	//iNode
	*node

	id string
}

const (
	rectTag    = `<rect`
	rectEndTag = `</rect>`
)

// Constructor
func Rect(x, y, width, height float64, s ...style.STYLE) *RECT {
	rect := &RECT{
		node: newNode(s...),
	}

	rect.X(x)
	rect.Y(y)
	rect.Width(width)
	rect.Height(height)

	return rect
}

// Source() returns svg implementation of RECT element
func (rect *RECT) Source() string {
	return _Source(rect, rectTag, rectEndTag, rect.node.inner)
}
