package pascal

func Triangle(n int) [][]int {

	var m = make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, i+1)

		for j := 0; j < len(m[i]); j++ {
			if j == 0 || j == len(m[i])-1 {
				m[i][j] = 1
			} else {
				m[i][j] = m[i-1][j] + m[i-1][j-1]
			}
		}
	}

	return m
}
