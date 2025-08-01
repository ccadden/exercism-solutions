package bob

import (
	"strings"
)

const Letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsQuestion(remark string) bool {
	if len(remark) < 1 {
		return false
	}

	return remark[len(remark) - 1] == '?'
}

func IsAllCaps(remark string) bool {
	upcase := strings.ToUpper(remark)
	return remark == upcase && strings.ContainsAny(upcase, Letters)
}

func IsYellQuestion(remark string) bool {
	return IsAllCaps(remark) && IsQuestion(remark)
}

func IsSilence(remark string) bool {
	return len(strings.TrimSpace(remark)) == 0
}

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case IsYellQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case IsQuestion(remark):
		return "Sure."
	case IsAllCaps(remark):
		return "Whoa, chill out!"
	case IsSilence(remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
