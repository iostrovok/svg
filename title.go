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
