// Package diffsquares calculates the difference between the square of sums and sum of squares
package diffsquares

// SquareOfSums totals all numbers from 1 to the given number (inclusive), squares the sum, then returns the result
func SquareOfSum(in int) int {
	var result int
	for i := 1; i <= in; i++ {
		result += i
	}
	return result * result
}

// SumOfSquares multiplies every number from 1 to the given number (inclusive), sums those results, and returns the sum
func SumOfSquares(in int) int {
	var result int
	for i := 1; i <= in; i++ {
		result += i * i
	}
	return result
}

// Difference calculates the SquareOfSums and the SumOfSquares and determines the difference between the two
func Difference(in int) int {
	return SquareOfSum(in) - SumOfSquares(in)
}
