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

	l := Group(style.Style().StrokeRGB(10, 20, 30))
	res := l.Source()
	c.Assert(string(res), Equals, "<g style=\"stroke:rgb(10,20,30)\"/>\n")
}

func (s GROUPTestsSuite) Test_group_Append_Title(c *C) {
	//c.Skip("Not now")

	l := Group()
	t := Title("WWW")
	l = l.Append(t)
	res := l.Source()
	c.Assert(string(res), Equals, "<g >\n<title>WWW</title></g>\n")
}

func (s GROUPTestsSuite) Test_group_Style(c *C) {
	//c.Skip("Not now")

	l := Group()
	st := style.Style().FillRGB(1, 2, 3)
	l = l.Style(st).Append(Title("WWW"))
	c.Assert(string(l.Source()), Equals, "<g style=\"fill:rgb(1,2,3)\">\n<title>WWW</title></g>\n")
}

func (s GROUPTestsSuite) Test_group_Transform(c *C) {
	//c.Skip("Not now")

	l := Group()
	st := transform.Transform().Matrix(12, 34, 45, 34.5, 212, 4.45)
	l = l.Transform(st).Append(Title("WWW"))
	c.Assert(string(l.Source()), Equals, "<g transform=\"matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)\">\n<title>WWW</title></g>\n")
}
