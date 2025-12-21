package arrays_hashing_01

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		res  []int
	}{
		{name: "Example 1", nums: []int{1, 2, 2, 3, 3, 3}, k: 2, res: []int{2, 3}},
		{name: "Example 2", nums: []int{7, 7}, k: 1, res: []int{7}},
	}

	runTest := func(t *testing.T, name string, f func([]int, int) []int) {
		for _, tc := range testCases {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				expected := make([]int, len(tc.res))
				copy(expected, tc.res)

				result := f(tc.nums, tc.k)

				sort.Ints(result)
				sort.Ints(expected)

				if !reflect.DeepEqual(result, expected) {
					t.Errorf("got %v; want %v", result, expected)
				}
			})
		}
	}

	t.Run("topKFrequentSorting", func(t *testing.T) { runTest(t, "topKFrequentSorting", topKFrequentSorting) })
	t.Run("topKFrequentMaxKey", func(t *testing.T) { runTest(t, "topKFrequentMaxKey", topKFrequentMaxKey) })
	t.Run("topKFrequentPriorityQueue", func(t *testing.T) { runTest(t, "topKFrequentPriorityQueue", topKFrequentPriorityQueue) })
	t.Run("topKFrequentBucketSort", func(t *testing.T) { runTest(t, "topKFrequentBucketSort", topKFrequentBucketSort) })
}
