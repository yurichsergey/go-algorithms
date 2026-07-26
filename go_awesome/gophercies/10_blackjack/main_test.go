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
		{
			// Ace+Six+Ten = 17 but Ace is demoted to 1 → hard 17, not soft
			hand:     Hand{deck.Card{Suit: deck.Spades, Value: deck.Ace}, deck.Card{Suit: deck.Spades, Value: deck.Six}, deck.Card{Suit: deck.Spades, Value: deck.Ten}},
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
	// player: Queen + Jack = 20, dealer: King + Seven = 17 → dealer stands; player wins
	cards := []deck.Card{
		{Suit: deck.Spades, Value: deck.King},  // dealer card 1
		{Suit: deck.Spades, Value: deck.Queen}, // player card 1
		{Suit: deck.Spades, Value: deck.Seven}, // dealer card 2
		{Suit: deck.Spades, Value: deck.Jack},  // player card 2
	}
	result, _ := playRound(cards, func(_ Hand, _ deck.Card) bool { return false })
	if result != Win {
		t.Errorf("expected Win, got %v", result)
	}
}

func TestPlayRoundDealerBusts(t *testing.T) {
	// player: King + Eight = 18 stands; dealer: Six + Seven = 13 → hits → King = 23 → bust
	cards := []deck.Card{
		{Suit: deck.Spades, Value: deck.Six},   // dealer card 1
		{Suit: deck.Spades, Value: deck.King},  // player card 1
		{Suit: deck.Spades, Value: deck.Seven}, // dealer card 2
		{Suit: deck.Spades, Value: deck.Eight}, // player card 2
		{Suit: deck.Spades, Value: deck.King},  // dealer hits → 23 bust
	}
	result, _ := playRound(cards, func(_ Hand, _ deck.Card) bool { return false })
	if result != Win {
		t.Errorf("expected Win (dealer bust), got %v", result)
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

func TestPlayRoundDealerBlackjack(t *testing.T) {
	// dealer: Ace + King = blackjack, player: Ten + Eight = 18
	cards := []deck.Card{
		{Suit: deck.Spades, Value: deck.Ace},   // dealer card 1
		{Suit: deck.Spades, Value: deck.Ten},   // player card 1
		{Suit: deck.Spades, Value: deck.King},  // dealer card 2
		{Suit: deck.Spades, Value: deck.Eight}, // player card 2
	}
	result, _ := playRound(cards, func(_ Hand, _ deck.Card) bool { return false })
	if result != Lose {
		t.Errorf("expected Lose (dealer blackjack), got %v", result)
	}
}

func TestPlayRoundBothBlackjack(t *testing.T) {
	// both: Ace + King = blackjack → tie
	cards := []deck.Card{
		{Suit: deck.Spades, Value: deck.Ace},  // dealer card 1
		{Suit: deck.Hearts, Value: deck.Ace},  // player card 1
		{Suit: deck.Spades, Value: deck.King}, // dealer card 2
		{Suit: deck.Hearts, Value: deck.King}, // player card 2
	}
	result, _ := playRound(cards, func(_ Hand, _ deck.Card) bool { return false })
	if result != Tie {
		t.Errorf("expected Tie (both blackjack), got %v", result)
	}
}

func TestPlayRoundPlayerBusts(t *testing.T) {
	// player: Ten + Eight = 18, hits → gets Four = 22 → bust
	cards := []deck.Card{
		{Suit: deck.Spades, Value: deck.Two},   // dealer card 1
		{Suit: deck.Spades, Value: deck.Ten},   // player card 1
		{Suit: deck.Spades, Value: deck.Three}, // dealer card 2
		{Suit: deck.Spades, Value: deck.Eight}, // player card 2
		{Suit: deck.Spades, Value: deck.Four},  // player hits → 22 bust
		{Suit: deck.Spades, Value: deck.Ten},   // extra card for dealer if needed
	}
	result, _ := playRound(cards, func(_ Hand, _ deck.Card) bool { return true }) // always hit
	if result != Lose {
		t.Errorf("expected Lose (player bust), got %v", result)
	}
}

func TestPlayRoundTie(t *testing.T) {
	// player: King + Eight = 18, dealer: King + Eight = 18 → tie
	cards := []deck.Card{
		{Suit: deck.Spades, Value: deck.King},  // dealer card 1
		{Suit: deck.Hearts, Value: deck.King},  // player card 1
		{Suit: deck.Spades, Value: deck.Eight}, // dealer card 2
		{Suit: deck.Hearts, Value: deck.Eight}, // player card 2
	}
	result, _ := playRound(cards, func(_ Hand, _ deck.Card) bool { return false })
	if result != Tie {
		t.Errorf("expected Tie, got %v", result)
	}
}

func TestApplyResult(t *testing.T) {
	cases := []struct {
		balance  int
		bet      int
		result   Result
		expected int
	}{
		{balance: 100, bet: 20, result: Win, expected: 120},
		{balance: 100, bet: 20, result: Lose, expected: 80},
		{balance: 100, bet: 20, result: Tie, expected: 100},
		{balance: 20, bet: 20, result: Lose, expected: 0},
	}
	for _, tc := range cases {
		got := applyResult(tc.balance, tc.bet, tc.result)
		if got != tc.expected {
			t.Errorf("applyResult(%d, %d, %v) = %d, want %d", tc.balance, tc.bet, tc.result, got, tc.expected)
		}
	}
}

func TestBasicStrategy(t *testing.T) {
	anyCard := deck.Card{Suit: deck.Spades, Value: deck.Two}
	cases := []struct {
		player        Hand
		dealerVisible deck.Card
		expectHit     bool
	}{
		{
			// score ≤ 11 → always hit
			player:        Hand{{Suit: deck.Spades, Value: deck.Five}, {Suit: deck.Spades, Value: deck.Four}},
			dealerVisible: anyCard,
			expectHit:     true,
		},
		{
			// score ≥ 17 → always stand
			player:        Hand{{Suit: deck.Spades, Value: deck.King}, {Suit: deck.Spades, Value: deck.Seven}},
			dealerVisible: anyCard,
			expectHit:     false,
		},
		{
			// score 14, dealer shows 9 (≥7) → hit
			player:        Hand{{Suit: deck.Spades, Value: deck.Eight}, {Suit: deck.Spades, Value: deck.Six}},
			dealerVisible: deck.Card{Suit: deck.Spades, Value: deck.Nine},
			expectHit:     true,
		},
		{
			// score 14, dealer shows 5 (<7) → stand
			player:        Hand{{Suit: deck.Spades, Value: deck.Eight}, {Suit: deck.Spades, Value: deck.Six}},
			dealerVisible: deck.Card{Suit: deck.Spades, Value: deck.Five},
			expectHit:     false,
		},
	}
	for _, tc := range cases {
		got := BasicStrategy(tc.player, tc.dealerVisible)
		if got != tc.expectHit {
			t.Errorf("BasicStrategy(score=%d, dealer=%v) = %v, want %v",
				Score(tc.player), tc.dealerVisible, got, tc.expectHit)
		}
	}
}
