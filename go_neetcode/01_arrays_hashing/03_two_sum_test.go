package _01_arrays_hashing

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		res    []int
	}{
		{name: "Example 1", nums: []int{3, 4, 5, 6}, target: 7, res: []int{0, 1}},
		{name: "Example 2", nums: []int{4, 5, 6}, target: 10, res: []int{0, 2}},
		{name: "Example 3", nums: []int{5, 5}, target: 10, res: []int{0, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := twoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(result, tc.res) {
				t.Errorf("twoSum(%v, %d) = %v; want %v", tc.nums, tc.target, result, tc.res)
			}
		})
	}
}
