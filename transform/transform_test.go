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
	c.Assert(Transform().SkewX(12.342).Source(), Equals, `transform="skewX=(12.34)"`)
}

func (s TRANSFORMTestsSuite) Test_Matrix(c *C) {
	//c.Skip("Not now")
	c.Assert(Transform().Matrix(12, 34, 45, 34.5, 212, 4.45).Source(), Equals, `transform="matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)"`)
}

func (s TRANSFORMTestsSuite) Test_ScaleX(c *C) {
	//c.Skip("Not now")
	c.Assert(Transform().Scale(21.45).Source(), Equals, `transform="scale(21.45)"`)
}

func (s TRANSFORMTestsSuite) Test_ScaleXY(c *C) {
	//c.Skip("Not now")
	c.Assert(Transform().Scale(212, 4.45).Source(), Equals, `transform="scale(212.00, 4.45)"`)
}

func (s TRANSFORMTestsSuite) Test_TranslateX(c *C) {
	//c.Skip("Not now")
	c.Assert(Transform().Translate(21.45).Source(), Equals, `transform="translate(21.45)"`)
}

func (s TRANSFORMTestsSuite) Test_TranslateXY(c *C) {
	//c.Skip("Not now")
	c.Assert(Transform().Translate(212, 4.45).Source(), Equals, `transform="translate(212.00, 4.45)"`)
}

func (s TRANSFORMTestsSuite) Test_Rotate(c *C) {
	//c.Skip("Not now")
	c.Assert(Transform().Rotate(21.45).Source(), Equals, `transform="rotate(21.45)"`)
}

func (s TRANSFORMTestsSuite) Test_RotateX(c *C) {
	//c.Skip("Not now")
	c.Assert(Transform().Rotate(21.45, 23.45).Source(), Equals, `transform="rotate(21.45, 23.45)"`)
}

func (s TRANSFORMTestsSuite) Test_RotateXY(c *C) {
	//c.Skip("Not now")
	c.Assert(Transform().Rotate(212, 4.45, 23.45).Source(), Equals, `transform="rotate(212.00, 4.45, 23.45)"`)
}
