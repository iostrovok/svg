package transform

import (
	"fmt"
)

/*
	translate(<tx> [<ty>]), which specifies a translation by tx and ty. If <ty> is not provided, it is assumed to be zero.
*/

const (
	lineTranslateX  = "translate(%.2f)"
	lineTranslateXY = "translate(%.2f, %.2f)"
)

type TRANSLATE struct {
	ONETRANSFORM
	sx, sy float64
}

// Constructor
func Translate(sx float64, sy ...float64) TRANSLATE {
	translate := TRANSLATE{
		sx: sx,
	}
	if len(sy) > 0 {
		translate.sy = sy[0]
	}
	return translate
}

// Source() returns svg implementation of "translate" transformation
func (translate TRANSLATE) Source() string {
	if translate.sy == 0 {
		return fmt.Sprintf(lineTranslateX, translate.sx)
	}
	return fmt.Sprintf(lineTranslateXY, translate.sx, translate.sy)
}
