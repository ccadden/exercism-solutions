package house

import (
	"fmt"
	"strings"
)

// This is the house that Jack built.
//
// This is the malt
// that lay in the house that Jack built.
//
// This is the rat
// that ate the malt
// that lay in the house that Jack built.
//
// This is the cat
// that killed the rat
// that ate the malt
// that lay in the house that Jack built.
//
// This is the dog
// that worried the cat
// that killed the rat
// that ate the malt
// that lay in the house that Jack built.
//
// This is the cow with the crumpled horn
// that tossed the dog
// that worried the cat
// that killed the rat
// that ate the malt
// that lay in the house that Jack built.
//
// This is the maiden all forlorn
// that milked the cow with the crumpled horn
// that tossed the dog
// that worried the cat
// that killed the rat
// that ate the malt
// that lay in the house that Jack built.
//
// This is the man all tattered and torn
// that kissed the maiden all forlorn
// that milked the cow with the crumpled horn
// that tossed the dog
// that worried the cat
// that killed the rat
// that ate the malt
// that lay in the house that Jack built.
//
// This is the priest all shaven and shorn
// that married the man all tattered and torn
// that kissed the maiden all forlorn
// that milked the cow with the crumpled horn
// that tossed the dog
// that worried the cat
// that killed the rat
// that ate the malt
// that lay in the house that Jack built.
//
// This is the rooster that crowed in the morn
// that woke the priest all shaven and shorn
// that married the man all tattered and torn
// that kissed the maiden all forlorn
// that milked the cow with the crumpled horn
// that tossed the dog
// that worried the cat
// that killed the rat
// that ate the malt
// that lay in the house that Jack built.
//
// This is the farmer sowing his corn
// that kept the rooster that crowed in the morn
// that woke the priest all shaven and shorn
// that married the man all tattered and torn
// that kissed the maiden all forlorn
// that milked the cow with the crumpled horn
// that tossed the dog
// that worried the cat
// that killed the rat
// that ate the malt
// that lay in the house that Jack built.
//
// This is the horse and the hound and the horn
// that belonged to the farmer sowing his corn
// that kept the rooster that crowed in the morn
// that woke the priest all shaven and shorn
// that married the man all tattered and torn
// that kissed the maiden all forlorn
// that milked the cow with the crumpled horn
// that tossed the dog
// that worried the cat
// that killed the rat
// that ate the malt
// that lay in the house that Jack built.

type Thing struct {
	name string
	verb string
}

var Things []Thing = []Thing{
	Thing{name: "house that Jack built", verb: ""},
	Thing{name: "malt", verb: "lay in"},
	Thing{name: "rat", verb: "ate"},
	Thing{name: "cat", verb: "killed"},
	Thing{name: "dog", verb: "worried"},
	Thing{name: "cow with the crumpled horn", verb: "tossed"},
	Thing{name: "maiden all forlorn", verb: "milked"},
	Thing{name: "man all tattered and torn", verb: "kissed"},
	Thing{name: "priest all shaven and shorn", verb: "married"},
	Thing{name: "rooster that crowed in the morn", verb: "woke"},
	Thing{name: "farmer sowing his corn", verb: "kept"},
	Thing{name: "horse and the hound and the horn", verb: "belonged to"},
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
