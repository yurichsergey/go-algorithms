package _01_arrays_hashing

import (
	"testing"
)

func TestHasDuplicate(t *testing.T) {
	testCases := []struct {
		name string
		data []int
		res  bool
	}{
		{name: "Example 1", data: []int{1, 2, 3, 3}, res: true},
		{name: "Example 2", data: []int{1, 2, 3, 4}, res: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := hasDuplicate(tc.data)
			if result != tc.res {
				t.Errorf("hasDuplicate(%v) = %v; want %v", tc.data, result, tc.res)
			}
		})
	}
}
