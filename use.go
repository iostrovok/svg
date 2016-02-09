package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
)

/*
‘class’
‘style’
‘externalResourcesRequired’
‘transform’
‘x’
‘y’
‘width’
‘height’
‘xlink:href’
*/

const (
	useTag    = `<use xlink:href="%s" `
	useEndTag = `</use>`
)

type USE struct {
	//iNode
	*node
	href string
	use  string
}

// Constructor of "use" object
func Use(s ...style.STYLE) *USE {
	return &USE{
		node: newNode(s...),
	}
}

// XYWH sets  X coordinate for element.
func (t *USE) Href(href string) *USE {
	t.href = href
	return t
}

// Source() returns svg implementation of USE element
func (use *USE) Source() string {
	if use.href == "" {
		return ""
	}
	body := fmt.Sprintf(useTag, use.href)
	return _Source(use, body, useEndTag, use.node.inner)
}
