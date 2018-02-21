package grains

import (
	"errors"
	"math"
)

func Square(x int) (uint64, error) {
	if x <= 0 || x > 64 {
		return 0, errors.New("")
	}

	return uint64(math.Pow(float64(2), float64(x-1))), nil
}

func Total() uint64 {
	max := uint64(math.Pow(float64(2), float64(64)))
	return max + max - 1
}
