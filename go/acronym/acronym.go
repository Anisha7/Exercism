// Package acronym abbreviates a string
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// existsInList checks if a letter exists in the given array
func existsInList(letter string, arr [3]string) bool {
	for _, l := range arr {
		if letter == l {
			return true
		}
	}
	return false
}

// Abbreviate gives the abbreviation of a word
func Abbreviate(s string) string {
	splitter := [3]string{" ", "-", "_"}
	var result string
	prev := " "
	for _, char := range s {
		c := string(char)
		if existsInList(c, splitter) {
			prev = c
			continue
		} else if existsInList(prev, splitter) {
			result += c
		}
		prev = c
	}
	return strings.ToUpper(result)
}
