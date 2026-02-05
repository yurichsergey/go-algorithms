package _2_two_pointers

import (
	"sort"
)

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

/*
Time & Space Complexity
Time complexity: O(n3)
Space complexity: O(m)
Where m is the number of triplets and n is the length of the given array.
*/
func threeSumBruteForce(nums []int) [][]int {
	sort.Ints(nums)

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

	resList := [][]int{}
	for triplet := range res {
		resList = append(resList, triplet[:])
	}
	return resList
}

/*
*
Complexity Analysis
Time complexity: O(n2)
Space complexity: O(n)
n is the number of elements in the input array nums.
*/
func threeSumHashMap(nums []int) [][]int {
	sort.Ints(nums)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	res := map[[3]int]struct{}{}
	//fmt.Printf("freq = %v\n", freq)
	//fmt.Printf("nums = %v\n", nums)
	for i := range nums {
		freq[nums[i]]--
		//fmt.Printf("\niterate by i = %v, nums[i] = %v, freq = %v\n", i, nums[i], freq)
		for j := i + 1; j < len(nums); j++ {
			freq[nums[j]]--
			target := -(nums[i] + nums[j])

			//fmt.Printf(
			//	"nums[i] = %v;  nums[j] = %v; target = %v; freq = %v\n",
			//	nums[i],
			//	nums[j],
			//	target,
			//	freq,
			//)
			if freqVal, ok := freq[target]; ok && freqVal > 0 && target >= nums[j] {
				//fmt.Printf(
				//	"FOUND = %v, freqVal = %v, ok = %v\n",
				//	[3]int{nums[i], nums[j], target},
				//	freqVal,
				//	ok,
				//)
				res[[3]int{nums[i], nums[j], target}] = struct{}{}
			}
			freq[nums[j]]++
		}
	}
	//fmt.Printf("\nres = %v\n", res)

	resList := [][]int{}
	for triplet := range res {
		resList = append(resList, triplet[:])
	}
	return resList
}
