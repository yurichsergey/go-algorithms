package arrays_hashing_01

import (
	"container/heap"
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

/*
Complexity Analysis
Time complexity: O(n)
Space complexity: O(n)

n is the number of elements in the input
*/
func topKFrequentBucketSort(nums []int, k int) []int {
	freq := map[int]int{}
	for _, i := range nums {
		freq[i] += 1
	}

	revert := make([][]int, len(nums))

	for num, freq := range freq {
		revert[freq-1] = append(revert[freq-1], num)
	}

	res := make([]int, 0, k)
	for i := len(nums); i >= 0; i-- {
		if len(revert[i-1]) == 0 {
			continue
		}
		for _, num := range revert[i-1] {
			res = append(res, num)
			if len(res) >= k {
				return res
			}
		}
	}

	return res
}

type FreqHeap [][2]int // [number, frequency]

func (h FreqHeap) Len() int           { return len(h) }
func (h FreqHeap) Less(i, j int) bool { return h[i][1] < h[j][1] } // min heap
func (h FreqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *FreqHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *FreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
Complexity Analysis
Time complexity: O(n*log(k))
Space complexity: O(n)

n is the number of elements in the input nums array.
k is the number of frequent elements to return.
*/
func topKFrequentPriorityQueue(nums []int, k int) []int {
	freq := map[int]int{}
	for _, i := range nums {
		freq[i] += 1
	}

	h := &FreqHeap{}
	heap.Init(h)

	for num, freq := range freq {
		heap.Push(h, [2]int{num, freq})
		if len(*h) > k {
			heap.Pop(h) // Remove the smallest frequency
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).([2]int)[0]
	}
	return res
}

/*
Time complexity: O(n⋅k)
Space complexity: O(n)

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
Time complexity: O(n*log(m))
Space complexity: O(n)

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
