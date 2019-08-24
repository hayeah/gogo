package gogo

import "time"

func Today() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}

func UTCToday() time.Time {
	return Today().In(time.UTC)
}
