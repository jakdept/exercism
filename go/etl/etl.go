// Package etl converts one scrabble encoding method to another
package etl

import "strings"

// Transform does the actual transormation
func Transform(input map[int][]string) map[string]int {
	output := map[string]int{}
	for value, characters := range input {
		for _, char := range characters {
			output[strings.ToLower(char)] = value
		}
	}
	return output
}