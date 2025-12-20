package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Group Anagrams
Solved
Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.

An anagram is a string that contains the exact same characters as another string, but the order of the characters
can be different.

Example 1:

Input: strs = ["act","pots","tops","cat","stop","hat"]

Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
Example 2:

Input: strs = ["x"]

Output: [["x"]]
Example 3:

Input: strs = [""]

Output: [[""]]
Constraints:

1 <= strs.length <= 1000.
0 <= strs[i].length <= 100
strs[i] is made up of lowercase English letters.

*/

func main() {
	testCases := []struct {
		data []string
		res  [][]string
	}{
		{
			data: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			res:  [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
		{
			data: []string{"x"},
			res:  [][]string{{"x"}},
		},
		{
			data: []string{""},
			res:  [][]string{{""}},
		},
	}

	runTest := func(name string, f func([]string) [][]string) {
		for _, tc := range testCases {
			// Since the order of sublists and elements within sublists doesn't matter for correctness but matters
			// for equality check we just print it here as requested in the example.
			// For a more robust test we would need to sort the result and expected.
			fmt.Printf("%s. data: %v, expected: %v, got: %v\n", name, tc.data, tc.res, f(tc.data))
		}
	}

	runTest("Sorting", groupAnagramsBySorting)
	runTest("Alphabet", groupAnagramsByAlphabet)
}

func groupAnagramsByAlphabet(strs []string) [][]string {

	createFrequencyKey := func(s string) string {
		m := [26]int{}
		for _, r := range s {
			m[r-'a'] += 1
		}

		var res strings.Builder
		for i, count := range m {
			res.WriteByte(byte('a' + i))
			res.WriteString(strconv.Itoa(count))
		}
		return res.String()
	}

	return groupAnagrams(strs, createFrequencyKey)
}

func groupAnagramsBySorting(strs []string) [][]string {

	createFrequencyKey := func(s string) string {
		m := map[rune]int{}
		for _, r := range s {
			m[r] += 1
		}

		runes := make([]rune, 0, len(m))
		for r := range m {
			runes = append(runes, r)
		}

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		res := ""
		for _, r := range runes {
			res += string(r) + strconv.Itoa(m[r])
		}
		return res
	}

	return groupAnagrams(strs, createFrequencyKey)
}

func groupAnagrams(strs []string, createFrequencyKey func(s string) string) [][]string {
	m := map[string][]string{}
	for _, s := range strs {
		k := createFrequencyKey(s)
		m[k] = append(m[k], s)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

/*

Recommended Time & Space Complexity
You should aim for a solution with O(m * n) time and O(m) space, where m is the number of strings and n is the length
of the longest string.


Hint 1
A naive solution would be to sort each string and group them using a hash map. This would be an O(m * nlogn) solution.
Though this solution is acceptable, can you think of a better way without sorting the strings?


Hint 2
By the definition of an anagram, we only care about the frequency of each character in a string. How is this helpful
in solving the problem?


Hint 3
We can simply use an array of size O(26), since the character set is a through z (26 continuous characters), to count
the frequency of each character in a string. Then, we can use this array as the key in the hash map to group the strings.

https://neetcode.io/problems/anagram-groups/question?list=neetcode150

*/
