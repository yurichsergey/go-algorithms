package blackjack

import (
	"09_deck_of_cards/deck"
	"testing"
)

type mockAI struct {
	move Move
}

func (a mockAI) Bet(_ int) int            { return 1 }
func (a mockAI) Play(_ Hand, _ Card) Move { return a.move }
func (a mockAI) Result(_ Hand, _ Result)  {}

func TestScore(t *testing.T) {
	cases := []struct {
		hand     Hand
		expected int
	}{
		{
			hand:     Hand{{Suit: deck.Spades, Value: deck.Ten}, {Suit: deck.Spades, Value: deck.Eight}},
			expected: 18,
		},
		{
			hand:     Hand{{Suit: deck.Spades, Value: deck.King}, {Suit: deck.Spades, Value: deck.Queen}},
			expected: 20,
		},
		{
			hand:     Hand{{Suit: deck.Spades, Value: deck.Ace}, {Suit: deck.Spades, Value: deck.Five}},
			expected: 16,
		},
		{
			// Ace drops from 11 to 1 to avoid bust
			hand:     Hand{{Suit: deck.Spades, Value: deck.Ace}, {Suit: deck.Spades, Value: deck.King}, {Suit: deck.Spades, Value: deck.Five}},
			expected: 16,
		},
		{
			// Two aces: one stays 11, one drops to 1
			hand:     Hand{{Suit: deck.Spades, Value: deck.Ace}, {Suit: deck.Spades, Value: deck.Ace}},
			expected: 12,
		},
		{
			// bust — no aces to rescue
			hand:     Hand{{Suit: deck.Spades, Value: deck.King}, {Suit: deck.Spades, Value: deck.Queen}, {Suit: deck.Spades, Value: deck.Five}},
			expected: 25,
		},
	}
	for _, tc := range cases {
		got := Score(tc.hand)
		if got != tc.expected {
			t.Errorf("Score(%v) = %d, want %d", tc.hand, got, tc.expected)
		}
	}
}

func TestIsSoft17(t *testing.T) {
	cases := []struct {
		hand     Hand
		expected bool
	}{
		{
			hand:     Hand{{Suit: deck.Spades, Value: deck.Ace}, {Suit: deck.Spades, Value: deck.Six}},
			expected: true,
		},
		{
			hand:     Hand{{Suit: deck.Spades, Value: deck.Ten}, {Suit: deck.Spades, Value: deck.Seven}},
			expected: false,
		},
		{
			hand:     Hand{{Suit: deck.Spades, Value: deck.Ace}, {Suit: deck.Spades, Value: deck.King}},
			expected: false,
		},
		{
			// Ace+Six+Ten = 17 but Ace is demoted to 1 → hard 17, not soft
			hand:     Hand{{Suit: deck.Spades, Value: deck.Ace}, {Suit: deck.Spades, Value: deck.Six}, {Suit: deck.Spades, Value: deck.Ten}},
			expected: false,
		},
	}
	for _, tc := range cases {
		got := isSoft17(tc.hand)
		if got != tc.expected {
			t.Errorf("isSoft17(%v) = %v, want %v", tc.hand, got, tc.expected)
		}
	}
}

func TestIsBlackjack(t *testing.T) {
	cases := []struct {
		hand     Hand
		expected bool
	}{
		{
			hand:     Hand{{Suit: deck.Spades, Value: deck.Ace}, {Suit: deck.Spades, Value: deck.King}},
			expected: true,
		},
		{
			hand:     Hand{{Suit: deck.Spades, Value: deck.Ace}, {Suit: deck.Spades, Value: deck.Ten}},
			expected: true,
		},
		{
			// 3 cards summing to 21 → not blackjack
			hand:     Hand{{Suit: deck.Spades, Value: deck.Seven}, {Suit: deck.Spades, Value: deck.Seven}, {Suit: deck.Spades, Value: deck.Seven}},
			expected: false,
		},
		{
			hand:     Hand{{Suit: deck.Spades, Value: deck.Ten}, {Suit: deck.Spades, Value: deck.Eight}},
			expected: false,
		},
	}
	for _, tc := range cases {
		got := isBlackjack(tc.hand)
		if got != tc.expected {
			t.Errorf("isBlackjack(%v) = %v, want %v", tc.hand, got, tc.expected)
		}
	}
}

