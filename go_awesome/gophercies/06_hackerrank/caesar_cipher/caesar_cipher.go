package caesar_cipher

import "strings"

func CaesarCipher(s string, k int) string {
	result := strings.Builder{}

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c = decipher(c, 'a', k)
		} else if c >= 'A' && c <= 'Z' {
			c = decipher(c, 'A', k)
		}
		result.WriteRune(c)
	}

	return result.String()
}

func decipher(c rune, base rune, k int) rune {
	return rune((int(c-base)+k)%26) + base
}
