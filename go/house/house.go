// Package house recites the poem "The House that Jack Built""
package house

import "strings"

const testVersion = 1

var verseHead = []string{
	"This is the house that Jack built.",
	"This is the malt",
	"This is the rat",
	"This is the cat",
	"This is the dog",
	"This is the cow with the crumpled horn",
	"This is the maiden all forlorn",
	"This is the man all tattered and torn",
	"This is the priest all shaven and shorn",
	"This is the rooster that crowed in the morn",
	"This is the farmer sowing his corn",
	"This is the horse and the hound and the horn",
}

var verseTail = []string{
	"that lay in the house that Jack built.",
	"that ate the malt",
	"that killed the rat",
	"that worried the cat",
	"that tossed the dog",
	"that milked the cow with the crumpled horn",
	"that kissed the maiden all forlorn",
	"that married the man all tattered and torn",
	"that woke the priest all shaven and shorn",
	"that kept the rooster that crowed in the morn",
	"that belonged to the farmer sowing his corn",
}

// Verse returns the given numbeed verse from the song
func Verse(verse int) string {
	if verse < 1 || verse > 12 {
		return ""
	}

	verse--
	verseLyrics := []string{verseHead[verse]}
	verse--

	for ; verse >= 0; verse-- {
		verseLyrics = append(verseLyrics, verseTail[verse])
	}

	return strings.Join(verseLyrics, "\n")
}

// Song returns the lyrics for the entire song
func Song() string {
	var songLyrics []string

	for i := 1; i < 13; i++ {
		songLyrics = append(songLyrics, Verse(i))
	}

	return strings.Join(songLyrics, "\n\n")
}