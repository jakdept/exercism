// Package luhn does luhn verification on input (CC and SS# validation)
package luhn

import (
	"fmt"
	"strconv"
	"unicode"
)

const testVersion = 1

// Valid determines if a given input string is valid or not with the luhn algorithm
func Valid(s string) bool {
	if len(s) < 2 {
		return false
	}
	s = filterString(s)
	s = luhnModify(s)
	sum, err := sumIntString(s)
	if err == nil && sum%10 == 0 {
		return true
	}
	return false
}

var doubleTable map[string]string

// this ends up being - in code:
// var doubleTable = map[string]string{
// 	"0": "0",
// 	"1": "2",
// 	"2": "4",
// 	"3": "6",
// 	"4": "8",
// 	"5": "1",
// 	"6": "3",
// 	"7": "5",
// 	"8": "7",
// 	"9": "9",
// }

func init() {
	doubleTable = createLuhnTable()
}

func createLuhnTable() map[string]string {
	output := make(map[string]string)
	for i := 0; i < 10; i++ {
		if 2*i > 9 {
			output[strconv.Itoa(i)] = strconv.Itoa((2 * i) - 9)
		} else {
			output[strconv.Itoa(i)] = strconv.Itoa(2 * i)
		}
	}
	return output
}

func luhnModify(s string) (out string) {
	for i := 0; i < len(s); i++ {
		switch {
		// figure out if the current position is odd, corrected for odd len strings
		case (i+len(s))%2 == 0:
			out += doubleTable[string(s[i])]
		case (i+len(s))%2 == 1:
			out += string(s[i])
		default:
			fmt.Printf("%d - i[%d] s[%q] len(s)[%d]\n", (i+len(s))%2, i, s, len(s))
		}
	}
	return
}

// removes all non-number characters from a string
func filterString(s string) string {
	for i := 0; i < len(s); {
		switch {
		case unicode.IsNumber(rune(s[i])):
			i++
		default:
			if i < len(s)-1 {
				s = s[:i] + s[i+1:]
			} else {
				s = s[:i]
			}
		}
	}
	return s
}

// converts a string of digits to ints and adds those integers
func sumIntString(s string) (int, error) {
	var out int
	for i := 0; i < len(s); i++ {
		n, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return 0, err
		}
		out += n
	}
	return out, nil
}