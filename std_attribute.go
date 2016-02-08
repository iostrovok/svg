package svg
import "github.com/iostrovok/svg/transform"
import "github.com/iostrovok/svg/style"

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (t *GROUP) Append(nodes ...iNode) *GROUP {
	t.node.Append(nodes...)
	return t
}

// AppendTo is interface function
func (t *GROUP) AppendTo(n iNode) *GROUP {
	n.appendIn(t)
	return t
}

// Attr adds any user attribute.
func (t *GROUP) Attr(attr, values string) *GROUP {
	t.node.attrs[attr] = values
	return t
}

// Class(string) set element class.
func (t *GROUP) Class(id string) *GROUP {
	t.node.class = id
	return t
}

// GetID() returns element id class for string.
func (t *GROUP) GetClass() string {
	return t.node.class
}

// ID(string) set element id.
func (t *GROUP) ID(id string) *GROUP {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *GROUP) GetID() string {
	return t.node.id
}

// nodes returns inner node object
func (t *GROUP) nodes() *node {
	return t.node
}

// Style sets the "style.STYLE" object
func (t *GROUP) Style(st style.STYLE) *GROUP {
	t.node.Style(st)
	return t
}

// Transform sets the "transform.TRANSFORM" object
func (t *GROUP) Transform(tr transform.TRANSFORM) *GROUP {
	t.node.Transform(tr)
	return t
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (t *LINE) Append(nodes ...iNode) *LINE {
	t.node.Append(nodes...)
	return t
}

// AppendTo is interface function
func (t *LINE) AppendTo(n iNode) *LINE {
	n.appendIn(t)
	return t
}

// Attr adds any user attribute.
func (t *LINE) Attr(attr, values string) *LINE {
	t.node.attrs[attr] = values
	return t
}

// Class(string) set element class.
func (t *LINE) Class(id string) *LINE {
	t.node.class = id
	return t
}

// GetID() returns element id class for string.
func (t *LINE) GetClass() string {
	return t.node.class
}

// ID(string) set element id.
func (t *LINE) ID(id string) *LINE {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *LINE) GetID() string {
	return t.node.id
}

// nodes returns inner node object
func (t *LINE) nodes() *node {
	return t.node
}

// Style sets the "style.STYLE" object
func (t *LINE) Style(st style.STYLE) *LINE {
	t.node.Style(st)
	return t
}

// Transform sets the "transform.TRANSFORM" object
func (t *LINE) Transform(tr transform.TRANSFORM) *LINE {
	t.node.Transform(tr)
	return t
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (t *PATH) Append(nodes ...iNode) *PATH {
	t.node.Append(nodes...)
	return t
}

// AppendTo is interface function
func (t *PATH) AppendTo(n iNode) *PATH {
	n.appendIn(t)
	return t
}

// Attr adds any user attribute.
func (t *PATH) Attr(attr, values string) *PATH {
	t.node.attrs[attr] = values
	return t
}

// Class(string) set element class.
func (t *PATH) Class(id string) *PATH {
	t.node.class = id
	return t
}

// GetID() returns element id class for string.
func (t *PATH) GetClass() string {
	return t.node.class
}

// ID(string) set element id.
func (t *PATH) ID(id string) *PATH {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *PATH) GetID() string {
	return t.node.id
}

// nodes returns inner node object
func (t *PATH) nodes() *node {
	return t.node
}

// Style sets the "style.STYLE" object
func (t *PATH) Style(st style.STYLE) *PATH {
	t.node.Style(st)
	return t
}

// Transform sets the "transform.TRANSFORM" object
func (t *PATH) Transform(tr transform.TRANSFORM) *PATH {
	t.node.Transform(tr)
	return t
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (t *RECT) Append(nodes ...iNode) *RECT {
	t.node.Append(nodes...)
	return t
}

// AppendTo is interface function
func (t *RECT) AppendTo(n iNode) *RECT {
	n.appendIn(t)
	return t
}

// Attr adds any user attribute.
func (t *RECT) Attr(attr, values string) *RECT {
	t.node.attrs[attr] = values
	return t
}

// Class(string) set element class.
func (t *RECT) Class(id string) *RECT {
	t.node.class = id
	return t
}

// GetID() returns element id class for string.
func (t *RECT) GetClass() string {
	return t.node.class
}

// ID(string) set element id.
func (t *RECT) ID(id string) *RECT {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *RECT) GetID() string {
	return t.node.id
}

// nodes returns inner node object
func (t *RECT) nodes() *node {
	return t.node
}

// Style sets the "style.STYLE" object
func (t *RECT) Style(st style.STYLE) *RECT {
	t.node.Style(st)
	return t
}

// Transform sets the "transform.TRANSFORM" object
func (t *RECT) Transform(tr transform.TRANSFORM) *RECT {
	t.node.Transform(tr)
	return t
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (t *TEXT) Append(nodes ...iNode) *TEXT {
	t.node.Append(nodes...)
	return t
}

// AppendTo is interface function
func (t *TEXT) AppendTo(n iNode) *TEXT {
	n.appendIn(t)
	return t
}

// Attr adds any user attribute.
func (t *TEXT) Attr(attr, values string) *TEXT {
	t.node.attrs[attr] = values
	return t
}

// Class(string) set element class.
func (t *TEXT) Class(id string) *TEXT {
	t.node.class = id
	return t
}

// GetID() returns element id class for string.
func (t *TEXT) GetClass() string {
	return t.node.class
}

// ID(string) set element id.
func (t *TEXT) ID(id string) *TEXT {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *TEXT) GetID() string {
	return t.node.id
}

// nodes returns inner node object
func (t *TEXT) nodes() *node {
	return t.node
}

// Style sets the "style.STYLE" object
func (t *TEXT) Style(st style.STYLE) *TEXT {
	t.node.Style(st)
	return t
}

// Transform sets the "transform.TRANSFORM" object
func (t *TEXT) Transform(tr transform.TRANSFORM) *TEXT {
	t.node.Transform(tr)
	return t
}

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (t *USE) Append(nodes ...iNode) *USE {
	t.node.Append(nodes...)
	return t
}

// AppendTo is interface function
func (t *USE) AppendTo(n iNode) *USE {
	n.appendIn(t)
	return t
}

// Attr adds any user attribute.
func (t *USE) Attr(attr, values string) *USE {
	t.node.attrs[attr] = values
	return t
}

// Class(string) set element class.
func (t *USE) Class(id string) *USE {
	t.node.class = id
	return t
}

// GetID() returns element id class for string.
func (t *USE) GetClass() string {
	return t.node.class
}

// ID(string) set element id.
func (t *USE) ID(id string) *USE {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *USE) GetID() string {
	return t.node.id
}

// nodes returns inner node object
func (t *USE) nodes() *node {
	return t.node
}

// Style sets the "style.STYLE" object
func (t *USE) Style(st style.STYLE) *USE {
	t.node.Style(st)
	return t
}

// Transform sets the "transform.TRANSFORM" object
func (t *USE) Transform(tr transform.TRANSFORM) *USE {
	t.node.Transform(tr)
	return t
}
