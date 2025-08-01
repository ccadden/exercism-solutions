package spiralmatrix

type Matrix [][]int

func (m Matrix) Contains(number int) bool {
  for _, p := range(m) {
    for _, q := range(p) {
      if q == number {
        return true
      }
    }
  }

  return false
}

func SpiralMatrix(size int) [][]int {
  matrix := initializeMatrix(size)

  if size == 0 {
    return matrix
  }

  x := 0
  y := 0
  dx := 1
  dy := 0

  counter := 1

  for {
    matrix[y][x] = counter

    if y + dy == size || y + dy < 0 || x + dx == size || x + dx < 0 || matrix[y+dy][x+dx] != 0 {
      if !matrix.Contains(0) {
        break
      }

      dy, dx = dx, -dy
    }

    x += dx
    y += dy

    counter++
  }

  return matrix
}

func initializeMatrix(size int) Matrix {
  matrix := make([][]int, size)

  for row := 0; row < size; row++ {
    matrix[row] = make([]int, size)
  }

  return matrix
}
