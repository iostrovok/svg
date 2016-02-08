package svg

import (
	"fmt"
)

type TITLE struct {
	iNode
	*node
	body string
}

const (
	titleTag = `<title>%s</title>`
)

// Constructor
func Title(text string) *TITLE {
	return &TITLE{
		node: newNode(),
		body: fmt.Sprintf(titleTag, text),
	}
}

// Source() returns svg implementation of TITLE element
func (title *TITLE) Source() string {
	return title.body
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

// Class(string) set element class.
func (t *TITLE) Class(id string) *TITLE {
	t.node.class = id
	return t
}

// GetID() returns element id class for string.
func (t *TITLE) GetClass() string {
	return t.node.class
}

// Attr adds any user attribute.
func (title *TITLE) Attr(attr, values string) *TITLE {
	title.node.attrs[attr] = values
	return title
}
