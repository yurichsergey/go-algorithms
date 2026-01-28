package _2_two_pointers

import (
	"reflect"
	"testing"
)

func TestTwoIntegersSum2(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "example_basic",
			numbers: []int{1, 2, 3, 4},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "classic_two_sum_ii",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "includes_negative_values",
			numbers: []int{-5, -2, 0, 1, 3, 9},
			target:  -4,
			want:    []int{1, 4},
		},
		{
			name:    "duplicate_values_pair",
			numbers: []int{0, 0, 3, 4},
			target:  0,
			want:    []int{1, 2},
		},
		{
			name:    "pair_in_middle",
			numbers: []int{1, 2, 3, 4, 8},
			target:  7,
			want:    []int{3, 4},
		},
		{
			name:    "pair_is_duplicates",
			numbers: []int{1, 2, 3, 4, 4, 9, 56, 90},
			target:  8,
			want:    []int{4, 5},
		},
		{
			name:    "length_two",
			numbers: []int{5, 7},
			target:  12,
			want:    []int{1, 2},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := twoIntegersSum2(tc.numbers, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("twoIntegersSum2(%v, %d) = %v; want %v", tc.numbers, tc.target, got, tc.want)
			}
		})
	}
}
