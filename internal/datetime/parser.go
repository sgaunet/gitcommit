package datetime

import (
	"errors"
	"fmt"
	"time"
)

var (
	// ErrInvalidCalendarDate is returned when a date doesn't exist in the calendar.
	ErrInvalidCalendarDate = errors.New("invalid calendar date")
)

// ParseDate parses a date string in the format "YYYY-MM-DD HH:MM:SS" and returns a time.Time.
// The date is parsed in the local timezone.
//
// Example input: "2025-02-05 20:19:19"
//
// Returns an error if the date format is invalid or the date doesn't exist (e.g., Feb 30).
func ParseDate(dateStr string) (time.Time, error) {
	// Parse using the input date layout with local timezone
	parsedTime, err := time.ParseInLocation(InputDateLayout, dateStr, time.Local)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date %q: %w", dateStr, err)
	}

	// Additional validation: check if the parsed date matches the input
	// This catches cases like "2025-02-30" which might parse but be invalid
	formatted := parsedTime.Format(InputDateLayout)
	if formatted != dateStr {
		return time.Time{}, fmt.Errorf("%w: %q", ErrInvalidCalendarDate, dateStr)
	}

	return parsedTime, nil
}
