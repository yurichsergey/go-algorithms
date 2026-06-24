package deck

type Suit int

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
)

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
)

type Card struct {
	Suit  Suit
	Value Value
}
