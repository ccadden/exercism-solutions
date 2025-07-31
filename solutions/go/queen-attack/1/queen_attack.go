package queenattack

import "errors"
import "fmt"

type Square struct {
	rank byte
	file byte
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == blackPosition {
		return false, errors.New("Pieces cannot share the same square")
	}

	whiteSquare, whiteOk := parsePosition(whitePosition)

	if !whiteOk {
		return false, errors.New("White position invalid")
	}

	blackSquare, blackOk := parsePosition(blackPosition)

	if !blackOk {
		return false, errors.New("Black position invalid")
	}

	return (whiteSquare.sharesRank(blackSquare) ||
		whiteSquare.sharesFile(blackSquare) ||
		whiteSquare.sharesDiagonal(blackSquare)), nil
}

func parsePosition(position string) (Square, bool) {
	if len(position) != 2 {
		return Square{}, false
	}

	file := position[0]

	if file < 97 || file > 104 {
		return Square{}, false
	}

	rank := position[1]

	if rank < 49 || rank > 56 {
		return Square{}, false
	}

	return Square{rank: rank, file: file}, true
}

func (s *Square) sharesRank(os Square) bool {
	return s.rank == os.rank
}

func (s *Square) sharesFile(os Square) bool {
	return s.file == os.file
}

func (s *Square) sharesDiagonal(os Square) bool {
	return abs(s.rank-os.rank) == abs(s.file-os.file)
}

func (s *Square) String() string {
	return fmt.Sprintf("%c%c", s.file, s.rank)
}

func abs(a byte) byte {
	if a < 0 {
		return 0 - a
	}

	return a
}
