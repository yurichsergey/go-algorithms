package _2_two_pointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example_true_phrase_with_punctuation",
			s:    "Was it a car or a cat I saw?",
			want: true,
		},
		{
			name: "example_false_phrase",
			s:    "tab a cat",
			want: false,
		},
		{
			name: "classic_true_panama",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "classic_false_race_a_car",
			s:    "race a car",
			want: false,
		},
		{
			name: "only_spaces",
			s:    "     ",
			want: true,
		},
		{
			name: "punctuation_only",
			s:    ".,,!!",
			want: true,
		},
		{
			name: "numeric_and_letters_false",
			s:    "0P",
			want: false,
		},
		{
			name: "single_character",
			s:    "a",
			want: true,
		},
		{
			name: "even_length_palindrome",
			s:    "abba",
			want: true,
		},
		{
			name: "odd_length_palindrome",
			s:    "abcba",
			want: true,
		},
		{
			name: "near_palindrome_false",
			s:    "abca",
			want: false,
		},
		{
			name: "mixed_case_true",
			s:    "No 'x' in Nixon",
			want: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := isPalindrome(tc.s)
			if got != tc.want {
				t.Fatalf("isPalindrome(%q) = %v; want %v", tc.s, got, tc.want)
			}
		})
	}
}
