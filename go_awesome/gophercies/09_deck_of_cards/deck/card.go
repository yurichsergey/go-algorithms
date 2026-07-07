package deck

import "fmt"

// Suit represents one of the four suits in a standard deck of cards,
// plus a special Joker value used for jokers.
//
//go:generate go tool stringer -type=Suit
type Suit int

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
	Joker
)

// Value represents the rank of a card, from Ace through King, plus a
// special Joker value used for jokers so they can't be mistaken for a
// real rank (e.g. Ace, which shares Value's zero value).
//
//go:generate go tool stringer -type=Value
type Value int

const (
	Ace Value = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	JokerRank
)

// Card represents a single playing card, defined by a Suit and a Value.
// A Card with Suit == Joker represents a joker, in which case its Value
// is meaningless and should be ignored.
type Card struct {
	Suit  Suit
	Value Value
}

// String returns a human-readable representation of the card, such as
// "Ace of Spades", or "Joker" for a joker.
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %s", c.Value, c.Suit)
}
