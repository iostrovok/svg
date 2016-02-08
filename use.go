package svg

import ()

/*
‘class’
‘style’
‘externalResourcesRequired’
‘transform’
‘x’
‘y’
‘width’
‘height’
‘xlink:href’
*/

type USE struct {
	//iNode
	*node
	id  string
	use string
}

// Constructor of "use" object
func Use(use string, attr ...map[string]string) *USE {
	t := &USE{
		node: newNode(),
		use:  use,
	}

	if len(attr) > 0 {
		for k, v := range attr[0] {
			t.Attr(k, v)
		}
	}

	return t
}

// Source() returns svg implementation of USE element
func (use *USE) Source() string {
	body := `<` + use.use
	return _Source(use, body, `</`+use.use+`>`, use.node.inner)
}
