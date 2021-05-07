// Package bob provides a teenager's conversational skills in functional format
package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Hey encapsulates Bob's conversational skills.
func Hey(in string) string {
	in = strings.TrimSpace(in)
	switch {
	case strings.TrimSpace(in) == "":
		return "Fine. Be that way!"
	case strings.ToUpper(in) == in && strings.IndexFunc(in, unicode.IsLetter) >= 0:
		return "Whoa, chill out!"
	case strings.HasSuffix(in, "?"):
		return "Sure."
	default:
		return "Whatever."
	}
}