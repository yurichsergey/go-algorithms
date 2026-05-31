package camelcase

import "unicode"

func CountWords(s string) int {
	count := 0
	if len(s) == 0 {
		return count
	}

	for _, l := range s {
		//if l >= 'A' && l <= 'Z' {
		if unicode.IsUpper(l) {
			count++
		}
	}
	return count + 1
}
