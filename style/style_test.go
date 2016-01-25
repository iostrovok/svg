package style

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func TestSTYLE(t *testing.T) {
	TestingT(t)
}

type STYLETestsSuite struct{}

var _ = Suite(&STYLETestsSuite{})

func (s STYLETestsSuite) Test_Style(c *C) {
	//c.Skip("Not now")

	l1 := Style()
	c.Assert(l1, NotNil)
}

func (s STYLETestsSuite) Test_Source(c *C) {
	//c.Skip("Not now")

	l1 := Style()
	c.Assert(fmt.Sprintf("%T", l1.Source()), Equals, "string")
	c.Assert(len(l1.Source()), Equals, 0)
}

func (s STYLETestsSuite) Test_StrokeWidth(c *C) {
	//c.Skip("Not now")

	l1 := Style()
	c.Assert(l1.StrokeWidth(12.12).Source(), Equals, "style=\"stroke-width:12.11\"")
}

func (s STYLETestsSuite) Test_Stroke(c *C) {
	//c.Skip("Not now")

	l1 := Style()
	c.Assert(l1.Stroke(12, 34, 45).Source(), Equals, "style=\"stroke:rgb(12,34,45)\"")
}

func (s STYLETestsSuite) Test_FillRGB(c *C) {
	//c.Skip("Not now")

	l1 := Style()
	c.Assert(l1.FillRGB(12, 34, 45).Source(), Equals, "style=\"fill:rgb(12,34,45)\"")
}

func (s STYLETestsSuite) Test_Fill(c *C) {
	//c.Skip("Not now")

	l1 := Style()
	c.Assert(l1.Fill("black").Source(), Equals, "style=\"fill:black\"")
}
