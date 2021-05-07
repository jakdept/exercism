// Package pythagorean implements some utilities for right triangles
package pythagorean

import "math"

// Triplet represents a right triangle with given side lengths
// When returned from a function, the three elements of each returned triplet
// will be in order such that t[0] <= t[1] <= t[2].
type Triplet [3]int

// Perimeter returns the perimiter of the Triplet
func (t Triplet) Perimeter() int {
	return t[0] + t[1] + t[2]
}

// Area returns the area of the Triplet
func (t Triplet) Area() int {
	return t[0] * t[1] / 2
}

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive. The three elements of each returned triplet must
// be in order, t[0] <= t[1] <= t[2], and the list of triplets must be in
// lexicographic order.
func Range(min, max int) (results []Triplet) {
	if max <= 0 {
		return []Triplet{}
	}
	if min <= 0 {
		min = 1
	}

	for a := min; a <= max-2; a++ {
		for b := a + 1; b <= max-1; b++ {
			c := math.Hypot(float64(a), float64(b))
			if c == math.Floor(c) {
				results = append(results, Triplet{a, b, int(c)})
			}
		}
	}
	return
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p. The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order.
func Sum(p int) (results []Triplet) {
	// {a,b,c} < p ∴ a + b < c ∴ c < p/2
	for _, each := range Range(1, p/2) {
		if each.Perimeter() == p {
			results = append(results, each)
		}
	}
	return
}