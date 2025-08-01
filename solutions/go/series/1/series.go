package series

func All(n int, s string) []string {
	result := []string{}

	for i, _ := range(s) {
		if i + n > len(s) {
			break
		}

		result = append(result, s[i: i + n])
	}

	return result
}

func UnsafeFirst(n int, s string) (string) {
	if n > len(s) {
		return ""
	}

	return s[0:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}

	return s[0:n], true
}
