package transform

import (
	"fmt"
)

/*
	matrix(<a> <b> <c> <d> <e> <f>), which specifies a transformation in the form of a transformation matrix of six values.
	matrix(a,b,c,d,e,f) is equivalent to applying the transformation matrix [a b c d e f].
*/

const (
	lineMatrix = "matrix(%.2f, %.2f, %.2f, %.2f, %.2f, %.2f)"
)

type MATRIX struct {
	ONETRANSFORM
	a, b, c, d, e, f float64
}

// Constructor
func Matrix(a, b, c, d, e, f float64) MATRIX {
	return MATRIX{
		a: a, b: b, c: c,
		d: d, e: e, f: f,
	}
}

// Source() returns svg implementation of "matrix" transformation
func (matrix MATRIX) Source() string {
	return fmt.Sprintf(lineMatrix, matrix.a, matrix.b, matrix.c, matrix.d, matrix.e, matrix.f)
}
