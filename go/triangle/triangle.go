// Package triangle identifies the type of a given triangle
package triangle

import "math"

const testVersion = 3

// Kind is a type for triangle identification
type Kind string

const (
	NaT = "NotaTriangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
	/* uncomment this bit (and lower) to enable degenerate deteection
	Deg = "degenerate"
	*/
)

// KindFromSides checks each side to make sure it's a valid triangle, and counts the equal sides to determine the type of triangle.
func KindFromSides(a, b, c float64) Kind {
	// pile the three inputs into an array so I can address them without writing out a bunch of stupid code
	side := []float64{a, b, c}
	var equalCount int // will track how many sides are equal here

	// because I'm lazy and do not want to write a bunch of silly checks, doing some stuff to
	for offset := 0; offset < 3; offset++ {
		// pick the three indexes to look at - in order
		x, y, z := offset, offset+1, offset+2
		x, y, z = x%3, y%3, z%3

		switch {
		case side[x] <= 0.0:
			// toss out negative or 0 length sides
			return NaT
		case math.IsNaN(side[x]):
			// toss out sides with NaN length
			return NaT
		case math.IsInf(side[x], 1):
			// toss out sides with infinite length sides
			return NaT
		case math.IsInf(side[x], -1):
			// toss out sides with infinite length sides
			return NaT
		case side[x]+side[y] < side[z]:
			// toss out impossible triangles
			return NaT
		/* uncomment this line (and above) to enable degernate detection
		case side[x]+side[y] == side[z]:
			return Deg
		*/
		case side[x] == side[y]:
			// if the sides are equal, increment the equal sides count
			equalCount++
		}
	}

	switch {
	// of note, it's impossible for just 2 sides to be equal - if so, it'd be 3
	case equalCount == 3:
		return Equ
	case equalCount == 1:
		return Iso
	default:
		return Sca
	}
}