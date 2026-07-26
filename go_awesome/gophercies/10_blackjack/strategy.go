package main

import (
	"09_deck_of_cards/deck"
	"fmt"
)

// PlayerStrategy decides whether the player should hit (true) or stand (false).
type PlayerStrategy func(player Hand, dealerVisible deck.Card) bool

// HumanStrategy asks the user to type h (hit) or s (stand).
func HumanStrategy(_ Hand, _ deck.Card) bool {
	var input string
	fmt.Printf("Hit or Stand? (h/s): ")
	_, err := fmt.Scan(&input)
	if err != nil {
		return false
	}
	return input == "h"
}

// BasicStrategy uses simplified optimal blackjack strategy.
func BasicStrategy(player Hand, dealerVisible deck.Card) bool {
	score := Score(player)
	if score <= 11 {
		return true
	}
	if score >= 17 {
		return false
	}
	// score 12-16: hit if dealer shows 7 or higher
	return scores[dealerVisible.Value] >= 7
}
