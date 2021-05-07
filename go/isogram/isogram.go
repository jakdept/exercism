// Package isogram provides a function to detect Isograms.
package isogram

import (
	"sort"
	"strings"
	"unicode"
)

const testVersion = 1

// IsIsogram detects if a given string is an Isogram.
// It does this by looking if any of the characters within that string are repeated.
func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	s = filterString(s)
	parts := strings.SplitN(s, "", -1)
	sort.Strings(parts)

	for i := 0; i < len(parts)-1; i++ {
		if parts[i] == parts[i+1] {
			return false
		}
	}
	return true
}

// removes all non-letter characters from a string
func filterString(s string) (cleaned string) {
	for _, char := range s {
		if unicode.IsLetter(char) {
			cleaned += string(char)
		}
	}
	return
}