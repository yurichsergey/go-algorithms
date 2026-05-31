package camelcase

import "testing"

func TestCountWords(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"saveChangesInTheEditor", 5},
		{"oneTwo", 2},
		{"save", 1},
		{"", 0},
		{"aA", 2},
	}

	for _, tc := range cases {
		got := CountWords(tc.input)
		if got != tc.expected {
			t.Errorf("CountWords(%q) = %d, want %d", tc.input, got, tc.expected)
		}
	}
}
