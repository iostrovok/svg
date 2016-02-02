package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"fmt"
)

func TestSVG(t *testing.T) {
	TestingT(t)
}

type SVGTestsSuite struct{}

var _ = Suite(&SVGTestsSuite{})

var TestSVGWithLine string = "<line x1=\"10\" y1=\"10\" x2=\"20\" y2=\"20\"  />\n"
var TestSVGWithRect string = "<rect x=\"10\" y=\"10\" width=\"20\" height=\"20\" />\n"
var TestSVGWithPath string = "<path d=\"M0,0 V100 v200 H300 h400\" />\n"

func (s SVGTestsSuite) Test_new(c *C) {
	//c.Skip("Not now")

	// Empty SVG
	m := New(100, 100)
	c.Assert(m, NotNil)
}

func (s SVGTestsSuite) Test_SVG_Source(c *C) {
	//c.Skip("Not now")

	m := New(100, 100)
	t := len(m.Source()) > 10
	c.Assert(t, Equals, true)
}

func (s SVGTestsSuite) Test_SVG_ViewBox(c *C) {
	//c.Skip("Not now")

	m := New(100, 100).ViewBox(0, 0, 200, 300)

	c.Assert(m.Height(), Equals, 100)
	m.Height(200)
	c.Assert(m.Height(), Equals, 200)

	c.Assert(m.Width(), Equals, 100)
	m.Width(600)
	c.Assert(m.Width(), Equals, 600)

	fmt.Printf("\n\n---------------------------\n\n%s\n\n--------------------------------\n", m.Source())

	line := `<?xml version="1.0"?>
<svg
	width="600" height="200" viewBox="0 0 200 300"
	xmlns="http://www.w3.org/2000/svg"
	xmlns:xlink="http://www.w3.org/1999/xlink"
/>
`

	c.Assert(m.Source(), Equals, line)
}

func (s SVGTestsSuite) Test_AppendTo_Line(c *C) {
	//c.Skip("Not now")
	m := New(100, 100)
	Line(10, 10, 20, 20).AppendTo(m)
	c.Assert(m.InnerSource(), Equals, TestSVGWithLine)
}

func (s SVGTestsSuite) Test_AppendTo_Rect(c *C) {
	//c.Skip("Not now")
	m := New(100, 100)
	Rect(10, 10, 20, 20).AppendTo(m)
	c.Assert(m.InnerSource(), Equals, TestSVGWithRect)
}

func (s SVGTestsSuite) Test_AppendTo_Path(c *C) {
	c.Skip("Not now")
	m := New(100, 100)
	Path().M(0, 0).V(100).Vv(200).H(300).Hh(400).AppendTo(m)
	c.Assert(m.InnerSource(), Equals, TestSVGWithPath)
}

func (s SVGTestsSuite) Test_Append_Line(c *C) {
	//c.Skip("Not now")
	l := Line(10, 10, 20, 20)
	m := New(100, 100).Append(l)
	c.Assert(m.InnerSource(), Equals, TestSVGWithLine)
}

func (s SVGTestsSuite) Test_Append_Rect(c *C) {
	//c.Skip("Not now")
	r := Rect(10, 10, 20, 20)
	m := New(100, 100).Append(r)
	c.Assert(m.InnerSource(), Equals, TestSVGWithRect)
}

func (s SVGTestsSuite) Test_Append_Path(c *C) {
	//c.Skip("Not now")
	p := Path().M(0, 0).V(100).Vv(200).H(300).Hh(400)
	m := New(1000, 1000).Append(p)
	c.Assert(m.InnerSource(), Equals, TestSVGWithPath)
}
