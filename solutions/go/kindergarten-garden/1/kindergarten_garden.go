package kindergarten

import (
	"errors"
	"regexp"
	"strings"
	"slices"
)

var InvalidPlants = regexp.MustCompile(`[^VCGR\n]`)

type Garden struct {
	WindowRow string
	BackRow string
	Children []string
}

func parseDiagram(diagram string, numChildren int) ([]string, bool) {
	if InvalidPlants.MatchString(diagram) {
		return []string{}, false
	}

	rows := strings.Split(diagram, "\n")

	if len(rows) != 3 {
		return []string{}, false
	}

	windowRow := strings.TrimSpace(rows[1])
	backRow := strings.TrimSpace(rows[2])

	if len(windowRow) != len(backRow) {
		return []string{}, false
	}

	if len(windowRow) % 2 != 0 || len(backRow) % 2 != 0 {
		return []string{}, false
	}

	if len(windowRow) != 2 * numChildren || len(backRow) != 2 * numChildren {
		return []string{}, false
	}

	return []string{windowRow, backRow}, true
}

func parseChildren(children []string) ([]string, bool) {
	childrenCopy := slices.Clone(children)
	slices.Sort(childrenCopy)

	for i:=1; i<len(childrenCopy); i++ {
		if childrenCopy[i] == childrenCopy[i-1] {
			return []string{}, false
		}
	}

	return childrenCopy, true
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	rows, valid := parseDiagram(diagram, len(children))

	if !valid {
		return nil, errors.New("Invalid diagram")
	}

	kids, valid := parseChildren(children)

	if !valid {
		return nil, errors.New("Invalid children list")
	}

	return &Garden{
		WindowRow: rows[0],
		BackRow: rows[1],
		Children: kids,
	}, nil
}

func translatePlants(s string) ([]string, bool) {
	res := []string{}

	for _, char := range(s) {
		switch char {
		case 'V':
			res = append(res, "violets")
		case 'C':
			res = append(res, "clover")
		case 'G':
			res = append(res, "grass")
		case 'R':
			res = append(res, "radishes")
		default:
			return []string{}, false
		}
	}

	return res, true
}

func (g *Garden) Plants(child string) ([]string, bool) {
	i := slices.Index(g.Children, child)

	if i == -1 {
		return []string{}, false
	}

	start := i * 2
	end := (i * 2) + 2

	windowRow, ok := translatePlants(g.WindowRow[start:end])

	if !ok {
		return []string{}, false
	}

	backRow, ok := translatePlants(g.BackRow[start:end])

	if !ok {
		return []string{}, false
	}

	return append(windowRow, backRow...), true
}
