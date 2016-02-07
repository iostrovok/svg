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
func (title *TITLE) ID(id string) *TITLE {
	title.node.id = id
	return title
}

// GetID() returns lement id.
func (title *TITLE) GetID() string {
	return title.node.id
}

// Attr adds any user attribute.
func (title *TITLE) Attr(attr, values string) *TITLE {
	title.node.attrs[attr] = values
	return title
}
