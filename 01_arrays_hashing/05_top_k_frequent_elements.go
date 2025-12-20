package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
Top K Frequent Elements

Medium

Given an integer array nums and an integer k, return the k most frequent elements within the array.

The test cases are generated such that the answer is always unique.

You may return the output in any order.

Example 1:

Input: nums = [1,2,2,3,3,3], k = 2

Output: [2,3]
Example 2:

Input: nums = [7,7], k = 1

Output: [7]
Constraints:

1 <= nums.length <= 10^4.
-1000 <= nums[i] <= 1000
1 <= k <= number of distinct elements in nums.

*/

func main() {
	testCases := []struct {
		nums []int
		k    int
		res  []int
	}{
		{nums: []int{1, 2, 2, 3, 3, 3}, k: 2, res: []int{2, 3}},
		{nums: []int{7, 7}, k: 1, res: []int{7}},
	}

	runTest := func(name string, f func([]int, int) []int) {
		for _, tc := range testCases {
			// Copy res to avoid modifying original test case during sort.Ints
			expected := make([]int, len(tc.res))
			copy(expected, tc.res)

			result := f(tc.nums, tc.k)

			sort.Ints(result)
			sort.Ints(expected)

			fmt.Printf(
				"%s. nums: %v, k: %d, expected: %v, got: %v, match: %v\n",
				name,
				tc.nums,
				tc.k,
				expected,
				result,
				reflect.DeepEqual(result, expected),
			)
		}
	}

	runTest("topKFrequentSorting", topKFrequentSorting)
	runTest("topKFrequentMaxKey", topKFrequentMaxKey)
}

/*
Time complexity:
O(n⋅k)
Space complexity:
O(n)

n is the number of elements in the nums array.
k is the number of frequent elements to find.
*/
func topKFrequentMaxKey(nums []int, k int) []int {
	freq := map[int]int{}
	for _, i := range nums {
		freq[i] += 1
	}

	maxKeyByValue := func(m map[int]int) int {
		r, rv := 0, 0
		for k, v := range m {
			if v > rv {
				r = k
				rv = v
			}
		}
		return r
	}

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		maxKey := maxKeyByValue(freq)
		res = append(res, maxKey)
		delete(freq, maxKey)
	}
	return res
}

/*
Time complexity:
O(n*log(m))
Space complexity:
O(n)

n is the number of elements in the input
m is the number of unique elements in the input
*/
func topKFrequentSorting(nums []int, k int) []int {

	freq := map[int]int{}
	for _, i := range nums {
		freq[i] += 1
	}

	keys := make([]int, 0, len(freq))
	for i := range freq {
		keys = append(keys, i)
	}

	sort.Slice(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})

	//fmt.Println("print freq")
	//for i := range(freq) {
	//	fmt.Println("%i - %i", i, freq[i])
	//}
	//fmt.Println("print keys")
	//for i := range(keys) {
	//	fmt.Println("%i", keys[i])
	//}

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, keys[i])
	}
	return res
}

/*
Recommended Time & Space Complexity
You should aim for a solution with O(n) time and O(n) space, where n is the size of the input array.


Hint 1
A naive solution would be to count the frequency of each number and then sort the array based on each element’s
frequency. After that, we would select the top k frequent elements. This would be an O(nlogn) solution.
Though this solution is acceptable, can you think of a better way?


Hint 2
Can you think of an algorithm which involves grouping numbers based on their frequency?


Hint 3
Use the bucket sort algorithm to create n buckets, grouping numbers based on their frequencies from 1 to n.
Then, pick the top k numbers from the buckets, starting from n down to 1.

https://neetcode.io/problems/top-k-elements-in-list/question?list=neetcode150
*/
