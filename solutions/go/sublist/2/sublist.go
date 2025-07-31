package sublist

// Relation type is defined in relations.go file.

func aContainsB(a, b []int) bool {
	lena := len(a)
	lenb := len(b)

	if lenb == 0 {
		return true
	}

	if lena <= lenb {
		return false
	}

	for i, val := range(a) {
		if i + lenb > lena {
			return false
		}
		

		if val == b[0] {
			contains := aEqualsB(a[i: i + lenb], b)

			if contains {
				return true
			}
		}
	}

	return false
}

func aEqualsB(a, b []int) bool {
	lena := len(a)
	lenb := len(b)

	if lena == 0 && lenb == 0 {
		return true
	}

	if lena != lenb {
		return false
	}

	for i, val := range(a) {
		if val != b[i] {
			return false
		}
	}

	return true
}

func Sublist(l1, l2 []int) Relation {
	len1 := len(l1)
	len2 := len(l2)

	if len1 == len2 {
		if aEqualsB(l1, l2) {
			return RelationEqual
		}
	} else if len1 < len2 {
		if aContainsB(l2, l1) {
			return RelationSublist
		}
	} else {
		if aContainsB(l1, l2) {
			return RelationSuperlist
		}
	}

	return RelationUnequal
}
