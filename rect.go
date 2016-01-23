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

func (rect RECT) Style(s style.STYLE) Node {
	rect.st = s
	return rect
}

func (rect RECT) Append(n ...Node) RECT {
	if len(n) > 0 {
		rect.inner = append(rect.inner, n...)
	}
	return rect
}

func (rect RECT) Inner() []Node {
	return rect.inner
}

func (rect RECT) Source() string {
	body := fmt.Sprintf(rectTag, rect.x, rect.y, rect.width, rect.height, rect.st.Source())
	return _Source(body, rectEndTag, rect.inner)
}
