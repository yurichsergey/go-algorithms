package _01_arrays_hashing

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 20, 4, 10, 3, 4, 5},
			want: 4,
		},
		{
			name: "Example 2",
			nums: []int{0, 3, 2, 5, 4, 6, 1, 1},
			want: 7,
		},
		{
			name: "Example 3",
			nums: []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6},
			want: 7,
		},
		{
			name: "Empty",
			nums: []int{},
			want: 0,
		},
		{
			name: "Single element",
			nums: []int{100},
			want: 1,
		},
	}

	functions := []struct {
		name string
		f    func([]int) int
	}{
		{"BruteForce", longestConsecutiveBruteForce},
		{"Sort", longestConsecutiveSort},
		{"Hash Set", longestConsecutiveHashSet},
		{"Hash Map", longestConsecutiveHashMap},
	}

	for _, tc := range testCases {
		for _, fn := range functions {
			t.Run(tc.name+"_"+fn.name, func(t *testing.T) {
				got := fn.f(tc.nums)
				if got != tc.want {
					t.Errorf("%s(%v) = %d; want %d", fn.name, tc.nums, got, tc.want)
				}
			})
		}
	}
}
