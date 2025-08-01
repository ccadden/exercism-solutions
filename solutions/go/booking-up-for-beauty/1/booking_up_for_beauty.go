package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	res, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}

	return res
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}

	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}

	return (t.Hour() > 11  && t.Hour() < 18)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date)
	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
