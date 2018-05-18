package diamond

import (
	"bytes"
	"errors"
)

func Gen(ch byte) (string, error) {
	if ch < 'A' || ch > 'Z' {
		return "", errors.New("nil input")
	}

	var buffer bytes.Buffer

	n := (int)(ch-'A')*2 + 1
	mid := n / 2
	offset := 0

	printChar := 'A'
	dummyChar := ' '

	for row := 0; row < n; row++ {
		for i := 0; i < n; i++ {
			if ((row == 0 || row == n-1) && i == mid) ||
				(i == mid-(row-offset) || i == mid+(row-offset)) {
				buffer.WriteRune(printChar)
			} else {
				buffer.WriteRune(dummyChar)
			}
		}
		buffer.WriteRune('\n')

		if row < mid {
			printChar += 1
		} else {
			printChar -= 1
			offset += 2
		}
	}

	return buffer.String(), nil
}
