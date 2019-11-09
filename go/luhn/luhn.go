package luhn

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// reverse returns its argument string reversed rune-wise left to right.
// cribbed from:
// https://github.com/golang/example/blob/master/stringutil/reverse.go
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// converts string of numbers to a list of numbers
func convertNumbers(numstr string) []int {
	var result []int
	for i, c := range reverse(numstr) {
		// fail test
		if !unicode.IsDigit(c) {
			return []int{2}
		}
		num, _ := strconv.Atoi(string(c))
		println(i % 2)
		if i%2 == 1 {
			num = num * 2
		}
		if num > 9 {
			num -= 9
		}
		result = append(result, num)
	}
	println("Converted:")
	fmt.Printf("%v\n", result)
	return result
}

// adds all numbers in a list
func sumList(nums []int) int {
	var result int
	for _, n := range nums {
		result += n
	}
	return result
}

// Valid uses Luhn algorithm to validate identification numbers
func Valid(numstr string) bool {
	numstr = strings.ReplaceAll(numstr, " ", "")
	if len(numstr) < 2 {
		return false
	}
	convertedNums := convertNumbers(numstr)
	summedNums := sumList(convertedNums)
	if summedNums%10 == 0 {
		return true
	}
	return false
}
