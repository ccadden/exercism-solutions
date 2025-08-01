package prime

func Factors(n int64) []int64 {
	factors := []int64{}
	for {
		if n % 2 == 0 {
			factors = append(factors, 2)
			n = n / 2
		} else {
			break
		}
	}

	var f int64 = 3

	for {
		if (f * f) > n {
			break
		}

		if n % f == 0 {
			factors = append(factors, f)
			n = n / f
		} else {
			f += 2
		}
	}

	if n != 1 {
		factors = append(factors, n)
	}

	return factors
}
