package cryptosquare

import (
	"unicode"
	"math"
	"strings"
)

func Encode(pt string) string {
    if len(pt) < 1 {
        return ""
    }

	  var b strings.Builder

    matrix := normalize(pt)

    rows := len(matrix)
    cols := len(matrix[0])

    for i:=0; i<cols; i++ {
        for j:=0; j<rows; j++ {
            b.WriteRune(matrix[j][i])
        }

        if i == cols - 1 {
            break
        }

        b.WriteRune(' ')
    }

    return b.String()
}

func normalize(pt string) [][]rune {
    runes := []rune{}

    for _, char := range(pt) {
        if unicode.IsLetter(char) || unicode.IsNumber(char) {
            runes = append(runes, unicode.ToLower(char))
        }
    }

    cols, rows := calcDimensions(len(runes))
    matrix := make([][]rune, rows)

    for i:=0; i < rows; i++ {
        if i == rows - 1 {
            matrix[i] = runes[i * cols:]
            for len(matrix[i]) < cols {
                matrix[i] = append(matrix[i], ' ')
            }
        } else {
            matrix[i] = runes[i * cols: (i + 1) * cols]
        }
    }

    return matrix
}

func calcDimensions(length int) (cols, rows int) {
    cols = int(math.Ceil(math.Sqrt(float64(length))))
    rows = int(math.Ceil(float64(length) / float64(cols)))

    return cols, rows
}
