// Package hamming provides a mechanism for calculating the hamming distance
package hamming

import (
	"fmt"
	"strings"
)

const testVersion = 5

// Distance takes two DNS sequences and calculates the difference between them.
// An error is returned if the strings are of different lengths.
func Distance(a, b string) (num int, err error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("input strings have differing length")
	}

	defer func() {
		if rErr := recover(); rErr != nil {
			num = 0
			err = fmt.Errorf("failed to read from reader")
		}
	}()

	aR := strings.NewReader(a)
	bR := strings.NewReader(b)

	for aR.Len() > 0 {
		if readRune(aR) != readRune(bR) {
			num++
		}
	}
	return
}

func readRune(r *strings.Reader) rune {
	each, _, err := r.ReadRune()
	if err != nil {
		panic("")
	}
	return each
}