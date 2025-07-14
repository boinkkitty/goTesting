package iteration

import "strings"

func Repeat(character string, n int) string {
	var repeated strings.Builder
	for i := 0; i < n; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
