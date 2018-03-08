package cryptosquare

import (
	"math"
	"strings"
)

func normalize(pt string) string {
	runes := make([]rune, 0)

	for _, c := range strings.ToLower(pt) {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			runes = append(runes, c)
		}
	}

	return string(runes)
}

func Encode(pt string) string {
	temp := normalize(pt)

	col := int(math.Ceil(math.Sqrt(float64(len(temp)))))

	result := make([]byte, 0)
	for i := 0; i < col; i++ {
		if i > 0 {
			result = append(result, ' ')
		}
		for j := i; j < len(temp); j += col {
			result = append(result, temp[j])
		}
	}

	return string(result)
}
