package main

import (
	"09_deck_of_cards/deck"
	"fmt"
)

func main() {
	fmt.Println("Blackjack!")
	balance := 100
	cards := deck.New(deck.WithShuffle())
	strategy := HumanStrategy // swap to BasicStrategy to use AI

	for balance > 0 {
		fmt.Printf("\nBalance: $%d\n", balance)
		fmt.Print("Your bet: $")
		var bet int
		_, err := fmt.Scan(&bet)
		if err != nil {
			break
		}
		if bet <= 0 || bet > balance {
			fmt.Println("Invalid bet.")
			continue
		}

		if len(cards) < 10 {
			cards = deck.New(deck.WithShuffle())
		}

		result, remaining := playRound(cards, strategy)
		cards = remaining

		switch result {
		case Win:
			balance += bet
		case Lose:
			balance -= bet
		}

		fmt.Printf("Balance: $%d\n", balance)
		fmt.Print("Play again? (y/n): ")
		var again string
		fmt.Scan(&again)
		if again != "y" {
			break
		}
	}

	fmt.Printf("\nGame over! Final balance: $%d\n", balance)
}
