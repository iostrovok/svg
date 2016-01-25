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

	color1 := RGB(1, 2, 3)
	c.Assert(color1, NotNil)
	color2 := Color("black")
	c.Assert(color2, NotNil)
}

func (s COLORTestsSuite) Test_Color_RGB(c *C) {
	//c.Skip("Not now")

	color1 := Color("black")
	c.Assert(color1.RGB(10, 20, 30).Source(), Equals, "rgb(10,20,30)")
}

func (s COLORTestsSuite) Test_Color_Source(c *C) {
	//c.Skip("Not now")

	c.Assert(RGB(1, 2, 3).Source(), Equals, "rgb(1,2,3)")
	c.Assert(Color("black").Source(), Equals, "black")
}
