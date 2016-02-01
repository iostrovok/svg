package svg

import (
	"github.com/iostrovok/svg/style"
)

type iNode interface {
	Source() string
	AppendIn(iNode)
}

type node struct {
	iNode
	st    style.STYLE
	inner []iNode
}

func newNode(s ...style.STYLE) *node {
	return &node{
		inner: []iNode{},
		st:    mstyle(s...),
	}
}

func mstyle(s ...style.STYLE) style.STYLE {
	if len(s) == 0 {
		return style.STYLE{}
	}
	return s[0]
}

// Style sets the "style.STYLE" object
func (n *node) Style(st style.STYLE) {
	n.st = st
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (n *node) Append(in ...iNode) {
	n.inner = append(n.inner, in...)
}
