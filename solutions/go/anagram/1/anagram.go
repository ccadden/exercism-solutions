package anagram

import (
	"strings"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

func Detect(subject string, candidates []string) []string {
	res := []string{}
	var add bool

	subjectRunes := sortRunes(strings.ToLower(subject))
	sort.Sort(subjectRunes)

	for _, candidate := range(candidates) {
		add = true
		if strings.ToLower(candidate) == strings.ToLower(subject) {
			continue
		}

		if len(candidate) != len(subject) {
			continue
		}

		candidateRunes := sortRunes(strings.ToLower(candidate))
		sort.Sort(candidateRunes)


		for i, _ := range(candidateRunes) {
			if candidateRunes[i] != subjectRunes[i] {
				add = false
				break
			}
		}

		if add {
			res = append(res, candidate)
		}
	}

	return res
}
