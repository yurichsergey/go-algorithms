package _01_arrays_hashing

import (
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
