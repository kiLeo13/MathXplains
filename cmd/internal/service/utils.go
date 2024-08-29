package service

import (
	"MathXplains/internal/utils"
	"time"
)

func FormatEpoch(sec int64) string {
	return time.Unix(sec, 0).
		UTC().
		Format(time.RFC3339)
}

func NowUTC() int64 {
	return time.Now().
		UTC().
		Unix()
}

// isDateInPast does not check for time, only the date
func isDateInPast(now *time.Time) bool {
	return now.Before(*utils.GetToday())
}
