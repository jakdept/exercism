// Package acronym provides a mechanism to find the acronym from a given string by some weird rules
package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Abbreviate takes a given string and creates an abbreviation according to some weird rules.
//  * Words are split based on symbols and whitespace.
//  * Then the first character of each word is included in the abbreviation.
//  * Additionally, any upper case letters in a mixed case word after the first is also included.
func Abbreviate(in string) string {
	var abbrev []rune
	// split based on non-letter characters
	parts := strings.FieldsFunc(in, isNonLetter)

	for _, each := range parts {
		// add the first character of each part
		abbrev = append(abbrev, unicode.ToUpper(rune(each[0])))
		if strings.ToUpper(each) != each {
			// if the given part is not all uppsercase (also implies length of at least 1)
			for _, char := range each[1:] {
				if unicode.IsUpper(char) {
					// go through the string after the first char, and append any upper case characters
					abbrev = append(abbrev, char)
				}
			}
		}
	}
	return string(abbrev)
}

func isNonLetter(c rune) bool {
	return !unicode.IsLetter(c)
}