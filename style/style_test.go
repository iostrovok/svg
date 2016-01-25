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

	st := Style()
	c.Assert(st, NotNil)
}

func (s STYLETestsSuite) Test_Source(c *C) {
	//c.Skip("Not now")

	st := Style()
	c.Assert(fmt.Sprintf("%T", st.Source()), Equals, "string")
	c.Assert(len(st.Source()), Equals, 0)
}

func (s STYLETestsSuite) Test_StrokeWidth(c *C) {
	//c.Skip("Not now")

	st := Style()
	c.Assert(st.StrokeWidth(12.12).Source(), Equals, "style=\"stroke-width:12.11\"")
}

func (s STYLETestsSuite) Test_StrokeRGB(c *C) {
	//c.Skip("Not now")

	st := Style()
	c.Assert(st.StrokeRGB(12, 34, 45).Source(), Equals, "style=\"stroke:rgb(12,34,45)\"")
}

func (s STYLETestsSuite) Test_Stroke(c *C) {
	//c.Skip("Not now")

	st := Style()
	c.Assert(st.Stroke("black").Source(), Equals, "style=\"stroke:black\"")
}

func (s STYLETestsSuite) Test_FillRGB(c *C) {
	//c.Skip("Not now")

	st := Style()
	c.Assert(st.FillRGB(12, 34, 45).Source(), Equals, "style=\"fill:rgb(12,34,45)\"")
}

func (s STYLETestsSuite) Test_Fill(c *C) {
	//c.Skip("Not now")

	st := Style()
	c.Assert(st.Fill("black").Source(), Equals, "style=\"fill:black\"")
}
