// Package summultiples provides a way to calculate ths sum of all of the multiples of a given number, or slice of ints
package summultiples

import "sort"

// SumMultiples takes an upper limit, plus a list of numbers. It determines all of the multiples of that list of numbers, up to the limit.
// Then, all of those multiples are summed, and the sum of the multiples of the input numbers are returned.
func SumMultiples(limit int, divisors ...int) int {
	var multis []int
	for _, each := range divisors {
		multis = append(multis, multiples(each, limit)...)
	}
	multis = dedup(multis)
	return sum(multis)
}

// standard function to list out multiples of a given number
func multiples(n, limit int) (multiples []int) {
	for i := n; i < limit; i++ {
		if i%n == 0 {
			multiples = append(multiples, i)
		}
	}
	return
}

// deduplicate a slice of ints
func dedup(arr []int) []int {
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; {
		switch {
		case arr[i] == arr[i+1]:
			arr = append(arr[:i], arr[i+1:]...)
		default:
			i++
		}
	}
	return arr
}

// sum a slice of ints
func sum(arr []int) (result int) {
	for _, each := range arr {
		result += each
	}
	return
}