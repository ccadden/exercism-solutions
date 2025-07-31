package reverse

func Reverse(input string) string {
	runes := []rune(input)

	for i := 0; i < len(runes) / 2; i++ {
        var j = len(runes) - i - 1
        runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
