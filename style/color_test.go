package style

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestCOLOR(t *testing.T) {
	TestingT(t)
}

type COLORTestsSuite struct{}

var _ = Suite(&COLORTestsSuite{})

func (s COLORTestsSuite) Test_Color(c *C) {
	//c.Skip("Not now")

	l1 := Color(1, 2, 3)
	c.Assert(l1, NotNil)
	l2 := ColorStr("black")
	c.Assert(l2, NotNil)
}

func (s COLORTestsSuite) Test_Color_RGB(c *C) {
	//c.Skip("Not now")

	l1 := ColorStr("black")
	c.Assert(l1.RGB(10, 20, 30).Source(), Equals, "rgb(10,20,30)")
}

func (s COLORTestsSuite) Test_Color_Source(c *C) {
	//c.Skip("Not now")

	c.Assert(Color(1, 2, 3).Source(), Equals, "rgb(1,2,3)")
	c.Assert(ColorStr("black").Source(), Equals, "black")
}
