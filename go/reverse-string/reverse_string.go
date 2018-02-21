package reverse

import (
	"unicode/utf8"
)

func String(x string) string {
	result := make([]rune, utf8.RuneCountInString(x))
	i := len(result)

	for _, c := range x {
		i--
		result[i] = c
	}

	return string(result)
}
