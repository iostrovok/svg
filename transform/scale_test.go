package transform

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestSCALE(t *testing.T) {
	TestingT(t)
}

type SCALETestsSuite struct{}

var _ = Suite(&SCALETestsSuite{})

func (s SCALETestsSuite) Test_Scale(c *C) {
	//c.Skip("Not now")
	t1 := Scale(1, 2)
	c.Assert(t1, NotNil)
}

func (s SCALETestsSuite) Test_ScaleX_Source(c *C) {
	//c.Skip("Not now")
	c.Assert(Scale(12).Source(), Equals, "scale(12.00)")
}

func (s SCALETestsSuite) Test_ScaleXY_Source(c *C) {
	//c.Skip("Not now")
	c.Assert(Scale(12, 34).Source(), Equals, "scale(12.00, 34.00)")
}
