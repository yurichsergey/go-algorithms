package arrays_hashing_01

/**
Longest Consecutive Sequence
Given an array of integers nums, return the length of the longest consecutive sequence of elements that can be formed.

A consecutive sequence is a sequence of elements in which each element is exactly 1 greater than the previous element. The elements do not have to be consecutive in the original array.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [2,20,4,10,3,4,5]

Output: 4
Explanation: The longest consecutive sequence is [2, 3, 4, 5].

Example 2:

Input: nums = [0,3,2,5,4,6,1,1]

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

/**
Recommended Time & Space Complexity
You should aim for a solution as good or better than O(n) time and O(n) space, where n is the size of the input array.


Hint 1
A brute force solution would be to consider every element from the array as the start of the sequence and
count the length of the sequence formed with that starting element. This would be an O(n^2) solution.
Can you think of a better way?


Hint 2
Is there any way to identify the start of a sequence? For example, in [1, 2, 3, 10, 11, 12],
only 1 and 10 are the beginning of a sequence. Instead of trying to form a sequence for every number,
we should only consider numbers like 1 and 10.


Hint 3
We can consider a number num as the start of a sequence if and only if num - 1 does not exist in the given array.
We iterate through the array and only start building the sequence if it is the start of a sequence.
This avoids repeated work. We can use a hash set for O(1) lookups by converting the array to a hash set.
*/
