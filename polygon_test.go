package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

func TestPOLYGON(t *testing.T) {
	TestingT(t)
}

type POLYGONTestsSuite struct{}

var _ = Suite(&POLYGONTestsSuite{})

func (s POLYGONTestsSuite) Test_polygon(c *C) {
	//c.Skip("Not now")

	l := Polygon()
	c.Assert(l, NotNil)
	l2 := Polygon(style.Style())
	c.Assert(l2, NotNil)
}

func (s POLYGONTestsSuite) Test_polygon2(c *C) {
	//c.Skip("Not now")

	l := Polygon(style.Style().StrokeRGB(10, 20, 30))
	res := l.Source()
	c.Assert(cSL(res), Equals, "")
}

func (s POLYGONTestsSuite) Test_Append_Title(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<polygon points="1,2 3,4"><title>WWW</title></polygon>`)

	l := Polygon().Points(1, 2, 3, 4)
	t := Title("WWW")
	l = l.Append(t)
	res := l.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s POLYGONTestsSuite) Test_Style(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<polygon points="1.1,2.22 3.3,4" style="fill:rgb(1,2,3)"><title>WWW</title></polygon>`)

	l := Polygon().Points(1.1, 2.22, 3.30, 4)
	st := style.Style().FillRGB(1, 2, 3)
	l = l.Style(st).Append(Title("WWW"))
	c.Assert(cSL(l.Source()), Equals, obtained)
}

func (s POLYGONTestsSuite) Test_Transform(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<polygon points="1.1,2.22 3.3,4" transform="matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)">
<title>WWW</title>
</polygon>`)

	l := Polygon().Points(1.1, 2.22, 3.30, 4)
	st := transform.Transform().Matrix(12, 34, 45, 34.5, 212, 4.45)
	l = l.Transform(st).Append(Title("WWW"))
	c.Assert(cSL(l.Source()), Equals, obtained)
}

func (s POLYGONTestsSuite) Test_ID(c *C) {
	//c.Skip("Not now")

	check := cSL(`<polygon points="1.1,2.22 3.3,4" id="id-1"/>`)

	t := Polygon().Points(1.1, 2.22, 3.30, 4)
	t = t.ID("id-1")

	c.Assert("id-1", Equals, t.GetID())

	res := t.Source()
	c.Assert(cSL(res), Equals, check)
}
