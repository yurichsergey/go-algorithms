package _2_two_pointers

import "sort"

/*

# Three Number Sum

**Difficulty:** Medium

## Problem Description

Given an array of integers `nums`, find all unique combinations of three numbers `[nums[i], nums[j], nums[k]]` that sum to zero, where `i`, `j`, and `k` are three different positions in the array.

Return all such triplets without duplicates. The order of triplets in the result and the order of elements within each triplet can be arbitrary.

## Examples

**Example 1:**

```java
Input: nums = [-3,1,2,-3,4,0]

Output: [[-3,-3,6] is not valid, but [-3,1,2],[-3,0,3] would be if 3 existed]
Actually: [[−3,−3,6]] - wait, let me recalculate
Output: [[-3,1,2]]
```

**Explanation:** The array contains the triplet where `nums[0] + nums[1] + nums[2] = (-3) + 1 + 2 = 0`. This is the only valid combination that sums to zero.

**Example 2:**

```java
Input: nums = [5,2,-1]

Output: []
```

**Explanation:** No combination of three numbers from this array can sum to zero.

**Example 3:**

```java
Input: nums = [1,1,-2]

Output: [[1,1,-2]]
```

**Explanation:** The triplet `[1,1,-2]` sums to `1 + 1 + (-2) = 0`, which is the only valid combination.

## Constraints

* The array contains at least 3 elements and at most 1000 elements: `3 <= nums.length <= 1000`
* Each element is within the range: `-10^5 <= nums[i] <= 10^5`

*/

func threeSumBruteForce(nums []int) [][]int {
	sort.Ints(nums)
	// freq := make(map[int]int)
	// for _, num := range(nums) {
	// 	freq[num]++
	// }

	res := map[[3]int]struct{}{}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				}
			}
		}
	}

	resTriplet := [][]int{}
	for triplet := range res {
		resTriplet = append(resTriplet, []int{triplet[0], triplet[1], triplet[2]})
	}
	return resTriplet
}

/*
# Recommended Complexity Target

Aim for a solution with **O(n²) time complexity** and **O(1) space complexity**, where `n` represents the length of the
input array.

## Hint 1
The naive approach of examining every possible combination of three elements would result in O(n³) complexity.
Is there a more efficient method?

## Hint 2
Consider what advantages sorting the array might provide. How could you transform the problem equation to make it
easier to solve?

## Hint 3
After sorting, iterate through the array using index `i`. Rearranging the sum equation `nums[i] + nums[j] + nums[k] = 0`
gives us `nums[j] + nums[k] = -nums[i]`. For each position `i`, we need an efficient method to find valid `j` and `k`
pairs while avoiding duplicates. What technique works well for finding two numbers that sum to a target in a sorted
array?

## Hint 4
Use the two-pointer technique on elements after index `i` in the sorted array. Set `j` at the start and `k` at the end
of this subarray, with `target = -nums[i]`. If `nums[j] + nums[k] < target`, move `j` forward to increase the sum.
If `nums[j] + nums[k] > target`, move `k` backward to decrease the sum. How can you handle duplicate triplets?

## Hint 5
When you find `nums[j] + nums[k] == target`, add the triplet to your results. Then, advance `j` or retreat `k` while
`j < k` and the values remain the same. This prevents adding duplicate triplets to your answer.
*/
