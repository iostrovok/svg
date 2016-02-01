package transform

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestROTATE(t *testing.T) {
	TestingT(t)
}

type ROTATETestsSuite struct{}

var _ = Suite(&ROTATETestsSuite{})

func (s ROTATETestsSuite) Test_Rotate(c *C) {
	//c.Skip("Not now")
	t1 := Rotate(1)
	c.Assert(t1, NotNil)

	t2 := Rotate(1, 2)
	c.Assert(t2, NotNil)

	t3 := Rotate(1, 2, 3)
	c.Assert(t3, NotNil)
}

func (s ROTATETestsSuite) Test_Rotate_Source(c *C) {
	//c.Skip("Not now")
	c.Assert(Rotate(12).Source(), Equals, "rotate(12.00)")
}

func (s ROTATETestsSuite) Test_RotateX_Source(c *C) {
	//c.Skip("Not now")
	c.Assert(Rotate(12, 34).Source(), Equals, "rotate(12.00, 34.00)")
}

func (s ROTATETestsSuite) Test_RotateXY_Source(c *C) {
	//c.Skip("Not now")
	c.Assert(Rotate(12, 34, 45.11).Source(), Equals, "rotate(12.00, 34.00, 45.11)")
}
