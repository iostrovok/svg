package svg

import (
	"github.com/iostrovok/svg/style"
)

type ELLIPSE struct {
	*node
}

const (
	ellipseTag    = `<ellipse`
	ellipseEndTag = `</ellipse>`
)

// Constructor
func Ellipse(cx, cy, rx, ry float64, s ...style.STYLE) *ELLIPSE {
	ellipse := &ELLIPSE{
		node: newNode(s...),
	}

	ellipse.CX(cx)
	ellipse.CY(cy)
	ellipse.RX(rx)
	ellipse.RY(ry)

	return ellipse
}

// Source() returns svg implementation of RECT element
func (ellipse *ELLIPSE) Source() string {
	return _Source(ellipse, ellipseTag, ellipseEndTag, ellipse.inner)
}
