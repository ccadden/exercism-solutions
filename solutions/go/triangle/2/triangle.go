// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle


// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT Kind = iota
    Equ
    Iso
    Sca
)

func IsTriangle(a, b, c float64) bool {
	return a + b > c && a + c > b && b + c > a
}

func IsEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

func IsIsosceles(a, b, c float64) bool {
	return a == b || b == c || a == c
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if !IsTriangle(a, b, c) {
		return NaT
	} else if IsEquilateral(a, b, c) {
		return Equ
	} else if IsIsosceles(a, b, c) {
		return Iso
	} else {
		return Sca
	}
}
