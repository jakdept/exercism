// Package pangram provides a way to determine if a string is a pangram
package pangram

import (
	"sort"
	"strings"
	"unicode"
)

const testVersion = 1

// the characters to match against/complete
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// IsPangram determines if the given input string is a pangram
func IsPangram(s string) bool {
	s = strings.ToLower(s)
	s = filterString(s)
	parts := strings.SplitN(s, "", -1)
	sort.Strings(parts)
	parts = uniqStringSlice(parts)

	solutionParts := strings.SplitN(alphabet, "", -1)
	return equalStringSlice(parts, solutionParts)
}

// removes all non-ascii and non-letter characters from a string
func filterString(s string) string {
	for i := 1; i < len(s)-1; {
		if !unicode.IsLetter(rune(s[i])) || rune(s[i]) > unicode.MaxASCII {
			s = s[:i] + s[i+1:]
		} else {
			i++
		}
	}
	if len(s) > 2 && !unicode.IsLetter(rune(s[0])) {
		s = s[1:]
	}
	if len(s) > 2 && !unicode.IsLetter(rune(s[len(s)-1])) {
		s = s[:len(s)-1]
	}
	return s
}

// deduplicates a sorted sort.StringSlice
func uniqStringSlice(arr []string) []string {
	for i := 1; i < len(arr)-1; {
		if arr[i] == arr[i+1] {
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			i++
		}
	}
	if len(arr) > 2 && arr[0] == arr[1] {
		arr = arr[1:]
	}
	if len(arr) > 2 && arr[len(arr)-1] == arr[len(arr)-2] {
		arr = arr[:len(arr)-1]
	}
	return arr
}

func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}