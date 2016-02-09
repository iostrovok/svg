package svg

import (
	"fmt"
	"strings"

	"github.com/iostrovok/svg/style"
)

type POLYGON struct {
	*node
	points []float64
}

const (
	polygonTag    = `<polygon points="%s"`
	polygonEndTag = `</polygon>`
)

// Constructor
func Polygon(s ...style.STYLE) *POLYGON {
	polygon := &POLYGON{
		node:   newNode(s...),
		points: []float64{},
	}

	return polygon
}

// Source() returns svg implementation of POLYGON element
func (pl *POLYGON) Points(p ...float64) *POLYGON {

	if len(p) > 0 && len(p)%2 == 0 {
		pl.points = append(pl.points, p...)
	}

	return pl
}

// Source() returns svg implementation of POLYGON element
func (pl *POLYGON) Source() string {

	coods := []string{}
	for i := 0; i < len(pl.points)-1; i += 2 {
		coods = append(coods, printNumber(pl.points[i])+","+printNumber(pl.points[i+1]))
	}

	body := fmt.Sprintf(polygonTag, strings.Join(coods, " "))
	return _Source(pl, body, polygonEndTag, pl.node.inner)
}
