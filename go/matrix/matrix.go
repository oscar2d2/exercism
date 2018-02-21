package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type matrix [][]int

func New(in string) (m matrix, err error) {
	rows := strings.Split(in, "\n")
	rowSize := len(rows)
	colSize := len(strings.Split(rows[0], " "))

	m = make([][]int, rowSize)
	for i := range m {
		m[i] = make([]int, colSize)
	}

	for i, r := range rows {
		cols := strings.Split(strings.TrimSpace(r), " ")
		if len(cols) != colSize {
			err = errors.New("row size not match")
			return
		}
		for j, c := range cols {
			m[i][j], err = strconv.Atoi(c)
			if err != nil {
				return
			}
		}
	}

	return m, nil
}

func (m matrix) Rows() [][]int {
	m2 := make([][]int, len(m))
	for i, row := range m {
		m2[i] = make([]int, len(row))
		for j, col := range row {
			m2[i][j] = col
		}
	}
	return m2
}

func (m matrix) Cols() [][]int {
	m2 := make([][]int, len(m[0]))
	for i := range m2 {
		m2[i] = make([]int, len(m))
	}

	for i, row := range m {
		for j, col := range row {
			m2[j][i] = col
		}
	}

	return m2
}
func (m *matrix) Set(r, c, val int) bool {
	mat := *m

	if r < 0 || r >= len(mat) || c < 0 || c >= len(mat[0]) {
		return false
	}

	mat[r][c] = val
	return true

}
