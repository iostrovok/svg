package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

func TestRECT(t *testing.T) {
	TestingT(t)
}

type RECTTestsSuite struct{}

var _ = Suite(&RECTTestsSuite{})

func (s RECTTestsSuite) Test_rect(c *C) {
	//c.Skip("Not now")

	l := Rect(1, 2, 3, 4)
	c.Assert(l, NotNil)
	l2 := Rect(1, 2, 3, 4, style.Style())
	c.Assert(l2, NotNil)
}

func (s RECTTestsSuite) Test_rect2(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<rect height="4" width="3" x="1" y="1" style="stroke:rgb(10,20,30)"/>`)

	l := Rect(1, 2, 3, 4, style.Style().StrokeRGB(10, 20, 30))
	res := l.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s RECTTestsSuite) Test_Append_Title(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<rect height="4" width="3" x="1" y="1"><title>WWW</title></rect>`)

	l := Rect(1, 2, 3, 4)
	t := Title("WWW")
	l = l.Append(t)
	res := l.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s RECTTestsSuite) Test_Style(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<rect height="4" width="3" x="1" y="1" style="fill:rgb(1,2,3)"><title>WWW</title></rect>`)

	l := Rect(1, 2, 3, 4)
	st := style.Style().FillRGB(1, 2, 3)
	l = l.Style(st).Append(Title("WWW"))
	c.Assert(cSL(l.Source()), Equals, obtained)
}

func (s RECTTestsSuite) Test_Transform(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<rect height="4" width="3" x="1" y="1" transform="matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)">
<title>WWW</title>
</rect>`)

	l := Rect(1, 2, 3, 4)
	st := transform.Transform().Matrix(12, 34, 45, 34.5, 212, 4.45)
	l = l.Transform(st).Append(Title("WWW"))
	c.Assert(cSL(l.Source()), Equals, obtained)
}

func (s RECTTestsSuite) Test_ID(c *C) {
	//c.Skip("Not now")

	check := cSL(`<rect id="id-1" height="4" width="3" x="1" y="1"/>`)

	t := Rect(1, 2, 3, 4)
	t = t.ID("id-1")

	c.Assert("id-1", Equals, t.GetID())

	res := t.Source()
	c.Assert(cSL(res), Equals, check)
}
