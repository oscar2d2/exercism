package luhn

import (
	"strings"
	"unicode"
)

func Valid(x string) bool {
	sum := 0

	x = strings.Replace(x, " ", "", -1)

	if x[0] == '0' {
		if len(x) == 1 {
			return false
		} else {
			x = x[1:]
		}
	}

	for i, c := range x {
		if unicode.IsDigit(c) {
			value := int(c - '0')
			if i%2 == 0 {
				value *= 2
				if value > 9 {
					value -= 9
				}
				sum += value
			} else {
				sum += value
			}
		} else {
			return false
		}
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
