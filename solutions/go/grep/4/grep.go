package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Grep struct {
	flags Flags
}

type Flags struct {
	multiple_files    bool
	show_line_numbers bool
	show_filenames    bool
	case_insensitive  bool
	invert            bool
	whole_line        bool
}

func parseFlags(input []string, lenFiles int) Flags {
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

	if lenFiles > 1 {
		flags.multiple_files = true
	}

	return flags
}

func (grep *Grep) match(pattern string, text string) bool {
	ml := grep.matchLine(pattern, text)

	return (ml || grep.flags.invert) && !(ml && grep.flags.invert)
}

func (grep *Grep) matchLine(pattern string, text string) bool {
	var p string = pattern

	if grep.flags.whole_line {
		p = fmt.Sprintf("^%s$", pattern)
	}

	if grep.flags.case_insensitive {
		p = fmt.Sprintf("(?i)%s", p)
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

func newGrep(flags []string, lenFiles int) *Grep {
	grep := Grep{}

	grep.flags = parseFlags(flags, lenFiles)

	return &grep
}

func Search(pattern string, flags, files []string) []string {
	result := []string{}
	grep := newGrep(flags, len(files))

	for _, file := range files {
		lineNum := 1

		f, _ := os.Open(file)

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			if grep.match(pattern, scanner.Text()) {
				result = append(result, resultEntry(scanner.Text(), file, lineNum, &grep.flags))
				if grep.flags.show_filenames {
					break
				}

			}

			lineNum++
		}

		f.Close()
	}

	return result
}
