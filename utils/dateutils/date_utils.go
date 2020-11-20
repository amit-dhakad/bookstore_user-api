package dateutils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

// GetNow return UTC time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString return current UTC time in string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat get database date format
func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
