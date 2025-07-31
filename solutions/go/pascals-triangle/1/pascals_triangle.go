package pascal

func Triangle(n int) [][]int {
	result := make([][]int, n)

	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			row[j] = pascal(i, j)
		}

		result[i] = row
	}

	return result
}

func pascal(row, column int) int {
	if column == 0 {
		return 1
	}

	if row == 0 {
		return column
	}

	return (row * pascal(row-1, column-1)) / column
}
