package etl

import (
	"strings"
)

func Transform(input map[int][]string) map[string]int {
	var output = make(map[string]int, 0)

	for point, cs := range input {
		for _, c := range cs {
			output[strings.ToLower(c)] = point
		}
	}

	return output
}
