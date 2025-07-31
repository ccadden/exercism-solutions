package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Invalid chess square")
	}
	product := 1

	for i:= 1; i < number; i++ {
		product *= 2
	}

	return uint64(product), nil
}

func Total() uint64 {
	// Geometric series convergence
	return uint64((1 - math.Pow(2, 64)) / (1 - 2))
}
