package gigasecond

import "time"

const Giga = 1000000000

// Returns the given time after 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * Giga)
}
