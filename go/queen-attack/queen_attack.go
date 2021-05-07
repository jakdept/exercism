// Package queenattack determines if two queens are in danger of each other
package queenattack

import (
	"errors"
	"strconv"
)

const testVersion = 2

type position struct {
	x, y int
}

var horzTranslation = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
}

// CanQueenAttack takes two Chess board positions and reports whether they can attack each other, if they are both queens.
func CanQueenAttack(wStr, bStr string) (bool, error) {
	var w, b position
	var ok bool

	if w, ok = parsePosition(wStr); !ok {
		return false, errors.New("white piece invalid position")
	}
	if b, ok = parsePosition(bStr); !ok {
		return false, errors.New("black piece invalid position")
	}

	switch {
	case w == b:
		return true, errors.New("you cannot have two pieces in the same place")
	case w.x == b.x:
		return true, nil
	case w.y == b.y:
		return true, nil
	case w.x+w.y == b.x+b.y:
		return true, nil
	case w.x-w.y == b.x-b.y:
		return true, nil
	default:
		return false, nil
	}
}

// parses the position of the piece from the input notation
// output 0 based coords, but the inputs are 1 based
func parsePosition(loc string) (position, bool) {
	if len(loc) != 2 {
		return position{}, false
	}
	var result position
	var err error
	var ok bool
	if result.x, ok = horzTranslation[loc[:1]]; !ok {
		return position{}, false
	}
	if result.y, err = strconv.Atoi(loc[1:]); err != nil {
		return position{}, false
	}
	if result.y > 8 || result.y < 1 {
		return position{}, false
	}
	result.y--
	return result, true
}