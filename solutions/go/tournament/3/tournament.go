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

const ColumnPadding int = -31

type Team struct {
	name          string
	matchesPlayed int
	wins          int
	losses        int
	draws         int
	points        int
}

type Teams map[string]*Team

func (t *Team) MatchPlayed() {
	t.matchesPlayed++
}

func (t *Team) AwardPoints(p int) {
	t.points += p
}

func (t *Team) Win() {
	t.MatchPlayed()
	t.wins++
	t.AwardPoints(3)
}

func (t *Team) Loss() {
	t.MatchPlayed()
	t.losses++
}

func (t *Team) Draw() {
	t.MatchPlayed()
	t.draws++
	t.AwardPoints(1)
}

func Tally(reader io.Reader, writer io.Writer) error {
	var teamNames []string

	teams, err := newTeams(reader)

	if err != nil {
		return err
	}

	for team := range teams {
		teamNames = append(teamNames, team)
	}

	slices.SortFunc(teamNames, func(a, b string) int {
		return cmp.Or(
			cmp.Compare(teams[b].points, teams[a].points),
			cmp.Compare(a, b),
		)
	})

	_, err = writer.Write(
		[]byte(
			fmt.Sprintf("%*s| MP |  W |  D |  L |  P\n",
				ColumnPadding,
				"Team"),
		),
	)

	if err != nil {
		return err
	}

	for _, name := range teamNames {
		team := teams[name]
		if _, err := writer.Write(
			[]byte(fmt.Sprintf(
				"%*s|  %v |  %v |  %v |  %v |  %v\n",
				ColumnPadding,
				name,
				team.matchesPlayed,
				team.wins,
				team.draws,
				team.losses,
				team.points,
			)),
		); err != nil {
			return err
		}
	}

	return nil
}

func newTeams(reader io.Reader) (Teams, error) {
	scanner := bufio.NewScanner(reader)

	teams := Teams{}

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

		firstTeamName := line[0]
		secondTeamName := line[1]
		result := line[2]

		firstTeam, ok := teams[firstTeamName]

		if !ok {
			firstTeam = newTeam(firstTeamName)
			teams[firstTeamName] = firstTeam
		}

		secondTeam, ok := teams[secondTeamName]

		if !ok {
			secondTeam = newTeam(secondTeamName)
			teams[secondTeamName] = secondTeam
		}

		switch result {
		case "win":
			firstTeam.Win()
			secondTeam.Loss()
		case "loss":
			firstTeam.Loss()
			secondTeam.Win()
		case "draw":
			firstTeam.Draw()
			secondTeam.Draw()
		default:
			return nil, errors.New("Invalid match result")
		}
	}

	return teams, nil
}

func newTeam(name string) *Team {
	return &Team{
		name: name,
	}
}
