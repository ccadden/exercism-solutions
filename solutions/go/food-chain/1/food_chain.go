package foodchain

import (
	"fmt"
	"strings"
)

type Animal struct {
	name string
	line string
}

var animals []*Animal = []*Animal{
	newAnimal("fly", "I don't know why she swallowed the fly. Perhaps she'll die."),
	newAnimal("spider", "It wriggled and jiggled and tickled inside her."),
	newAnimal("bird", "How absurd to swallow a bird!"),
	newAnimal("cat", "Imagine that, to swallow a cat!"),
	newAnimal("dog", "What a hog, to swallow a dog!"),
	newAnimal("goat", "Just opened her throat and swallowed a goat!"),
	newAnimal("cow", "I don't know how she swallowed a cow!"),
	newAnimal("horse", "She's dead, of course!"),
}

func newAnimal(name, line string) *Animal {
	return &Animal{name: name, line: line}
}

func Verse(v int) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("I know an old lady who swallowed a %s.\n%s", animals[v-1].name, animals[v-1].line))

	if v > 1 && v < 8 {
		for i := v; i > 1; i-- {
			var postfix string
			if i == 3 {
				postfix = " that wriggled and jiggled and tickled inside her."
			} else {
				postfix = "."

			}
			b.WriteString("\n")
			b.WriteString(
				fmt.Sprintf(
					"She swallowed the %s to catch the %s%s",
					animals[i-1].name,
					animals[i-2].name,
					postfix,
				))
		}
		b.WriteString("\n")
		b.WriteString(animals[0].line)
	}

	return b.String()
}

func Verses(start, end int) string {
	var b strings.Builder

	for i := start; i <= end; i++ {
		if i > start {
			b.WriteString("\n\n")
		}
		b.WriteString(Verse(i))
	}

	return b.String()
}

func Song() string {
	return Verses(1, 8)
}
