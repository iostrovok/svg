package svg

import (
	. "gopkg.in/check.v1"
	"testing"

	"github.com/iostrovok/svg/style"
)

func TestTEXT(t *testing.T) {
	TestingT(t)
}

type TEXTTestsSuite struct{}

var _ = Suite(&TEXTTestsSuite{})

func (s TEXTTestsSuite) Test_text(c *C) {
	//c.Skip("Not now")

	t1 := Text()
	c.Assert(t1, NotNil)
	t2 := Text(style.Style())
	c.Assert(t2, NotNil)
}

func (s TEXTTestsSuite) Test_text2(c *C) {
	//c.Skip("Not now")

	check := `<text style="font-family:Verdana;font-size:55;fill:blue" x="250" y="150">
Hello, out there
</text>
`
	f := style.Font("font-family:Verdana;font-size:55")
	text := Text(style.Style().Fill("blue").Font(f))
	text = text.XY(250, 150).String("Hello, out there\n")
	res := text.Source()
	c.Assert(string(res), Equals, check)
}

// func (s TEXTTestsSuite) Test_text_Append_Title(c *C) {
// 	//c.Skip("Not now")

// 	l := Text(1, 2, 3, 4)
// 	t := Title("WWW")
// 	l = l.Append(t)
// 	res := l.Source()
// 	c.Assert(string(res), Equals, "<text x1=\"1\" y1=\"2\" x2=\"3\" y2=\"4\"  >\n<title>WWW</title></text>\n")
// }

// func (s TEXTTestsSuite) Test_text_Style(c *C) {
// 	//c.Skip("Not now")

// 	l := Text(1, 2, 3, 4)
// 	st := style.Style().FillRGB(1, 2, 3)
// 	l = l.Style(st).Append(Title("WWW"))
// 	c.Assert(string(l.Source()), Equals, "<text x1=\"1\" y1=\"2\" x2=\"3\" y2=\"4\" style=\"fill:rgb(1,2,3)\" >\n<title>WWW</title></text>\n")
// }
