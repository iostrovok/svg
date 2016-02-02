package svg

import (
	"strings"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

type iNode interface {
	Source() string
	appendIn(iNode)
}

type node struct {
	iNode
	st    style.STYLE
	tr    transform.TRANSFORM
	inner []iNode
}

func newNode(st ...style.STYLE) *node {
	n := &node{
		inner: []iNode{},
		st:    style.STYLE{},
		tr:    transform.TRANSFORM{},
	}
	if len(st) != 0 {
		n.st = st[0]
	}
	return n
}

// Style sets the "style.STYLE" object
func (n *node) mSource() string {
	out := []string{}

	if st := n.st.Source(); st != "" {
		out = append(out, st)
	}

	if tr := n.tr.Source(); tr != "" {
		out = append(out, tr)
	}

	return strings.Join(out, " ")
}

// Style sets the "style.STYLE" object
func (n *node) Style(st style.STYLE) {
	n.st = st
}

// Style sets the "style.STYLE" object
func (n *node) Transform(tr transform.TRANSFORM) {
	n.tr = tr
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (n *node) Append(in ...iNode) {
	n.inner = append(n.inner, in...)
}
