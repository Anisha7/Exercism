package isogram

import (
	"strings"
)

// IsIsogram tells us if a word is an isogram ("a nonpattern word")
func IsIsogram(word string) bool {
	seen := make(map[rune]string)
	word = strings.ToLower(word)
	for _, c := range word {
		if _, found := seen[c]; found {
			if string(c) != " " && string(c) != "-" {
				return false
			}
		}
		seen[c] = ""
	}
	// check for repeating letters except spaces and hyphens
	return true
}
