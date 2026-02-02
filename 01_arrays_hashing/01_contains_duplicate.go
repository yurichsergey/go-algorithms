package arrays_hashing_01

/*
Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

Example 1:

Input: nums = [1, 2, 3, 3]

Output: true

Example 2:

Input: nums = [1, 2, 3, 4]

Output: false

*/

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
