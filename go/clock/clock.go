// Package clock provides a type clock that manually tracks a given time.
package clock

import "fmt"

const testVersion = 4

// Clock provides a time tracking clock.
// This clock does not actually track time and self update over time.
// Instead, minutes can be manually added/removed.
type Clock struct {
	hour   int
	minute int
}

// New creates a new clock and normalizes it
func New(hour, minute int) Clock {
	newClock := Clock{hour: hour, minute: minute}
	newClock.update()
	return newClock
}

// String prints out the current numbers on the clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds the given minutes to the clock, then normalizes it
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	c.update()
	return c
}

func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-1 * minutes)
}

// internal function that brings the clock back into normal bounds
func (c *Clock) update() {
	for c.minute < 0 {
		c.minute += 60
		c.hour--
	}
	for c.minute >= 60 {
		c.minute -= 60
		c.hour++
	}
	for c.hour < 0 {
		c.hour += 24
	}
	for c.hour >= 24 {
		c.hour -= 24
	}
	return
}
