package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
	totalFodder, err := fc.FodderAmount(numCows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fc.FatteningFactor()

	if err != nil {
		return 0, err
	}

	totalFodder *= fatteningFactor
	return totalFodder / float64(numCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
	if numCows > 0 {
		return DivideFood(fc, numCows)
	} else {
		return 0, errors.New("invalid number of cows")
	}
}

type InvalidCowsError struct {
	numCows int
	msg string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.msg)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{numCows: numCows, msg: "there are no negative cows"}
	} else if numCows == 0 {
		return &InvalidCowsError{numCows: numCows, msg: "no cows don't need food"}
	} else {
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
