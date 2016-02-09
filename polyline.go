package svg

import (
	"fmt"
	"strings"

	"github.com/iostrovok/svg/style"
)

type POLYLINE struct {
	*node
	points []float64
}

const (
	polylineTag    = `<polyline points="%s"`
	polylineEndTag = `</polyline>`
)

// Constructor
func Polyline(s ...style.STYLE) *POLYLINE {
	polyline := &POLYLINE{
		node:   newNode(s...),
		points: []float64{},
	}

	return polyline
}

// Source() returns svg implementation of POLYLINE element
func (pl *POLYLINE) Points(p ...float64) *POLYLINE {

	if len(p) > 0 && len(p)%2 == 0 {
		pl.points = append(pl.points, p...)
	}

	return pl
}

func printNumber(n float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", n), "0"), ".")
}

// Source() returns svg implementation of POLYLINE element
func (pl *POLYLINE) Source() string {

	coods := []string{}
	for i := 0; i < len(pl.points)-1; i += 2 {
		coods = append(coods, printNumber(pl.points[i])+","+printNumber(pl.points[i+1]))
	}

	body := fmt.Sprintf(polylineTag, strings.Join(coods, " "))
	return _Source(pl, body, polylineEndTag, pl.node.inner)
}
