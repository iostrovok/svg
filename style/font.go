package style

import (
	"fmt"
	"log"
	"strings"
)

const (
	lineFont    = "%s:%s"
	lineFontInt = "%d"
)

var (
	font_weights map[string]bool = map[string]bool{
		`100`: true, `200`: true, `300`: true,
		`400`: true, `500`: true, `600`: true,
		`700`: true, `800`: true, `900`: true,
		`bold`: true, `bolder`: true,
		`lighter`: true, `normal`: true,
	}

	font_style map[string]bool = map[string]bool{
		`normal`: true, `italic`: true, `oblique`: true,
	}
	font_variant map[string]bool = map[string]bool{
		`normal`: true, `small-caps`: true,
	}
	font_stretch map[string]bool = map[string]bool{
		`normal`: true, `wider`: true, `narrower`: true,
		`ultra-condensed`: true, `extra-condensed`: true,
		`condensed`: true, `extra-expanded`: true,
		`expanded`: true, `ultra-expanded`: true,
	}
)

type FONT struct {
	ONESTYLE
	body             string
	size, family     string
	weight, style    string
	variant, stretch string
}

// Constructor
func Font(body ...string) FONT {
	out := FONT{}
	if len(body) > 0 {
		if len(body[0]) > 0 {
			out.body = body[0]
		}
	}
	return out
}

// Weight - setter
func (font FONT) Weight(weight string) FONT {
	f := strings.ToLower(weight)
	if font_weights[f] {
		font.weight = f
	} else {
		log.Printf("FONT.Weight: %s is not valid", weight)
	}
	return font
}

// Style - setter
func (font FONT) Stretch(stretch string) FONT {
	s := strings.ToLower(stretch)
	if font_stretch[s] {
		font.stretch = s
	} else {
		log.Printf("FONT.Stretch: %s is not valid", stretch)
	}
	return font
}

// Style - setter
func (font FONT) Variant(variant string) FONT {
	s := strings.ToLower(variant)
	if font_variant[s] {
		font.variant = s
	} else {
		log.Printf("FONT.Variant: %s is not valid", variant)
	}
	return font
}

// Size - setter
func (font FONT) Size(size string) FONT {
	font.size = size
	return font
}

// Size - setter
func (font FONT) Family(family string) FONT {
	font.family = family
	return font
}

// Source() returns svg implementation of FONT style
func (font FONT) Source() string {
	out := []string{}
	if font.body != "" {
		return font.body
	}
	if font.family != "" {
		out = append(out, fmt.Sprintf(lineFont, "font-family", font.family))
	}
	if font.size != "" {
		out = append(out, fmt.Sprintf(lineFont, "font-size", font.size))
	}
	if font.weight != "" {
		out = append(out, fmt.Sprintf(lineFont, "font-weight", font.weight))
	}
	if font.style != "" {
		out = append(out, fmt.Sprintf(lineFont, "font-style", font.style))
	}
	if font.stretch != "" {
		out = append(out, fmt.Sprintf(lineFont, "font-stretch", font.stretch))
	}
	if font.variant != "" {
		out = append(out, fmt.Sprintf(lineFont, "font-variant", font.variant))
	}

	return strings.Join(out, ";")
}
