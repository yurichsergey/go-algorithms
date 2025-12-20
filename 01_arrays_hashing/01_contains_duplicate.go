// package _1_arrays_hashing
package main

import "fmt"

/*
Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

Example 1:

Input: nums = [1, 2, 3, 3]

Output: true

Example 2:

Input: nums = [1, 2, 3, 4]

Output: false

*/

func main() {
	testCases := []struct {
		data []int
		res  bool
	}{
		{data: []int{1, 2, 3, 3}, res: true},
		{data: []int{1, 2, 3, 4}, res: false},
	}

	runTest := func(name string, f func([]int) bool) {
		for _, tc := range testCases {
			result := f(tc.data)
			fmt.Printf("%s. data: %v, expected: %v, got: %v\n", name, tc.data, tc.res, result)
		}
	}

	runTest("hasDuplicate", hasDuplicate)
}

func hasDuplicate(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		_, ok := set[v]
		if ok {
			return true
		} else {
			set[v] = struct{}{}
		}
	}
	return false
}

/*
https://neetcode.io/problems/duplicate-integer/question

Time complexity: O(n)
Space complexity: O(n)
n is the number of elements in the nums slice.
*/
