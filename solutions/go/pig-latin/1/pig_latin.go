package piglatin

import (
	"regexp"
	"strings"
)

var startsWithVowel = regexp.MustCompile("^([aeiou]|yt|xr)")
var consonantSound = regexp.MustCompile("^[^aeiouy]{2,}")
var qu = regexp.MustCompile("^[^aeiou]?qu")

func Sentence(sentence string) string {
	words := strings.Fields(sentence)

	for i, word := range words {
		words[i] = toPigLatin(word)
	}

	return strings.Join(words, " ")
}

// Rule 1: If a word begins with a vowel sound, add an "ay" sound to the end of the word. Please note that "xr" and "yt" at the beginning of a word make vowel sounds (e.g. "xray" -> "xrayay", "yttria" -> "yttriaay").
// Rule 2: If a word begins with a consonant sound, move it to the end of the word and then add an "ay" sound to the end of the word. Consonant sounds can be made up of multiple consonants, such as the "ch" in "chair" or "st" in "stand" (e.g. "chair" -> "airchay").
// Rule 3: If a word starts with a consonant sound followed by "qu", move it to the end of the word, and then add an "ay" sound to the end of the word (e.g. "square" -> "aresquay").
// Rule 4: If a word contains a "y" after a consonant cluster or as the second letter in a two letter word it makes a vowel sound (e.g. "rhythm" -> "ythmrhay", "my" -> "ymay").
func toPigLatin(word string) string {
	quIndex := qu.FindStringIndex(word)

	if quIndex != nil {
		return word[quIndex[1]:] + word[:quIndex[1]] + "ay"
	}
	if startsWithVowel.MatchString(word) {
		return word + "ay"
	}

	consonants := consonantSound.FindStringIndex(word)

	if consonants != nil {
		return word[consonants[1]:] + word[:consonants[1]] + "ay"
	}

	return word[1:] + word[0:1] + "ay"
}
