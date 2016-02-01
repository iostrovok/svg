package transform

import (
	"fmt"
)

/*
	scale(<sx> [<sy>]), which specifies a scale operation by sx and sy. If <sy> is not provided, it is assumed to be equal to <sx>.
*/

const (
	lineScaleX  = "scale(%.2f)"
	lineScaleXY = "scale(%.2f, %.2f)"
)

type SCALE struct {
	ONETRANSFORM
	sx, sy float64
}

// Constructor
func Scale(sx float64, sy ...float64) SCALE {
	scale := SCALE{
		sx: sx,
	}
	if len(sy) > 0 {
		scale.sy = sy[0]
	}
	return scale
}

// Source() returns svg implementation of "scale" transformation
func (scale SCALE) Source() string {
	if scale.sy == 0 {
		return fmt.Sprintf(lineScaleX, scale.sx)
	}
	return fmt.Sprintf(lineScaleXY, scale.sx, scale.sy)
}
