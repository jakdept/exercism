// Package letter is written to ineficiently total the letter count from a given input
package letter

import "sync"

// FreqMap is used to return the rune totals from a string
type FreqMap map[rune]int

// Frequency totals the runes in a string in a single-threaded fashion
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency totals the letts in a slice of strings concurrentl
func ConcurrentFrequency(inputs []string) FreqMap {
	c := make(chan rune)
	var wg sync.WaitGroup

	for _, each := range inputs {
		wg.Add(1)
		go reader(each, c, &wg)
	}
	go closer(c, &wg)
	return totaler(c)
}

func reader(s string, out chan<- rune, wg *sync.WaitGroup) {
	for _, r := range s {
		out <- r
	}
	wg.Done()
}

func closer(c chan<- rune, wg *sync.WaitGroup) {
	wg.Wait()
	close(c)
}

func totaler(inputChan <-chan rune) FreqMap {
	m := make(FreqMap)
	for char := range inputChan {
		m[char]++
	}
	return m
}