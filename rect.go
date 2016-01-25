package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
)

type RECT struct {
	Node
	x, y   int
	width  int
	height int
	inner  []Node
	st     style.STYLE
}

const (
	rectTag    = `<rect x="%d" y="%d" width="%d" height="%d" %s`
	rectEndTag = `</rect>`
)

func Rect(x, y, width, height int, s ...style.STYLE) RECT {

	return RECT{
		x:      x,
		y:      y,
		width:  width,
		height: height,
		st:     mstyle(s...),
		inner:  []Node{},
	}
}

// Setter
func (rect RECT) Style(s style.STYLE) Node {
	rect.st = s
	return rect
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (rect RECT) Append(n ...Node) RECT {
	if len(n) > 0 {
		rect.inner = append(rect.inner, n...)
	}
	return rect
}

// Inner() returns inner elements of RECT
func (rect RECT) Inner() []Node {
	return rect.inner
}

// Source() returns svg implementation of RECT element
func (rect RECT) Source() string {
	body := fmt.Sprintf(rectTag, rect.x, rect.y, rect.width, rect.height, rect.st.Source())
	return _Source(body, rectEndTag, rect.inner)
}
