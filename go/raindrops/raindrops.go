package raindrops

import (
	"strconv"
)

func Convert(x int) string {
	var output string
	if x%3 == 0 {
		output += "Pling"
	}
	if x%5 == 0 {
		output += "Plang"
	}
	if x%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		output = strconv.Itoa(x)
	}
	return output
}
