package main

import (
	"09_deck_of_cards/deck"
	"fmt"
)

type Hand []deck.Card

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

func (h Hand) String() string {
	result := ""
	for _, c := range h {
		result += c.String() + "  "
	}
	return result
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

func isBlackjack(h Hand) bool {
	return len(h) == 2 && Score(h) == 21
}

func isSoft17(h Hand) bool {
	if Score(h) != 17 {
		return false
	}
	for _, c := range h {
		if c.Value == deck.Ace {
			return true
		}
	}
	return false
}

func printStat(player Hand, dealer Hand) {
	fmt.Printf("Player: %s (score %d)\n", player, Score(player))
	fmt.Printf("Dealer: %s (score: %d)\n", dealer, Score(dealer))
}
