// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func hasNoLetters(remark string) bool {
	for i := 0; i < len(remark); i++ {
		if unicode.IsLetter(rune(remark[i])) {
			return false
		}
	}
	return true
}

// Hey function
func Hey(remark string) string {
	remark = strings.TrimFunc(strings.TrimSpace(remark), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && string(r) != "?"
	})
	noLetters := hasNoLetters(string(remark))
	if strings.ToUpper(remark) == remark && strings.Contains(remark, "?") && !noLetters {
		return "Calm down, I know what I'm doing!"
	} else if strings.HasSuffix(remark, "?") {
		return "Sure."
	} else if strings.ToUpper(remark) == remark && !noLetters {
		return "Whoa, chill out!"
	} else if strings.EqualFold(remark, "bob") || strings.Trim(remark, " ") == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
