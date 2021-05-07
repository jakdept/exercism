// Package grains runs through the classic "graints of wheat on a chess board" proverb
package grains

import (
	"errors"
	"math"
)

const testVersion = 1

// Square returns the number of grains on a given square
func Square(n int) (uint64, error) {
	n--
	if n < 0 || n > 63 {
		return 0, errors.New("square address out of range")
	}
	return uint64(math.Pow(2, float64(n))), nil
}

// Total calculates the total number of grains on the chessboard and returns them
func Total() uint64 {
	return fasterTotal()
}

func slowerTotal() uint64 { //averages 365 ns/op
	var square, total uint64
	square = 1
	for i := 0; i < 64; i++ {
		total += square
		square = square * 2
	}
	return total
}

func fasterTotal() uint64 { // averages 352ns/op
	var total uint64
	for i := 0; i < 66; i++ {
		total += uint64(math.Pow(2, float64(i)))
	}
	return total
}