package _2_two_pointers

import (
	"regexp"
	"strings"
)

/**
# Palindrome Verification
**Difficulty: Easy**

Given a string `s`, determine whether it forms a palindrome and return `true` if it does, or `false` otherwise.

A palindrome is defined as a string that remains identical when read in both directions
(left-to-right and right-to-left). The comparison should be case-insensitive and should
exclude all non-alphanumeric characters.

Note: Alphanumeric characters include both letters `(A-Z, a-z)` and digits `(0-9)`.

**Example 1:**

```java
Input: s = "Was it a car or a cat I saw?"

Output: true
```

Explanation: When we filter for only alphanumeric characters, we get "wasitacaroracatisaw",
which reads the same forwards and backwards.

**Example 2:**

```java
Input: s = "tab a cat"

Output: false
```

Explanation: After filtering, we get "tabacat", which is not the same when reversed.

**Constraints:**
* `1 <= s.length <= 1000`
* `s` consists of printable ASCII characters only.

*/

/*
Complexity Analysis
Time complexity: O(n)
Space complexity: O(n)

n is the length of the input string s
*/

func isPalindrome(s string) bool {
	clearedString := regexp.MustCompile("[^a-zA-Z0-9]").ReplaceAllString(s, "")
	clearedString = strings.ToLower(clearedString)
	runes := []rune(clearedString)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}

/*
*

Complexity Analysis
Time complexity: O(n)
Space complexity: O(1)

n is the number of characters in the input string s
*/
func isPalindromeEfficient(s string) bool {
	n := len(s)
	i := 0
	j := n - 1
	re := regexp.MustCompile("[a-zA-Z0-9]")
	for i <= j {
		if !re.MatchString(string(s[i])) {
			i++
			continue
		}
		if !re.MatchString(string(s[j])) {
			j--
			continue
		}
		//print(string(s[i]) + " -- " + string(s[j]) + "\n")
		if strings.ToLower(string(s[i])) == strings.ToLower(string(s[j])) {
			i++
			j--
			continue
		}
		return false
	}
	return true
}

/**
**Recommended Time & Space Complexity**

Your solution should target `O(n)` time complexity and `O(1)` space complexity, where `n` represents the length of
the string.

**Hint 1**
The straightforward approach involves creating a reversed copy of the string and comparing it with the original.
While this achieves `O(n)` time, it requires `O(n)` additional space. Can you solve this without using extra space
proportional to the input size?

**Hint 2**
What pattern can you identify from the palindrome definition or from analyzing the brute force approach?

**Hint 3**
Since a palindrome reads identically from both ends, characters at corresponding positions from the start and end
should match. The two-pointer technique can efficiently verify this property.
*/
