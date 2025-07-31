package parsinglogfiles

import "regexp"
import "fmt"

var ValidLog = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
var CustomSeparator = regexp.MustCompile(`<[~*=-]*>`)
var CaseInsensitivePassword = regexp.MustCompile(`(?i)".*password.*"`)
var EndOfLine = regexp.MustCompile(`end-of-line\d*`)
var UserTag = regexp.MustCompile(`User\s+(\w+)`)

func IsValidLine(text string) bool {
	return ValidLog.MatchString(text)
}

func SplitLogLine(text string) []string {
	return CustomSeparator.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int

	for _, line := range(lines) {
		if CaseInsensitivePassword.MatchString(line) {
			count ++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	return string(EndOfLine.ReplaceAllString(text, ""))
}

func TagWithUserName(lines []string) []string {
	for i, line := range(lines) {
		if UserTag.MatchString(line) {
			submatch := UserTag.FindStringSubmatch(line)

			lines[i] = fmt.Sprintf("[USR] %s %s", submatch[1], line)
		}
	}
	return lines
}
