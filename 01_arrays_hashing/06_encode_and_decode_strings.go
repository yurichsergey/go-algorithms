package _01_arrays_hashing

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*

Encode and Decode Strings
Medium

Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the
original list of strings.

Please implement encode and decode

Example 1:

Input: ["neet","code","love","you"]

Output:["neet","code","love","you"]
Example 2:

Input: ["we","say",":","yes"]

Output: ["we","say",":","yes"]
Constraints:

0 <= strs.length < 100
0 <= strs[i].length < 200
strs[i] contains only UTF-8 characters.

Complexity Analysis
Time complexity: O(N)
Space complexity:O(N)

N is the total number of characters across all strings in the input for Encode, and the total number of characters
in the encoded string for Decode.
*/

type SolutionEncodedDecodedString struct{}

func (s *SolutionEncodedDecodedString) Encode(strs []string) string {
	var res strings.Builder
	for _, s := range strs {
		l := utf8.RuneCountInString(s)
		res.WriteString(strconv.Itoa(l) + "#" + s)
	}
	fmt.Println(res.String())
	return res.String()
}

func (s *SolutionEncodedDecodedString) Decode(encoded string) []string {
	readSize := true
	var sizeStr strings.Builder
	size := 0
	count := 0

	var word strings.Builder

	res := []string{}
	for _, r := range encoded {
		if readSize && r != '#' {
			sizeStr.WriteRune(r)
			continue
		}

		if readSize {
			num, err := strconv.Atoi(sizeStr.String())
			if err != nil {
				fmt.Printf("Error strconv.Atoi(sizeStr)")
				return []string{}
			}
			size = num
			readSize = false
			sizeStr.Reset()

			if size == 0 {
				res = append(res, "")
				readSize = true
			}

			continue
		}

		count++
		word.WriteRune(r)
		if count >= size {
			readSize = true
			res = append(res, word.String())
			word.Reset()
			count = 0
		}
	}

	return res
}
