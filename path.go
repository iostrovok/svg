package svg

import (
	"fmt"
	"strings"

	"github.com/iostrovok/svg/style"
)

/*
	https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/d

	"A", "a":
	The declaration of an arcto is relatively complicated: "A rx,ry xAxisRotate LargeArcFlag,SweepFlag x,y".
	To deconstruct, rx and ry are the radius in the x and y directions respectively,
	The LargeArcFlag has a value of either 0 or 1, and determines whether the smallest (0) or largest (1) arc possible is drawn. T
	he SweepFlag is either 0 or 1, and determines if the arc should be swept in a clockwise (1) or anti-clockwise (0) direction.
	x and y are the destination coordinates.
*/

const (
	formatPoint         = "%s,%s"
	formatPartOneX      = "%s"
	formatEllipticalArc = "%s,%s %s %s,%s %s,%s"
	format3Point        = "%s,%s %s,%s %s,%s"
	formatPartStart     = "%s%s"

	pathTag    = `<path d="%s"`
	pathEndTag = `</path>`
)

type pathPart struct {
	/*
	   M, m = moveto
	   L, l = lineto
	   H, h = horizontal lineto
	   V, v = vertical lineto
	   C, c = curveto
	   S, s = smooth curveto
	   Q, q = quadratic Bézier curve
	   T, t = smooth quadratic Bézier curveto
	   A, a = elliptical Arc
	   Z, z = closepath
	*/
	Typ string // M, m, L, l, H, h, V, v, C, c, S, s, Q, q, T, t, A, a
	xy  []float64
}

type PATH struct {
	//iNode
	*node

	id string

	parts []*pathPart
}

// Constructor
func Path(s ...style.STYLE) *PATH {
	p := &PATH{
		node:  newNode(s...),
		parts: []*pathPart{},
	}
	return p
}

// Source() returns svg implementation of *PATH element
func (path *PATH) Source() string {

	body := []string{}
	var last *pathPart
	for _, p := range path.parts {
		leadSymbol := false

		if p.Typ == "Z" || p.Typ == "z" {
			body = append(body, p.Typ)
			last = p
			continue
		}

		s := p.drawPart()
		if p.Typ == "V" || p.Typ == "v" || p.Typ == "H" || p.Typ == "h" {
			leadSymbol = true
		} else if (last != nil && last.Typ != p.Typ) || (last == nil) {
			leadSymbol = true
		}

		if leadSymbol {
			s = fmt.Sprintf(formatPartStart, p.Typ, s)
		}

		body = append(body, s)
		last = p
	}

	bodyLine := fmt.Sprintf(pathTag, strings.Join(body, " "))
	return _Source(path, bodyLine, pathEndTag, path.inner)
}

func (p *pathPart) drawPart() string {

	strs := []interface{}{}
	for _, x := range p.xy {
		strs = append(strs, printNumber(x))
	}

	tmpl := ""

	switch p.Typ {
	// M = moveto
	case "M", "m", "L", "l", "T", "t":
		tmpl = formatPoint
	case "V", "v", "H", "h":
		tmpl = formatPartOneX
	case "Q", "q", "S", "s":
		tmpl = formatPoint + " " + formatPoint
	case "C", "c":
		tmpl = format3Point
	case "A", "a":
		tmpl = formatEllipticalArc
	default:
		return ""
	}

	return fmt.Sprintf(tmpl, strs...)
}

// Z = closepath
func (path *PATH) Z() *PATH {
	path.parts = append(path.parts, &pathPart{
		Typ: "Z",
	})
	return path
}

func (path *PATH) z() *PATH {
	path.parts = append(path.parts, &pathPart{
		Typ: "z",
	})
	return path
}

// AddPart
func (path *PATH) AddPart(t string, x ...float64) *PATH {

	// Checks true creating path
	if len(path.parts) == 0 && t != "M" && t != "m" {
		panic(`For *PATH the first element must be "M" or "m". See https://www.w3.org/TR/SVG/paths.html`)
	}

	path.parts = append(path.parts, &pathPart{
		Typ: t,
		xy:  x,
	})
	return path
}

// M (absolute) moveto
func (path *PATH) M(x, y float64) *PATH {
	return path.AddPart("M", x, y)
}

// m (relative) moveto
func (path *PATH) Mm(x, y float64) *PATH {
	return path.AddPart("m", x, y)
}

// L (absolute) lineto
func (path *PATH) L(x, y float64) *PATH {
	return path.AddPart("L", x, y)
}

// l (relative) lineto
func (path *PATH) Ll(x, y float64) *PATH {
	return path.AddPart("l", x, y)
}

// H (absolute) horizontal linetov
func (path *PATH) H(x float64) *PATH {
	return path.AddPart("H", x)
}

// h (relative) horizontal linetov
func (path *PATH) Hh(x float64) *PATH {
	return path.AddPart("h", x)
}

// V (absolute) vertical lineto
func (path *PATH) V(y float64) *PATH {
	return path.AddPart("V", y)
}

// v (relative) vertical lineto
func (path *PATH) Vv(y float64) *PATH {
	return path.AddPart("v", y)
}

// Q  (absolute)  quadratic Bézier curve
func (path *PATH) Q(x1, y1, x2, y2 float64) *PATH {
	return path.AddPart("Q", x1, y1, x2, y2)
}

// q  (relative) quadratic Bézier curve
func (path *PATH) Qq(x1, y1, x2, y2 float64) *PATH {
	return path.AddPart("q", x1, y1, x2, y2)
}

// T (absolute)  smooth quadratic Bézier curveto
func (path *PATH) T(x1, y1 float64) *PATH {
	return path.AddPart("T", x1, y1)
}

// t (relative)  smooth quadratic Bézier curveto
func (path *PATH) Tt(x1, y1 float64) *PATH {
	return path.AddPart("t", x1, y1)
}

// A (absolute) elliptical Arc
// a25,100 -30 1,0 50,-25
func (path *PATH) A(rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y float64) *PATH {
	return path.AddPart("A", rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

// a (relative) elliptical Arc
func (path *PATH) Aa(rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y float64) *PATH {
	return path.AddPart("a", rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

// C (absolute) curveto
func (path *PATH) C(x1, y1, x2, y2, x, y float64) *PATH {
	return path.AddPart("C", x1, y1, x2, y2, x, y)
}

// c (relative) curveto
func (path *PATH) Cc(x1, y1, x2, y2, x, y float64) *PATH {
	return path.AddPart("c", x1, y1, x2, y2, x, y)
}

// S (absolute) shorthand/smooth curveto
func (path *PATH) S(x2, y2, x, y float64) *PATH {
	return path.AddPart("S", x2, y2, x, y)
}

// s (relative) shorthand/smooth curveto
func (path *PATH) Ss(x2, y2, x, y float64) *PATH {
	return path.AddPart("s", x2, y2, x, y)
}
