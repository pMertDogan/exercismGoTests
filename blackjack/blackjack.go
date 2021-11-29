package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {

	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "other":
		return 0
	default:
		return 10

	}
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack && dealerScore >= 10 {
		return "S"
	} else if dealerScore ==11{
		return "P"
	}else {
		return "W"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch handScore {
		
	case 16, 15, 14, 13, 12:
		if dealerScore >= 7 {
			return "H"
		} else {
			return "S"
		}

	case 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0:
		return "H"
	default:
		return "S"
	}

}
