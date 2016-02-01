package transform

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func TestTRANSFORM(t *testing.T) {
	TestingT(t)
}

type TRANSFORMTestsSuite struct{}

var _ = Suite(&TRANSFORMTestsSuite{})

func (s TRANSFORMTestsSuite) Test_Transform(c *C) {
	// c.Skip("Not now")

	st := Transform()
	c.Assert(st, NotNil)
}

func (s TRANSFORMTestsSuite) Test_Source(c *C) {
	// c.Skip("Not now")

	st := Transform()
	c.Assert(fmt.Sprintf("%T", st.Source()), Equals, "string")
	c.Assert(len(st.Source()), Equals, 0)
}

func (s TRANSFORMTestsSuite) Test_StrokeWidth(c *C) {
	// c.Skip("Not now")

	st := Transform()
	c.Assert(st.SkewX(12.342).Source(), Equals, `transform="skewX=(12.34)"`)
}

func (s TRANSFORMTestsSuite) Test_Matrix(c *C) {
	//c.Skip("Not now")

	st := Transform()
	c.Assert(st.Matrix(12, 34, 45, 34.5, 212, 4.45).Source(), Equals, `transform="matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)"`)
}

func (s TRANSFORMTestsSuite) Test_ScaleX(c *C) {
	//c.Skip("Not now")
	st := Transform()
	c.Assert(st.Scale(21.45).Source(), Equals, `transform="scale(21.45)"`)
}

func (s TRANSFORMTestsSuite) Test_ScaleXY(c *C) {
	//c.Skip("Not now")
	st := Transform()
	c.Assert(st.Scale(212, 4.45).Source(), Equals, `transform="scale(212.00, 4.45)"`)
}

// func (s TRANSFORMTestsSuite) Test_Stroke(c *C) {
// 	//c.Skip("Not now")

// 	st := Transform()
// 	c.Assert(st.Stroke("black").Source(), Equals, "style=\"stroke:black\"")
// }

// func (s TRANSFORMTestsSuite) Test_FillRGB(c *C) {
// 	//c.Skip("Not now")

// 	st := Transform()
// 	c.Assert(st.FillRGB(12, 34, 45).Source(), Equals, "style=\"fill:rgb(12,34,45)\"")
// }

// func (s TRANSFORMTestsSuite) Test_Fill(c *C) {
// 	//c.Skip("Not now")

// 	st := Transform()
// 	c.Assert(st.Fill("black").Source(), Equals, "style=\"fill:black\"")
// }

// func (s TRANSFORMTestsSuite) Test_Font(c *C) {
// 	//c.Skip("Not now")

// 	st := Transform()
// 	f := Font("---")
// 	c.Assert(st.Font(f).Source(), Equals, "style=\"---\"")
// }
