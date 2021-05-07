// Package secret allows decryption of handshakes per a set of given rules.
package secret

const testVersion = 1

const (
	wink        = 1 << iota // 1
	doubleBlink             // 2
	closeEyes               // 4
	jump                    // 8
	reverse                 // 16
)

// Handshake allows the decryption of handshakes from a coded uint per the following rules:
//
// First, convert the number to binary. THen, if each position is present, reading from rtl, add the following action to the handshake (in order).
// Powers of 2 for every given step has also been included - you can instead (uint % 2^x) to determine whether the action should be taken.
//  2^0 -> 1  wink
//  2^1 -> 2  double blink
//  2^2 -> 4  close your eyes
//  2^3 -> 8  jump
//  2^4 -> 16 Reverse the order of the handshake
func Handshake(dirtyCode uint) []string {
	code := int(dirtyCode)
	var output []string

	if code&wink != 0 {
		output = append(output, "wink")
	}
	if code&doubleBlink != 0 {
		output = append(output, "double blink")
	}
	if code&closeEyes != 0 {
		output = append(output, "close your eyes")
	}
	if code&jump != 0 {
		output = append(output, "jump")
	}

	if code&reverse != 0 {
		oldOutput := output
		output = []string{}
		for _, step := range oldOutput {
			output = append([]string{step}, output...)
		}
	}
	return output
}