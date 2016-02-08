package svg

import (
	"github.com/iostrovok/svg/style"
)

type GROUP struct {
	//iNode
	*node

	id string
}

const (
	groupTag    = `<g`
	groupEndTag = `</g>`
)

// Constructor of "group" object
func Group(s ...style.STYLE) *GROUP {
	return &GROUP{
		node: newNode(s...),
	}
}

// Source() returns svg implementation of GROUP element
func (group *GROUP) Source() string {
	return _Source(group, groupTag, groupEndTag, group.node.inner)
}
