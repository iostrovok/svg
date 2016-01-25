package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/iostrovok/svg/style"
)

func TestLINE(t *testing.T) {
	TestingT(t)
}

type LINETestsSuite struct{}

var _ = Suite(&LINETestsSuite{})

func (s LINETestsSuite) Test_line(c *C) {
	//c.Skip("Not now")

	l := Line(1, 2, 3, 4)
	c.Assert(l, NotNil)
	l2 := Line(1, 2, 3, 4, style.Style())
	c.Assert(l2, NotNil)
}

func (s LINETestsSuite) Test_line2(c *C) {
	//c.Skip("Not now")

	l := Line(1, 2, 3, 4, style.Style().StrokeRGB(10, 20, 30))
	res := l.Source()
	c.Assert(string(res), Equals, "<line x1=\"1\" y1=\"2\" x2=\"3\" y2=\"4\" style=\"stroke:rgb(10,20,30)\" />\n")
}
