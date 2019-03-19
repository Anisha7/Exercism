package hamming

import "errors"

// Distance calculates the hamming distance between strands a and b
func Distance(a, b string) (int, error) {
	var dist int
	var i int
	var err error
	// check if valid strands
	if len(a) != len(b) {
		return 0, errors.New("can't compare strands of unequal lengths")
	}
	// find unequal DNA It is found by comparing two DNA strands and counting how many of the
	for i < len(a) {
		if a[i] != b[i] {
			dist++
		}
		i++
	}

	return dist, err
}
