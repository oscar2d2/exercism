package pangram

import "strings"

func IsPangram(input string) bool {
	var m = make(map[rune]bool, 0)

	for _, c := range strings.ToLower(input) {
		if c >= 'a' && c <= 'z' {
			m[c] = true
		}
	}

	if len(m) == 26 {
		return true
	}

	return false
}
