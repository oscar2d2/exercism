package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(x int) (count int, err error) {
	count = 0

	if x <= 0 {
		err = errors.New("Invalid input")
		return
	}

	for x != 1 {
		if x%2 == 0 {
			x = x / 2
		} else {
			x = 3*x + 1
		}
		count++
	}

	return
}
