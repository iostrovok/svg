package svg

import (
	"strings"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

type TAG struct {
	iNode
	*node
	tag  string
	attr map[string]string
}

// Constructor of "tag" object
func Tag(tag string, attr ...map[string]string) *TAG {
	t := &TAG{
		node: newNode(),
		tag:  tag,
		attr: map[string]string{},
	}

	if len(attr) > 0 {
		t.attr = attr[0]
	}
	return t

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

// Attr sets the "transform.TRANSFORM" object
func (tag *TAG) Attr(k, v string) *TAG {
	tag.attr[k] = v
	return tag
}

// Source() returns svg implementation of TAG element
func (tag *TAG) Source() string {
	out := []string{`<` + tag.tag}
	has := map[string]bool{}

	for k, v := range tag.attr {
		if k != "" && v != "" {
			out = append(out, k+`="`+v+`"`)
			has[strings.ToLower(k)] = true
		}
	}

	if !has[`style`] {
		if s := tag.node.styleSource(); s != "" {
			out = append(out, s)
		}
	}

	if !has[`transform`] {
		if s := tag.node.transSource(); s != "" {
			out = append(out, s)
		}
	}

	body := strings.Join(out, ` `)
	return _Source(body, `</`+tag.tag+`>`, tag.node.inner)
}
