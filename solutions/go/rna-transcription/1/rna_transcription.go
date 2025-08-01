package strand

import "strings"

func ToRNA(dna string) string {
	var b strings.Builder

	for _, nucleotide := range(dna) {
		switch nucleotide {
		case 'G':
			b.WriteString("C")
		case 'C':
			b.WriteString("G")
		case 'T':
			b.WriteString("A")
		case 'A':
			b.WriteString("U")
		}
	}

	return b.String()
}
