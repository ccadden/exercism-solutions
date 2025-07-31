package sieve

func Sieve(limit int) []int {
	/*
	algorithm Sieve of Eratosthenes is
	input: an integer n > 1.
	output: all prime numbers from 2 through n.

	let A be an array of Boolean values, indexed by integers 2 to n,
	initially all set to true.

	for i = 2, 3, 4, ..., not exceeding √n do
	if A[i] is true
	for j = i2, i2+i, i2+2i, i2+3i, ..., not exceeding n do
	set A[j] := false

	return all i such that A[i] is true.

	*/

	boolSlice := make([]bool, limit + 1)

	for index, _ := range(boolSlice) {
		boolSlice[index] = true
	}

	primes := []int{}

	for i:=2; i <= limit; i++ {
		if boolSlice[i] {
			primes = append(primes, i)

			j := 2
			for {
				ij := j * i
				if ij > limit {
					break
				}

				boolSlice[ij] = false

				j++
			}
		}
	}

	return primes
}
