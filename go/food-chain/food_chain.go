package foodchain

import (
	"bytes"
	"strings"
	"text/template"
)

const testVersion = 3

const verse = "I know an old lady who swallowed a {{ .Animal }}.\n{{ with .Second }}{{.}}\n{{ end }}{{ range $index, $element := .Chunks }}{{if $index}}\n{{end}}{{ $element }}{{ end }}"

var animal = []string{
	"fly",
	"spider",
	"bird",
	"cat",
	"dog",
	"goat",
	"cow",
	"horse",
}

var secondLine = []string{
	"I don't know why she swallowed the fly. Perhaps she'll die.",
	"It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
	"She's dead, of course!",
}

var verseLine = []string{
	"I don't know why she swallowed the fly. Perhaps she'll die.",
	"She swallowed the spider to catch the fly.",
	"She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
	"She swallowed the cat to catch the bird.",
	"She swallowed the dog to catch the cat.",
	"She swallowed the goat to catch the dog.",
	"She swallowed the cow to catch the goat.",
}

var verseTempl = template.Must(template.New("verse").Parse(verse))

type verseStructure struct {
	Animal string
	Second string
	Chunks []string
}

// Verse returns a specific verse from the song
func Verse(verse int) string {
	// limits on the verse number
	if verse < 1 || verse > 8 {
		return ""
	}

	// decrement verse number so it matches up
	verse--

	data := verseStructure{Animal: animal[verse], Second: secondLine[verse]}

	if verse > 0 && verse < 7 {
		for ; verse >= 0; verse-- {
			if verseLine[verse] != "" {
				data.Chunks = append(data.Chunks, verseLine[verse])
			}
		}
	}

	buf := bytes.NewBufferString("")
	verseTempl.Execute(buf, data)
	return strings.TrimSpace(buf.String())
}

// Verses returns numbered verses start - stop inclusive from the song
func Verses(start, stop int) string {
	var songLyrics []string

	for i := start; i < stop+1; i++ {
		songLyrics = append(songLyrics, Verse(i))
	}

	return strings.Join(songLyrics, "\n\n")
}

// Song returns the lyrics for the entire song
func Song() string {
	return Verses(1, 8)
}