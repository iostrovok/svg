package svg

import (
	"fmt"
	"strings"

	"github.com/iostrovok/svg/style"
)

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
	INode
	*Node
	parts   []*pathPart
	isClose bool
}

// Constructor
func Path(s ...style.STYLE) *PATH {
	p := &PATH{
		Node:  NewNode(s...),
		parts: []*pathPart{},
	}
	return p
}

func (path *PATH) AppendTo(n INode) *PATH {
	n.AppendIn(path)
	return path
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (path *PATH) Append(nodes ...INode) *PATH {
	path.Node.Append(nodes...)
	return path
}

// Setter
func (path *PATH) Style(st style.STYLE) *PATH {
	path.Node.Style(st)
	return path
}

// Source() returns svg implementation of *PATH element
func (path *PATH) Source() string {

	body := []string{}
	var last *pathPart
	for _, p := range path.parts {

		s := p.drawPart()

		leadSymbol := false
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

	if path.isClose {
		body = append(body, "Z")
	}

	bodyLine := fmt.Sprintf(pathTag, strings.Join(body, " "), path.st.Source())
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
func (path *PATH) Z(is ...bool) *PATH {
	if len(is) > 0 {
		path.isClose = is[0]
	} else {
		path.isClose = true
	}
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
func (path *PATH) m(x, y int) *PATH {
	return path.AddPart("m", x, y)
}

// L (absolute) lineto
func (path *PATH) L(x, y int) *PATH {
	return path.AddPart("L", x, y)
}

// l (relative) lineto
func (path *PATH) l(x, y int) *PATH {
	return path.AddPart("l", x, y)
}

// H (absolute) horizontal linetov
func (path *PATH) H(x int) *PATH {
	return path.AddPart("H", x)
}

// h (relative) horizontal linetov
func (path *PATH) h(x int) *PATH {
	return path.AddPart("h", x)
}

// V (absolute) vertical lineto
func (path *PATH) V(y int) *PATH {
	return path.AddPart("V", y)
}

// v (relative) vertical lineto
func (path *PATH) v(y int) *PATH {
	return path.AddPart("v", y)
}

// Q  (absolute)  quadratic Bézier curve
func (path *PATH) Q(x1, y1, x2, y2 int) *PATH {
	return path.AddPart("Q", x1, y1, x2, y2)
}

// q  (relative) quadratic Bézier curve
func (path *PATH) q(x1, y1, x2, y2 int) *PATH {
	return path.AddPart("q", x1, y1, x2, y2)
}

// T (absolute)  smooth quadratic Bézier curveto
func (path *PATH) T(x1, y1 int) *PATH {
	return path.AddPart("T", x1, y1)
}

// t (relative)  smooth quadratic Bézier curveto
func (path *PATH) t(x1, y1 int) *PATH {
	return path.AddPart("t", x1, y1)
}

// A (absolute) elliptical Arc
// a25,100 -30 1,0 50,-25
func (path *PATH) A(rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y int) *PATH {
	return path.AddPart("A", rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

// a (relative) elliptical Arc
func (path *PATH) a(rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y int) *PATH {
	return path.AddPart("a", rx, ry, xAxisRotation, largeArcFlag, sweepFlag, x, y)
}

// C (absolute) curveto
func (path *PATH) C(x1, y1, x2, y2, x, y int) *PATH {
	return path.AddPart("C", x1, y1, x2, y2, x, y)
}

// c (relative) curveto
func (path *PATH) c(x1, y1, x2, y2, x, y int) *PATH {
	return path.AddPart("c", x1, y1, x2, y2, x, y)
}

// S (absolute) shorthand/smooth curveto
func (path *PATH) S(x2, y2, x, y int) *PATH {
	return path.AddPart("S", x2, y2, x, y)
}

// s (relative) shorthand/smooth curveto
func (path *PATH) s(x2, y2, x, y int) *PATH {
	return path.AddPart("s", x2, y2, x, y)
}
