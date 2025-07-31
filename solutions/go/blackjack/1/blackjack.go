package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "king" || card == "queen" || card == "jack" || card == "ten":
		return 10
	case card == "nine":
		return 9
	case card == "eight":
		return 8
	case card == "seven":
		return 7
	case card == "six":
		return 6
	case card == "five":
		return 5
	case card == "four":
		return 4
	case card == "three":
		return 3
	case card == "two":
		return 2
	case card == "one":
		return 1
	default:
		return 0
	}

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardSum := ParseCard(card1) + ParseCard(card2) 
	dealerValue := ParseCard(dealerCard) 

	switch {
	case cardSum == 22:
		return "P"
	case cardSum == 21:
		if dealerValue < 10 {
			return "W"
		} else {
			return "S"
		}
	case cardSum > 16 && cardSum < 21:
		return "S"
	case cardSum > 11 && cardSum < 17:
		if dealerValue > 6 {
			return "H"
		} else {
			return "S"
		}
	default:
		return "H"
	}
}
