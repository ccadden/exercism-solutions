package cipher

import (
	"unicode"
	"strings"
)

type Shift struct {
	Distance int
}

type Vigenere struct {
	Key []rune
}

type Caesar struct {}

func rotateChar(r rune, offset rune) rune {
	mod := ((r - 97 + offset) % 26)
	if mod < 0 {
		mod += 26
	}
	return mod + 97
}

func rotate(input string, offset int) string {
	var res strings.Builder

	for _, char := range(input) {
		if unicode.IsLetter(char) {
			shiftedChar := rotateChar(char, rune(offset))
			res.WriteRune(shiftedChar)
		}
	}

	return res.String()
}

func (c Caesar) Encode(input string) string {
	input = strings.ToLower(input)

	return rotate(input, 3)
}

func (c Caesar) Decode(input string) string {
	input = strings.ToLower(input)

	return rotate(input, -3)
}

func NewCaesar() Cipher {
	return Caesar{}
}

func NewShift(distance int) Cipher {
	if !distanceValid(distance) {
		return nil
	}

	return Shift{Distance: distance}
}

func distanceValid(distance int) bool {
	return distance != 0 &&
	((distance >= -25 && distance <= -1) || (distance >= 1 && distance <= 25))
}

func (c Shift) Encode(input string) string {
	input = strings.ToLower(input)
	
	return rotate(input, c.Distance)
}

func (c Shift) Decode(input string) string {
	input = strings.ToLower(input)

	return rotate(input, -c.Distance)
}

func validKey(key string) bool {
	allA := true

	for _, char := range(key) {
		if char < 97 || char > 122 {
			return false
		}
		if char != 97 {
			allA = false
		}
	}

	return !allA
}

func NewVigenere(key string) Cipher {
	if !validKey(key) {
		return nil
	}
	return Vigenere{Key: []rune(key)}
}

func rotateVig(v Vigenere, input string, decode bool) string {
	input = strings.ToLower(input)
	var res strings.Builder
	var i int
	var offset rune

	for _, char := range(input) {
		if unicode.IsLetter(char) {
			if decode {
				offset = 97 - v.Key[i % len(v.Key)]
			} else {
				offset = v.Key[i % len(v.Key)] - 97
			}

			shiftedChar := rotateChar(char, offset)

			res.WriteRune(shiftedChar)

			i++
		}
	}

	return res.String()
}

func (v Vigenere) Encode(input string) string {
	return rotateVig(v, input, false)
}

func (v Vigenere) Decode(input string) string {
	return rotateVig(v, input, true)
}
