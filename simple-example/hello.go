package hello

import "strings"

// []string is a slice of strings - variable length array
func SayHello (names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}
	return "Hello " + strings.Join(names, ", ") + "!" 
}