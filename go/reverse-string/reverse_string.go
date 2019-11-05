package reverse

import (
	"bytes"
)

// Version 1: does not work on multi-byte test case
// Reverse reverses a string
// func ReverseV1(s string) string {
// 	// r := []rune(s)
// 	var result []string

// 	for i := range s {
// 		result = append(result, string(s[len(s)-i-1]))
// 	}

// 	return strings.Join(result, "")
// }

// Reverse reverses a string
func Reverse(s string) string {
	var result [][]byte
	var bytesArr [][]byte

	// create a byte array
	for _, c := range s {
		b := []byte(string(c))
		bytesArr = append(bytesArr, b)
	}

	// reverse the byte array
	for i := range bytesArr {
		result = append(result, bytesArr[len(bytesArr)-i-1])
	}

	// convert the byte array to string
	return string(bytes.Join(result, []byte("")))
}
