package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	var b strings.Builder

	count := 1
	i := 0
	j := 1

	for i < len(input) {
		for j < len(input) {
			if input[i] == input[j] {
				count++
				j++
			} else {
				break
			}
		}

		if count > 1 {
			b.WriteString(strconv.Itoa(count))
		}

		count = 1

		b.WriteByte(input[i])
		i = j
		j = i + 1
	}

	return b.String()
}

func RunLengthDecode(input string) string {
	var b strings.Builder

	i := 0
	j := 1

	for i < len(input) {
		if !unicode.IsDigit(rune(input[i])) {
			b.WriteByte(input[i])
			i++
			j++
			continue
		}
		for j < len(input) {
			if unicode.IsDigit(rune(input[j])) {
				j++
			} else {
				break
			}
		}

		val, _ := strconv.Atoi(input[i:j])

		for k := 0; k < val; k++ {
			b.WriteByte(input[j])
		}

		i = j + 1
		j = i + 1
	}

	return b.String()
}
