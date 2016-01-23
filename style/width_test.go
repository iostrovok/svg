package style

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestWIDTH(t *testing.T) {
	TestingT(t)
}

type WIDTHTestsSuite struct{}

var _ = Suite(&WIDTHTestsSuite{})

func (s WIDTHTestsSuite) Test_Width(c *C) {
	//c.Skip("Not now")

	l1 := Width(1.0)
	c.Assert(l1, NotNil)
	l2 := Width(10)
	c.Assert(l2, NotNil)
}

func (s WIDTHTestsSuite) Test_Width_Source(c *C) {
	//c.Skip("Not now")

	c.Assert(Width(1).Source(), Equals, "1")
	c.Assert(Width(1.0).Source(), Equals, "1")
	c.Assert(Width(23.12).Source(), Equals, "23.12")
	c.Assert(Width(233.1223).Source(), Equals, "233.12")

}
