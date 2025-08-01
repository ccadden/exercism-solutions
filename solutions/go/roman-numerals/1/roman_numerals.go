package romannumerals

import (
	"errors"
	"strings"
)

var RomanNumeralsMap = map[int] string {
	1000 : "M",
	900 : "CM",
	500 : "D",
	400 : "CD",
	100 : "C",
	90 : "XC",
	50 : "L",
	40 : "XL",
	10 : "X",
	9 : "IX",
	5 : "V",
	4 : "IV",
	1 : "I",
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 {
		return "", errors.New("Roman Numeral must be a positive integer")
	}

	if input > 3999 {
		return "", errors.New("Romans can't count past MMMCMIX")
	}

	var multiple int
	var remainder int

	var b strings.Builder
	for _, factor := range([]int { 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}) {
		if input == 0 {
			break
		}
		if input < factor {
			continue
		}

		multiple = input / factor
		remainder = input % factor

		b.WriteString(strings.Repeat(RomanNumeralsMap[factor], multiple))
		input = remainder
	}

	return b.String(), nil
}
