package pythagorean

import "math"

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	res := []Triplet{}
	for i := min; i < max; i++ {
		for j := i + 1; j < max; j++ {
			k := math.Sqrt(float64(i*i + j*j))
			if k == math.Floor(k) {
				res = append(res, [3]int{i, j, int(k)})
			}
		}
	}

	return res
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	res := []Triplet{}

	for i := 1; i < p-2; i++ {
		for j := i; j < p-2; j++ {
			k := p - i - j

			if k*k == i*i+j*j {
				res = append(res, [3]int{i, j, int(k)})
			}
		}
	}

	return res
}
