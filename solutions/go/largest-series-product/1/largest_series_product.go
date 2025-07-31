package lsproduct

import (
	"errors"
	"slices"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	products := []int64{}
	var product int64

	if span > len(digits) || span < 0 {
		return 0, errors.New("Invalid span given")
	}

	for i := 0; i <= len(digits) - span; i ++ {
		product = 1

		for _, digit := range(digits[i: i+span]) {
			val, err := strconv.ParseInt(string(digit), 10, 64)
			if err != nil {
				return 0, errors.New("Problem converting string to int64")
			}

			product *= val
		}


		products = append(products, product)
	}

	return slices.Max(products), nil
}
