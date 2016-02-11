package svg

import (
	"fmt"
	"strings"

	"github.com/iostrovok/svg/style"
)

const (
	textTag    = `<text `
	textEndTag = `</text>`

	tspanTag    = `<tspan %s`
	tspanEndTag = `</tspan>`

	trefTag    = ` <tref xlink:href="%s"`
	trefEndTag = `</tspan>`
)

type textType int

const (
	isText  textType = iota
	isTSpan textType = iota
	isTRef  textType = iota
)

type textBody struct {
	//iNode
	*node
	str string
}

// Source() returns svg implementation of TEXT element
func (text textBody) Source() string {
	return text.str
}

// Source() returns svg implementation of TEXT element
func (text textBody) GetID() string {
	return ""
}

// GetID() returns element id class for string.
func (t textBody) GetClass() string {
	return ""
}

type TEXT struct {
	iNode
	*node

	typ textType

	id string

	x, y   float64
	dim_xy string
	has_xy bool

	dx, dy  float64
	dim_dxy string

	anchor       string
	href         string
	rotate       []float64
	textLength   float64
	lengthAdjust float64
}

// Text is the constructor of text object
func Text(s ...style.STYLE) *TEXT {
	return &TEXT{
		typ:  isText,
		node: newNode(s...),
	}
}

// Tspan is the constructor of tspan object
func Tspan(s ...style.STYLE) *TEXT {
	return &TEXT{
		typ:  isTSpan,
		node: newNode(s...),
	}
}

// Tref is the constructor of tref object
func Tref(s ...style.STYLE) *TEXT {
	return &TEXT{
		typ:  isTRef,
		node: newNode(s...),
	}
}

// Source() returns svg implementation of TEXT element
func (text *TEXT) Source() string {
	body := ""
	end := ""
	switch text.typ {
	case isTRef:
		end = trefEndTag
		body = fmt.Sprintf(trefTag, text.href)
	case isTSpan:
		end = tspanEndTag
		body = fmt.Sprintf(tspanTag, text.href)
	default:
		body = textTag
		end = textEndTag
	}

	body = text.bodyTags(body)

	return _Source(text, body, end, text.node.inner)
}

func (text *TEXT) bodyTags(body string) string {

	res := []string{body}
	if text.dx != 0 || text.dy != 0 {
		res = append(res, fmt.Sprintf(`dx="%s" dy="%s"`, printNumber(text.dx, text.dim_dxy), printNumber(text.dy, text.dim_dxy)))
	}
	if text.has_xy {
		res = append(res, fmt.Sprintf(`x="%s" y="%s"`, printNumber(text.x, text.dim_xy), printNumber(text.y, text.dim_xy)))
	}
	if len(text.rotate) > 0 {
		a := make([]string, len(text.rotate))
		for i, v := range text.rotate {
			a[i] = printNumber(v)
		}
		res = append(res, fmt.Sprintf(`rotate="%s"`, strings.Join(a, ",")))
	}

	return strings.Join(res, " ")
}

// Style sets the content of text/tspan object
func (text *TEXT) String(str string) *TEXT {

	t := &textBody{
		str: str,
	}

	text.Append(t)
	return text
}

// Style sets the content of text/tspan object
func (text *TEXT) AddString(str string) *TEXT {

	last := len(text.node.inner)

	if last == 0 {
		return text.String(str)
	}

	switch v := text.node.inner[last-1].(type) {
	case textBody:
		v.str += str
	default:
		text.String(str)
	}

	return text
}

// XY sets the absolute coordinates of object
func (text *TEXT) XY(x, y float64, dim ...string) *TEXT {
	text.has_xy = true
	text.x = x
	text.y = y
	if len(dim) > 0 {
		text.dim_xy = dim[0]
	}
	return text
}

// XY sets the relative coordinates of object
func (text *TEXT) DXY(dx, dy float64, dim ...string) *TEXT {
	text.dx = dx
	text.dy = dy
	if len(dim) > 0 {
		text.dim_dxy = dim[0]
	}
	return text
}

// Rotate set the rotate attribute
func (text *TEXT) Rotate(r ...float64) *TEXT {
	text.rotate = r
	return text
}

// Href set the href attribute
func (text *TEXT) Href(str string) *TEXT {
	text.href = str
	return text
}

func (text *TEXT) Anchor(str string) *TEXT {
	text.anchor = str
	return text
}
