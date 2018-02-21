package clock

import (
	"fmt"
)

type Clock struct {
	hour   int
	minute int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func New(hour, minute int) Clock {
	for ; minute < 0; minute += 60 {
		hour -= 1
	}
	for ; hour < 0; hour += 24 {
	}
	return Clock{hour: (hour + minute/60) % 24, minute: minute % 60}
}

func (c Clock) Add(minutes int) Clock {
	num := c.hour*60 + c.minute
	num += minutes
	return New(num/60, num%60)
}

func (c Clock) Subtract(minute int) Clock {
	return c.Add(-1 * minute)
}
