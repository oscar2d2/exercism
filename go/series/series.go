package series

func All(n int, s string) []string {
	var temp = make([]byte, n)
	var buf = make([]string, len(s)-n+1)

	for i := 0; i < len(s)-n+1; i++ {
		for j := 0; j < n; j++ {
			temp[j] = s[i+j]
		}
		buf[i] = string(temp)
	}

	return buf
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		ok = false
		return
	}

	first = UnsafeFirst(n, s)
	ok = true
	return
}
