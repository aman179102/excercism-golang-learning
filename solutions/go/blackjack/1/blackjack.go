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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn.
func FirstTurn(card1, card2, dealerCard string) string {
	p1 := ParseCard(card1)
	p2 := ParseCard(card2)
	playerTotal := p1 + p2
	dealer := ParseCard(dealerCard)

	// Rule 1: Always split aces.
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	// Rule 2: Blackjack conditions.
	if playerTotal == 21 {
		// If dealer does NOT have ace, face card, or 10 → Automatic win.
		if dealer != 11 && dealer != 10 {
			return "W"
		}
		// Else → Stand.
		return "S"
	}

	// Rule 3: Stand on 17–20
	if playerTotal >= 17 && playerTotal <= 20 {
		return "S"
	}

	// Rule 4: 12–16
	if playerTotal >= 12 && playerTotal <= 16 {
		if dealer >= 7 {
			return "H"
		}
		return "S"
	}

	// Rule 5: 11 or lower → Hit
	return "H"
}
