package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	tags = map[string][]string{
		`PATH`:  []string{`node`, `id`, `class`, `attr`, `append`, `style`, `transform`},
		`GROUP`: []string{`node`, `id`, `class`, `attr`, `append`, `style`, `transform`},
		`TEXT`:  []string{`node`, `id`, `class`, `attr`, `append`, `style`, `transform`},
		`LINE`:  []string{`node`, `id`, `class`, `attr`, `append`, `style`, `transform`},
		`RECT`:  []string{`node`, `id`, `class`, `attr`, `append`, `style`, `transform`},
		`USE`:   []string{`node`, `id`, `class`, `attr`, `append`, `style`, `transform`, `x`, `y`, `width`, `height`},
	}

	tagReg       = regexp.MustCompile(`<TAG>`)
	fileTemplate = `package svg
import "github.com/iostrovok/svg/transform"
import "github.com/iostrovok/svg/style"

%s
`

	maps map[string]string = map[string]string{

		`x`: `// XYWH sets  X coordinate for element.
func (t *USE) X(x float64, dim ...string) *USE {
	t.node.XYWH("x", x, dim...)
	return t
}`,

		`y`: `// Y sets  y coordinate for element.
func (t *USE) Y(x float64, dim ...string) *USE {
	t.node.XYWH("y", x, dim...)
	return t
}`,

		`width`: `// Width sets  width for element.
func (t *USE) Width(x float64, dim ...string) *USE {
	t.node.XYWH("width", x, dim...)
	return t
}`,

		`height`: `// Height sets height for element.
func (t *USE) Height(x float64, dim ...string) *USE {
	t.node.XYWH("height", x, dim...)
	return t
}`,

		`node`: `// nodes returns inner node object
func (t *<TAG>) nodes() *node {
	return t.node
}`,

		`id`: `// ID(string) set element id.
func (t *<TAG>) ID(id string) *<TAG> {
	t.node.id = id
	return t
}

// GetID() returns lement id.
func (t *<TAG>) GetID() string {
	return t.node.id
}`,

		`class`: `// Class(string) set element class.
func (t *<TAG>) Class(id string) *<TAG> {
	t.node.class = id
	return t
}

// GetID() returns element id class for string.
func (t *<TAG>) GetClass() string {
	return t.node.class
}`,

		`attr`: `// Attr adds any user attribute.
func (t *<TAG>) Attr(attr, values string) *<TAG> {
	t.node.attrs[attr] = values
	return t
}`,
		`append`: `// Append() inserts content, specified by the parameter, to the end of each element in the set of matched elements.
func (t *<TAG>) Append(nodes ...iNode) *<TAG> {
	t.node.Append(nodes...)
	return t
}

// AppendTo is interface function
func (t *<TAG>) AppendTo(n iNode) *<TAG> {
	n.appendIn(t)
	return t
}`,

		`style`: `// Style sets the "style.STYLE" object
func (t *<TAG>) Style(st style.STYLE) *<TAG> {
	t.node.Style(st)
	return t
}`,

		`transform`: `// Transform sets the "transform.TRANSFORM" object
func (t *<TAG>) Transform(tr transform.TRANSFORM) *<TAG> {
	t.node.Transform(tr)
	return t
}`,
	}
)

func main() {
	flag.Parse()

	args := flag.Args()
	path := args[0]
	filelName := filepath.Join(path, "/", "std_attribute.go")

	os.Remove(filelName)

	lines := []string{}

	tagsList := []string{}
	for one := range tags {
		tagsList = append(tagsList, one)
	}

	sort.Strings(tagsList)

	for _, one := range tagsList {
		lines = append(lines, "// >>>>>>> START "+one)
		list := tags[one]
		sort.Strings(list)
		for _, k := range list {
			res := tagReg.ReplaceAllString(maps[k], one)
			lines = append(lines, res)
		}
		lines = append(lines, "// >>>>>>> FINISH "+one)
	}

	fmt.Printf("\n: %s\n", filelName)
	str := fmt.Sprintf(fileTemplate, strings.Join(lines, "\n\n"))

	err := ioutil.WriteFile(filelName, []byte(str), 0666)
	if err != nil {
		log.Fatalln(err)
	}
}
