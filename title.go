package svg

import (
	"fmt"

	"github.com/iostrovok/svg/style"
)

type TITLE struct {
	Node
	body string
}

const (
	titleTag = `<title>%s</title>`
)

// Constructor
func Title(text string) TITLE {
	return TITLE{
		body: fmt.Sprintf(titleTag, text),
	}
}

// Setter
func (title TITLE) Style(s style.STYLE) Node {
	return title
}

/*
	Inner() returns inner elements of TITLE.
	For TITLE it is always empty list
*/
func (title TITLE) Inner() []Node {
	return []Node{}
}

// Source() returns svg implementation of TITLE element
func (title TITLE) Source() string {
	return title.body
}
