package service

import (
	"time"
)

func FormatEpoch(sec int64) string {
	return time.Unix(sec, 0).
		UTC().
		Format(time.RFC3339)
}

func FormatDate(sec int64) string {
	return time.Unix(sec, 0).
		UTC().
		Format(time.DateOnly)
}

func ToEpoch(layout, seq string) (int64, error) {
	date, err := time.Parse(layout, seq)
	if err != nil {
		return 0, err
	}
	return date.UTC().Unix(), nil
}

func NowUTC() int64 {
	return time.Now().
		UTC().
		Unix()
}

// isDateInPast does not check for time, only the date
func isDateInPast(sec int64) bool {
	now := time.Now().UTC()
	dateNow := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return sec < dateNow.Unix()
}
