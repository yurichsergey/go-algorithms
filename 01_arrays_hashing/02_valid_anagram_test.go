package arrays_hashing_01

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		tt   string
		res  bool
	}{
		{name: "Example 1", s: "racecar", tt: "carrace", res: true},
		{name: "Example 2", s: "jar", tt: "jam", res: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isAnagram(tc.s, tc.tt)
			if result != tc.res {
				t.Errorf("isAnagram(%q, %q) = %v; want %v", tc.s, tc.tt, result, tc.res)
			}
		})
	}
}
