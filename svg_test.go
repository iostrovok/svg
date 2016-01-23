package svg

import (
	"fmt"

	. "gopkg.in/check.v1"
	"testing"
)

func TestSVG(t *testing.T) {
	TestingT(t)
}

type SVGTestsSuite struct{}

var _ = Suite(&SVGTestsSuite{})

func (s SVGTestsSuite) Test_new(c *C) {
	//c.Skip("Not now")

	// Empty SVG
	m := New(100, 100)
	c.Assert(m, NotNil)
}

func (s SVGTestsSuite) Test_SVG_Source(c *C) {
	//c.Skip("Not now")

	m := New(100, 100)
	fmt.Printf("%s\n", m.Source())
	t := len(m.Source()) > 10
	c.Assert(t, Equals, true)
}
