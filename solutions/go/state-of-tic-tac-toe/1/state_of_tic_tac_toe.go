package stateoftictactoe

import (
	"errors"
	"strings"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	xCount, oCount := count(&board)

	if oCount > xCount || xCount > oCount+1 {
		return "", errors.New("Invalid board")
	}

	xWins, oWins := 0, 0

	for _, row := range board {
		switch row {
		case "XXX":
			xWins++
		case "OOO":
			oWins++
		}
	}

	for i := 0; i < 3; i++ {
		var col strings.Builder

		col.WriteByte(board[0][i])
		col.WriteByte(board[1][i])
		col.WriteByte(board[2][i])

		switch col.String() {
		case "XXX":
			xWins++
		case "OOO":
			oWins++
		}
	}

	var fallingDiag strings.Builder
	var ascendingDiag strings.Builder

	fallingDiag.WriteByte(board[0][0])
	fallingDiag.WriteByte(board[1][1])
	fallingDiag.WriteByte(board[2][2])

	ascendingDiag.WriteByte(board[0][2])
	ascendingDiag.WriteByte(board[1][1])
	ascendingDiag.WriteByte(board[2][0])

	switch fallingDiag.String() {
	case "XXX":
		xWins++
	case "OOO":
		oWins++
	}

	switch ascendingDiag.String() {
	case "XXX":
		xWins++
	case "OOO":
		oWins++
	}

	if xWins > 0 && oWins > 0 {
		return "", errors.New("Invalid game result: both players won")
	}

	if (xWins == 0 && oWins == 0) && (xCount == 5 && oCount == 4) {
		return Draw, nil
	}

	if xWins > 0 || oWins > 0 {
		return Win, nil
	}

	return Ongoing, nil

}

func count(board *[]string) (int, int) {
	xCount, oCount := 0, 0

	for _, row := range *board {
		xCount += strings.Count(row, "X")
		oCount += strings.Count(row, "O")
	}

	return xCount, oCount
}
