package svg

import (
	"fmt"
	"strings"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

const (
	classLine = `class="%s"`
)

type iNode interface {
	Source() string
	GetID() string
	appendIn(iNode)
	nodes() *node
}

type node struct {
	iNode
	class, id string
	st        style.STYLE
	tr        transform.TRANSFORM
	inner     []iNode
	attrs     map[string]string
}

func newNode(st ...style.STYLE) *node {
	n := &node{
		inner: []iNode{},
		st:    style.STYLE{},
		tr:    transform.TRANSFORM{},
		attrs: map[string]string{},
	}
	if len(st) != 0 {
		n.st = st[0]
	}
	return n
}

// Style sets the "style.STYLE" object
func (n *node) mSource() string {
	out := []string{}

	if st := n.styleSource(); st != "" {
		out = append(out, st)
	}

	if tr := n.transSource(); tr != "" {
		out = append(out, tr)
	}

	if class := n.GetClass(); class != "" {
		out = append(out, fmt.Sprintf(classLine, class))
	}

	return strings.Join(out, " ")
}

// ID(string) set element id.
func (n *node) ID(id string) {
	n.id = id
}

// GetID() returns element id.
func (n *node) GetID() string {
	return n.id
}

// Class(string) set element class.
func (n *node) Class(class string) {
	n.class = class
}

// GetID() returns element id class for string.
func (n *node) GetClass() string {
	return n.class
}

// Style sets the "style.STYLE" object
func (n *node) styleSource() string {
	return n.st.Source()
}

func (n *node) transSource() string {
	return n.tr.Source()
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

// attrSource() returns attrebutes of tag as line
func (n *node) attrSource() string {
	out := []string{}
	has := map[string]bool{}

	for k, v := range n.attrs {
		if k != "" && v != "" {
			out = append(out, k+`="`+v+`"`)
			has[strings.ToLower(k)] = true
		}
	}

	if !has[`style`] {
		if s := n.styleSource(); s != "" {
			out = append(out, s)
		}
	}

	if !has[`transform`] {
		if s := n.transSource(); s != "" {
			out = append(out, s)
		}
	}

	if !has[`class`] {
		if s := n.GetClass(); s != "" {
			out = append(out, fmt.Sprintf(classLine, s))
		}
	}

	return strings.Join(out, ` `)
}
