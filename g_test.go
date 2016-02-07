package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

func TestGROUP(t *testing.T) {
	TestingT(t)
}

type GROUPTestsSuite struct{}

var _ = Suite(&GROUPTestsSuite{})

func (s GROUPTestsSuite) Test_group(c *C) {
	//c.Skip("Not now")

	l := Group()
	c.Assert(l, NotNil)
	l2 := Group(style.Style())
	c.Assert(l2, NotNil)
}

func (s GROUPTestsSuite) Test_group2(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<g style="stroke:rgb(10,20,30)"/>`)

	l := Group(style.Style().StrokeRGB(10, 20, 30))
	res := l.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s GROUPTestsSuite) Test_Append_Title(c *C) {
	//c.Skip("Not now")
	obtained := cSL(`<g ><title>WWW</title></g>`)

	l := Group()
	t := Title("WWW")
	l = l.Append(t)
	res := l.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s GROUPTestsSuite) Test_Style(c *C) {
	//c.Skip("Not now")
	obtained := cSL(`<g style="fill:rgb(1,2,3)"><title>WWW</title></g>`)

	l := Group()
	st := style.Style().FillRGB(1, 2, 3)
	l = l.Style(st).Append(Title("WWW"))
	c.Assert(cSL(l.Source()), Equals, obtained)
}

func (s GROUPTestsSuite) Test_Transform(c *C) {
	//c.Skip("Not now")
	obtained := cSL(`<g transform="matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)"><title>WWW</title></g>`)

	l := Group()
	st := transform.Transform().Matrix(12, 34, 45, 34.5, 212, 4.45)
	l = l.Transform(st).Append(Title("WWW"))
	c.Assert(cSL(l.Source()), Equals, obtained)
}

func (s GROUPTestsSuite) Test_ID(c *C) {
	//c.Skip("Not now")

	check := cSL(`<g id="id-1"/>`)

	t := Group()
	t = t.ID("id-1")

	c.Assert("id-1", Equals, t.GetID())

	res := t.Source()
	c.Assert(cSL(res), Equals, check)
}
