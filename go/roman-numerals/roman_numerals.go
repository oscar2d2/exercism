package romannumerals

import (
	"bytes"
	"errors"
)

var interval = [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

var m = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func ToRomanNumeral(arabic int) (roman string, err error) {
	if arabic <= 0 || arabic > 3000 {
		err = errors.New("Out of range")
		return
	}

	var buffer bytes.Buffer
	ind := 0

	for arabic > 0 {
		bound := interval[ind]

		if arabic >= bound {
			buffer.WriteString(m[bound])
			arabic -= bound
		} else {
			ind++
		}
	}

	return buffer.String(), nil
}
