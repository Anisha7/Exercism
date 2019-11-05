package raindrops

import (
	"strconv"
)

// Convert converts a number to a string
func Convert(n int) string {
	result := ""
	if (n%3 != 0) && (n%5 != 0) && (n%7 != 0) {
		return strconv.Itoa(n)
	}

	if n%3 == 0 {
		result += "Pling"
	}
	if n%5 == 0 {
		result += "Plang"
	}
	if n%7 == 0 {
		result += "Plong"
	}
	return result
}
