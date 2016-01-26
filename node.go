package svg

import (
	"github.com/iostrovok/svg/style"
)

type INode interface {
	Source() string
	AppendIn(INode)
}

type Node struct {
	INode
	st    style.STYLE
	inner []INode
}

func NewNode(s ...style.STYLE) *Node {
	return &Node{
		inner: []INode{},
		st:    mstyle(s...),
	}
}

func mstyle(s ...style.STYLE) style.STYLE {
	if len(s) == 0 {
		return style.STYLE{}
	}
	return s[0]
}

// Setter
func (n *Node) Style(s style.STYLE) {
	n.st = s
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (n *Node) Append(in ...INode) {
	n.inner = append(n.inner, in...)
}
