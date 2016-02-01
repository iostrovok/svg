package transform

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestTRANSLATE(t *testing.T) {
	TestingT(t)
}

type TRANSLATETestsSuite struct{}

var _ = Suite(&TRANSLATETestsSuite{})

func (s TRANSLATETestsSuite) Test_Translate(c *C) {
	//c.Skip("Not now")
	t1 := Translate(1, 2)
	c.Assert(t1, NotNil)
}

func (s TRANSLATETestsSuite) Test_TranslateX_Source(c *C) {
	//c.Skip("Not now")
	c.Assert(Translate(12).Source(), Equals, "translate(12.00)")
}

func (s TRANSLATETestsSuite) Test_TranslateXY_Source(c *C) {
	//c.Skip("Not now")
	c.Assert(Translate(12, 34).Source(), Equals, "translate(12.00, 34.00)")
}
