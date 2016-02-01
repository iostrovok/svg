package transform

import (
	"fmt"
	"math"
)

// skewX(<skew-angle>), which specifies a skew transformation along the x-axis.
// skewY(<skew-angle>), which specifies a skew transformation along the y-axis.

const (
	lineSkewX   = "skewX=(%s)"
	lineSkewY   = "skewY=(%s)"
	lineSkew    = "%.2f"
	lineSkewInt = "%d"
)

type SKEW struct {
	ONETRANSFORM
	angle float64
	typ   string
}

// Constructor
func Skew(typ string, angle float64) SKEW {
	return SKEW{
		angle: angle,
		typ:   typ,
	}
}

// Source() returns svg implementation of "skewX" or "skewY" transformation
func (skew SKEW) Source() string {
	if skew.typ == "x" {
		return fmt.Sprintf(lineSkewX, skew._source())
	}
	return fmt.Sprintf(lineSkewY, skew._source())
}

func (skew SKEW) _source() string {
	if math.Floor(skew.angle) == skew.angle {
		return fmt.Sprintf(lineSkewInt, int(skew.angle))
	}

	return fmt.Sprintf(lineSkew, skew.angle)
}
