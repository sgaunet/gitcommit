package datetime

import (
	"strings"
	"testing"
	"time"
)

// TestFormatForGit tests that dates are formatted correctly for Git.
func TestFormatForGit(t *testing.T) {
	tests := []struct {
		name            string
		input           time.Time
		expectedPattern string // Pattern to match instead of exact string (due to timezone)
	}{
		{
			name:            "standard date",
			input:           time.Date(2025, 2, 5, 20, 19, 19, 0, time.Local),
			expectedPattern: "5 Feb 2025 20:19:19",
		},
		{
			name:            "midnight",
			input:           time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
			expectedPattern: "1 Jan 2025 00:00:00",
		},
		{
			name:            "end of year",
			input:           time.Date(2025, 12, 31, 23, 59, 59, 0, time.Local),
			expectedPattern: "31 Dec 2025 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatForGit(tt.input)

			// Check that the result contains the expected date/time components
			if !strings.Contains(result, tt.expectedPattern) {
				t.Errorf("FormatForGit(%v) = %q, want to contain %q", tt.input, result, tt.expectedPattern)
			}

			// Verify it follows Git's format: "Dow DD Mon YYYY HH:MM:SS TZ"
			parts := strings.Fields(result)
			if len(parts) != 6 {
				t.Errorf("FormatForGit(%v) = %q, expected 6 space-separated parts (Day DD Mon YYYY HH:MM:SS TZ)", tt.input, result)
			}
		})
	}
}

// TestFormatForGitTimezone tests that the timezone is included.
func TestFormatForGitTimezone(t *testing.T) {
	now := time.Now()
	result := FormatForGit(now)

	// Result should end with a timezone abbreviation (e.g., "CET", "EST", "PST")
	parts := strings.Fields(result)
	if len(parts) < 6 {
		t.Fatalf("FormatForGit result should have at least 6 parts, got: %s", result)
	}

	timezone := parts[len(parts)-1]
	if len(timezone) < 2 {
		t.Errorf("FormatForGit should include timezone abbreviation, got: %s", result)
	}
}
