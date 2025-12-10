package lib

import "time"

// StringPtr is a helper function that returns a pointer to the given string.
func StringPtr(s string) *string {
	return &s
}

// TimePtr is a helper function that returns a pointer to the given time.Time.
func TimePtr(t time.Time) *time.Time {
	return &t
}
