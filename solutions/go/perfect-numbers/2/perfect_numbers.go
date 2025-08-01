package perfect

import (
	"errors"
)

type Classification int

const (
	ClassificationDeficient = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("Can only classify positive integers")

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return 0, ErrOnlyPositive
	}

	factors := findFactors(n)

	var sum int64 = 0

	for _, factor := range factors {
		sum += factor
	}

	switch {
	case sum > n:
		return ClassificationAbundant, nil
	case sum < n:
		return ClassificationDeficient, nil
	default:
		return ClassificationPerfect, nil
	}
}

func findFactors(n int64) []int64 {
	factors := []int64{}
	var i int64

	for i = 1; i <= n/2; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}

	}
	return factors
}
