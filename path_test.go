package svg

import (
	"fmt"
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

	l := Path()
	c.Assert(l, NotNil)
	l2 := Path(style.Style())
	c.Assert(l2, NotNil)
}

func (s PATHTestsSuite) Test_path2(c *C) {
	//c.Skip("Not now")

	l := Path(style.Style().Stroke(10, 20, 30))
	res := l.Source()
	fmt.Printf("%s\n", res)
	c.Assert(string(res), Equals, "<path d=\"\" style=\"stroke:rgb(10,20,30)\"/>\n")
}
