package svg

import (
	"regexp"
)

var (
	deleteSpace      = regexp.MustCompile(`(\s|\t|\n)+`)
	deleteLastSpace  = regexp.MustCompile(`[ 	\n]+$`)
	deleteLastSpace3 = regexp.MustCompile(`[ 	\n]+/>`)
	deleteLastSpace2 = regexp.MustCompile(`[ 	\n]+>`)

	deleteFirstSpace = regexp.MustCompile(`>[ 	\n]+`)
)

// clean space line
func cSL(s string) string {
	s = deleteLastSpace3.ReplaceAllString(s, `/>`)
	s = deleteLastSpace2.ReplaceAllString(s, `>`)
	s = deleteSpace.ReplaceAllString(s, ` `)
	s = deleteLastSpace.ReplaceAllString(s, ``)
	s = deleteFirstSpace.ReplaceAllString(s, `>`)

	return s

}
