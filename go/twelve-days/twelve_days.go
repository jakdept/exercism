// Package twelve provides the lyrics for the song "The Twelve Days of Christmas"
package twelve

import (
	"bytes"
	"html/template"
)

const testVersion = 1

const verse = "On the {{ .Day }} day of Christmas my true love gave to me,{{ range .Chunks }}{{ . }}{{ end }}"

var numberedDays = []string{
	"nonce",
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = []string{
	"On the first day of Christmas my true love gave to me, a Partridge in a Pear Tree.",
	" and a Partridge in a Pear Tree.",
	" two Turtle Doves,",
	" three French Hens,",
	" four Calling Birds,",
	" five Gold Rings,",
	" six Geese-a-Laying,",
	" seven Swans-a-Swimming,",
	" eight Maids-a-Milking,",
	" nine Ladies Dancing,",
	" ten Lords-a-Leaping,",
	" eleven Pipers Piping,",
	" twelve Drummers Drumming,",
}

var verseTempl = template.Must(template.New("verse").Parse(verse))

// Verse returns one specific verse from the song - numbered in natural order
func Verse(verse int) string {
	if verse == 1 {
		return gifts[0]
	}
	if verse > 1 && verse <= 12 {
		var data struct {
			Day    string
			Chunks []string
		}

		data.Day = numberedDays[verse]
		for i := verse; i > 0; i-- {
			data.Chunks = append(data.Chunks, gifts[i])
		}

		buf := bytes.NewBufferString("")

		verseTempl.Execute(buf, data)

		return buf.String()
	}
	return ""
}

// Song returns the entire song, plus an extra newline as that's required for the test cases
func Song() string {
	song := gifts[0]
	for verse := 2; verse < 14; verse++ {
		song += "\n"
		song += Verse(verse)
	}
	return song
}