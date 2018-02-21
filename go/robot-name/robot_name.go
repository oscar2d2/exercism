package robotname

import (
	"math/rand"
	"time"
)

var inUse = make(map[string]bool)

type Robot struct {
	name string
}

func (r *Robot) Name() string {
	if r.name == "" {
		r.name = genRobotName()
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = genRobotName()
}

func genRobotName() string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const nums = "0123456789"
	const prefixSize = 2
	const postfixSize = 3

	result := make([]byte, prefixSize+postfixSize)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		for i := 0; i < prefixSize; i++ {
			result[i] = chars[r.Intn(len(chars))]
		}
		for i := prefixSize; i < prefixSize+postfixSize; i++ {
			result[i] = nums[r.Intn(len(nums))]
		}

		if _, prs := inUse[string(result)]; !prs {
			break
		}
	}
	inUse[string(result)] = true

	return string(result)
}
