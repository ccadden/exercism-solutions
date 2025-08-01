package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var b strings.Builder
	counter := 1

	for _, char := range(strings.ToLower(s)) {
		if counter == 6 {
			counter = 1
			b.WriteRune(' ')
		}

		if !unicode.IsSpace(char) && !unicode.IsPunct(char) {
			counter += 1
			if unicode.IsLetter(char) {
				var transform int32 = 122-(char-97)
				b.WriteRune(transform)
			} else if unicode.IsNumber(char) {
				b.WriteRune(char)
			}
		}
	}

	return strings.TrimSpace(b.String())
}
