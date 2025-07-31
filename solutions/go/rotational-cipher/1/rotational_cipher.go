package rotationalcipher

import (
	"strings"
	"unicode"
)

func rotate(char int32, offset int32, start int32) int32 {
	return ((char - start + offset) % 26) + start
}

func RotationalCipher(plain string, shiftKey int) string {
	var b strings.Builder
	var start int32

	for _, char := range(plain) {
		if unicode.IsLetter(char) {
			switch {
			case char > 64 && char < 91:
				start = 65
			case char > 96 && char < 123:
				start = 97
			}
			b.WriteRune(rotate(char, int32(shiftKey), start))
		} else {
			b.WriteRune(char)
		}
	}

	return b.String()
}
