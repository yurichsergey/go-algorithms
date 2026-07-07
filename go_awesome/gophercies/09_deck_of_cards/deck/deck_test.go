package deck

import "testing"

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 52 {
		t.Fatalf("len(cards) = %d, want 52", len(cards))
	}

	if want := (Card{Suit: Spades, Value: Ace}); cards[0] != want {
		t.Errorf("cards[0] = %v, want %v", cards[0], want)
	}

	if want := (Card{Suit: Hearts, Value: King}); cards[len(cards)-1] != want {
		t.Errorf("cards[len(cards)-1] = %v, want %v", cards[len(cards)-1], want)
	}
}

func TestWithShuffle(t *testing.T) {
	// TODO: build an unshuffled deck (New()) and a shuffled one
	// (New(WithShuffle())). Assert they're the same length, then
	// assert they contain the same multiset of cards by sorting
	// both with WithSort(DefaultLess) and comparing.
	// Hint: reflect.DeepEqual works on slices of comparable structs.
}

func TestWithFilter(t *testing.T) {
	// TODO: build a deck filtering out Two and Three, then assert
	// no card in the result has Value == Two or Value == Three,
	// and that the length is 52 minus the number removed (4 suits
	// x 2 values = 8).
}

func TestWithSortDefaultLess(t *testing.T) {
	// TODO: build a shuffled-then-sorted deck
	// (New(WithShuffle(), WithSort(DefaultLess))) and assert it's
	// deeply equal to a plain New() deck.
}

func TestWithJokers(t *testing.T) {
	// TODO: build a deck with New(WithJokers(3)), assert the total
	// length is 52+3, and that exactly 3 cards have Suit == Joker.
}

func TestWithMultipleDecks(t *testing.T) {
	// TODO: table-driven test over n = 1, 2, 3, 5, asserting
	// len(New(WithMultipleDecks(n))) == n*52 for each.
}

func TestCardString(t *testing.T) {
	// TODO: table-driven test with a few Card values (including
	// one with Suit: Joker) mapped to their expected String() output.
}
