package utils

import "time"

// GetTodayf returns the current UTC date as a string in the format of "yyyy-mm-dd"
func GetTodayf() string {
	return GetToday().Format(time.DateOnly)
}

func GetToday() *time.Time {
	y, m, d := time.Now().UTC().Date()
	today := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	return &today
}

func DateFromFormat(val string) (*time.Time, error) {
	date, err := time.Parse(time.DateOnly, val)
	if err != nil {
		return nil, err
	}
	utc := date.UTC()
	return &utc, nil
}
