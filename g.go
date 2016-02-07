package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

type GROUP struct {
	iNode
	*node

	id string
}

const (
	groupTag    = `<g %s`
	groupEndTag = `</g>`
)

// Constructor of "group" object
func Group(s ...style.STYLE) *GROUP {
	return &GROUP{
		node: newNode(s...),
	}
}

// ID(string) set element id.
func (group *GROUP) ID(id string) *GROUP {
	group.node.id = id
	return group
}

// GetID() returns lement id.
func (group *GROUP) GetID() string {
	return group.node.id
}

// Attr adds any user attribute.
func (group *GROUP) Attr(attr, values string) *GROUP {
	group.node.attrs[attr] = values
	return group
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (group *GROUP) Append(nodes ...iNode) *GROUP {
	group.node.Append(nodes...)
	return group
}

// AppendTo is interface function
func (group *GROUP) AppendTo(n iNode) *GROUP {
	n.appendIn(group)
	return group
}

// Style sets the "style.STYLE" object
func (group *GROUP) Style(st style.STYLE) *GROUP {
	group.node.Style(st)
	return group
}

// Transform sets the "transform.TRANSFORM" object
func (group *GROUP) Transform(tr transform.TRANSFORM) *GROUP {
	group.node.Transform(tr)
	return group
}

// Source() returns svg implementation of GROUP element
func (group *GROUP) Source() string {
	body := fmt.Sprintf(groupTag, group.node.mSource())
	return _Source(group, body, groupEndTag, group.node.inner)
}
