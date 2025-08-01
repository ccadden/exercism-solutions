package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, ok := cb[file]
	if !ok {
		return 0
	}
	pieces := 0

	for _, x := range(f) {
		if x {
			pieces ++
		}
	}
	
	return pieces
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank <= 0 {
		return 0
	}

	pieces := 0

	for _, file := range(cb) {
		if rank > len(file) {
			return 0
		}

		val := file[rank - 1]

		if val {
			pieces ++
		}
	}

	return pieces
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return 64 // It is known
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	pieces := 0

	for _, file := range(cb) {
		for _, sq := range(file) {
			if sq {
				pieces ++
			}
		}
	}

	return pieces
}
