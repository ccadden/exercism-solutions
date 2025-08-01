package raindrops

import (
	"fmt"
	"math"
	"strings"
)

func Convert(number int) string {
	numFloat := float64(number)
	var b strings.Builder

	if math.Mod(numFloat, 3) == 0 {
		b.WriteString("Pling")
	}
	if math.Mod(numFloat, 5) == 0 {
		b.WriteString("Plang")
	}
	if math.Mod(numFloat, 7) == 0 {
		b.WriteString("Plong")
	}

	output := b.String()

	if len(output) == 0 {
		return fmt.Sprint(number)
	} else {
		return output
	}
}
