package svg

import (
	"github.com/iostrovok/svg/style"
)

type CIRCLE struct {
	*node
}

const (
	circleTag    = `<circle`
	circleEndTag = `</circle>`
)

// Constructor
func Circle(cx, cy, r float64, s ...style.STYLE) *CIRCLE {
	circle := &CIRCLE{
		node: newNode(s...),
	}

	circle.CX(cx)
	circle.CY(cy)
	circle.R(r)

	return circle
}

// Source() returns svg implementation of RECT element
func (circle *CIRCLE) Source() string {
	return _Source(circle, circleTag, circleEndTag, circle.inner)
}
