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
	obtained := cSL(`<g/>`)

	res := Tag("g").Source()
	c.Assert(cSL(res), Equals, obtained)

	obtained = cSL(`<g x="1"/>`)

	res = Tag("g", date_tag_list).Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s TAGTestsSuite) Test_Attr(c *C) {
	//c.Skip("Not now")
	obtained := cSL(`<g xxx="123"/>`)

	l := Tag("g").Attr("xxx", "123")
	c.Assert(string(cSL(l.Source())), Equals, obtained)
}

func (s TAGTestsSuite) Test_Append_Title(c *C) {
	//c.Skip("Not now")
	obtained := cSL(`<g><title>WWW</title></g>`)

	t := Title("WWW")
	res := Tag("g").Append(t).Source()
	c.Assert(cSL(res), Equals, obtained)
}

func (s TAGTestsSuite) Test_Style(c *C) {
	//c.Skip("Not now")
	obtained := cSL(`<g style="fill:rgb(1,2,3)"/>`)

	st := style.Style().FillRGB(1, 2, 3)

	l := Tag("g").Style(st)
	c.Assert(cSL(l.Source()), Equals, obtained)
}

func (s TAGTestsSuite) Test_Transform(c *C) {
	//c.Skip("Not now")

	obtained := cSL(`<g x="1" transform="matrix(12.00, 34.00, 45.00, 34.50, 212.00, 4.45)"/>`)

	st := transform.Transform().Matrix(12, 34, 45, 34.5, 212, 4.45)
	l := Tag("g", date_tag_list).Transform(st)
	c.Assert(cSL(l.Source()), Equals, obtained)
}

func (s TAGTestsSuite) Test_ID(c *C) {
	//c.Skip("Not now")

	check := cSL(`<g id="id-1" x="1"/>`)

	t := Tag("g", date_tag_list)
	t = t.ID("id-1")

	c.Assert("id-1", Equals, t.GetID())

	res := t.Source()
	c.Assert(cSL(res), Equals, check)
}
