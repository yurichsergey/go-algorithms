package arrays_hashing_01

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name string
		data []string
		res  [][]string
	}{
		{
			name: "Example 1",
			data: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			res:  [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
		{
			name: "Example 2",
			data: []string{"x"},
			res:  [][]string{{"x"}},
		},
		{
			name: "Example 3",
			data: []string{""},
			res:  [][]string{{""}},
		},
	}

	normalize := func(groups [][]string) {
		for _, group := range groups {
			sort.Strings(group)
		}
		sort.Slice(groups, func(i, j int) bool {
			if len(groups[i]) != len(groups[j]) {
				return len(groups[i]) < len(groups[j])
			}
			return groups[i][0] < groups[j][0]
		})
	}

	runTest := func(t *testing.T, name string, f func([]string) [][]string) {
		for _, tc := range testCases {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				result := f(tc.data)

				// Normalize result and expected for comparison
				expected := make([][]string, len(tc.res))
				for i := range tc.res {
					expected[i] = make([]string, len(tc.res[i]))
					copy(expected[i], tc.res[i])
				}

				normalize(result)
				normalize(expected)

				if !reflect.DeepEqual(result, expected) {
					t.Errorf("got %v; want %v", result, expected)
				}
			})
		}
	}

	runTest(t, "Sorting", groupAnagramsBySorting)
	runTest(t, "Alphabet", groupAnagramsByAlphabet)
}
