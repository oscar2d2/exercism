package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	seconds := 1000000000
	t = t.Add(time.Duration(seconds) * time.Second)

	return t
}
