package svg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

const (
	textTag    = `<text %s`
	textEndTag = `</text>`

	tspanTag    = `<tspan %s %s`
	tspanEndTag = `</tspan>`

	trefTag    = ` <tref xlink:href="%s" %s`
	trefEndTag = `</tspan>`
)

type textType int

const (
	isText  textType = iota
	isTSpan textType = iota
	isTRef  textType = iota
)

type textBody struct {
	iNode
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

// // Attr adds any user attribute.
// func (text *textBody) GetAttr() map[string]string {
// }

type TEXT struct {
	iNode
	*node

	typ textType

	id string

	x, y   int
	dim_xy string
	has_xy bool

	dx, dy  int
	dim_dxy string

	anchor       string
	href         string
	rotate       []int
	textLength   int
	lengthAdjust int
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

// ID(string) set element id.
func (text *TEXT) ID(id string) *TEXT {
	text.node.id = id
	return text
}

// GetID() returns lement id.
func (text *TEXT) GetID() string {
	return text.node.id
}

// Attr adds any user attribute.
func (text *TEXT) Attr(attr, values string) *TEXT {
	text.node.attrs[attr] = values
	return text
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (text *TEXT) Append(nodes ...iNode) *TEXT {
	text.node.Append(nodes...)
	return text
}

// AppendTo is interface function
func (text *TEXT) AppendTo(n iNode) *TEXT {
	n.appendIn(text)
	return text
}

// Style sets the "style.STYLE" object
func (text *TEXT) Style(st style.STYLE) *TEXT {
	text.node.Style(st)
	return text
}

// Transform sets the "transform.TRANSFORM" object
func (text *TEXT) Transform(tr transform.TRANSFORM) *TEXT {
	text.node.Transform(tr)
	return text
}

// Source() returns svg implementation of TEXT element
func (text *TEXT) Source() string {
	body := ""
	end := ""
	switch text.typ {
	case isTRef:
		end = trefEndTag
		body = fmt.Sprintf(trefTag, text.href, text.node.mSource())
	case isTSpan:
		end = tspanEndTag
		body = fmt.Sprintf(tspanTag, text.href, text.node.mSource())
	default:
		body = fmt.Sprintf(textTag, text.node.mSource())
		end = textEndTag
	}

	body = text.bodyTags(body)

	return _Source(text, body, end, text.node.inner)
}

func (text *TEXT) bodyTags(body string) string {
	res := []string{body}
	if text.dx != 0 {
		res = append(res, fmt.Sprintf(`dx="%d%s"`, text.dx, text.dim_dxy))
	}
	if text.dy != 0 {
		res = append(res, fmt.Sprintf(`dy="%d%s"`, text.dy, text.dim_dxy))
	}
	if text.has_xy {
		res = append(res, fmt.Sprintf(`x="%d%s" y="%d%s"`, text.x, text.dim_xy, text.y, text.dim_xy))
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
func (text *TEXT) XY(x, y int, dim ...string) *TEXT {
	text.has_xy = true
	text.x = x
	text.y = y
	if len(dim) > 0 {
		text.dim_xy = dim[0]
	}
	return text
}

// XY sets the relative coordinates of object
func (text *TEXT) DX(dx, dy int, dim ...string) *TEXT {
	text.x = dx
	text.y = dy
	if len(dim) > 0 {
		text.dim_dxy = dim[0]
	}
	return text
}

// Rotate set the rotate attribute
func (text *TEXT) Rotate(r ...int) *TEXT {
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
