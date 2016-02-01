package transform

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestSKEW(t *testing.T) {
	TestingT(t)
}

type SKEWTestsSuite struct{}

var _ = Suite(&SKEWTestsSuite{})

func (s SKEWTestsSuite) Test_Skew(c *C) {
	//c.Skip("Not now")
	t1 := Skew("x", 12.45)
	c.Assert(t1, NotNil)
	t2 := Skew("y", 12)
	c.Assert(t2, NotNil)
}

func (s SKEWTestsSuite) Test_Skew_Source(c *C) {
	//c.Skip("Not now")
	c.Assert(Skew("x", 12.45).Source(), Equals, "skewX=(12.45)")
}
