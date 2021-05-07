// Package greeting is used to provide greetings.
package greeting

import "fmt"

const testVersion = 3

// HelloWorld is used to generate the greeting.
// If a name is provided, the greeting will be personalized.
// If no name is provided, a generic greeting will be returned.
func HelloWorld(s ...string) string {
	if len(s) > 0 && s[0] != "" {
		return fmt.Sprintf("Hello, %s!", s)
	}
	return "Hello, World!"
}
