// Package cryptosquare provides a cryptosquare encryption algorithm
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

// Encode encodes a given string with the crypto-square encryption technique.
// The input is sanitized to english lowercase characters without puncuation or spaces.
// Text is then put into a grid horizontally, and read from the grid vertically.
func Encode(s string) string {
	s = strings.ToLower(s)
	s = filterString(s)
	s = matrixFlip(s)
	s = strings.TrimSpace(s)
	return s
}

// removes all non-ascii and non-letter characters from a string
func filterString(s string) string {
	for i := 0; i < len(s); {
		switch {
		case unicode.IsLetter(rune(s[i])):
			i++
		case unicode.IsNumber(rune(s[i])):
			i++
		default:
			if i < len(s)-1 {
				s = s[:i] + s[i+1:]
			} else {
				s = s[:i]
			}
		}
	}
	return s
}

func matrixFlip(s string) (out string) {
	// c - r <= 1 (so they have to be within 1 of each other)
	// c >= r (so you want more columns than rows)
	// (c-1) x r < len < c x r
	// however, I only care about the column width - so i know how much to offset my inner loop
	c := int(math.Sqrt(float64(len(s))))
	if c*c < len(s) {
		// c was rounded down prev, so this only happens when the last row of chars is not full
		c++
	}

	// however, ironically, rather than put this into a matrix and read it out that way
	// i will instead just read them from the string
	for i := 0; i < c; i++ {
		// go "across"" the first "row" of characters
		for j := 0; j < len(s); j += c {
			// go "down" the characters for that "column" by skipping the col size each time
			if i+j < len(s) {
				// if that position is valid, add it to the output string
				out += string(s[i+j])
			}
		}
		// the format specifies there should be spaces between columns in output
		out += " "
	}
	return
}