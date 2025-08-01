package summultiples

import "slices"

func SumMultiples(limit int, divisors ...int) int {
	if limit == 0 {
		return 0
	}

	multiples := []int{}
	sum := 0

	for _, val := range(divisors) {
		i := 1

		for {
			if val == 0 {
				break
			}
			multiple := i * val

			if multiple >= limit {
				break
			}

			if !slices.Contains(multiples, multiple) {
				multiples = append(multiples, multiple)
				sum += multiple
			}

			i++
		}
	}

	return sum
}
