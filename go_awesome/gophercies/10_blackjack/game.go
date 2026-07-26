package main

import (
	"09_deck_of_cards/deck"
	"fmt"
)

type Result int

const (
	Win Result = iota
	Lose
	Tie
)

func applyResult(balance, bet int, result Result) int {
	switch result {
	case Win:
		return balance + bet
	case Lose:
		return balance - bet
	default:
		return balance
	}
}

func playRound(cards []deck.Card, strategy PlayerStrategy) (Result, []deck.Card) {
	var player, dealer Hand
	player = append(player, cards[1], cards[3])
	dealer = append(dealer, cards[0], cards[2])
	cards = cards[4:]

	var result Result

	switch {
	case isBlackjack(player) && isBlackjack(dealer):
		fmt.Println("Both have blackjack! It's a tie.")
		result = Tie
	case isBlackjack(player):
		fmt.Println("Blackjack! You win!")
		result = Win
	case isBlackjack(dealer):
		fmt.Println("Dealer has blackjack! Dealer wins.")
		result = Lose
	default:
		result, cards = playTurns(player, dealer, cards, strategy)
	}

	return result, cards
}

func playTurns(player, dealer Hand, cards []deck.Card, strategy PlayerStrategy) (Result, []deck.Card) {
	var result Result

	for Score(player) < 21 {
		fmt.Printf("Player: %s (score %d)\n", player, Score(player))
		fmt.Printf("Dealer: %s [hidden]\n", dealer[0])
		if !strategy(player, dealer[0]) {
			break
		}
		player = append(player, cards[0])
		cards = cards[1:]
	}

	if Score(player) > 21 {
		fmt.Printf("Player: %s (score %d)\n", player, Score(player))
		fmt.Println("You busted! Dealer wins")
		result = Lose
	} else {
		dealerScore := Score(dealer)
		for dealerScore <= 16 || isSoft17(dealer) {
			dealer = append(dealer, cards[0])
			cards = cards[1:]
			dealerScore = Score(dealer)
		}
		printStat(player, dealer)
		playerScore, dealerScore := Score(player), Score(dealer)
		switch {
		case dealerScore > 21:
			fmt.Println("Dealer busted! You win!")
			result = Win
		case playerScore > dealerScore:
			fmt.Println("You win!")
			result = Win
		case playerScore < dealerScore:
			fmt.Println("Dealer wins!")
			result = Lose
		default:
			fmt.Println("It's a tie!")
			result = Tie
		}
	}

	return result, cards
}
