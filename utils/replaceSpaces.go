package utils

import (
	"unicode"
)

func replaceSpaces(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		if unicode.IsSpace(rune(s[i])) {
			var j, spaceCount int
			for j = i; unicode.IsSpace(rune(s[j])); j++ {
				spaceCount++
			}
			for k := i + 1; j < len(s); k++ {
				s[k] = s[j]
				j++
			}
			s = s[:len(s) - spaceCount + 1]
			i -= spaceCount - 1
		}
	}
	return s
}
