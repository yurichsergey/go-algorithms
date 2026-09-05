package blackjack

import "09_deck_of_cards/deck"

type Card = deck.Card

type Hand []Card

var scores = map[deck.Value]int{
	deck.Ace:   11,
	deck.Two:   2,
	deck.Three: 3,
	deck.Four:  4,
	deck.Five:  5,
	deck.Six:   6,
	deck.Seven: 7,
	deck.Eight: 8,
	deck.Nine:  9,
	deck.Ten:   10,
	deck.Jack:  10,
	deck.Queen: 10,
	deck.King:  10,
}

func isBlackjack(h Hand) bool {
	return len(h) == 2 && Score(h) == 21
}

func isSoft17(h Hand) bool {
	if Score(h) != 17 {
		return false
	}
	minTotal := 0
	for _, c := range h {
		if c.Value == deck.Ace {
			minTotal++
		} else {
			minTotal += scores[c.Value]
		}
	}
	return minTotal != 17
}

func CardPoints(c Card) int {
	return scores[c.Value]
}

func Score(h Hand) int {
	total := 0
	aces := 0
	for _, card := range h {
		total += scores[card.Value]
		if card.Value == deck.Ace {
			aces++
		}
	}
	for total > 21 && aces > 0 {
		total -= 10
		aces--
	}
	return total
}
