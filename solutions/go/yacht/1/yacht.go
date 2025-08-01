package yacht

import (
	"slices"
)

func Score(dice []int, category string) int {
	var score int

	switch category {
	case "ones":
		score = count(dice, 1)
	case "twos":
		score = count(dice, 2) * 2
	case "threes":
		score = count(dice, 3) * 3
	case "fours":
		score = count(dice, 4) * 4
	case "fives":
		score = count(dice, 5) * 5
	case "sixes":
		score = count(dice, 6) * 6
	case "full house":
		score = fullHouse(dice)
	case "four of a kind":
		score = fourOfAKind(dice)
	case "little straight":
		slices.Sort(dice)
		if slices.Equal(dice, []int{1, 2, 3, 4, 5}) {
			score = 30
		}
	case "big straight":
		slices.Sort(dice)
		if slices.Equal(dice, []int{2, 3, 4, 5, 6}) {
			score = 30
		}
	case "choice":
		score = sum(dice)
	case "yacht":
		if len(keys(tally(dice))) == 1 {
			score = 50
		}
	}

	return score
}

func count(dice []int, value int) int {
	count := 0

	for _, die := range dice {
		if die == value {
			count++
		}
	}

	return count
}

func sum(dice []int) int {
	count := 0

	for _, die := range dice {
		count += die
	}

	return count
}

func tally(dice []int) map[int]int {
	t := map[int]int{}

	for _, die := range dice {
		t[die] += 1
	}

	return t
}

func keys(diceCount map[int]int) []int {
	res := []int{}

	for k := range diceCount {
		res = append(res, k)
	}

	return res
}

func fourOfAKind(dice []int) int {
	score := 0

	for k, v := range tally(dice) {
		if v >= 4 {
			score = k * 4
		}
	}

	return score
}

func fullHouse(dice []int) int {
	score := 0

	hasThree := false
	hasTwo := false

	t := tally(dice)

	for _, v := range t {
		if v == 2 {
			hasTwo = true
		}

		if v == 3 {
			hasThree = true
		}
	}

	if hasThree && hasTwo {
		score = sum(dice)
	}

	return score
}
