package armstrong

import "strconv"

func pow(base, power int) int {
	product := 1

	for i := 1; i <= power; i++{
		product *= base
	}

	return product
}

func IsNumber(n int) bool {
	numStr := strconv.Itoa(n)
	lenNum := len(numStr)
	sum := 0

	for _, digit := range(numStr) {
		sum += pow(int(digit - 48), lenNum)
	}

	return sum == n
}
