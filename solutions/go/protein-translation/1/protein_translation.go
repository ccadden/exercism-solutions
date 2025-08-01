package protein

import "errors"

var ErrStop error = errors.New("Stop sequence encountered")
var ErrInvalidBase error = errors.New("InvalidBase")

func FromRNA(rna string) ([]string, error) {
	codonCount := len(rna) / 3
	res := []string{}
	var codon string
	var err error
	var startIndex int
	var endIndex int

	for i:=0; i < codonCount; i++ {
		startIndex =  i * 3
		endIndex = (i + 1) * 3

		codon, err = FromCodon(rna[startIndex:endIndex])

		if err == ErrStop {
			return res, nil
		} else if err == ErrInvalidBase {
			return []string{}, err
		}

		res = append(res, codon)
	}

	return res, nil
}

func FromCodon(codon string) (string, error) {
	switch {
	case codon == "AUG":
		return "Methionine", nil
	case codon == "UUU" || codon == "UUC":
		return "Phenylalanine", nil
	case codon == "UUA" || codon == "UUG":
		return "Leucine", nil
	case codon == "UCU" || codon == "UCC" || codon == "UCA" || codon == "UCG":
		return "Serine", nil
	case codon == "UAU" || codon == "UAC":
		return "Tyrosine", nil
	case codon == "UGU" || codon == "UGC":
		return "Cysteine", nil
	case codon == "UGG":
		return "Tryptophan", nil
	case codon == "UGU" || codon == "UGC":
		return "Cysteine", nil
	case codon == "UAA" || codon == "UAG" || codon == "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
