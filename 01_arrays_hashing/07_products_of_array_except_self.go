package arrays_hashing_01

import (
	"encoding/json"
	"fmt"
)

/*
Products of Array Except Self

Medium

Given an integer array nums, return an array output where output[i] is
the product of all the elements of nums except nums[i].

Each product is guaranteed to fit in a 32-bit integer.

Follow-up: Could you solve it in O(n) time without using the division operation?

Example 1:

Input: nums = [1,2,4,6]

Output: [48,24,12,8]
Example 2:

Input: nums = [-1,0,1,2,3]

Output: [0,-6,0,0,0]
Constraints:

2 <= nums.length <= 1000
-20 <= nums[i] <= 20

Complexity Analysis
Time complexity: O(n)
Space complexity: O(n)

n is the number of elements in the input

*/

func productExceptSelf(nums []int) []int {
	right := make([]int, len(nums)-1)
	right[len(nums)-2] = nums[len(nums)-1]
	for i := len(nums) - 2; i > 0; i-- {
		right[i-1] = right[i] * nums[i]
	}

	left := make([]int, len(nums)-1)
	left[0] = nums[0]
	for i := 1; i < len(nums)-1; i++ {
		left[i] = left[i-1] * nums[i]
	}

	// nums    = [   2,   3,   4,   5,]
	// nums_i  = [   0,   1,   2,  ..,]
	// nums_i  = [ i-1,   i, i+2,   3,]
	//
	// right   = [  60,  20,   5,     ]
	// right_i = [   0,   1,   2,     ]
	// left    = [        2,   6,  24,]
	// left_i  = [        0,   1,   2,]
	res := make([]int, len(nums))
	res[0] = right[0]
	res[len(nums)-1] = left[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		res[i] = right[i] * left[i-1]
	}

	leftJson, _ := json.Marshal(left)
	fmt.Println(string(leftJson))
	return res
}

/*

Recommended Time & Space Complexity
You should aim for a solution as good or better than O(n) time and O(n) space,
where n is the size of the input array.


Hint 1
A brute-force solution would be to iterate through the array with index i and compute
the product of the array except for that index element. This would be an O(n^2) solution.
Can you think of a better way?


Hint 2
Is there a way to avoid the repeated work? Maybe we can store the results of the repeated
work in an array.


Hint 3
We can use the prefix and suffix technique. First, we iterate from left to right and store
the prefix products for each index in a prefix array, excluding the current index's number.
Then, we iterate from right to left and store the suffix products for each index in a suffix
array, also excluding the current index's number. Can you figure out the solution from here?


Hint 4
We can use the stored prefix and suffix products to compute the result array by iterating
through the array and simply multiplying the prefix and suffix products at each index.

*/
