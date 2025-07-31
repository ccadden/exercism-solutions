package acronym

import (
	"regexp"
	"strings"
)

var firstLetter = regexp.MustCompile(`\b[A-Z]`)

func Abbreviate(s string) string {
	var b strings.Builder
	s = strings.ToUpper(s)

	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, "_", "")

	indices := firstLetter.FindAllStringIndex(s, -1)

	for _, index := range(indices) {
		b.WriteString(s[index[0]:index[1]])
	}

	return b.String()
}
