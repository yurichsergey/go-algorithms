package _2_two_pointers

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Multiple triplets",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
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
		{
			name: "Comprehensive with many valid triplets",
			nums: []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			want: [][]int{
				{-4, 0, 4},
				{-4, 1, 3},
				{-3, -1, 4},
				{-3, 0, 3},
				{-3, 1, 2},
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}

	funcs := []struct {
		name string
		f    func([]int) [][]int
	}{
		{"BruteForce", threeSumBruteForce},
		{"HashMap", threeSumHashMap},
		{"TwoPointers", threeSumTwoPointers},
		{"TwoPointersEfficient", threeSumTwoPointersEfficient},
	}

	for _, fn := range funcs {
		testHelper(t, fn, tests)
	}
}

func testHelper(t *testing.T, fn struct {
	name string
	f    func([]int) [][]int
}, tests []struct {
	name string
	nums []int
	want [][]int
}) bool {
	return t.Run(fn.name, func(t *testing.T) {
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				// We need to pass a copy of nums because some implementations might sort it in place
				numsCopy := make([]int, len(tc.nums))
				copy(numsCopy, tc.nums)

				got := fn.f(numsCopy)

				// Normalize results for comparison
				sortTriplets(got)

				// We must NOT sort tc.want in place because it's shared between tests
				wantCopy := make([][]int, len(tc.want))
				for i := range tc.want {
					wantCopy[i] = make([]int, len(tc.want[i]))
					copy(wantCopy[i], tc.want[i])
				}
				sortTriplets(wantCopy)

				if len(got) == 0 && len(wantCopy) == 0 {
					return
				}

				if !reflect.DeepEqual(got, wantCopy) {
					t.Errorf("%s(%v) = %v; want %v", fn.name, tc.nums, got, wantCopy)
				}
			})
		}
	})
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
