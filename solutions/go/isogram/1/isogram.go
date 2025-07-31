package isogram

import (
	"sort"
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	s := []rune(strings.ToUpper(word))

	sort.Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})

	for i := 1; i < len(s); i++ {
		if !unicode.IsSpace(s[i]) && s[i] != '-' {
			if s[i] == s[i - 1] {
				return false
			}
		}
	}

	return true
}
