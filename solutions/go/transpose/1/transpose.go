package transpose

import "strings"

func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}

	maxLen := CalcMaxLen(input)
	result := []string{}

	for i := 0; i < maxLen; i++ {
		var r strings.Builder
		for _, row := range input {
			if len(row) <= i {
				r.WriteString(" ")
			} else {
				r.WriteByte(row[i])
			}

		}
		result = append(result, r.String())
	}

	return result
}

func CalcMaxLen(input []string) int {
	max := 0

	for _, row := range input {
		if len(row) > max {
			max = len(row)
		}
	}

	return max
}
