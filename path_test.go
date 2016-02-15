package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

func TestPATH(t *testing.T) {
	TestingT(t)
}

type PATHTestsSuite struct{}

var _ = Suite(&PATHTestsSuite{})

func (s PATHTestsSuite) Test_path(c *C) {
	//c.Skip(`Not now")

	c.Assert(Path(), NotNil)
	c.Assert(Path(style.Style()), NotNil)
}

func (s PATHTestsSuite) Test_path2(c *C) {
	//c.Skip(`Not now")

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	c.Assert(cSL(p.Source()), Equals, cSL(`<path d="" style="stroke:rgb(10,20,30)"/>`))
}

func (s PATHTestsSuite) Test_path_ML(c *C) {
	//c.Skip(`Not now")

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.M(100, 200)
	c.Assert(cSL(p.Source()), Equals, cSL(`<path d="M100,200" style="stroke:rgb(10,20,30)"/>`))

	p = Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.M(0, 0).L(100, 200)
	c.Assert(cSL(p.Source()), Equals, cSL(`<path d="M0,0 L100,200" style="stroke:rgb(10,20,30)"/>`))

	p = Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.M(0, 0).Ll(100, 200)
	c.Assert(cSL(p.Source()), Equals, cSL(`<path d="M0,0 l100,200" style="stroke:rgb(10,20,30)"/>`))

	p = Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.Mm(100, 200)
	c.Assert(cSL(p.Source()), Equals, cSL(`<path d="m100,200" style="stroke:rgb(10,20,30)"/>`))

	p = Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.Mm(100, 200).L(11, 12)
	c.Assert(cSL(p.Source()), Equals, cSL(`<path d="m100,200 L11,12" style="stroke:rgb(10,20,30)"/>`))
}

func (s PATHTestsSuite) Test_path_Lxxx(c *C) {
	//c.Skip(`Not now")

	obtained := cSL(`<path d="M2,2 L0,1 1,2 2,3 3,4 4,5 5,6 6,7 7,8 8,9 9,10" style="stroke:rgb(10,20,30)"/>`)

	p := Path(style.Style().StrokeRGB(10, 20, 30)).M(2, 2)
	for i := 0; i < 10; i++ {
		p = p.L(float64(i), float64(i+1))
	}
	c.Assert(cSL(p.Source()), Equals, obtained)
}

func (s PATHTestsSuite) Test_path_HhVv(c *C) {
	//c.Skip(`Not now")

	obtained := cSL(`<path d="M0,0 V100 v200 H300 h400" style="stroke:rgb(10,20,30)"/>`)

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.M(0, 0).V(100).Vv(200).H(300).Hh(400)
	c.Assert(cSL(p.Source()), Equals, obtained)
}

func (s PATHTestsSuite) Test_path_Qq(c *C) {
	//c.Skip(`Not now")

	obtained := cSL(`<path d="M0,0 Q1,2 3,4 q5,6 7,8 S9,0 10,11 s12,13 14,15" style="stroke:rgb(10,20,30)"/>`)

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.M(0, 0).Q(1, 2, 3, 4).Qq(5, 6, 7, 8).S(9, 0, 10, 11).Ss(12, 13, 14, 15)
	c.Assert(cSL(p.Source()), Equals, obtained)
}

func (s PATHTestsSuite) Test_path_Aa(c *C) {
	//c.Skip(`Not now")

	obtained := cSL(`<path d="M0,0 A10,20 30 1,1 40,50 a11,21 31 0,0 41,51" style="stroke:rgb(10,20,30)"/>`)
	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.M(0, 0).A(10, 20, 30, 1, 1, 40, 50).Aa(11, 21, 31, 0, 0, 41, 51)
	c.Assert(cSL(p.Source()), Equals, obtained)
}

func (s PATHTestsSuite) Test_path_Cc(c *C) {
	//c.Skip(`Not now")

	obtained := cSL(`<path d="M0,0 C10,20 30,40 50,60 c11,21 31,41 51,61" style="stroke:rgb(10,20,30)"/>`)

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.M(0, 0).C(10, 20, 30, 40, 50, 60).Cc(11, 21, 31, 41, 51, 61)
	c.Assert(cSL(p.Source()), Equals, obtained)
}

func (s PATHTestsSuite) Test_Append_Title(c *C) {
	//c.Skip(`Not now")

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.M(0, 0).L(100, 200)
	t := Title("WWW")
	p = p.Append(t)
	obtained := cSL(`<path d="M0,0 L100,200" style="stroke:rgb(10,20,30)"><title>WWW</title></path>`)

	c.Assert(cSL(p.Source()), Equals, obtained)
}

func (s PATHTestsSuite) Test_path_Style(c *C) {
	// c.Skip(`Not now")

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.M(0, 0).L(100, 200)
	st := style.Style().FillRGB(1, 2, 3)
	p = p.Style(st).Append(Title("WWW"))
	obtained := cSL(`<path d="M0,0 L100,200" style="fill:rgb(1,2,3)"><title>WWW</title></path>`)

	c.Assert(cSL(p.Source()), Equals, obtained)
}

func (s PATHTestsSuite) Test_Transform(c *C) {
	// c.Skip(`Not now")

	p := Path()
	p = p.M(0, 0).L(100, 200)
	st := transform.Transform().Matrix(12, 34, 45, 34.5, 212, 4.45)
	p = p.Transform(st).Append(Title(`WWW`))
	obtained := cSL(`<path d="M0,0 L100,200" transform="matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)">
<title>WWW</title></path>`)

	c.Assert(cSL(p.Source()), Equals, obtained)
}

func (s PATHTestsSuite) Test_ID(c *C) {
	//c.Skip(`Not now")

	check := cSL(`<path d="M0,0 L100,200" id="id-1"/>`)

	t := Path().M(0, 0).L(100, 200)
	t = t.ID("id-1")

	c.Assert("id-1", Equals, t.GetID())

	res := t.Source()
	c.Assert(cSL(res), Equals, check)
}
