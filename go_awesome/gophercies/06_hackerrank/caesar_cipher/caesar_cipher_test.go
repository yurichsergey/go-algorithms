package caesar_cipher

import "testing"

func TestCaesarCipher(t *testing.T) {
	cases := []struct {
		input    string
		k        int
		expected string
	}{
		{"abc", 1, "bcd"},
		{"xyz", 3, "abc"},
		{"ABC", 1, "BCD"},
		{"a-b-c", 1, "b-c-d"},
		{"a", 27, "b"},
		{"middle-Outz", 2, "okffng-Qwvb"},
	}

	for _, tc := range cases {
		got := CaesarCipher(tc.input, tc.k)
		if got != tc.expected {
			t.Errorf("CaesarCipher(%q, %d) = %q; want %q", tc.input, tc.k, got, tc.expected)
		}
	}
}
