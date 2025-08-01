package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var Word = regexp.MustCompile(`\b\w+'?\w*\b`)

func WordCount(phrase string) Frequency {
	res := make(map[string]int)
	matches := Word.FindAllString(phrase, -1)

	for _, word := range(matches) {
		res[strings.ToLower(word)] += 1
	}

	return res
}
