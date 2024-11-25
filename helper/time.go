package helper

import "time"

const (
	CacheDuration time.Duration = 15 * time.Minute

	NotExistCacheDuration time.Duration = 30 * time.Second
)

func GetPrimitiveTime() time.Time {
	return time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func IsActionExpired(period time.Time, duration time.Duration) bool {
	return time.Now().After(period.Add(duration))
}
