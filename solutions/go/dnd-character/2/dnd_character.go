package dndcharacter

import (
	"math/rand"
	"math"
)

var BASE_HITPOINTS int = 10

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	sum := 0
	minVal := 6

	for i:=0; i<4; i++ {
		val := randInt(1, 7)
		sum += val
		if val < minVal {
			minVal = val
		}
	}

	return sum - minVal
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()
	hitpoints := Modifier(constitution) + BASE_HITPOINTS

	return Character{
		Strength: Ability(),
		Dexterity: Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom: Ability(),
		Charisma: Ability(),
		Hitpoints: hitpoints,
	}
}

func randInt(min, max int) int {
	return min + rand.Intn(max - min)
}
