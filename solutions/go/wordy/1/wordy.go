package wordy

import (
	"strconv"
	"strings"
	"regexp"
)

var Numbers = regexp.MustCompile(`-?\d+`)

var Operations = regexp.MustCompile(`\bplus\b|\bminus\b|\bmultiplied by\b|\bdivided by\b`)

func parseNumber(s string) (int, bool) {
	val, err := strconv.Atoi(s)

	if err != nil {
		return 0, false
	}

	return val, true
}

func convertStringToNum(s string) (int, bool) {
	val, err := strconv.Atoi(s)

	if err != nil {
		return 0, false
	}

	return val, true
}

func getNums(question string, indices [][]int) []int {
	nums := []int{}

	for _, i := range(indices) {
		num, ok := convertStringToNum(question[i[0]:i[1]])

		if !ok {
			panic("something went wrong converting numbers")
		}

		nums = append(nums, num)
	}

	return nums
}

func calculate(a int, b int, operation string) int {
	switch operation {
	case "plus":
		return a + b
	case "minus":
		return a - b
	case "multiplied by":
		return a * b 
	case "divided by":
		return a / b
	default:
		return 0
	}
}

func indicesAdjacent(a, b []int) bool {
	if a[1] + 1 == b[0] {
		return true
	}

	return false
}

func validOperations(operations [][]int, question string) bool {
	for i:=1;i<len(operations);i++ {
		if indicesAdjacent(operations[i-1], operations[i]) {
			return false
		}
	}

	if operations[len(operations) - 1][1] == len(question) {
		return false
	}

	return true
}

func validNumbers(numbers [][]int) bool {
	for i:=1;i<len(numbers);i++ {
		if indicesAdjacent(numbers[i-1], numbers[i]) {
			return false
		}
	}

	return true
}

func Answer(question string) (int, bool) {
	question = strings.ReplaceAll(question, "What is ", "")
	question = strings.ReplaceAll(question, "?", "")

	numIndices := Numbers.FindAllStringIndex(question, -1)
	operations := Operations.FindAllStringIndex(question, -1)

	if numIndices == nil {
		return 0, false
	}

	if !validNumbers(numIndices) {
		return 0, false
	}

	if len(operations) == 0 {
		return parseNumber(question)
	} else {
		if !validOperations(operations, question) {
			return 0, false
		}
	}

	nums := getNums(question, numIndices)
	answer := nums[0]


	for i, operation := range(operations) {
		left := operation[0]
		right := operation[1]

		answer = calculate(answer, nums[i + 1], question[left:right])
	}

	return answer, true
}
