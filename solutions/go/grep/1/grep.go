package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var MULTIPLE_FILES bool = false

// - `-n` Prepend the line number and a colon (':') to each line in the output, placing the number after the filename (if present).
var SHOW_LINE_NUMS bool = false

// - `-l` Output only the names of the files that contain at least one matching line.
var SHOW_FILENAMES bool = false

// - `-i` Match using a case-insensitive comparison.
var CASE_INSENSITIVE bool = false

// - `-v` Invert the program -- collect all lines that fail to match.
var INVERT bool = false

// - `-x` Search only for lines where the search string matches the entire line.
var WHOLE_LINE bool = false

func parseFlags(flags []string) {
	for _, flag := range flags {
		switch flag {
		case "-n":
			SHOW_LINE_NUMS = true
		case "-l":
			SHOW_FILENAMES = true
		case "-i":
			CASE_INSENSITIVE = true
		case "-v":
			INVERT = true
		case "-x":
			WHOLE_LINE = true
		}
	}
}

func match(pattern, text string) bool {
	ml := matchLine(pattern, text)
	return (ml || INVERT) && !(ml && INVERT)
}

func matchLine(pattern, text string) bool {
	var p string = pattern

	if WHOLE_LINE {
		p = "^" + pattern + "$"
	}

	if CASE_INSENSITIVE {
		p = "(?i)" + p
	}

	ok, _ := regexp.MatchString(p, text)

	return ok
}

func resultEntry(line, filename string, lineNum int) string {
	if SHOW_FILENAMES {
		return filename
	}

	var result string = line

	if SHOW_LINE_NUMS {
		result = fmt.Sprintf("%v:%s", lineNum, result)
	}

	if MULTIPLE_FILES {
		result = fmt.Sprintf("%s:%s", filename, result)
	}

	return result

}

func unsetFlags() {
	MULTIPLE_FILES = false
	SHOW_LINE_NUMS = false
	SHOW_FILENAMES = false
	CASE_INSENSITIVE = false
	INVERT = false
	WHOLE_LINE = false
}

func Search(pattern string, flags, files []string) []string {
	result := []string{}
	if len(files) > 1 {
		MULTIPLE_FILES = true
	}

	parseFlags(flags)
	defer unsetFlags()

	for _, file := range files {
		lineNum := 1

		f, _ := os.Open(file)

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			if match(pattern, scanner.Text()) {
				result = append(result, resultEntry(scanner.Text(), file, lineNum))
				if SHOW_FILENAMES {
					break
				}

			}

			lineNum++
		}

		f.Close()
	}

	return result
}
