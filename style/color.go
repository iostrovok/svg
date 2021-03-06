package style

import (
	"fmt"
	"strings"
)

const (
	lineColor = "rgb(%d,%d,%d)"
)

type COLOR struct {
	ONESTYLE
	r, g, b int
	line    string
}

// RGB is the color constructor which gets red, green and blue color componetns as integer.
func RGB(r, g, b int) COLOR {
	color := COLOR{
		r: r,
		g: g,
		b: b,
	}
	return color.RGB(r, g, b)
}

func (color COLOR) RGB(r, g, b int) COLOR {
	color.line = fmt.Sprintf(lineColor, r, g, b)
	return color
}

// RGB is the color constructor which gets color as string.
func Color(c string) COLOR {
	return COLOR{
		line: strings.ToLower(c),
	}
}

func (color COLOR) Source() string {
	return color.line
}
