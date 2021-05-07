// Package scrabble blah
package scrabble

import "strings"

const testVersion = 4

// using this scoreboard since i'm required to
var oldScoreBoard = map[int][]string{
	1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	2:  {"D", "G"},
	3:  {"B", "C", "M", "P"},
	4:  {"F", "H", "V", "W", "Y"},
	5:  {"K"},
	8:  {"J", "X"},
	10: {"Q", "Z"},
}

var scoreBoard map[string]int

func init() {
	scoreBoard = Transform(oldScoreBoard)
}

// Transform transforms the style of the scoreboard to something that works better.
func Transform(input map[int][]string) map[string]int {
	// using the previous exercise to transform the scoreboard to something more sensible
	output := map[string]int{}
	for value, characters := range input {
		for _, char := range characters {
			output[strings.ToLower(char)] = value
		}
	}
	return output
}

// Score determines the score of a word in Scrabble
func Score(word string) int {
	var score int
	word = strings.ToLower(word)
	for _, char := range word {
		score += scoreBoard[string(char)]
	}
	return score
}