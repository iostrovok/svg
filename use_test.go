package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

func TestUSE(t *testing.T) {
	TestingT(t)
}

type USETestsSuite struct{}

var _ = Suite(&USETestsSuite{})

func (s USETestsSuite) Test_use(c *C) {
	//c.Skip("Not now")

	l := Use()
	c.Assert(l, NotNil)
	l2 := Use(style.Style())
	c.Assert(l2, NotNil)
}

func (s USETestsSuite) Test_use2(c *C) {
	//c.Skip("Not now")

	l := Use(style.Style().StrokeRGB(10, 20, 30))
	res := l.Source()
	c.Assert(cSL(res), Equals, ``)
}

func (s USETestsSuite) Test_use_href(c *C) {
	//c.Skip("Not now")
	obtained := cSL(`<use xlink:href="#MySymbol" style="stroke:rgb(10,20,30)"/>`)

	l := Use(style.Style().StrokeRGB(10, 20, 30)).Href(`#MySymbol`)
	res := l.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s USETestsSuite) Test_Append_Title(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<use xlink:href="#MySymbol"><title>WWW</title></use>`)

	l := Use().Href(`#MySymbol`)
	t := Title("WWW")
	l = l.Append(t)
	res := l.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s USETestsSuite) Test_Style(c *C) {
	// c.Skip("Not now")

	obtained := cSL(`<use xlink:href="#MySymbol" style="fill:rgb(1,2,3)" ><title>WWW</title></use>`)

	l := Use().Href(`#MySymbol`)
	st := style.Style().FillRGB(1, 2, 3)
	l = l.Style(st).Append(Title("WWW"))
	c.Assert(cSL(l.Source()), Equals, obtained)
}

func (s USETestsSuite) Test_Transform(c *C) {
	// c.Skip("Not now")

	obtained := cSL(`<use xlink:href="#MySymbol" transform="matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)" ><title>WWW</title></use>`)

	l := Use().Href(`#MySymbol`)
	st := transform.Transform().Matrix(12, 34, 45, 34.5, 212, 4.45)
	l = l.Transform(st).Append(Title("WWW"))
	c.Assert(cSL(l.Source()), Equals, obtained)
}

func (s USETestsSuite) Test_ID(c *C) {
	// c.Skip("Not now")

	check := cSL(`<use xlink:href="#MySymbol" id="id-1"/>`)

	t := Use().Href(`#MySymbol`)
	t = t.ID("id-1")

	c.Assert("id-1", Equals, t.GetID())

	res := t.Source()
	c.Assert(cSL(res), Equals, check)
}

func (s USETestsSuite) Test_X(c *C) {
	// c.Skip("Not now")

	check := cSL(`<use xlink:href="#MySymbol" x="10pt"/>`)

	t := Use().Href(`#MySymbol`)
	t = t.X(10, "pt")
	res := t.Source()
	c.Assert(cSL(res), Equals, check)
}

func (s USETestsSuite) Test_Y(c *C) {
	// c.Skip("Not now")

	check := cSL(`<use xlink:href="#MySymbol" y="10"/>`)

	t := Use().Href(`#MySymbol`)
	t = t.Y(10)
	res := t.Source()
	c.Assert(cSL(res), Equals, check)
}

func (s USETestsSuite) Test_Width(c *C) {
	// c.Skip("Not now")

	check := cSL(`<use xlink:href="#MySymbol" width="10.20"/>`)

	t := Use().Href(`#MySymbol`)
	t = t.Width(10.2)
	res := t.Source()
	c.Assert(cSL(res), Equals, check)
}

func (s USETestsSuite) Test_Height(c *C) {
	// c.Skip("Not now")

	check := cSL(`<use xlink:href="#MySymbol" height="10.23"/>`)

	t := Use().Href(`#MySymbol`)
	t = t.Height(10.231)
	res := t.Source()
	c.Assert(cSL(res), Equals, check)
}
