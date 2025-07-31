package meetup

import "time"

type WeekSchedule int

const (
	First = iota
	Second
	Teenth
	Third
	Fourth
	Last
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	switch wSched {
	case First:
		return findDate(wDay, month, year, 0)
	case Second:
		return findDate(wDay, month, year, 1)
	case Teenth:
		day := findDate(wDay, month, year, 2)

		if day > 19 {
			return day - 7
		}

		return day

	case Third:
		return findDate(wDay, month, year, 2)
	case Fourth:
		return findDate(wDay, month, year, 3)
	case Last:
		return findDate(wDay, month+1, year, -1)
	default:
		return 0
	}
}

func findDate(wDay time.Weekday, month time.Month, year int, numWks int) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	for date.Weekday() != wDay {
		date = date.AddDate(0, 0, 1)
	}

	date = date.AddDate(0, 0, 7*numWks)
	return date.Day()
}
