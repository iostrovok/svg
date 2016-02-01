// Package svg provides an API for generating Scalable Vector Graphics (SVG)
package svg

import (
	"bytes"
	"fmt"
	"io"

	"github.com/iostrovok/svg/style"
)

// SVG defines the location of the generated SVG
type SVG struct {
	iNode
	*node
	w, h int
	dim  string
	body string

	// viewBox
	vx, vy int
	vw, vh int
}

var (
	svgTop = `<?xml version="1.0"?>
<svg
	width="%d%s" height="%d%s" %s
	xmlns="http://www.w3.org/2000/svg"
	xmlns:xlink="http://www.w3.org/1999/xlink"
`
	svgEnd     = `</svg>`
	vbfmt      = `viewBox="%d %d %d %d"`
	emptyclose = "/>\n"
	br         = "\n"
)

/*
	New is the SVG constructor
	example:
	s := New(100, 100, "px")
*/
func New(width, height int, dim ...string) *SVG {
	s := &SVG{
		node: newNode(),
		w:    width,
		h:    height,
	}

	if len(dim) > 0 {
		s.dim = dim[0]
	}

	return s
}

// Width(width ...int) set/get width of SVG element
func (svg *SVG) Width(width ...int) int {
	if len(width) > 0 {
		svg.w = width[0]
	}
	return svg.w
}

// Height(height ...int) set/get height of SVG element
func (svg *SVG) Height(height ...int) int {
	if len(height) > 0 {
		svg.h = height[0]
	}
	return svg.h
}

// Dim(dim ...string) set/get dimension of SVG element
func (svg *SVG) Dim(dim ...string) string {
	if len(dim) > 0 {
		svg.dim = dim[0]
	}
	return svg.dim
}

// ViewBox(x, y, width, height) set the viewBox attribute
func (svg *SVG) ViewBox(x, y, width, height int) *SVG {
	svg.vx = x
	svg.vy = y
	svg.vw = width
	svg.vh = height

	return svg
}

// GetViewBox() get the details of viewBox attribute
func (svg *SVG) GetViewBox() (x, y, width, height int) {
	x = svg.vx
	y = svg.vy
	width = svg.vw
	height = svg.vh
	return
}

// Source() returns svg implementation of SVG element
func (svg *SVG) Source() string {
	vb := ""
	if svg.vx != 0 || svg.vy != 0 || svg.vw != 0 || svg.vh != 0 {
		vb = fmt.Sprintf(vbfmt, svg.vx, svg.vy, svg.vw, svg.vh)
	}
	body := fmt.Sprintf(svgTop, svg.w, svg.dim, svg.h, svg.dim, vb)

	return _Source(body, svgEnd, svg.inner)
}

func _Source(body, tagEnd string, inner []iNode) string {

	if len(inner) == 0 {
		return body + emptyclose
	}

	return body + ">\n" + _innerSource(inner) + tagEnd + br
}

// Source() returns svg implementation of SVG element
func (svg *SVG) InnerSource() string {
	return _innerSource(svg.inner)
}

func _innerSource(inner []iNode) string {

	out := ""
	for _, n := range inner {
		out += n.Source()
	}

	return out
}

// Save() saves content of SVG to io.Writer
func (svg *SVG) Save(Writer io.Writer) error {
	src := []byte(svg.Source())
	_, err := io.Copy(Writer, bytes.NewReader(src))
	return err
}

// // Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (svg *SVG) Append(n iNode) *SVG {
	svg.node.inner = append(svg.node.inner, n)
	return svg
}

func (svg *SVG) AppendIn(n iNode) {
	svg.node.Append(n)
}

// Rect() adds rect element to SVG
func (svg *SVG) Rect(x1, y1, width, height int, st ...style.STYLE) *SVG {
	r := Rect(x1, y1, width, height, st...)
	return svg.Append(r)
}

// Line() adds line element to SVG
func (svg *SVG) Line(x1, y1, x2, y2 int, st ...style.STYLE) *SVG {
	n := Line(x1, y1, x2, y2, st...)
	return svg.Append(n)
}
