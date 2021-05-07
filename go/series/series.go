// Package slice combines numbers to determine the shortest continous number sequence
package series

// UnsafeFirst is an unsafe function asked to write as a part of the assignment.
// It does not include input sanitization.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First is safe function that returns the first substring of a string, with input sanitization.
func First(n int, s string) (string, bool) {
	if len(s) < n {
		return "", false
	}
	return s[:n], true
}

// All returns all possible substrings with a given length from a string
func All(n int, s string) []string {
	var result []string
	for i := 0; i < len(s)-n+1; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}
