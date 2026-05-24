package _01_arrays_hashing

import (
	"reflect"
	"testing"
)

func TestProductsOfArrayExceptSelf(t *testing.T) {
	testCases := []struct {
		name     string
		data     []int
		expected []int
	}{
		{
			name:     "Example 1",
			data:     []int{1, 2, 4, 6},
			expected: []int{48, 24, 12, 8},
		},
		{
			name:     "Example 2",
			data:     []int{-1, 0, 1, 2, 3},
			expected: []int{0, -6, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := productExceptSelf(tc.data)

			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf(
					"Test %s failed: \nInput: %v\nActual: %v\nExpected: %v",
					tc.name,
					tc.data,
					actual,
					tc.expected,
				)
			}
		})
	}
}
