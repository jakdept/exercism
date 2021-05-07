// Package lsproduct finds the sequence of consecutive digits in a string with the highest product
package lsproduct

import (
	"fmt"
	"strconv"
)

const testVersion = 4

// LargestSeriesProduct finds the looks for the largest product in a string
func LargestSeriesProduct(s string, length int) (int, error) {
	if length > len(s) {
		return 0, fmt.Errorf("you cannot ask for more digits than you give")
	}
	if length < 0 {
		return 0, fmt.Errorf("you cannot ask for a negative size")
	}
	var largestProduct int
	for i := 0; i <= len(s)-length; i++ {
		nums, err := stringToInts(s[i : i+length])
		if err != nil {
			return 0, err
		}
		product := productOfInts(nums)
		if product > largestProduct {
			largestProduct = product
		}
	}

	return largestProduct, nil
}

// converts a string to a slice of digits, or errors if anything other than a digit is found
func stringToInts(s string) ([]int, error) {
	var digits []int
	for _, char := range s {
		i, err := strconv.Atoi(string(char))
		if err != nil {
			return []int{}, fmt.Errorf("[%q] - %v", s, err)
		}
		digits = append(digits, i)
	}
	return digits, nil
}

// finds the product of a slice of ints
func productOfInts(numbers []int) int {
	result := 1
	for _, n := range numbers {
		result *= n
	}
	return result
}