func TestPlayHandPlayerWins(t *testing.T) {
	// player: Queen+Jack=20, dealer: King+Seven=17 → player wins
	player := Hand{{Suit: deck.Spades, Value: deck.Queen}, {Suit: deck.Spades, Value: deck.Jack}}
	dealer := Hand{{Suit: deck.Spades, Value: deck.King}, {Suit: deck.Spades, Value: deck.Seven}}
	_, result, _ := playHand(player, dealer, nil, mockAI{move: Stand})
	if result != Win {
		t.Errorf("expected Win, got %v", result)
	}
}

func TestPlayHandDealerWins(t *testing.T) {
	// player: Two+Three=5 stands, dealer: King+Nine=19 → dealer wins
	player := Hand{{Suit: deck.Spades, Value: deck.Two}, {Suit: deck.Spades, Value: deck.Three}}
	dealer := Hand{{Suit: deck.Spades, Value: deck.King}, {Suit: deck.Spades, Value: deck.Nine}}
	_, result, _ := playHand(player, dealer, nil, mockAI{move: Stand})
	if result != Lose {
		t.Errorf("expected Lose, got %v", result)
	}
}

func TestPlayHandTie(t *testing.T) {
	// player: King+Eight=18, dealer: King+Eight=18 → tie
	player := Hand{{Suit: deck.Spades, Value: deck.King}, {Suit: deck.Spades, Value: deck.Eight}}
	dealer := Hand{{Suit: deck.Hearts, Value: deck.King}, {Suit: deck.Hearts, Value: deck.Eight}}
	_, result, _ := playHand(player, dealer, nil, mockAI{move: Stand})
	if result != Tie {
		t.Errorf("expected Tie, got %v", result)
	}
}

func TestPlayHandPlayerBusts(t *testing.T) {
	// player: Ten+Eight=18, hits Four → 22 bust
	player := Hand{{Suit: deck.Spades, Value: deck.Ten}, {Suit: deck.Spades, Value: deck.Eight}}
	dealer := Hand{{Suit: deck.Spades, Value: deck.Two}, {Suit: deck.Spades, Value: deck.Three}}
	extra := []Card{{Suit: deck.Spades, Value: deck.Four}}
	_, result, _ := playHand(player, dealer, extra, mockAI{move: Hit})
	if result != Lose {
		t.Errorf("expected Lose (player bust), got %v", result)
	}
}

func TestPlayHandDealerBusts(t *testing.T) {
	// player: King+Eight=18 stands; dealer: Six+Seven=13 hits King → 23 bust
	player := Hand{{Suit: deck.Spades, Value: deck.King}, {Suit: deck.Spades, Value: deck.Eight}}
	dealer := Hand{{Suit: deck.Spades, Value: deck.Six}, {Suit: deck.Spades, Value: deck.Seven}}
	extra := []Card{{Suit: deck.Spades, Value: deck.King}}
	_, result, _ := playHand(player, dealer, extra, mockAI{move: Stand})
	if result != Win {
		t.Errorf("expected Win (dealer bust), got %v", result)
	}
}

func TestPlayHandPlayerBlackjack(t *testing.T) {
	// player: Ace+King = blackjack
	player := Hand{{Suit: deck.Spades, Value: deck.Ace}, {Suit: deck.Spades, Value: deck.King}}
	dealer := Hand{{Suit: deck.Spades, Value: deck.Two}, {Suit: deck.Spades, Value: deck.Five}}
	_, result, _ := playHand(player, dealer, nil, mockAI{move: Stand})
	if result != Blackjack {
		t.Errorf("expected Blackjack, got %v", result)
	}
}

func TestPlayHandDealerBlackjack(t *testing.T) {
	// dealer: Ace+King = blackjack, player: Ten+Eight=18
	player := Hand{{Suit: deck.Spades, Value: deck.Ten}, {Suit: deck.Spades, Value: deck.Eight}}
	dealer := Hand{{Suit: deck.Spades, Value: deck.Ace}, {Suit: deck.Spades, Value: deck.King}}
	_, result, _ := playHand(player, dealer, nil, mockAI{move: Stand})
	if result != Lose {
		t.Errorf("expected Lose (dealer blackjack), got %v", result)
	}
}

func TestPlayHandBothBlackjack(t *testing.T) {
	// both Ace+King → tie
	player := Hand{{Suit: deck.Hearts, Value: deck.Ace}, {Suit: deck.Hearts, Value: deck.King}}
	dealer := Hand{{Suit: deck.Spades, Value: deck.Ace}, {Suit: deck.Spades, Value: deck.King}}
	_, result, _ := playHand(player, dealer, nil, mockAI{move: Stand})
	if result != Tie {
		t.Errorf("expected Tie (both blackjack), got %v", result)
	}
}
