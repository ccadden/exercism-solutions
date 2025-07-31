package robotname

import (
	"fmt"
	"math/rand"
	"strings"
)

var issuedNames map[string]bool = map[string]bool{}

// Define the Robot type here.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	var b strings.Builder

	b.WriteRune(randRune())
	b.WriteRune(randRune())

	num := randRange(0, 999)

	name := fmt.Sprintf("%s%03d", b.String(), num)

	_, alreadyIssued := issuedNames[name]

	if alreadyIssued {
		return r.Name()
	}

	r.name = name
	issuedNames[name] = true

	return name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func randRange(minimum, maximum int) int {
	return rand.Intn(maximum-minimum) + minimum
}

func randRune() rune {
	return rune(randRange(65, 90))
}
