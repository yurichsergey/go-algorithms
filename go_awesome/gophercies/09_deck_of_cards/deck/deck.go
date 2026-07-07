package deck

import (
	"math/rand/v2"
	"slices"
	"sort"
)

// Option is a function that transforms a slice of Cards. New applies a
// series of Options, in order, to customize the deck it returns.
type Option func([]Card) []Card

// New builds a standard 52-card deck in the default order (all Spades
// A-K, then Diamonds, Clubs, and Hearts), then applies each of the
// given options, in order, to the result.
func New(opts ...Option) []Card {
	var cards []Card

	// TODO: this bound relies on Joker being declared after Hearts in
	// the Suit const block (card.go) to stay excluded from the
	// standard deck. Reordering that block would silently include
	// Joker here with no compiler or test error.
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

// WithShuffle returns an Option that randomly shuffles the deck.
func WithShuffle() Option {
	return func(cards []Card) []Card {
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
		return cards
	}
}

// WithFilter returns an Option that keeps only the cards for which f
// returns true, discarding the rest.
func WithFilter(f func(Card) bool) Option {
	return func(cards []Card) []Card {
		filteredCards := []Card{}
		for _, card := range cards {
			if f(card) {
				filteredCards = append(filteredCards, card)
			}
		}
		return filteredCards
	}
}

// WithSort returns an Option that sorts the deck using the given less
// function, which should report whether c1 belongs before c2.
func WithSort(less func(c1, c2 Card) bool) Option {
	return func(cards []Card) []Card {
		sort.Slice(cards, func(i, j int) bool {
			return less(cards[i], cards[j])
		})
		return cards
	}
}

// DefaultLess orders cards the way a new deck comes: by Suit first
// (Spades, Diamonds, Clubs, Hearts), then by Value within each suit.
//
// TODO: jokers sorting last is a side effect of Joker's iota value
// being greater than Hearts's, not an explicit rule here. Reordering
// the Suit const block would silently change joker sort position.
func DefaultLess(c1, c2 Card) bool {
	return c1.Suit < c2.Suit || (c1.Suit == c2.Suit && c1.Value < c2.Value)
}

// WithJokers returns an Option that appends n jokers to the deck.
func WithJokers(n int) Option {
	return func(cards []Card) []Card {
		return append(cards, slices.Repeat([]Card{{Suit: Joker, Value: JokerRank}}, n)...)
	}
}

// WithMultipleDecks returns an Option that replaces the deck with n
// copies of itself concatenated together, useful for games like
// blackjack that are played with several decks shuffled as one.
func WithMultipleDecks(n int) Option {
	return func(cards []Card) []Card {
		return slices.Repeat(cards, n)
	}
}
