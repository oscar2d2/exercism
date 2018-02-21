package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("length mismatch")
	}

	distance := 0
	for i, c := range a {
		if c != rune(b[i]) {
			distance++
		}
	}
	return distance, nil
}
