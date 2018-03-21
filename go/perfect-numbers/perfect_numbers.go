package perfect

import (
	"errors"
)

type Classification string

const (
	ClassificationPerfect   Classification = "ClassificationPerfect"
	ClassificationDeficient Classification = "ClassificationDeficient"
	ClassificationAbundant  Classification = "ClassificationAbundant"
)

var ErrOnlyPositive error = errors.New("positive only")

func aliquotSum(num int64) int64 {
	var sum int64 = 1

	for i := int64(2); i*i < num; i++ {
		if num%i == 0 {
			sum += i
			sum += num / i
		}
	}

	return sum
}

func Classify(num int64) (cat Classification, err error) {
	if num <= 0 {
		err = ErrOnlyPositive
		return
	}

	sum := aliquotSum(num)

	if sum == 1 || sum < num {
		cat = ClassificationDeficient
	} else if sum == num {
		cat = ClassificationPerfect
	} else if sum > num {
		cat = ClassificationAbundant
	}

	return
}
