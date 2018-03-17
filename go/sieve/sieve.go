package sieve

func Sieve(limit int) []int {
	result := make([]int, 0)
	m := make(map[int]bool, 0)

	for i := 2; i <= limit; i++ {
		if _, exist := m[i]; !exist {
			result = append(result, i)
			for j := 2 * i; j <= limit; j += i {
				m[j] = true
			}
		}
	}

	return result
}
