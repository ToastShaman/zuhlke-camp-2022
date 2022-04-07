package clock

import "time"

type Clockwork = func() time.Time

var SystemClockUTC = Clockwork(func() time.Time {
	return time.Now().UTC()
})

var EPOCH = Clockwork(func() time.Time {
	return time.Unix(0, 0)
})

func Fixed(now time.Time) Clockwork {
	return Clockwork(func() time.Time {
		return now
	})
}
