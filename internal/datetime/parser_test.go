package datetime

import (
	"testing"
	"time"
)

// TestParseDateValidFormats tests parsing of valid date formats.
func TestParseDateValidFormats(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
	}{
		{
			name:     "standard date and time",
			input:    "2025-02-05 20:19:19",
			expected: time.Date(2025, 2, 5, 20, 19, 19, 0, time.Local),
		},
		{
			name:     "midnight",
			input:    "2025-01-01 00:00:00",
			expected: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name:     "end of day",
			input:    "2025-12-31 23:59:59",
			expected: time.Date(2025, 12, 31, 23, 59, 59, 0, time.Local),
		},
		{
			name:     "leap year date",
			input:    "2024-02-29 12:00:00",
			expected: time.Date(2024, 2, 29, 12, 0, 0, 0, time.Local),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDate(tt.input)
			if err != nil {
				t.Fatalf("ParseDate(%q) unexpected error: %v", tt.input, err)
			}

			if !result.Equal(tt.expected) {
				t.Errorf("ParseDate(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestParseDateInvalidFormats tests parsing of invalid date formats.
func TestParseDateInvalidFormats(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"wrong separator", "2025/02/05 20:19:19"},
		{"missing time", "2025-02-05"},
		{"missing seconds", "2025-02-05 20:19"},
		{"invalid format", "not a date"},
		{"invalid month", "2025-13-01 00:00:00"},
		{"invalid day", "2025-02-30 00:00:00"},
		{"non-leap year feb 29", "2025-02-29 00:00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseDate(tt.input)
			if err == nil {
				t.Errorf("ParseDate(%q) expected error, got nil", tt.input)
			}
		})
	}
}
