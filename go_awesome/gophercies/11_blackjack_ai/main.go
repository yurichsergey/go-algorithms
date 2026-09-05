package main

import (
	"fmt"

	"11_blackjack_ai/blackjack"
)

type AI struct{}

func (a AI) Bet(_ int) int { return 1 }
func (a AI) Play(hand blackjack.Hand, dealer blackjack.Card) blackjack.Move {
	score := blackjack.Score(hand)
	if score <= 11 {
		return blackjack.Hit
	}
	if score >= 17 {
		return blackjack.Stand
	}
	if blackjack.CardPoints(dealer) >= 7 {
		return blackjack.Hit
	}
	return blackjack.Stand
}
func (a AI) Result(_ blackjack.Hand, _ blackjack.Result) {}

func main() {
	opts := blackjack.Options{Hands: 1000, Decks: 3}
	game := blackjack.New(opts)
	winnings := game.Play(AI{})
	fmt.Println("Winnings:", winnings)
}
