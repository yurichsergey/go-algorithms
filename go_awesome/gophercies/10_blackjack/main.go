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

func main() {
	fmt.Println("Blackjack!")
	cards := deck.New(deck.WithShuffle())

	var player Hand
	var dealer Hand
	player = append(player, cards[1], cards[3])
	dealer = append(dealer, cards[0], cards[2])
	cards = cards[4:]

	var input string
	for Score(player) <= 21 {
		fmt.Printf("Player: %s (score %d)\n", player, Score(player))
		fmt.Printf("Dealer: %s [hidden]\n", dealer[0])
		fmt.Printf("Hit or Stand? (h/s)")
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}

		if input == "s" {
			break
		}
		player = append(player, cards[0])
		cards = cards[1:]

		if Score(player) > 21 {
			fmt.Println("You busted! Dealer wins")
			return
		}
	}

	for Score(dealer) <= 16 || isSoft17(dealer) {
		dealer = append(dealer, cards[0])
		cards = cards[1:]
	}

	fmt.Printf("Dealer: %s (score: %d)\n", dealer, Score(dealer))
	playerScore := Score(player)
	dealerScore := Score(dealer)

	switch {
	case dealerScore > 21:
		fmt.Printf("Dealer busted! You win!\n")
	case playerScore > dealerScore:
		fmt.Printf("You win!\n")
	case playerScore < dealerScore:
		fmt.Printf("Dealer wins!\n")
	default:
		fmt.Printf("It's a tie!\n")

	}
}
