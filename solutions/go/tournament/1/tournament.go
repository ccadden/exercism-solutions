package tournament

import (
	"bufio"
	"cmp"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
)

var ColumnPadding int = -31

func Tally(reader io.Reader, writer io.Writer) error {
	scores, err := createMap(reader)

	if err != nil {
		return err
	}

	buffer := bufio.NewWriter(writer)


	teams := []string{}
	var longestTeamName int

	for team := range(scores) {
		longestTeamName = MaxInt(longestTeamName, len(team))
		teams = append(teams, team)
	}

	slices.SortFunc(teams, func(a, b string) int {
		return cmp.Or(
			cmp.Compare(scores[b]["P"], scores[a]["P"]),
			cmp.Compare(a, b),
		)
	})

	_, err = buffer.WriteString(fmt.Sprintf("%*s| MP |  W |  D |  L |  P\n", ColumnPadding, "Team"))
	if err != nil {
		return err
	}


	if err = buffer.Flush(); err != nil {
		return err
	}


	for _, team := range(teams) {
		score := scores[team]
		if _, err := buffer.WriteString(
			fmt.Sprintf(
				"%*s|  %v |  %v |  %v |  %v |  %v\n",
				ColumnPadding,
				team,
				score["MP"],
				score["W"],
				score["D"],
				score["L"],
				score["P"],
			),
		); err != nil {
			return err
		}
		if err = buffer.Flush(); err != nil {
			return err
		}
	}

	return nil
}

func createMap(reader io.Reader) (map[string]map[string]int, error) {
	scanner := bufio.NewScanner(reader)
	results := map[string]map[string]int{}

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}

		if strings.HasPrefix(scanner.Text(), "#") {
			continue
		}

		line := strings.Split(scanner.Text(), ";")

		if len(line) != 3 {
			return nil, errors.New("Invalid format")
		}
		_, ok := results[line[0]]
		if !ok {
			results[line[0]] = initializeTeam()
		}

		_, ok = results[line[1]]
		if !ok {
			results[line[1]] = initializeTeam()
		}

		switch line[2] {
		case "win":
			results[line[0]]["MP"]++
			results[line[0]]["W"]++
			results[line[0]]["P"] += 3
			results[line[1]]["MP"]++
			results[line[1]]["L"]++
		case "loss":
			results[line[1]]["MP"]++
			results[line[1]]["W"]++
			results[line[1]]["P"] += 3
			results[line[0]]["MP"]++
			results[line[0]]["L"]++
		case "draw":
			results[line[0]]["MP"]++
			results[line[0]]["D"]++
			results[line[0]]["P"] ++
			results[line[1]]["MP"]++
			results[line[1]]["D"]++
			results[line[1]]["P"]++
		default:
			return nil, errors.New("Invalid match result")
		}
	}

	return results, nil
}

func initializeTeam() map[string]int {
	return map[string]int{
		"MP": 0,
		"W": 0,
		"L": 0,
		"D": 0,
		"P": 0,
	}
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
