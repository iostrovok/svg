package transform

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestMATRIX(t *testing.T) {
	TestingT(t)
}

type MATRIXTestsSuite struct{}

var _ = Suite(&MATRIXTestsSuite{})

func (s MATRIXTestsSuite) Test_Matrix(c *C) {
	//c.Skip("Not now")
	t1 := Matrix(1, 2, 3, 4, 5, 6)
	c.Assert(t1, NotNil)
}

func (s MATRIXTestsSuite) Test_Matrix_Source(c *C) {
	//c.Skip("Not now")
	c.Assert(Matrix(12, 34, 45, 34.5, 212, 4.45).Source(), Equals, "matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)")
}
