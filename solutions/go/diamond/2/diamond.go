package diamond

import (
	"errors"
	"strings"
)

func genLine(char byte, op int, ip int, newLine bool) string {
	var b strings.Builder

	if newLine {
		b.WriteRune('\n')
	}

	b.WriteString(strings.Repeat(" ", op))

	if ip != 0 {
		b.WriteByte(char)
		b.WriteString(strings.Repeat(" ", ip-1))
	}

	b.WriteByte(char)
	b.WriteString(strings.Repeat(" ", op))

	return b.String()
}

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("illegal character")
	}

	var b strings.Builder

	var letter byte = 'A'
	var innerPadding int = 0
	var outterPadding int = int(char) - 65

	for letter <= char {
		b.WriteString(
			genLine(letter,
				outterPadding,
				innerPadding,
				letter != 'A'),
		)

		letter++
		innerPadding += 2
		outterPadding -= 1
	}

	letter -= 2
	innerPadding -= 4
	outterPadding += 2

	for letter >= 'A' {
		b.WriteString(
			genLine(letter,
				outterPadding,
				innerPadding,
				true),
		)

		letter--
		innerPadding -= 2
		outterPadding += 1
	}

	return b.String(), nil
}
