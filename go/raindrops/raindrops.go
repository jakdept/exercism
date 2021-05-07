// Package raindrops is a cheap fizzbuzz clone. It converts a number to a string.
package raindrops

import "strconv"

const testVersion = 2

// Convert takes a number, and based on the factors of the string, prints a given output string.
func Convert(in int) (out string) {
	modify(in, 3, "Pling", &out)
	modify(in, 5, "Plang", &out)
	modify(in, 7, "Plong", &out)
	if out == "" {
		out = strconv.Itoa(in)
	}
	return
}

func modify(in, divisor int, append string, out *string) {
	if in%divisor == 0 {
		*out += append
	}
	return
}

// The test program has a benchmark too.  How fast does your Convert convert?