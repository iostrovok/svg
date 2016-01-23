// Package svg provides an API for generating Scalable Vector Graphics (SVG)
package svg

import (
	"bytes"
	"fmt"
	"io"

	"github.com/iostrovok/svg/style"
)

type Node interface {
	Inner() []Node
	Source() string
	Style(style.STYLE) Node
}

// SVG defines the location of the generated SVG
type SVG struct {
	Node
	w, h  int
	body  string
	inner []Node
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
		w:     width,
		h:     height,
		body:  fmt.Sprintf(svgTop, width, height),
		inner: []Node{},
	}
}

func (svg *SVG) Width() int {
	return svg.w
}

func (svg *SVG) Height() int {
	return svg.h
}

func (svg *SVG) Inner() []Node {
	return svg.inner
}

func (svg *SVG) Append(n Node) *SVG {
	svg.inner = append(svg.inner, n)
	return svg
}

func (svg *SVG) Source() string {
	return _Source(svg.body, svgEnd, svg.inner)
}

func _Source(body, tagEnd string, inner []Node) string {

	if len(inner) == 0 {
		return body + emptyclose
	}

	out := body + ">\n"
	for _, n := range inner {
		out += n.Source() + br
	}

	return out + tagEnd + br
}

func mstyle(s ...style.STYLE) style.STYLE {
	if len(s) == 0 {
		return style.STYLE{}
	}
	return s[0]
}

func (svg *SVG) SaveToFile(Writer io.Writer) error {

	src := []byte(svg.Source())
	_, err := io.Copy(Writer, bytes.NewReader(src))
	return err
}

func (svg *SVG) Rect(x1, y1, width, height int, style ...style.STYLE) *SVG {
	return svg.Append(Rect(x1, y1, width, height, style...))
}

func (svg *SVG) Line(x1, y1, x2, y2 int, style ...style.STYLE) *SVG {
	n := Line(x1, y1, x2, y2, style...)
	return svg.Append(n)
}
