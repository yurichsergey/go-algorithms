package main

/*
Two Sum
Solved
Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target
and i != j.

You may assume that every input has exactly one pair of indices i and j that satisfy the condition.

Return the answer with the smaller index first.

Example 1:

Input:
nums = [3,4,5,6], target = 7

Output: [0,1]
Explanation: nums[0] + nums[1] == 7, so we return [0, 1].

Example 2:

Input: nums = [4,5,6], target = 10

Output: [0,2]
Example 3:

Input: nums = [5,5], target = 10

Output: [0,1]
Constraints:

2 <= nums.length <= 1000
-10,000,000 <= nums[i] <= 10,000,000
-10,000,000 <= target <= 10,000,000
Only one valid answer exists.

Recommended Time & Space Complexity
You should aim for a solution with O(n) time and O(n) space, where n is the size of the input array.

https://neetcode.io/problems/two-integer-sum/question?list=neetcode150

*/

import (
	"fmt"
	"reflect"
)

func main() {
	testCases := []struct {
		nums   []int
		target int
		res    []int
	}{
		{nums: []int{3, 4, 5, 6}, target: 7, res: []int{0, 1}},
		{nums: []int{4, 5, 6}, target: 10, res: []int{0, 2}},
		{nums: []int{5, 5}, target: 10, res: []int{0, 1}},
	}

	for _, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		fmt.Printf("nums: %v, target: %d, expected: %v, got: %v, match: %v\n",
			tc.nums, tc.target, tc.res, result, reflect.DeepEqual(tc.res, result))
	}
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for ki, i := range nums {
		diff := target - i
		kj, ok := m[diff]
		if ok {
			return []int{kj, ki}
		}
		m[i] = ki
	}
	return []int{-1, -1}
}
