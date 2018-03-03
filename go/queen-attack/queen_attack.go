package queenattack

import (
	"errors"
	"regexp"
)

func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func CanQueenAttack(wpos, bpos string) (canAttack bool, err error) {
	if len(wpos) > 2 || len(wpos) == 0 || len(bpos) > 2 || len(bpos) == 0 {
		return false, errors.New("length mismatch")
	}

	r, _ := regexp.Compile("[a-h]{1}[1-8]{1}")
	wMatch := r.MatchString(wpos)
	bMatch := r.MatchString(bpos)
	if !wMatch || !bMatch {
		return false, errors.New("Invalid position")
	}

	if wpos == bpos {
		return false, errors.New("Same postition")
	}

	rowDiff := Abs(int(wpos[0]) - int(bpos[0]))
	colDiff := Abs(int(wpos[1]) - int(bpos[1]))
	if rowDiff != colDiff && wpos[0] != bpos[0] && wpos[1] != bpos[1] {
		return false, nil
	}

	return true, nil
}
