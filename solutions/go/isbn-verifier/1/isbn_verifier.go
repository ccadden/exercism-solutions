package isbn

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.TrimSpace(isbn)
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, num := range(isbn) {
		if num == 'X' {
			if i == 9 {
				sum += 10
				continue
			} else {
				return false
			}
		} else if num < 48 || num > 57 {
			return false
		}

		sum += (10 - i) * int(num - 48)
	}

	return sum % 11 == 0
}
