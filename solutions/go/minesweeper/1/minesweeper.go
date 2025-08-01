package minesweeper

import (
	"fmt"
	"strings"
)

func Annotate(board []string) []string {
	res := []string{}

	for y, row := range board {
		var b strings.Builder

		for x, char := range row {

			if char == '*' {
				b.WriteRune('*')
				continue
			}

			b.WriteString(countAdjacentMines(board, x, y))
		}
		res = append(res, b.String())
	}

	return res
}

func countAdjacentMines(board []string, x, y int) string {
	count := 0

	startX := Max(x-1, 0)
	endX := Min(x+2, len(board[0]))

	startY := Max(0, y-1)
	endY := Min(y+2, len(board))

	for j := startY; j < endY; j++ {
		for i := startX; i < endX; i++ {
			if board[j][i] == '*' {
				count++
			}
		}
	}

	if count > 0 {
		return fmt.Sprintf("%v", count)
	}

	return " "
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
