package main

import (
	"09_deck_of_cards/deck"
	"testing"
)

func TestScore(t *testing.T) {
	cases := []struct {
		hand     Hand
		expected int
	}{
		{
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Ten}, deck.Card{Suit: deck.Spades, Value: deck.Eight}},
			expected: 18,
		},
		{
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.King}, deck.Card{Suit: deck.Spades, Value: deck.Queen}},
			expected: 20,
		},
		{
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Ace}, deck.Card{Suit: deck.Spades, Value: deck.Five}},
			expected: 16,
		},
		{
			// Ace drops from 11 to 1 to avoid bust
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Ace}, deck.Card{Suit: deck.Spades, Value: deck.King}, deck.Card{Suit: deck.Spades, Value: deck.Five}},
			expected: 16,
		},
		{
			// Two aces: one stays 11, one drops to 1
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Ace}, deck.Card{Suit: deck.Spades, Value: deck.Ace}},
			expected: 12,
		},
		{
			// bust — no aces to rescue
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.King}, deck.Card{Suit: deck.Spades, Value: deck.Queen}, deck.Card{Suit: deck.Spades, Value: deck.Five}},
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
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Ace}, deck.Card{Suit: deck.Spades, Value: deck.Six}},
			expected: true,
		},
		{
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Ten}, deck.Card{Suit: deck.Spades, Value: deck.Seven}},
			expected: false,
		},
		{
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Ace}, deck.Card{Suit: deck.Spades, Value: deck.King}},
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
			// Ace + King = 21 with 2 cards → blackjack
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Ace}, deck.Card{Suit: deck.Spades, Value: deck.King}},
			expected: true,
		},
		{
			// Ace + Ten = 21 → blackjack
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Ace}, deck.Card{Suit: deck.Spades, Value: deck.Ten}},
			expected: true,
		},
		{
			// 3 cards summing to 21 → not blackjack
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Seven}, deck.Card{Suit: deck.Spades, Value: deck.Seven}, deck.Card{Suit: deck.Spades, Value: deck.Seven}},
			expected: false,
		},
		{
			// 2 cards but not 21
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Ten}, deck.Card{Suit: deck.Spades, Value: deck.Eight}},
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

func TestPlayRoundPlayerWins(t *testing.T) {
	// player: King + Queen = 20, dealer: Seven + Eight = 15 → dealer hits → dealer: 15 + Nine = 24 → bust
	cards := []deck.Card{
		{Suit: deck.Spades, Value: deck.King},   // dealer card 1
		{Suit: deck.Spades, Value: deck.Queen},  // player card 1
		{Suit: deck.Spades, Value: deck.Seven},  // dealer card 2
		{Suit: deck.Spades, Value: deck.Jack},   // player card 2
		{Suit: deck.Spades, Value: deck.Nine},   // dealer hits
	}
	result, _ := playRound(cards, func(_ Hand, _ deck.Card) bool { return false }) // player stands
	if result != Win {
		t.Errorf("expected Win, got %v", result)
	}
}

func TestPlayRoundDealerWins(t *testing.T) {
	// player: Two + Three = 5, stands; dealer: King + Nine = 19
	cards := []deck.Card{
		{Suit: deck.Spades, Value: deck.King},  // dealer card 1
		{Suit: deck.Spades, Value: deck.Two},   // player card 1
		{Suit: deck.Spades, Value: deck.Nine},  // dealer card 2
		{Suit: deck.Spades, Value: deck.Three}, // player card 2
	}
	result, _ := playRound(cards, func(_ Hand, _ deck.Card) bool { return false }) // player stands
	if result != Lose {
		t.Errorf("expected Lose, got %v", result)
	}
}

func TestPlayRoundPlayerBlackjack(t *testing.T) {
	// player: Ace + King = blackjack
	cards := []deck.Card{
		{Suit: deck.Spades, Value: deck.Two},  // dealer card 1
		{Suit: deck.Spades, Value: deck.Ace},  // player card 1
		{Suit: deck.Spades, Value: deck.Five}, // dealer card 2
		{Suit: deck.Spades, Value: deck.King}, // player card 2
	}
	result, _ := playRound(cards, func(_ Hand, _ deck.Card) bool { return false })
	if result != Win {
		t.Errorf("expected Win (blackjack), got %v", result)
	}
}
