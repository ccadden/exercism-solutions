package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Flags struct {
	multiple_files    bool
	show_line_numbers bool
	show_filenames    bool
	case_insensitive  bool
	invert            bool
	whole_line        bool
}

func parseFlags(input []string) *Flags {
	flags := Flags{}

	for _, flag := range input {
		switch flag {
		case "-n":
			flags.show_line_numbers = true
		case "-l":
			flags.show_filenames = true
		case "-i":
			flags.case_insensitive = true
		case "-v":
			flags.invert = true
		case "-x":
			flags.whole_line = true
		}
	}

	return &flags
}

func match(pattern string, text string, flags *Flags) bool {
	ml := matchLine(pattern, text, flags)

	return (ml || flags.invert) && !(ml && flags.invert)
}

func matchLine(pattern string, text string, flags *Flags) bool {
	var p string = pattern

	if flags.whole_line {
		p = "^" + pattern + "$"
	}

	if flags.case_insensitive {
		p = "(?i)" + p
	}

	ok, _ := regexp.MatchString(p, text)

	return ok
}

func resultEntry(line, filename string, lineNum int, flags *Flags) string {
	if flags.show_filenames {
		return filename
	}

	var result string = line

	if flags.show_line_numbers {
		result = fmt.Sprintf("%v:%s", lineNum, result)
	}

	if flags.multiple_files {
		result = fmt.Sprintf("%s:%s", filename, result)
	}

	return result

}

func Search(pattern string, flags, files []string) []string {
	result := []string{}
	grep := parseFlags(flags)

	if len(files) > 1 {
		grep.multiple_files = true
	}

	for _, file := range files {
		lineNum := 1

		f, _ := os.Open(file)

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			if match(pattern, scanner.Text(), grep) {
				result = append(result, resultEntry(scanner.Text(), file, lineNum, grep))
				if grep.show_filenames {
					break
				}

			}

			lineNum++
		}

		f.Close()
	}

	return result
}
