package darts

import "math"

func Score(x, y float64) int {
	dist := math.Sqrt(x*x + y*y)

	switch {
	case dist <= 10 && dist > 5:
		return 1
	case dist <= 5 && dist > 1:
		return 5
	case dist <= 1:
		return 10
	default:
		return 0
	}
}
