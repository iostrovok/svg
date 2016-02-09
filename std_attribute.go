package svg

import (
	"github.com/iostrovok/svg/transform"
	"github.com/iostrovok/svg/style"
	"strconv"
)

// >>>>>>> START CIRCLE

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (t *CIRCLE) Append(nodes ...iNode) *CIRCLE {
	t.node.Append(nodes...)
	return t
}

// AppendTo is interface function
func (t *CIRCLE) AppendTo(n iNode) *CIRCLE {
	n.appendIn(t)
	return t
}

// Attr adds any user attribute.
func (t *CIRCLE) Attr(attr, value string) *CIRCLE {
	t.node.attrs[attr] = value
	return t
}

// Class(string) set element class.
func (t *CIRCLE) Class(id string) *CIRCLE {
	t.node.class = id
	return t
}

// GetClass() returns element id class for string.
func (t *CIRCLE) GetClass() string {
	return t.node.class
}

// XYWH sets  CX coordinate for element.
func (t *CIRCLE) CX(x float64, dim ...string) *CIRCLE {
	t.node.XYWH("cx", x, dim...)
	return t
}

// XYWH sets  CY coordinate for element.
func (t *CIRCLE) CY(x float64, dim ...string) *CIRCLE {
	t.node.XYWH("cy", x, dim...)
	return t
}

// ID(string) set element id.
func (t *CIRCLE) ID(id string) *CIRCLE {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *CIRCLE) GetID() string {
	return t.node.id
}

// genID() creates element id.
func (n *CIRCLE) genID() {
	idCounter++
	n.node.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
}

// nodes returns inner node object
func (t *CIRCLE) nodes() *node {
	return t.node
}

// XYWH sets radius for element.
func (t *CIRCLE) R(x float64, dim ...string) *CIRCLE {
	t.node.XYWH("r", x, dim...)
	return t
}

// Style sets the "style.STYLE" object
func (t *CIRCLE) Style(st style.STYLE) *CIRCLE {
	t.node.Style(st)
	return t
}

// Transform sets the "transform.TRANSFORM" object
func (t *CIRCLE) Transform(tr transform.TRANSFORM) *CIRCLE {
	t.node.Transform(tr)
	return t
}

// >>>>>>> FINISH CIRCLE

// >>>>>>> START ELLIPSE

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (t *ELLIPSE) Append(nodes ...iNode) *ELLIPSE {
	t.node.Append(nodes...)
	return t
}

// AppendTo is interface function
func (t *ELLIPSE) AppendTo(n iNode) *ELLIPSE {
	n.appendIn(t)
	return t
}

// Attr adds any user attribute.
func (t *ELLIPSE) Attr(attr, value string) *ELLIPSE {
	t.node.attrs[attr] = value
	return t
}

// Class(string) set element class.
func (t *ELLIPSE) Class(id string) *ELLIPSE {
	t.node.class = id
	return t
}

// GetClass() returns element id class for string.
func (t *ELLIPSE) GetClass() string {
	return t.node.class
}

// XYWH sets  CX coordinate for element.
func (t *ELLIPSE) CX(x float64, dim ...string) *ELLIPSE {
	t.node.XYWH("cx", x, dim...)
	return t
}

// XYWH sets  CY coordinate for element.
func (t *ELLIPSE) CY(x float64, dim ...string) *ELLIPSE {
	t.node.XYWH("cy", x, dim...)
	return t
}

// ID(string) set element id.
func (t *ELLIPSE) ID(id string) *ELLIPSE {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *ELLIPSE) GetID() string {
	return t.node.id
}

// genID() creates element id.
func (n *ELLIPSE) genID() {
	idCounter++
	n.node.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
}

// nodes returns inner node object
func (t *ELLIPSE) nodes() *node {
	return t.node
}

// XYWH sets  RX coordinate for element.
func (t *ELLIPSE) RX(x float64, dim ...string) *ELLIPSE {
	t.node.XYWH("rx", x, dim...)
	return t
}

// XYWH sets  RY coordinate for element.
func (t *ELLIPSE) RY(x float64, dim ...string) *ELLIPSE {
	t.node.XYWH("ry", x, dim...)
	return t
}

// Style sets the "style.STYLE" object
func (t *ELLIPSE) Style(st style.STYLE) *ELLIPSE {
	t.node.Style(st)
	return t
}

