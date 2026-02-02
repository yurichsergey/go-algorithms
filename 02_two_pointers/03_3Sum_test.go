package _2_two_pointers

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSumBruteForce(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{-3, 1, 2, -3, 4, 0},
			want: [][]int{{-3, 1, 2}},
		},
		{
			name: "Example 2",
			nums: []int{5, 2, -1},
			want: [][]int{},
		},
		{
			name: "Example 3",
			nums: []int{1, 1, -2},
			want: [][]int{{-2, 1, 1}},
		},
		{
			name: "Multiple triplets",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "All zeros",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Large values",
			nums: []int{100000, -100000, 0},
			want: [][]int{{-100000, 0, 100000}},
		},
		{
			name: "Empty input",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "Less than three elements",
			nums: []int{0, 0},
			want: [][]int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := threeSumBruteForce(tc.nums)

			// Normalize results for comparison
			sortTriplets(got)
			sortTriplets(tc.want)

			if len(got) == 0 && len(tc.want) == 0 {
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("threeSumBruteForce(%v) = %v; want %v", tc.nums, got, tc.want)
			}
		})
	}
}

func TestThreeSumHashMap(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{-3, 1, 2, -3, 4, 0},
			want: [][]int{{-3, 1, 2}},
		},
		{
			name: "Example 2",
			nums: []int{5, 2, -1},
			want: [][]int{},
		},
		{
			name: "Example 3",
			nums: []int{1, 1, -2},
			want: [][]int{{-2, 1, 1}},
		},
		{
			name: "Multiple triplets",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "All zeros",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Large values",
			nums: []int{100000, -100000, 0},
			want: [][]int{{-100000, 0, 100000}},
		},
		{
			name: "Empty input",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "Less than three elements",
			nums: []int{0, 0},
			want: [][]int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := threeSumHashMap(tc.nums)

			// Normalize results for comparison
			sortTriplets(got)
			sortTriplets(tc.want)

			if len(got) == 0 && len(tc.want) == 0 {
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("threeSumHashMap(%v) = %v; want %v", tc.nums, got, tc.want)
			}
		})
	}
}

func sortTriplets(triplets [][]int) {
	for i := range triplets {
		sort.Ints(triplets[i])
	}
	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < len(triplets[i]); k++ {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
}
