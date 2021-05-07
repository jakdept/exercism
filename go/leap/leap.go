// Package leap provies a leap year detection utility.
package leap

const testVersion = 2

// IsLeapYear takes a year and returns if it is a leap year.
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		// 400 is a product of 4 and 100 - if divisible all conditions are met
		return true
	}
	if year%100 == 0 {
		// if the year was not divisible by 4, it would've been rejected anyway
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}