// Package gigasecond returns a given time + 10^9th seconds
package gigasecond

import "time"

// Constant declaration.
// testVersion defines the test versions to compare against
const testVersion = 4

// gigasecond is a constant used to represent the value to add.
const gigasecond = 1000000000 * time.Second

// AddGigasecond returns the given time + 10^9th seconds.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}