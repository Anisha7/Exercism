package diffsquares

// SquareOfSum calculates the square of the sum
func SquareOfSum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result * result
}

// SumOfSquares calculates the sum of the squares
func SumOfSquares(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

// Difference calculates the square of the sum and the sum of the squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
