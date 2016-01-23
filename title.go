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

func Title(text string) TITLE {
	return TITLE{
		body: fmt.Sprintf(titleTag, text),
	}
}

func (title TITLE) Style(s style.STYLE) Node {
	return title
}

func (title TITLE) Inner() []Node {
	return []Node{}
}

func (title TITLE) Source() string {
	return title.body
}
