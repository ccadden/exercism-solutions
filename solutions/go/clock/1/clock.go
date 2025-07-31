package clock

import "fmt"

var MinutesInAnHour int = 60
var HoursInADay int = 24

// Define the Clock type here.
type Clock struct {
	Hour int
	Min int
}

func New(h, m int) Clock {
	hoursDelta := m / MinutesInAnHour

	m = m % 60

	if m < 0 {
		h -= 1
		m = 60 + m
	}

	h = (h + hoursDelta) % 24

	if h < 0 {
		h = 24 + h
	}

	return Clock{Hour: h, Min: m}
}

func (c Clock) Add(m int) Clock {
	newHour := c.Hour
	newMin := c.Min

	newMin += m

	if newMin >= MinutesInAnHour {
		newHour += newMin / MinutesInAnHour
		newMin = newMin % 60
	}

	return New(newHour, newMin)
}

func (c Clock) Subtract(m int) Clock {
	newHour := c.Hour
	newMin := c.Min

	newMin -= m

	if newMin < 0 {
		newHour += newMin / MinutesInAnHour
		newMin = newMin % 60
	}

	return New(newHour, newMin)

}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Min)
}
