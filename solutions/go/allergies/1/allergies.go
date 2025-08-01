package allergies

import (
	"slices"
)

var Allergens []string = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

var lenAllergens int = len(Allergens)

func Allergies(allergies uint) []string {
	res := []string{}
	for _, allergen := range Allergens {
		if AllergicTo(allergies, allergen) {
			res = append(res, allergen)
		}
	}

	return res
}

func AllergicTo(allergies uint, allergen string) bool {
	i := slices.Index(Allergens, allergen)

	if i == -1 {
		return false
	}

	return (uint(1<<i) & allergies) > 0
}
