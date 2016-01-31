package svg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iostrovok/svg/style"
)

const (
	textTag    = `<text %s`
	textEndTag = `</text>`

	tspanTag    = `<tspan %s %s`
	tspanEndTag = `</tspan>`

	trefTag    = ` <tref xlink:href="%s" %s`
	trefEndTag = `</tspan>`
)

type TextType int

const (
	isText  TextType = iota
	isTSpan TextType = iota
	isTRef  TextType = iota
)

type TEXTBody struct {
	INode
	*Node
	str string
}

// Source() returns svg implementation of TEXT element
func (text *TEXTBody) Source() string {
	return text.str
}

type TEXT struct {
	INode
	*Node

	typ TextType

	x, y   int
	has_xy bool

	dx, dy int

	anchor       string
	href         string
	rotate       []int
	textLength   int
	lengthAdjust int
}

// Constructor
func Text(s ...style.STYLE) *TEXT {
	return &TEXT{
		typ:  isText,
		Node: NewNode(s...),
	}
}

// Constructor
func Tspan(s ...style.STYLE) *TEXT {
	return &TEXT{
		typ:  isTSpan,
		Node: NewNode(s...),
	}
}

// Constructor
func Tref(s ...style.STYLE) *TEXT {
	return &TEXT{
		typ:  isTRef,
		Node: NewNode(s...),
	}
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (text *TEXT) Append(nodes ...INode) *TEXT {
	text.Node.Append(nodes...)
	return text
}

func (text *TEXT) AppendTo(n INode) *TEXT {
	n.AppendIn(text)
	return text
}

// Setter
func (text *TEXT) Style(st style.STYLE) *TEXT {
	text.Node.Style(st)
	return text
}

// Source() returns svg implementation of TEXT element
func (text *TEXT) Source() string {
	body := ""
	end := ""
	switch text.typ {
	case isTRef:
		end = trefEndTag
		body = fmt.Sprintf(trefTag, text.href, text.Node.st.Source())
	case isTSpan:
		end = tspanEndTag
		body = fmt.Sprintf(tspanTag, text.href, text.Node.st.Source())
	default:
		body = fmt.Sprintf(textTag, text.Node.st.Source())
		end = textEndTag
	}

	body = text.bodyTags(body)

	return _Source(body, end, text.Node.inner)
}

func (text *TEXT) bodyTags(body string) string {
	res := []string{body}
	if text.dx != 0 {
		res = append(res, fmt.Sprintf(`dx="%d"`, text.dx))
	}
	if text.dy != 0 {
		res = append(res, fmt.Sprintf(`dy="%d"`, text.dy))
	}
	if text.has_xy {
		res = append(res, fmt.Sprintf(`x="%d" y="%d"`, text.x, text.y))
	}
	if len(text.rotate) > 0 {
		a := make([]string, len(text.rotate))
		for i, v := range text.rotate {
			a[i] = strconv.Itoa(v)
		}
		res = append(res, fmt.Sprintf(`rotate="%s"`, strings.Join(a, ",")))
	}

	return strings.Join(res, " ")
}

func (text *TEXT) String(str string) *TEXT {

	t := &TEXTBody{
		str: str,
	}

	text.Append(t)
	return text
}

func (text *TEXT) XY(x, y int) *TEXT {
	text.has_xy = true
	text.x = x
	text.y = y
	return text
}

func (text *TEXT) DX(dx, dy int) *TEXT {
	text.x = dx
	text.y = dy
	return text
}

func (text *TEXT) Rotate(r ...int) *TEXT {
	text.rotate = r
	return text
}

func (text *TEXT) Href(str string) *TEXT {
	text.href = str
	return text
}

func (text *TEXT) Anchor(str string) *TEXT {
	text.anchor = str
	return text
}
