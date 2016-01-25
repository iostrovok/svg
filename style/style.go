package style

import (
	"strings"
)

const (
	Eq  = `:`
	Del = `;`
)

type ONESTYLE interface {
	Source() string
}

type STYLE struct {
	tags map[string]ONESTYLE
}

// New is the SVG constructor, specifying the io.Writer where the generated SVG is written.
func Style() STYLE {
	return STYLE{
		tags: map[string]ONESTYLE{},
	}
}

// Source() returns svg implementation of STYLE element
func (style STYLE) Source() string {

	if len(style.tags) == 0 {
		return ""
	}

	out := []string{}
	for k, v := range style.tags {
		// out = append(out, k+Eq+url.QueryEscape(v.Source()))
		out = append(out, k+Eq+v.Source())
	}
	return `style="` + strings.Join(out, Del) + `"`
}

func (style STYLE) StrokeWidth(width float64) STYLE {
	style.tags[`stroke-width`] = Width(width)
	return style
}

func (style STYLE) Stroke(r, g, b int) STYLE {
	style.tags[`stroke`] = Color(r, g, b)
	return style
}

func (style STYLE) StrokeStr(c string) STYLE {
	style.tags[`stroke`] = ColorStr(c)
	return style
}

func (style STYLE) FillRGB(r, g, b int) STYLE {
	style.tags[`fill`] = Color(r, g, b)
	return style
}

func (style STYLE) Fill(c string) STYLE {
	style.tags[`fill`] = ColorStr(c)
	return style
}
