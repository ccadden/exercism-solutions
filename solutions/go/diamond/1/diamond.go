package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 65 || char > 90 {
		return "", errors.New("illegal character")
	}
	var b strings.Builder

	var letter byte = 65
	var innerPadding int = 0
	var outterPadding int = int(char) - 65

	for letter <= char {
		if innerPadding != 0 {
			b.WriteRune(10)
		}
		for j := 0; j < outterPadding; j++ {
			b.WriteString(" ")
		}
		if innerPadding != 0 {
			b.WriteByte(letter)
		}
		for k := 1; k < innerPadding; k++ {
			b.WriteString(" ")
		}
		b.WriteByte(letter)
		for j := 0; j < outterPadding; j++ {
			b.WriteString(" ")
		}

		letter++
		innerPadding += 2
		outterPadding -= 1
	}

	letter -= 2
	innerPadding -= 4
	outterPadding += 2

	for letter >= 65 {
		b.WriteRune(10)

		for j := 0; j < outterPadding; j++ {
			b.WriteString(" ")
		}
		if innerPadding != 0 {
			b.WriteByte(letter)
		}
		for k := 1; k < innerPadding; k++ {
			b.WriteString(" ")
		}
		b.WriteByte(letter)
		for j := 0; j < outterPadding; j++ {
			b.WriteString(" ")
		}

		letter--
		innerPadding -= 2
		outterPadding += 1

	}

	return b.String(), nil
}
