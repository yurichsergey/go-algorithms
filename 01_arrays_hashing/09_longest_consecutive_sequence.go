package arrays_hashing_01

import (
	"slices"
)

/**
Longest Consecutive Sequence
Given an array of integers nums, return the length of the longest consecutive sequence of elements that can be formed.

A consecutive sequence is a sequence of elements in which each element is exactly 1 greater than the previous element.
The elements do not have to be consecutive in the original array.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [2,20,4,10,3,4,5]
Output: 4
Explanation: The longest consecutive sequence is [2, 3, 4, 5].

Example 2:

Input: nums = [0,3,2,5,4,6,1,1]
Output: 7

Example 3:

Input: nums = [9,1,4,7,3,-1,0,5,8,-1,6]
Output: 7

Constraints:

0 <= nums.length <= 1000
-10^9 <= nums[i] <= 10^9
*/

/*
*
Complexity Analysis
Time complexity: O(n2)
Space complexity: O(n)

n is the number of elements in the input
*/
func longestConsecutiveBruteForce(nums []int) int {
	uniq := map[int]struct{}{}
	for _, i := range nums {
		uniq[i] = struct{}{}
	}

	maxLen := 0
	for _, num := range nums {
		curLen := 1

		for i := 1; true; i++ {
			_, ok := uniq[num+i]
			if !ok {
				break
			}
			curLen++
		}

		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}

/*
*
Complexity Analysis
Time complexity: O(nlog(n))
Space complexity: O(n)
*/
func longestConsecutiveSort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uniq := map[int]struct{}{}
	for _, i := range nums {
		uniq[i] = struct{}{}
	}

	uniqNums := make([]int, 0, len(uniq))
	for num := range uniq {
		uniqNums = append(uniqNums, num)
	}
	slices.Sort(uniqNums)

	maxLen := 1
	curLen := 1

	currInd := 0
	for i := 1; i < len(uniqNums); i++ {
		if uniqNums[currInd]+(i-currInd) != uniqNums[i] {
			currInd = i
			curLen = 1
			continue
		}

		curLen++

		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}

/*
*
Complexity Analysis
Time complexity: O(n)
Space complexity: O(n)
*/
func longestConsecutiveHashSet(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}

	return longest
}

/*
*
Complexity Analysis
Time complexity: O(n)
Space complexity: O(n)
*/
func longestConsecutiveHashMap(nums []int) int {
	mp := make(map[int]int)
	res := 0
	for _, n := range nums {
		_, ok := mp[n]
		if ok {
			continue
		}
		length := mp[n-1] + mp[n+1] + 1
		mp[n] = length
		// update left and right boundaries
		mp[n-mp[n-1]] = length
		mp[n+mp[n+1]] = length
		if length > res {
			res = length
		}
	}
	return res
}
