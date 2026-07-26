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
