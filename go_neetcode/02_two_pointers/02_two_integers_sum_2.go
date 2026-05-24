package _2_two_pointers

/*
# Find Two Numbers That Sum to Target (Sorted Array)

**Difficulty:** Medium

## Problem Description

You are given a sorted array of integers `numbers` arranged in non-decreasing order. Find two distinct numbers
in this array that sum to a given `target` value.

Return the **1-indexed positions** `[index1, index2]` of these two numbers where `index1 < index2`.

**Key Requirements:**
- Each element can only be used once (index1 ≠ index2)
- Exactly one valid solution is guaranteed to exist
- Space complexity must be O(1) (constant extra space)

## Example

**Input:**
```java
numbers = [1,2,3,4]
target = 3
```

**Output:**
```java
[1,2]
```

**Explanation:**
The elements at positions 1 and 2 (values 1 and 2) sum to 3. Using 1-based indexing, we return `[1, 2]`.

## Constraints

- Array length: 2 ≤ numbers.length ≤ 1000
- Element values: -1000 ≤ numbers[i] ≤ 1000
- Target value: -1000 ≤ target ≤ 1000
*/

/*
Complexity Analysis
Time complexity: O(n)
Space complexity: O(1)

n is the number of elements in the numbers slice.
*/

func twoIntegersSum2(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j {
		s := numbers[i] + numbers[j]
		if s == target {
			break
		}
		if s < target {
			i++
			continue
		}
		if s > target {
			j--
			continue
		}
	}

	// array is 1-indexed
	i++
	j++

	return []int{i, j}
}
