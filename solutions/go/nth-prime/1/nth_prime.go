package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Need n > 0")
	}

	if n == 1 {
		return 2, nil
	}

	var currentNum int = 3
	var primeCount int = 2
	var prime bool

	for {
		if primeCount == n {
			break
		}

		currentNum += 2

		prime = true

		for i:=3; i < currentNum; i = i + 2 {
			if currentNum % i == 0 {
				prime = false
				break
			}
		}

		if prime {
			primeCount++
		}
	}

	return currentNum, nil
}
