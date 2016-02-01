package transform

import (
	"fmt"
)

/*
	rotate(<rotate-angle> [<cx> <cy>]), which specifies a rotation by <rotate-angle> degrees about a given point.
*/

const (
	lineRotate   = "rotate(%.2f)"
	lineRotateX  = "rotate(%.2f, %.2f)"
	lineRotateXY = "rotate(%.2f, %.2f, %.2f)"
)

type ROTATE struct {
	ONETRANSFORM
	angle, sx, sy float64
}

// Constructor
func Rotate(angle float64, s ...float64) ROTATE {
	rotate := ROTATE{
		angle: angle,
	}
	if len(s) > 0 {
		rotate.sx = s[0]
	}
	if len(s) > 1 {
		rotate.sy = s[1]
	}

	return rotate
}

// Source() returns svg implementation of "rotate" transformation
func (rotate ROTATE) Source() string {
	if rotate.sy == 0 && rotate.sx == 0 {
		return fmt.Sprintf(lineRotate, rotate.angle)
	}
	if rotate.sy == 0 {
		return fmt.Sprintf(lineRotateX, rotate.angle, rotate.sx)
	}
	return fmt.Sprintf(lineRotateXY, rotate.angle, rotate.sx, rotate.sy)
}
