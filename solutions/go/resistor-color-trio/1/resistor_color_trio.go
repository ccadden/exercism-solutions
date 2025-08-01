package resistorcolortrio

import (
	"fmt"
	"strconv"
	"strings"
)

var ResistorColors map[string]string = map[string]string{
	"black":  "0",
	"brown":  "1",
	"red":    "2",
	"orange": "3",
	"yellow": "4",
	"green":  "5",
	"blue":   "6",
	"violet": "7",
	"grey":   "8",
	"white":  "9",
}

func Label(colors []string) string {
	var b strings.Builder

	for i, color := range colors[0:2] {
		if color == "black" && i == 0 {
			continue
		}

		b.WriteString(ResistorColors[color])
	}

	trailingZeroes, _ := strconv.Atoi(ResistorColors[colors[2]])

	for i := 0; i < trailingZeroes; i++ {
		b.WriteString("0")
	}

	num, _ := strconv.Atoi(b.String())

	switch {
	case num > 1e9:
		num /= 1e9
		return fmt.Sprintf("%v gigaohms", num)
	case num > 1e6:
		num /= 1e6
		return fmt.Sprintf("%v megaohms", num)
	case num > 1e3:
		num /= 1e3
		return fmt.Sprintf("%v kiloohms", num)
	default:
		return fmt.Sprintf("%v ohms", num)
	}
}
