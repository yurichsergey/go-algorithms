package arrays_hashing_01

/**
Valid Anagram
Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can
be different.

Example 1:

Input: s = "racecar", t = "carrace"

Output: true
Example 2:

Input: s = "jar", t = "jam"

Output: false
Constraints:

s and t consist of lowercase English letters.

*/

func isAnagram(s string, t string) bool {

	buildStat := func(s string) map[rune]int {
		m := map[rune]int{}
		for _, r := range s {
			m[r] += 1
		}
		return m
	}

	ms := buildStat(s)
	mt := buildStat(t)

	for r, i := range ms {
		mt[r] -= i
	}

	res := true
	for _, j := range mt {
		if j != 0 {
			res = false
			break
		}
	}
	return res
}

/*
Recommended Time & Space Complexity
You should aim for a solution with O(n + m) time and O(1) space, where n is the length of the string s and m is the
length of the string t.

*/
