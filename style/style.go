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
	font *ONESTYLE
	tags map[string]ONESTYLE
}

// Style is the constructor
func Style() STYLE {
	return STYLE{
		tags: map[string]ONESTYLE{},
	}
}

// Ref() returns reference to style element
func (style STYLE) Ref() *STYLE {
	return &style
}

// Source() returns svg implementation of STYLE element
func (style STYLE) Source() string {

	if len(style.tags) == 0 && style.font == nil {
		return ""
	}

	out := []string{}
	if style.font != nil {
		f := *style.font
		out = append(out, f.Source())
	}

	for k, v := range style.tags {
		// out = append(out, k+Eq+url.QueryEscape(v.Source()))
		out = append(out, k+Eq+v.Source())
	}
	return `style="` + strings.Join(out, Del) + `"`
}

// Constructor
func (style STYLE) Font(font ONESTYLE) STYLE {
	style.font = &font
	return style
}

func (style STYLE) StrokeWidth(width float64) STYLE {
	style.tags[`stroke-width`] = Width(width)
	return style
}

func (style STYLE) StrokeRGB(r, g, b int) STYLE {
	style.tags[`stroke`] = RGB(r, g, b)
	return style
}

func (style STYLE) Stroke(c string) STYLE {
	style.tags[`stroke`] = Color(c)
	return style
}

func (style STYLE) FillRGB(r, g, b int) STYLE {
	style.tags[`fill`] = RGB(r, g, b)
	return style
}

func (style STYLE) Fill(c string) STYLE {
	style.tags[`fill`] = Color(c)
	return style
}
