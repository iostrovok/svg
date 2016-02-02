package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

var date_tag_list map[string]string = map[string]string{
	"x": "1",
}

func TestTAG(t *testing.T) {
	TestingT(t)
}

type TAGTestsSuite struct{}

var _ = Suite(&TAGTestsSuite{})

func (s TAGTestsSuite) Test_tag(c *C) {
	//c.Skip("Not now")

	l := Tag("mytag")
	c.Assert(l, NotNil)
	l2 := Tag("mytag", date_tag_list)
	c.Assert(l2, NotNil)
}

func (s TAGTestsSuite) Test_tag2(c *C) {
	//c.Skip("Not now")

	res := Tag("g").Source()
	c.Assert(res, Equals, "<g/>\n")

	res = Tag("g", date_tag_list).Source()
	c.Assert(res, Equals, "<g x=\"1\"/>\n")
}

func (s TAGTestsSuite) Test_tag_Attr(c *C) {
	//c.Skip("Not now")

	l := Tag("g").Attr("xxx", "123")
	c.Assert(string(l.Source()), Equals, "<g xxx=\"123\"/>\n")
}

func (s TAGTestsSuite) Test_tag_Append_Title(c *C) {
	//c.Skip("Not now")

	t := Title("WWW")
	res := Tag("g").Append(t).Source()
	c.Assert(res, Equals, "<g>\n<title>WWW</title></g>\n")
}

func (s TAGTestsSuite) Test_tag_Style(c *C) {
	//c.Skip("Not now")

	st := style.Style().FillRGB(1, 2, 3)

	l := Tag("g").Style(st)
	c.Assert(l.Source(), Equals, "<g style=\"fill:rgb(1,2,3)\"/>\n")
}

func (s TAGTestsSuite) Test_tag_Transform(c *C) {
	//c.Skip("Not now")
	st := transform.Transform().Matrix(12, 34, 45, 34.5, 212, 4.45)
	l := Tag("g", date_tag_list).Transform(st)
	c.Assert(string(l.Source()), Equals, "<g x=\"1\" transform=\"matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)\"/>\n")
}
