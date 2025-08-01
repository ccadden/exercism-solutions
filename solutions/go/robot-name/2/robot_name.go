package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

var issuedNames map[string]bool = map[string]bool{}
var numIssuedNames int = 0

const maxNumNames int = 26 * 26 * 10 * 10 * 10

// Define the Robot type here.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if numIssuedNames == maxNumNames {
		return "", errors.New("Issued maximum number of names")
	}

	if r.name != "" {
		return r.name, nil
	}

	var b strings.Builder

	b.WriteRune(randRune())
	b.WriteRune(randRune())

	num := randRange(0, 999)

	name := fmt.Sprintf("%s%03d", b.String(), num)

	issued, exists := issuedNames[name]

	if issued || !exists {
		return r.Name()
	}

	r.name = name
	issuedNames[name] = true
	numIssuedNames++

	return name, nil
}

func (r *Robot) Reset() {
	issuedNames[r.name] = false
	numIssuedNames--
	r.name = ""
}

func randRange(minimum, maximum int) int {
	return rand.Intn(maximum-minimum) + minimum
}

func randRune() rune {
	return rune(randRange(65, 90))
}
