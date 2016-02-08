package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
)

type RECT struct {
	//iNode
	*node

	id string

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

// Source() returns svg implementation of RECT element
func (rect *RECT) Source() string {
	body := fmt.Sprintf(rectTag, rect.x, rect.y, rect.width, rect.height, rect.st.Source())
	return _Source(rect, body, rectEndTag, rect.inner)

}
