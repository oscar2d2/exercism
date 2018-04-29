package say

import (
	"fmt"
	"strings"
)

//ones converts a digit to the english phrase
var ones = []string{
	"", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
}

//teens converts a ones digit in a teen number to the english phrase
var twos = []string{
	"ten", "eleven", "twelve", "thirteen", "fourteen",
	"fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

//tens converts a tens digits and the english phrase
var tens = []string{
	"", "ten", "twenty", "thirty", "forty",
	"fifty", "sixty", "seventy", "eighty", "ninty",
}

//powers coverts the position of a groups of digits to the english phrase
var powers = []string{
	"", "thousand", "million", "billion",
}

func sayHelper(n int64) string {
	if n == 0 {
		return "zero"
	}

	var groups []int64
	for n > 0 {
		groups = append(groups, n%1000)
		n = n / 1000
	}

	result := ""
	for i := len(groups) - 1; i >= 0; i-- {
		if groups[i] != 0 {
			result += sayHelper2(groups[i]) + powers[i] + " "
		}
	}

	return strings.TrimSpace(result)
}

func sayHelper2(n int64) string {
	words := ""
	hundred, ten, one := (n%1000)/100, (n%100)/10, n%10
	if 0 < hundred {
		words += ones[hundred] + " hundred "
	}
	switch {
	case ten == 1:
		words += twos[one] + " "
	case one == 0:
		words += tens[ten] + " "
	case ten == 0:
		words += ones[one] + " "
	default:
		words += fmt.Sprintf("%s-%s ", tens[ten], ones[one])
	}
	return words
}

func Say(n int64) (string, bool) {
	if n < 0 || n > 999999999999 {
		return "", false
	}

	return sayHelper(n), true
}
