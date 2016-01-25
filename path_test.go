package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/iostrovok/svg/style"
)

func TestPATH(t *testing.T) {
	TestingT(t)
}

type PATHTestsSuite struct{}

var _ = Suite(&PATHTestsSuite{})

func (s PATHTestsSuite) Test_path(c *C) {
	//c.Skip("Not now")

	c.Assert(Path(), NotNil)
	c.Assert(Path(style.Style()), NotNil)
}

func (s PATHTestsSuite) Test_path2(c *C) {
	//c.Skip("Not now")

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	c.Assert(string(p.Source()), Equals, "<path d=\"\" style=\"stroke:rgb(10,20,30)\"/>\n")
}

func (s PATHTestsSuite) Test_apth_ML(c *C) {
	//c.Skip("Not now")

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.M(100, 200)
	c.Assert(string(p.Source()), Equals, "<path d=\"M100,200\" style=\"stroke:rgb(10,20,30)\"/>\n")

	p = Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.L(100, 200)
	c.Assert(string(p.Source()), Equals, "<path d=\"L100,200\" style=\"stroke:rgb(10,20,30)\"/>\n")

	p = Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.l(100, 200)
	c.Assert(string(p.Source()), Equals, "<path d=\"l100,200\" style=\"stroke:rgb(10,20,30)\"/>\n")

	p = Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.m(100, 200)
	c.Assert(string(p.Source()), Equals, "<path d=\"m100,200\" style=\"stroke:rgb(10,20,30)\"/>\n")

	p = Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.m(100, 200).L(11, 12)
	c.Assert(string(p.Source()), Equals, "<path d=\"m100,200 L11,12\" style=\"stroke:rgb(10,20,30)\"/>\n")
}

func (s PATHTestsSuite) Test_apth_HhVv(c *C) {
	//c.Skip("Not now")

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.V(100).v(200).H(300).h(400)
	c.Assert(string(p.Source()), Equals, "<path d=\"V100 v200 H300 h400\" style=\"stroke:rgb(10,20,30)\"/>\n")
}

func (s PATHTestsSuite) Test_apth_Qq(c *C) {
	//c.Skip("Not now")

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.Q(1, 2, 3, 4).q(5, 6, 7, 8).S(9, 0, 10, 11).s(12, 13, 14, 15)
	c.Assert(string(p.Source()), Equals, "<path d=\"Q1,2 3,4 q5,6 7,8 S9,0 10,11 s12,13 14,15\" style=\"stroke:rgb(10,20,30)\"/>\n")
}

func (s PATHTestsSuite) Test_apth_Aa(c *C) {
	//c.Skip("Not now")

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.A(10, 20, 30, 1, 1, 40, 50).a(11, 21, 31, 0, 0, 41, 51)
	c.Assert(string(p.Source()), Equals, "<path d=\"A10,20 30 1,1 40,50 a11,21 31 0,0 41,51\" style=\"stroke:rgb(10,20,30)\"/>\n")
}

func (s PATHTestsSuite) Test_apth_Cc(c *C) {
	//c.Skip("Not now")

	p := Path(style.Style().StrokeRGB(10, 20, 30))
	p = p.C(10, 20, 30, 40, 50, 60).c(11, 21, 31, 41, 51, 61)
	c.Assert(string(p.Source()), Equals, "<path d=\"C10,20 30,40 50,60 c11,21 31,41 51,61\" style=\"stroke:rgb(10,20,30)\"/>\n")
}
