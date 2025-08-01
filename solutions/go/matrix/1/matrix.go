package matrix

import (
	"errors"
	// "fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	matrix := [][]int{}
	lines := strings.Split(s, "\n")
	var numCols int

	for i, line := range(lines) {
		row := []int{}
		nums := strings.Fields(line)

		for _, num := range(nums) {
			n, err := strconv.Atoi(num)

			if err != nil {
				return nil, errors.New("Problem parsing input")
			}
			row = append(row, n)
		}

		if i > 0 {
			if len(row) != numCols {
				return [][]int{}, errors.New("Uneven matrix provided")
			}
		} else {
			numCols = len(row)
		}

		matrix = append(matrix, row)
	}

	return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cols := [][]int{}
	for i := range(m[0]) {
		col := []int{}

		for _, row := range(m) {
			col = append(col, row[i])
		}

		cols = append(cols, col)
	}

	return cols
}

func (m Matrix) Rows() [][]int {
	rows := [][]int{}
	for _, row := range(m) {
		newRow := make([]int, len(row))
		copy(newRow, row)
		rows = append(rows, newRow)
	}

	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m) || col >= len(m[0]) {
		return false
	}

	m[row][col] = val
	return true
}
