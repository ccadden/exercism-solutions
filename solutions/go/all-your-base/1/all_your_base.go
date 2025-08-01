package allyourbase

import "errors"

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}

	var from int
	var err error
	var to []int

	from, err = fromBase(inputDigits, inputBase)

	if err != nil {
		return []int{}, err
	}

	to, err = toBase(from, outputBase)

	if err != nil {
		return []int{}, err
	}

	return to, nil
}

func toBase(number, base int) ([]int, error) {
	if number == 0 {
		return []int{0}, nil
	}

	result := []int{}
	for ; number > 0; number /= base {
		result = append([]int{number % base}, result...)
	}

	return result, nil
}

func fromBase(digits []int, base int) (int, error) {
	sum := 0

	for _, digit := range(digits) {
		if digit < 0 || digit >= base {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}
		sum = sum * base + digit
	}

	return sum, nil
}
