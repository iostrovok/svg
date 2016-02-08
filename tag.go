package svg

import (
	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

type TAG struct {
	//iNode
	*node
	id  string
	tag string
}

// Constructor of "tag" object
func Tag(tag string, attr ...map[string]string) *TAG {
	t := &TAG{
		node: newNode(),
		tag:  tag,
	}

	if len(attr) > 0 {
		for k, v := range attr[0] {
			t.Attr(k, v)
		}
	}

	return t
}

func (t *TAG) nodes() *node {
	return t.node
}

// ID(string) set element id.
func (t *TAG) ID(id string) *TAG {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *TAG) GetID() string {
	return t.node.id
}

// Class(string) set element class.
func (t *TAG) Class(id string) *TAG {
	t.node.class = id
	return t
}

// GetID() returns element id class for string.
func (t *TAG) GetClass() string {
	return t.node.class
}

// Attr adds any user attribute.
func (tag *TAG) Attr(attr, values string) *TAG {
	tag.node.attrs[attr] = values
	return tag
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (tag *TAG) Append(nodes ...iNode) *TAG {
	tag.node.Append(nodes...)
	return tag
}

// AppendTo is interface function
func (tag *TAG) AppendTo(n iNode) *TAG {
	n.appendIn(tag)
	return tag
}

// Style sets the "style.STYLE" object
func (tag *TAG) Style(st style.STYLE) *TAG {
	tag.node.Style(st)
	return tag
}

// Transform sets the "transform.TRANSFORM" object
func (tag *TAG) Transform(tr transform.TRANSFORM) *TAG {
	tag.node.Transform(tr)
	return tag
}

// Source() returns svg implementation of TAG element
func (tag *TAG) Source() string {
	body := `<` + tag.tag
	return _Source(tag, body, `</`+tag.tag+`>`, tag.node.inner)
}
