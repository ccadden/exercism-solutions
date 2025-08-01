package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

var issuedNames map[string]bool = map[string]bool{}
var numIssuedNames int = 0
var namesNotGenerated bool = true

const maxNumNames int = 26 * 26 * 10 * 10 * 10

var names []string = make([]string, maxNumNames, maxNumNames)

// Define the Robot type here.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if namesNotGenerated {
		generateNames()
		namesNotGenerated = false
	}

	if r.name != "" {
		return r.name, nil
	}

	if len(names) == 0 {
		return "", errors.New("Issued maximum number of names")
	}

	index := rand.Intn(len(names))

	result := names[index]
	r.name = result

	names = append(names[:index], names[index+1:]...)

	return result, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func generateNames() {
	counter := 0

	var i rune
	var j rune

	for i = 65; i < 91; i++ {
		for j = 65; j < 91; j++ {
			var b strings.Builder

			b.WriteRune(i)
			b.WriteRune(j)

			for k := 0; k < 1000; k++ {
				names[counter] = fmt.Sprintf("%s%03d", b.String(), k)
				counter++
			}
		}
	}
}
