package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
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

	obtained := cSL(`<line x1="1" y1="2" x2="3" y2="4" style="stroke:rgb(10,20,30)" />`)

	l := Line(1, 2, 3, 4, style.Style().StrokeRGB(10, 20, 30))
	res := l.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s LINETestsSuite) Test_line_Append_Title(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<line x1="1" y1="2" x2="3" y2="4"  ><title>WWW</title></line>`)

	l := Line(1, 2, 3, 4)
	t := Title("WWW")
	l = l.Append(t)
	res := l.Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s LINETestsSuite) Test_line_Style(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<line x1="1" y1="2" x2="3" y2="4" style="fill:rgb(1,2,3)" ><title>WWW</title></line>`)

	l := Line(1, 2, 3, 4)
	st := style.Style().FillRGB(1, 2, 3)
	l = l.Style(st).Append(Title("WWW"))
	c.Assert(cSL(l.Source()), Equals, obtained)
}

func (s LINETestsSuite) Test_line_Transform(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<line x1="1" y1="2" x2="3" y2="4" transform="matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)" ><title>WWW</title></line>`)

	l := Line(1, 2, 3, 4)
	st := transform.Transform().Matrix(12, 34, 45, 34.5, 212, 4.45)
	l = l.Transform(st).Append(Title("WWW"))
	c.Assert(cSL(l.Source()), Equals, obtained)
}
