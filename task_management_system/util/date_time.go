package util

import "time"

const (
	// DateTimeFormat const
	DateTimeFormat = "2006-01-02T15:04:05:000Z"
	// DateLayoutISO const
	DateLayoutISO = "2006-01-02"
)

// TimeFormat function
func TimeFormat(dateTime time.Time) (time.Time, error) {
	dateString := dateTime.Format(DateTimeFormat)
	dateFormat, err := time.Parse(DateTimeFormat, dateString)
	return dateFormat, err
}

// CurrentTime to get current time in UTC
func CurrentTime() time.Time {
	currentTime, _ := TimeFormat(time.Now())
	return currentTime
}

//StringToTime function
func StringToTime(dateString string) (time.Time, error) {
	dateTime, err := time.Parse(DateTimeFormat, dateString)
	return dateTime, err
}
