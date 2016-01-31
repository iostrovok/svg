package style

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestFONT(t *testing.T) {
	TestingT(t)
}

type FONTTestsSuite struct{}

var _ = Suite(&FONTTestsSuite{})

func (s FONTTestsSuite) Test_Font(c *C) {
	//c.Skip("Not now")

	t1 := Font("font-family:Verdana;font-size:55")
	c.Assert(t1, NotNil)
	t2 := Font()
	c.Assert(t2, NotNil)
}

func (s FONTTestsSuite) Test_Font_Source(c *C) {
	//c.Skip("Not now")
	c.Assert(Font("font-family:Verdana;font-size:55").Source(), Equals, "font-family:Verdana;font-size:55")
}

func (s FONTTestsSuite) Test_Font_Setters(c *C) {
	//c.Skip("Not now")
	t := Font()
	t = t.Family("sans-serif").Size("13em")
	c.Assert(t.Source(), Equals, "font-family:sans-serif;font-size:13em")
}

func (s FONTTestsSuite) Test_Font_Weight(c *C) {
	//c.Skip("Not now")
	t := Font()
	t = t.Weight("bold")
	c.Assert(t.Source(), Equals, "font-weight:bold")
}

func (s FONTTestsSuite) Test_Font_Variant(c *C) {
	//c.Skip("Not now")
	t := Font().Variant("small-caps")
	c.Assert(t.Source(), Equals, "font-variant:small-caps")
}

func (s FONTTestsSuite) Test_Font_Stretch(c *C) {
	//c.Skip("Not now")
	t := Font().Stretch("wider")
	c.Assert(t.Source(), Equals, "font-stretch:wider")
}
