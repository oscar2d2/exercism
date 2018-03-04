package summultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	seen := make(map[int]bool, 0)

	for _, e := range divisors {
		for i := e; i < limit; i += e {
			if !seen[i] {
				sum += i
				seen[i] = true
			}
		}
	}

	return sum
}
