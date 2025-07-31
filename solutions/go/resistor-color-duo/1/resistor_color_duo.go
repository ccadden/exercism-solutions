package resistorcolorduo

import (
	"strings"
	"strconv"
)

var ResistorColorMap = map[string] string {
	"black": "0",
	"brown": "1",
	"red": "2",
	"orange": "3",
	"yellow": "4",
	"green": "5",
	"blue": "6",
	"violet": "7",
	"grey": "8",
	"white": "9",
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	var b strings.Builder

	for i := 0; i < 2; i ++ {
		val, ok := ResistorColorMap[strings.ToLower(colors[i])]
		if !ok {
			panic("Invalid color provided")
		}

		b.WriteString(val)
	}

	res, err := strconv.Atoi(b.String())

	if err != nil {
		panic(err)
	}

	return res
}
