package bottlesong

import (
	"fmt"
	"strings"
)

var NumMap = map[int]string {
	0: "no",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
	10: "ten",
}

func formattedRepeatedLine(numBottles int) string {
	if numBottles == 1 {
		return fmt.Sprintf("%s green bottle hanging on the wall,", strings.Title(NumMap[numBottles]))
	}

	return fmt.Sprintf("%s green bottles hanging on the wall,", strings.Title(NumMap[numBottles]))
}

func formattedLastLine(numBottles int) string {
	if numBottles == 2 {
		return fmt.Sprintf("There'll be %s green bottle hanging on the wall.", NumMap[numBottles - 1])
	}

	return fmt.Sprintf("There'll be %s green bottles hanging on the wall.", NumMap[numBottles - 1])
}

func formattedVerse(numBottles int) []string {
	return []string{
		formattedRepeatedLine(numBottles),
		formattedRepeatedLine(numBottles),
		"And if one green bottle should accidentally fall,",
		formattedLastLine(numBottles),
		"",
	}
}

func Recite(startBottles, takeDown int) []string {
	res := []string{}

	for i := takeDown; i > 0; i-- {
		res = append(res, formattedVerse(startBottles)...)
		startBottles-- 
	}

	return res[:len(res) - 1]
}
