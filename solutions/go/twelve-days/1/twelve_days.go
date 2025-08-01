package twelve

import (
	"fmt"
	"strings"
)

var ordinals = []string {
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = []string {
	"twelve Drummers Drumming",
	"eleven Pipers Piping",
	"ten Lords-a-Leaping",
	"nine Ladies Dancing",
	"eight Maids-a-Milking",
	"seven Swans-a-Swimming",
	"six Geese-a-Laying",
	"five Gold Rings",
	"four Calling Birds",
	"three French Hens",
	"two Turtle Doves",
	"and a Partridge in a Pear Tree.",
}

func Gifts(i int) string {
	switch i {
	case 0:
		return gifts[11][4:]
	default:
		return strings.Join(gifts[11-i:], ", ")
	}
}

func Verse(i int) string {
	i -= 1
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", ordinals[i], Gifts(i))
}

func Song() string {
	var b strings.Builder
	for i:=1; i<13; i++ {
		b.WriteString(Verse(i))
		if i != 12 {
			b.WriteString("\n")
		}
	}

	return b.String()
}
