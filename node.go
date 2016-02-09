package svg

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/iostrovok/svg/style"
	"github.com/iostrovok/svg/transform"
)

var (
	idCounter      = 0
	coordinateKeys = map[string]bool{
		`x`:      true,
		`y`:      true,
		`width`:  true,
		`height`: true,
		`cx`:     true,
		`cy`:     true,
		`rx`:     true,
		`ry`:     true,
		`r`:      true,
	}
)

const (
	classLine        = `class="%s"`
	coordinateDim    = `%s="%.2f%s"`
	coordinateDimInt = `%s="%d%s"`
)

type coordinate struct {
	c        float64
	typ, dim string
}

// XYWH sets `coordinates` for element.
func (c *coordinate) Source() string {
	if c == nil {
		return ""
	}

	if math.Floor(c.c) == c.c {
		return fmt.Sprintf(coordinateDimInt, c.typ, int(c.c), c.dim)
	}

	return fmt.Sprintf(coordinateDim, c.typ, c.c, c.dim)
}

type iNode interface {
	Source() string
	GetID() string
	genID()
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
	xywh      map[string]*coordinate
}

func newNode(st ...style.STYLE) *node {
	n := &node{
		inner: []iNode{},
		st:    style.STYLE{},
		tr:    transform.TRANSFORM{},
		attrs: map[string]string{},
		xywh:  map[string]*coordinate{},
	}
	if len(st) != 0 {
		n.st = st[0]
	}
	return n
}

// Style sets the "style.STYLE" object
func (n *node) mSource() string {

	out := n.sourceXYWH([]string{}, map[string]bool{})

	for _, v := range n.xywh {
		if s := v.Source(); s != "" {
			out = append(out, s)
		}
	}

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

// genID() creates element id.
func (n *node) genID() {
	idCounter++
	n.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
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

func (n *node) sourceXYWH(out []string, has map[string]bool) []string {

	add := []string{}

	for k, v := range n.xywh {
		if !has[k] {
			if s := v.Source(); s != "" {
				add = append(add, s)
			}
		}
	}

	sort.Strings(add)

	return append(out, add...)
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

	out = n.sourceXYWH(out, has)

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

// XYWH sets `coordinates` for element.
func (n *node) XYWH(t string, c float64, dim ...string) {
	x := &coordinate{
		c:   c,
		typ: t,
	}
	if len(dim) > 0 {
		x.dim = dim[0]
	}

	if coordinateKeys[t] {
		n.xywh[t] = x
	}
}
