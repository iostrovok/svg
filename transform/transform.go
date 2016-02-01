package transform

import (
	"strings"
)

// matrix(<a> <b> <c> <d> <e> <f>), which specifies a transformation in the form of a transformation matrix of six values. matrix(a,b,c,d,e,f) is equivalent to applying the transformation matrix [a b c d e f].

// translate(<tx> [<ty>]), which specifies a translation by tx and ty. If <ty> is not provided, it is assumed to be zero.

// scale(<sx> [<sy>]), which specifies a scale operation by sx and sy. If <sy> is not provided, it is assumed to be equal to <sx>.

// rotate(<rotate-angle> [<cx> <cy>]), which specifies a rotation by <rotate-angle> degrees about a given point.
// If optional parameters <cx> and <cy> are not supplied, the rotate is about the origin of the current user coordinate system. The operation corresponds to the matrix [cos(a) sin(a) -sin(a) cos(a) 0 0].
// If optional parameters <cx> and <cy> are supplied, the rotate is about the point (cx, cy). The operation represents the equivalent of the following specification: translate(<cx>, <cy>) rotate(<rotate-angle>) translate(-<cx>, -<cy>).

// skewX(<skew-angle>), which specifies a skew transformation along the x-axis.

type ONETRANSFORM interface {
	Source() string
}

type TRANSFORM struct {
	tags []ONETRANSFORM
}

// Style is the constructor
func Transform() TRANSFORM {
	return TRANSFORM{
		tags: []ONETRANSFORM{},
	}
}

// Source() returns svg implementation of TRANSFORM element
func (trans TRANSFORM) Source() string {

	if len(trans.tags) == 0 {
		return ""
	}

	out := []string{}
	for _, v := range trans.tags {
		// out = append(out, k+Eq+url.QueryEscape(v.Source()))
		out = append(out, v.Source())
	}
	return `transform="` + strings.Join(out, " ") + `"`
}

func (trans TRANSFORM) Matrix(a, b, c, d, e, f float64) TRANSFORM {
	trans.tags = append(trans.tags, Matrix(a, b, c, d, e, f))
	return trans
}

func (trans TRANSFORM) SkewX(angle float64) TRANSFORM {
	trans.tags = append(trans.tags, Skew("x", angle))
	return trans
}

func (trans TRANSFORM) SkewY(angle float64) TRANSFORM {
	trans.tags = append(trans.tags, Skew("y", angle))
	return trans
}

func (trans TRANSFORM) Scale(sx float64, sy ...float64) TRANSFORM {
	trans.tags = append(trans.tags, Scale(sx, sy...))
	return trans
}

func (trans TRANSFORM) Translate(sx float64, sy ...float64) TRANSFORM {
	trans.tags = append(trans.tags, Translate(sx, sy...))
	return trans
}

func (trans TRANSFORM) Rotate(angle float64, xy ...float64) TRANSFORM {
	trans.tags = append(trans.tags, Rotate(angle, xy...))
	return trans
}
