package blackjack

import (
	"math"

	"09_deck_of_cards/deck"
)

type Result int

const (
	Win Result = iota
	Lose
	Tie
	Blackjack
)

type Options struct {
	Hands           int
	Decks           int
	StartingBalance int
}

type Game struct {
	opts Options
}

func New(opts Options) Game {
	return Game{opts: opts}
}

func (g Game) Play(ai AI) int {
	decks := g.opts.Decks
	if decks == 0 {
		decks = 1
	}
	cards := deck.New(deck.WithMultipleDecks(decks), deck.WithShuffle())
	balance := float64(g.opts.StartingBalance)
	if balance == 0 {
		balance = 100
	}
	startingBalance := balance
	for i := 0; i < g.opts.Hands; i++ {
		if len(cards) < 52 {
			cards = deck.New(deck.WithMultipleDecks(decks), deck.WithShuffle())
		}
		var player, dealer Hand
		player = append(player, cards[1], cards[3])
		dealer = append(dealer, cards[0], cards[2])
		cards = cards[4:]

		bet := min(float64(max(ai.Bet(int(math.Round(balance))), 1)), balance)
		var result Result
		player, result, cards = playHand(player, dealer, cards, ai)
		ai.Result(player, result)
		switch result {
		case Win:
			balance += bet
		case Blackjack:
			balance += bet * 1.5
		case Lose:
			balance -= bet
		case Tie:
			// balance unchanged
		}
	}
	return int(math.Round(balance - startingBalance))
}

func playHand(player, dealer Hand, cards []Card, ai AI) (Hand, Result, []Card) {
	switch {
	case isBlackjack(player) && isBlackjack(dealer):
		return player, Tie, cards
	case isBlackjack(player):
		return player, Blackjack, cards
	case isBlackjack(dealer):
		return player, Lose, cards
	}

	for Score(player) < 21 {
		move := ai.Play(player, dealer[0])
		if move == Stand || len(cards) == 0 {
			break
		}
		player = append(player, cards[0])
		cards = cards[1:]
	}

	playerScore := Score(player)
	if playerScore > 21 {
		return player, Lose, cards
	}

	dealerScore := Score(dealer)
	for dealerScore <= 16 || (dealerScore == 17 && isSoft17(dealer)) {
		if len(cards) == 0 {
			break
		}
		dealer = append(dealer, cards[0])
		cards = cards[1:]
		dealerScore = Score(dealer)
	}
	switch {
	case dealerScore > 21:
		return player, Win, cards
	case playerScore > dealerScore:
		return player, Win, cards
	case playerScore < dealerScore:
		return player, Lose, cards
	default:
		return player, Tie, cards
	}
}
