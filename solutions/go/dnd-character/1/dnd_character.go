package dndcharacter

import (
	"math/rand"
	"math"
	"sort"
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

type Dice []int

func (d Dice) Len() int {
	return len(d)
}

func (d Dice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Dice) Less(i, j int) bool {
	return d[i] < d[j]
}

func (d Dice) Sum() int {
	sum := 0
	for _, val := range d {
		sum += val
	}

	return sum
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	dice := Dice{}
	for i:=0; i<4; i++ {
		dice = append(dice, randInt(1, 7))
	}

	sort.Sort(dice)

	return dice[1:].Sum()
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
