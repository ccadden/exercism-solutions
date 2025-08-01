package diffsquares

func sqr(n int) int {
	return n * n
}

func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}

	return sqr(sum)
}

func SumOfSquares(n int) int {
	if n == 1 {
		return 1
	} else {
		return sqr(n) + SumOfSquares(n - 1)
	}
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
