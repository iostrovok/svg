package svg

import (
	"fmt"
	"strings"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
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
	formatPoint         = "%d,%d"
	formatPartOneX      = "%d"
	formatEllipticalArc = "%d,%d %d %d,%d %d,%d"
	format3Point        = "%d,%d %d,%d %d,%d"
	formatPartStart     = "%s%s"

	pathTag    = `<path d="%s" %s`
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
	xy  []int
}

type PATH struct {
	iNode
	*node
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

// AppendTo is interface function
func (path *PATH) AppendTo(n iNode) *PATH {
	n.appendIn(path)
	return path
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (path *PATH) Append(nodes ...iNode) *PATH {
	path.node.Append(nodes...)
	return path
}

// Style sets the "style.STYLE" object
func (path *PATH) Style(st style.STYLE) *PATH {
	path.node.Style(st)
	return path
}

// Transform sets the "transform.TRANSFORM" object
func (path *PATH) Transform(tr transform.TRANSFORM) *PATH {
	path.node.Transform(tr)
	return path
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

	bodyLine := fmt.Sprintf(pathTag, strings.Join(body, " "), path.node.mSource())
	return _Source(bodyLine, pathEndTag, path.inner)
}

func (p *pathPart) drawPart() string {
	switch p.Typ {
	// M = moveto
	case "M", "m", "L", "l", "T", "t":
		return fmt.Sprintf(formatPoint, p.xy[0], p.xy[1])
	case "V", "v", "H", "h":
		return fmt.Sprintf(formatPartOneX, p.xy[0])
	case "Q", "q", "S", "s":
		return fmt.Sprintf(formatPoint+" "+formatPoint, p.xy[0], p.xy[1], p.xy[2], p.xy[3])
	case "C", "c":
		return fmt.Sprintf(format3Point, p.xy[0], p.xy[1], p.xy[2], p.xy[3], p.xy[4], p.xy[5])
	case "A", "a":
		return fmt.Sprintf(formatEllipticalArc, p.xy[0], p.xy[1], p.xy[2], p.xy[3], p.xy[4], p.xy[5], p.xy[6])
	}
	return ""
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
func (path *PATH) AddPart(t string, x ...int) *PATH {

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
func (path *PATH) M(x, y int) *PATH {
	return path.AddPart("M", x, y)
}

// m (relative) moveto
func (path *PATH) Mm(x, y int) *PATH {
	return path.AddPart("m", x, y)
}

// L (absolute) lineto
func (path *PATH) L(x, y int) *PATH {
	return path.AddPart("L", x, y)
}

// l (relative) lineto
func (path *PATH) Ll(x, y int) *PATH {
	return path.AddPart("l", x, y)
}

// H (absolute) horizontal linetov
func (path *PATH) H(x int) *PATH {
	return path.AddPart("H", x)
}

// h (relative) horizontal linetov
func (path *PATH) Hh(x int) *PATH {
	return path.AddPart("h", x)
}

// V (absolute) vertical lineto
func (path *PATH) V(y int) *PATH {
	return path.AddPart("V", y)
}

// v (relative) vertical lineto
func (path *PATH) Vv(y int) *PATH {
	return path.AddPart("v", y)
}

// Q  (absolute)  quadratic Bézier curve
func (path *PATH) Q(x1, y1, x2, y2 int) *PATH {
	return path.AddPart("Q", x1, y1, x2, y2)
}

// q  (relative) quadratic Bézier curve
func (path *PATH) Qq(x1, y1, x2, y2 int) *PATH {
	return path.AddPart("q", x1, y1, x2, y2)
}

// T (absolute)  smooth quadratic Bézier curveto
func (path *PATH) T(x1, y1 int) *PATH {
	return path.AddPart("T", x1, y1)
}

// t (relative)  smooth quadratic Bézier curveto
func (path *PATH) Tt(x1, y1 int) *PATH {
	return path.AddPart("t", x1, y1)
}

// A (absolute) elliptical Arc
// a25,100 -30 1,0 50,-25
func (path *PATH) A(rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y int) *PATH {
	return path.AddPart("A", rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

// a (relative) elliptical Arc
func (path *PATH) Aa(rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y int) *PATH {
	return path.AddPart("a", rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

// C (absolute) curveto
func (path *PATH) C(x1, y1, x2, y2, x, y int) *PATH {
	return path.AddPart("C", x1, y1, x2, y2, x, y)
}

// c (relative) curveto
func (path *PATH) Cc(x1, y1, x2, y2, x, y int) *PATH {
	return path.AddPart("c", x1, y1, x2, y2, x, y)
}

// S (absolute) shorthand/smooth curveto
func (path *PATH) S(x2, y2, x, y int) *PATH {
	return path.AddPart("S", x2, y2, x, y)
}

// s (relative) shorthand/smooth curveto
func (path *PATH) Ss(x2, y2, x, y int) *PATH {
	return path.AddPart("s", x2, y2, x, y)
}
