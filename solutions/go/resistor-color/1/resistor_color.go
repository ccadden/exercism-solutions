package resistorcolor

import (
	"strings"
)

var ColorMap = map[string]int{
	"black": 0,
	"brown": 1,
	"red": 2,
	"orange": 3,
	"yellow": 4,
	"green": 5,
	"blue": 6,
	"violet": 7,
	"grey": 8,
	"white": 9,
}

func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}


// Colors returns the list of all colors.
func Colors() []string {
	return Keys(ColorMap)
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return ColorMap[strings.ToLower(color)]
}
