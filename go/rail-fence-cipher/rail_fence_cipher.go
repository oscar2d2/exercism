package railfence

import (
	"strings"
)

func Encode(message string, n int) string {
	str := strings.Replace(message, " ", "", -1)
	rails := make([][]rune, n)
	result := make([]rune, 0)

	for i, _ := range rails {
		rails[i] = make([]rune, 0)
	}

	i := 0
	inc := true
	for _, c := range str {
		rails[i] = append(rails[i], c)
		if inc {
			i++
		} else {
			i--
		}

		if i == n-1 || i == 0 {
			inc = !inc
		}
	}

	for _, rail := range rails {
		for _, c := range rail {
			result = append(result, c)
		}
	}

	return string(result)
}
func Decode(message string, n int) string {
	str := strings.Replace(message, " ", "", -1)
	railsCnt := make([]int, n)
	rails := make([][]rune, n)
	result := make([]rune, len(str))

	for i := 0; i < n; i++ {
		railsCnt[i] = 0
		rails[i] = make([]rune, 0)
	}

	ind := 0
	inc := true

	for range str {
		railsCnt[ind]++
		if inc {
			ind++
		} else {
			ind--
		}
		if ind == n-1 || ind == 0 {
			inc = !inc
		}
	}

	temp := 0
	for i, cnt := range railsCnt {
		for j := 0; j < cnt; j++ {
			rails[i] = append(rails[i], rune(str[temp+j]))
		}
		temp += cnt
	}

	ind = 0
	inc = true
	for i := 0; i < len(str); i++ {
		result[i] = rails[ind][0]
		rails[ind] = rails[ind][1:len(rails[ind])]

		if inc {
			ind++
		} else {
			ind--
		}

		if ind == n-1 || ind == 0 {
			inc = !inc
		}
	}

	return string(result)
}
