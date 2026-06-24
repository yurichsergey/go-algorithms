package deck

import (
	"math/rand/v2"
	"sort"
)

type Option func([]Card) []Card

func New(opts ...Option) []Card {
	var cards []Card

	for suit := Spades; suit <= Hearts; suit++ {
		for value := Ace; value <= King; value++ {
			cards = append(cards, Card{Suit: suit, Value: value})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

func WithShuffle() Option {
	return func(cards []Card) []Card {
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
		return cards
	}
}

func WithFilter(f func(Card) bool) Option {
	return func(cards []Card) []Card {
		var filteredCards []Card
		for _, card := range cards {
			if f(card) {
				filteredCards = append(filteredCards, card)
			}
		}
		return filteredCards
	}
}

func WithSort(less func(i, j int) bool) Option {
	return func(cards []Card) []Card {
		sort.Slice(cards, less)
		return cards
	}
}
