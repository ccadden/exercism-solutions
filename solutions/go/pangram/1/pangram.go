package pangram

import "strings"

func IsPangram(input string) bool {
	upper := strings.ToUpper(input)

	for _, char := range("ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		if !strings.ContainsRune(upper, char) {
			return false
		}
	}
	return true
}
