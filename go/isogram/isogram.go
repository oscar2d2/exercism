package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(x string) bool {
	m := make(map[rune]bool)

	for _, c := range strings.ToLower(x) {
		_, exist := m[c]

		if !unicode.IsLetter(c) {
			continue
		}

		if exist {
			return false
		}

		m[c] = true
	}

	return true
}
