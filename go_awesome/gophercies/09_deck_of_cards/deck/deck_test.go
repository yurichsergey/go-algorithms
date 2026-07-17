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
	unshuffled := New(WithSort(DefaultLess))
	shuffled := New(WithShuffle(), WithSort(DefaultLess))
	if len(unshuffled) != len(shuffled) {
		t.Fatalf("length mismatch: %d vs %d", len(unshuffled), len(shuffled))
	}
	for i := range unshuffled {
		if unshuffled[i] != shuffled[i] {
			t.Errorf("card mismatch at %d: %v vs %v", i, unshuffled[i], shuffled[i])
		}
	}
}

func TestWithFilter(t *testing.T) {
	cards := New(WithFilter(func(c Card) bool {
		return c.Value != Two && c.Value != Three
	}))
	if len(cards) != 44 {
		t.Errorf("expected 44 cards, got %d", len(cards))
	}
	for _, c := range cards {
		if c.Value == Two || c.Value == Three {
			t.Errorf("found filtered card: %v", c)
		}
	}
}

func TestWithSortDefaultLess(t *testing.T) {
	sorted := New(WithShuffle(), WithSort(DefaultLess))
	expected := New()
	if len(sorted) != len(expected) {
		t.Fatalf("length mismatch: %d vs %d", len(sorted), len(expected))
	}
	for i := range expected {
		if sorted[i] != expected[i] {
			t.Errorf("card mismatch at %d: got %v, want %v", i, sorted[i], expected[i])
		}
	}
}

func TestWithJokers(t *testing.T) {
	cards := New(WithJokers(3))
	if len(cards) != 55 {
		t.Errorf("expected 55 cards, got %d", len(cards))
	}
	jokers := 0
	for _, c := range cards {
		if c.Suit == Joker {
			jokers++
		}
	}
	if jokers != 3 {
		t.Errorf("expected 3 jokers, got %d", jokers)
	}
}

func TestWithMultipleDecks(t *testing.T) {
	cases := []int{1, 2, 3, 5}
	for _, n := range cases {
		cards := New(WithMultipleDecks(n))
		if len(cards) != n*52 {
			t.Errorf("WithMultipleDecks(%d): expected %d cards, got %d", n, n*52, len(cards))
		}
	}
}

func TestCardString(t *testing.T) {
	cases := []struct {
		card     Card
		expected string
	}{
		{Card{Suit: Spades, Value: Ace}, "Ace of Spades"},
		{Card{Suit: Hearts, Value: King}, "King of Hearts"},
		{Card{Suit: Diamonds, Value: Ten}, "Ten of Diamonds"},
		{Card{Suit: Joker, Value: JokerRank}, "Joker"},
	}
	for _, tc := range cases {
		if got := tc.card.String(); got != tc.expected {
			t.Errorf("Card(%v).String() = %q, want %q", tc.card, got, tc.expected)
		}
	}
}
