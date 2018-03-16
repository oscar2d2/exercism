package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (product int64, err error) {
	if len(digits) < span {
		err = errors.New("span is larger than the number of digits")
		return
	}

	if len(digits) < 0 {
		err = errors.New("empty string")
		return
	}

	if span < 0 {
		err = errors.New("span is invalid")
		return
	}

	if span == 0 || len(digits) == 0 {
		product = 1
		return
	}

	var x, temp int64
	for i := span - 1; i < len(digits); i++ {
		temp = 1
		for j := 0; j < span; j++ {
			x = int64(digits[i-j] - '0')
			if x < 0 || x > 9 {
				product = 0
				err = errors.New("Non numeric character detected")
				return
			}
			temp = temp * x
		}
		if temp > product {
			product = temp
		}
	}

	return
}