// Transform sets the "transform.TRANSFORM" object
func (t *ELLIPSE) Transform(tr transform.TRANSFORM) *ELLIPSE {
	t.node.Transform(tr)
	return t
}

// >>>>>>> FINISH ELLIPSE

// >>>>>>> START GROUP

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
func (t *GROUP) Attr(attr, value string) *GROUP {
	t.node.attrs[attr] = value
	return t
}

// Class(string) set element class.
func (t *GROUP) Class(id string) *GROUP {
	t.node.class = id
	return t
}

// GetClass() returns element id class for string.
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

// genID() creates element id.
func (n *GROUP) genID() {
	idCounter++
	n.node.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
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

// >>>>>>> FINISH GROUP

// >>>>>>> START LINE

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
func (t *LINE) Attr(attr, value string) *LINE {
	t.node.attrs[attr] = value
	return t
}

// Class(string) set element class.
func (t *LINE) Class(id string) *LINE {
	t.node.class = id
	return t
}

// GetClass() returns element id class for string.
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

// genID() creates element id.
func (n *LINE) genID() {
	idCounter++
	n.node.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
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

// >>>>>>> FINISH LINE

// >>>>>>> START PATH

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
func (t *PATH) Attr(attr, value string) *PATH {
	t.node.attrs[attr] = value
	return t
}

// Class(string) set element class.
func (t *PATH) Class(id string) *PATH {
	t.node.class = id
	return t
}

// GetClass() returns element id class for string.
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

// genID() creates element id.
func (n *PATH) genID() {
	idCounter++
	n.node.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
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

// >>>>>>> FINISH PATH

// >>>>>>> START POLYGON

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (t *POLYGON) Append(nodes ...iNode) *POLYGON {
	t.node.Append(nodes...)
	return t
}

// AppendTo is interface function
func (t *POLYGON) AppendTo(n iNode) *POLYGON {
	n.appendIn(t)
	return t
}

// Attr adds any user attribute.
func (t *POLYGON) Attr(attr, value string) *POLYGON {
	t.node.attrs[attr] = value
	return t
}

// Class(string) set element class.
func (t *POLYGON) Class(id string) *POLYGON {
	t.node.class = id
	return t
}

// GetClass() returns element id class for string.
func (t *POLYGON) GetClass() string {
	return t.node.class
}

// ID(string) set element id.
func (t *POLYGON) ID(id string) *POLYGON {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *POLYGON) GetID() string {
	return t.node.id
}

// genID() creates element id.
func (n *POLYGON) genID() {
	idCounter++
	n.node.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
}

// nodes returns inner node object
func (t *POLYGON) nodes() *node {
	return t.node
}

// Style sets the "style.STYLE" object
func (t *POLYGON) Style(st style.STYLE) *POLYGON {
	t.node.Style(st)
	return t
}

// Transform sets the "transform.TRANSFORM" object
func (t *POLYGON) Transform(tr transform.TRANSFORM) *POLYGON {
	t.node.Transform(tr)
	return t
}

// >>>>>>> FINISH POLYGON

// >>>>>>> START POLYLINE

// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (t *POLYLINE) Append(nodes ...iNode) *POLYLINE {
	t.node.Append(nodes...)
	return t
}

// AppendTo is interface function
func (t *POLYLINE) AppendTo(n iNode) *POLYLINE {
	n.appendIn(t)
	return t
}

// Attr adds any user attribute.
func (t *POLYLINE) Attr(attr, value string) *POLYLINE {
	t.node.attrs[attr] = value
	return t
}

// Class(string) set element class.
func (t *POLYLINE) Class(id string) *POLYLINE {
	t.node.class = id
	return t
}

// GetClass() returns element id class for string.
func (t *POLYLINE) GetClass() string {
	return t.node.class
}

// ID(string) set element id.
func (t *POLYLINE) ID(id string) *POLYLINE {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *POLYLINE) GetID() string {
	return t.node.id
}

// genID() creates element id.
func (n *POLYLINE) genID() {
	idCounter++
	n.node.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
}

// nodes returns inner node object
func (t *POLYLINE) nodes() *node {
	return t.node
}

// Style sets the "style.STYLE" object
func (t *POLYLINE) Style(st style.STYLE) *POLYLINE {
	t.node.Style(st)
	return t
}

// Transform sets the "transform.TRANSFORM" object
func (t *POLYLINE) Transform(tr transform.TRANSFORM) *POLYLINE {
	t.node.Transform(tr)
	return t
}

// >>>>>>> FINISH POLYLINE

// >>>>>>> START RECT

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
func (t *RECT) Attr(attr, value string) *RECT {
	t.node.attrs[attr] = value
	return t
}

// Class(string) set element class.
func (t *RECT) Class(id string) *RECT {
	t.node.class = id
	return t
}

// GetClass() returns element id class for string.
func (t *RECT) GetClass() string {
	return t.node.class
}

// Height sets height for element.
func (t *RECT) Height(x float64, dim ...string) *RECT {
	t.node.XYWH("height", x, dim...)
	return t
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

// genID() creates element id.
func (n *RECT) genID() {
	idCounter++
	n.node.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
}

// nodes returns inner node object
func (t *RECT) nodes() *node {
	return t.node
}

// XYWH sets  RX coordinate for element.
func (t *RECT) RX(x float64, dim ...string) *RECT {
	t.node.XYWH("rx", x, dim...)
	return t
}

// XYWH sets  RY coordinate for element.
func (t *RECT) RY(x float64, dim ...string) *RECT {
	t.node.XYWH("ry", x, dim...)
	return t
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

// Width sets  width for element.
func (t *RECT) Width(x float64, dim ...string) *RECT {
	t.node.XYWH("width", x, dim...)
	return t
}

// XYWH sets  X coordinate for element.
func (t *RECT) X(x float64, dim ...string) *RECT {
	t.node.XYWH("x", x, dim...)
	return t
}

// Y sets  y coordinate for element.
func (t *RECT) Y(x float64, dim ...string) *RECT {
	t.node.XYWH("y", x, dim...)
	return t
}

// >>>>>>> FINISH RECT

// >>>>>>> START TEXT

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
func (t *TEXT) Attr(attr, value string) *TEXT {
	t.node.attrs[attr] = value
	return t
}

// Class(string) set element class.
func (t *TEXT) Class(id string) *TEXT {
	t.node.class = id
	return t
}

// GetClass() returns element id class for string.
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

// genID() creates element id.
func (n *TEXT) genID() {
	idCounter++
	n.node.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
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

// >>>>>>> FINISH TEXT

// >>>>>>> START TITLE

// Attr adds any user attribute.
func (t *TITLE) Attr(attr, value string) *TITLE {
	t.node.attrs[attr] = value
	return t
}

// Class(string) set element class.
func (t *TITLE) Class(id string) *TITLE {
	t.node.class = id
	return t
}

// GetClass() returns element id class for string.
func (t *TITLE) GetClass() string {
	return t.node.class
}

// ID(string) set element id.
func (t *TITLE) ID(id string) *TITLE {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *TITLE) GetID() string {
	return t.node.id
}

// genID() creates element id.
func (n *TITLE) genID() {
	idCounter++
	n.node.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
}

// >>>>>>> FINISH TITLE

// >>>>>>> START USE

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
func (t *USE) Attr(attr, value string) *USE {
	t.node.attrs[attr] = value
	return t
}

// Class(string) set element class.
func (t *USE) Class(id string) *USE {
	t.node.class = id
	return t
}

// GetClass() returns element id class for string.
func (t *USE) GetClass() string {
	return t.node.class
}

// Height sets height for element.
func (t *USE) Height(x float64, dim ...string) *USE {
	t.node.XYWH("height", x, dim...)
	return t
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

// genID() creates element id.
func (n *USE) genID() {
	idCounter++
	n.node.id = "_auto_id_generate_" + strconv.Itoa(idCounter)
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

// Width sets  width for element.
func (t *USE) Width(x float64, dim ...string) *USE {
	t.node.XYWH("width", x, dim...)
	return t
}

// XYWH sets  X coordinate for element.
func (t *USE) X(x float64, dim ...string) *USE {
	t.node.XYWH("x", x, dim...)
	return t
}

// Y sets  y coordinate for element.
func (t *USE) Y(x float64, dim ...string) *USE {
	t.node.XYWH("y", x, dim...)
	return t
}

// >>>>>>> FINISH USE
