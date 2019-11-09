package scrabble

import "strings"

// existsInList checks if a letter exists in the given array
func existsInList(letter string, arr [10]string) bool {
	for _, l := range arr {
		if letter == l {
			return true
		}
	}
	return false
}

// getValue gets the scrabble value of a letter
func getValue(letter string) int {
	wordValues := map[[10]string]int{
		[10]string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}: 1,
		[10]string{"D", "G"}:                2,
		[10]string{"B", "C", "M", "P"}:      3,
		[10]string{"F", "H", "V", "W", "Y"}: 4,
		[10]string{"K"}:                     5,
		[10]string{"J", "X"}:                8,
		[10]string{"Q", "Z"}:                10,
	}

	for k := range wordValues {
		if existsInList(letter, k) {
			return wordValues[k]
		}
	}

	return 0
}

// Score calculates the scrabble score for a word
func Score(word string) int {
	value := 0
	for _, char := range word {
		value += getValue(strings.ToUpper(string(char)))
	}
	return value
}
