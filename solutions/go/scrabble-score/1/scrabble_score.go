package scrabble

import "strings"

func Score(word string) int {
	score := 0

	for _, char := range(strings.ToUpper(word)) {
		score += GetLetterScore(char)
	}

	return score
}

func GetLetterScore(r rune) int {
	switch {
	case strings.ContainsRune("AEIOULNRST", r):
		return 1
	case strings.ContainsRune("DG", r):
		return 2
	case strings.ContainsRune("BCMP", r):
		return 3
	case strings.ContainsRune("FHVWY", r):
		return 4
	case strings.ContainsRune("K", r):
		return 5
	case strings.ContainsRune("JX", r):
		return 8
	case strings.ContainsRune("QZ", r):
		return 10
	default:
		return 0
	}
}
