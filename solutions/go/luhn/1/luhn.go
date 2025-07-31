package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func convertString(id string) ([]int, bool) {
	res := []int{}

	if len(id) < 2 {
		return res, false
	}

	for _, char := range(id) {
		if unicode.IsSpace(char) {
			continue
		} else if !unicode.IsNumber(char) {
			return []int{}, false
		}

		val, err := strconv.Atoi(string(char))

		if err != nil {
			panic(err)
		}

		res = append(res, val)
	}

	return res, true
}
func Valid(id string) bool {
	nums, ok := convertString(strings.ReplaceAll(id, " ", ""))

	if !ok {
		return false
	}

	sum := 0
	for i := 1; i <= len(nums); i++ {
		j := len(nums) - i
		if i % 2 == 0 {
			nums[j] *= 2
			if nums[j] > 9 {
				nums[j] -= 9
			}
		} 

		sum += nums[j]
	}

	return (sum % 10) == 0
}
