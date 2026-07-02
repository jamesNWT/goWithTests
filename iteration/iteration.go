package iteration

import "strings"

// repeats the given characer repeatCount number of times
func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}
