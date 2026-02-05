package _2_two_pointers

/*
# Container With Most Water

## Problem Statement

Given an array of integers representing the heights of vertical bars, find two bars that can form a container holding the maximum amount of water.

## Key Details

- **Water capacity** is determined by: `width × min(height[i], height[j])` where `i` and `j` are two bar indices
  - Width = distance between the two bars
  - Height = the shorter of the two bars (water would overflow at the shorter one)

- **Goal**: Find the pair of bars that maximizes this product

## Examples

**Example 1:**

.10,............................................
..9,............................................
..8,.........*@..................S@.............
..7,.........*@::::::::::::::::::S@::::::::+@...
..6,.........*@::::::::::::::::::S@::::::::+@...
..5,.........*@::::::::S@::::::::S@::::::::+@...
..4,.........*@::::::::S@:::S@:::S@::::::::+@...
..3,.........*@::::::::S@:::S@:::S@:::S@:::+@...
..2,.........*@:::S@:::S@:::S@:::S@:::S@:::+@...
..1,....S@...*@:::S@:::S@:::S@:::S@:::S@:::+@...
..0,....--...--...--...--...--...--...--...--...
:::::::::0::::1::::2::::3::::4::::5::::6::::7:::

```
Input: [1,7,2,5,4,7,3,6]
Output: 36
```
The bars at indices 1 (height=7) and 5 (height=7) create a container with width=4 and height=7, holding 4×7=28 units. Actually, indices 1 (height=7) and 6 (height=3) give width=5 and height=3, or better: indices 5 and 7 (heights 7 and 6) with width=2 and height=6 gives 12. The maximum is actually formed by bars creating an area of 36.

**Example 2:**
```
Input: [2,2,2]
Output: 4
```
Any two bars have width=1 or 2. With bars of height 2, the maximum is 2×2=4 (using bars at distance 2).

## Constraints

- Array length: 2 to 1,000
- Bar heights: 0 to 1,000

## Optimal Approach

Use the **two-pointer technique**:
1. Start with pointers at both ends (maximum width)
2. Calculate water capacity at current positions
3. Move the pointer pointing to the shorter bar inward (moving the taller bar won't improve capacity)
4. Track the maximum capacity found

**Time Complexity**: O(n)
**Space Complexity**: O(1)
*/

/*
*
Time complexity: O(n2)
Space complexity: O(1)
n is the number of elements in the heights slice.
*/
func maxArea(heights []int) int {
	maxArea := 0
	for i := range heights {
		for j := i + 1; j < len(heights); j++ {
			currentArea := min(heights[i], heights[j]) * (j - i)
			if currentArea > maxArea {
				maxArea = currentArea
			}
		}
	}
	return maxArea
}
