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
	INode
	*Node
	w, h int
	body string
}

var (
	svgTop = `<?xml version="1.0"?>
<svg 
	width="%d" height="%d"
	xmlns="http://www.w3.org/2000/svg" 
	xmlns:xlink="http://www.w3.org/1999/xlink"
`
	svgEnd     = `</svg>`
	svginitfmt = `%s width="%d%s" height="%d%s"`
	vbfmt      = `viewBox="%d %d %d %d"`
	emptyclose = "/>\n"
	br         = "\n"
)

// New is the SVG constructor
func New(width, height int) *SVG {
	return &SVG{
		Node: NewNode(),
		w:    width,
		h:    height,
		body: fmt.Sprintf(svgTop, width, height),
	}
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

// Source() returns svg implementation of SVG element
func (svg *SVG) Source() string {
	return _Source(svg.body, svgEnd, svg.inner)
}

func _Source(body, tagEnd string, inner []INode) string {

	if len(inner) == 0 {
		return body + emptyclose
	}

	return body + ">\n" + _innerSource(inner) + tagEnd + br
}

// Source() returns svg implementation of SVG element
func (svg *SVG) InnerSource() string {
	return _innerSource(svg.inner)
}

func _innerSource(inner []INode) string {

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
func (svg *SVG) Append(n INode) *SVG {
	svg.Node.inner = append(svg.Node.inner, n)
	return svg
}

func (svg *SVG) AppendIn(n INode) {
	svg.Node.Append(n)
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
