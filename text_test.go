package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

func TestTEXT(t *testing.T) {
	TestingT(t)
}

type TEXTTestsSuite struct{}

var _ = Suite(&TEXTTestsSuite{})

func (s TEXTTestsSuite) Test_text(c *C) {
	//c.Skip("Not now")

	t1 := Text()
	c.Assert(t1, NotNil)
	t2 := Text(style.Style())
	c.Assert(t2, NotNil)
}

func (s TEXTTestsSuite) Test_text2(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<text x="250" y="150" style="font-family:Verdana;font-size:55;fill:blue">Hello, out there</text>`)
	f := style.Font("font-family:Verdana;font-size:55")
	text := Text(style.Style().Fill("blue").Font(f))
	text = text.XY(250, 150).String("Hello, out there")
	res := text.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s TEXTTestsSuite) Test_AddString(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<text x="250" y="150" style="font-family:Verdana;font-size:55;fill:blue" >Hello, out there Next hello, out there</text>`)
	f := style.Font("font-family:Verdana;font-size:55")
	text := Text(style.Style().Fill("blue").Font(f))
	text = text.XY(250, 150).String("Hello, out there")
	text = text.AddString(" Next hello, out there")
	res := text.Source()

	c.Assert(cSL(res), Equals, obtained)
}

func (s TEXTTestsSuite) Test_Transform(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<text x="250" y="150" style="font-family:Verdana;font-size:55;fill:blue" transform="matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)">Hello, out there</text>`)
	f := style.Font("font-family:Verdana;font-size:55")
	text := Text(style.Style().Fill("blue").Font(f))
	text = text.XY(250, 150).String("Hello, out there")
	tr := transform.Transform().Matrix(12, 34, 45, 34.5, 212, 4.45)
	text = text.Transform(tr)
	res := text.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s TEXTTestsSuite) Test_Attrs(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<text id="id-1">Next hello, out there</text>`)

	text := Text()
	text = text.ID("id-1").AddString("Next hello, out there")

	c.Assert("id-1", Equals, text.GetID())

	res := text.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s TEXTTestsSuite) Test_DXY(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<text dx="10px" dy="10px">Next hello, out there</text>`)

	text := Text().DXY(10, 10, "px").AddString("Next hello, out there")

	res := text.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s TEXTTestsSuite) Test_XY(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<text x="10.1px" y="110px">Next hello, out there</text>`)

	text := Text().XY(10.1, 110, "px").AddString("Next hello, out there")

	res := text.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s TEXTTestsSuite) Test_ID(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<text id="id-1">Next hello, out there</text>`)

	text := Text()
	text = text.ID("id-1").AddString("Next hello, out there")

	c.Assert("id-1", Equals, text.GetID())

	res := text.Source()
	c.Assert(cSL(res), Equals, obtained)
}
