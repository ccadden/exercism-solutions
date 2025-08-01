package house

import (
	"fmt"
	"strings"
)

type Thing struct {
	name string
	verb string
}

var Things []Thing = []Thing{
	{name: "house that Jack built", verb: ""},
	{name: "malt", verb: "lay in"},
	{name: "rat", verb: "ate"},
	{name: "cat", verb: "killed"},
	{name: "dog", verb: "worried"},
	{name: "cow with the crumpled horn", verb: "tossed"},
	{name: "maiden all forlorn", verb: "milked"},
	{name: "man all tattered and torn", verb: "kissed"},
	{name: "priest all shaven and shorn", verb: "married"},
	{name: "rooster that crowed in the morn", verb: "woke"},
	{name: "farmer sowing his corn", verb: "kept"},
	{name: "horse and the hound and the horn", verb: "belonged to"},
}

func Verse(v int) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("This is the %s", Things[v-1].name))

	for i := v; i > 1; i-- {
		b.WriteString(fmt.Sprintf("\nthat %s the %s", Things[i-1].verb, Things[i-2].name))
	}

	b.WriteRune('.')

	return b.String()
}

func Song() string {
	var b strings.Builder

	for v := 1; v <= len(Things); v++ {
		if v > 1 {
			b.WriteString("\n\n")
		}
		b.WriteString(Verse(v))
	}

	return b.String()
}
