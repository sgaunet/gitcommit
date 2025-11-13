package datetime

import (
	"testing"
	"time"
)

// TestValidateChronology tests the date chronology validation logic.
func TestValidateChronology(t *testing.T) {
	tests := []struct {
		name            string
		commitDate      time.Time
		lastCommitDate  *time.Time // nil if no previous commits
		expectValid     bool
		expectErrorType string
	}{
		{
			name:       "commit date after last commit",
			commitDate: parseTestDate("2025-02-05 20:19:19"),
			lastCommitDate: &[]time.Time{
				parseTestDate("2025-02-04 15:00:00"),
			}[0],
			expectValid:     true,
			expectErrorType: "",
		},
		{
			name:       "commit date before last commit",
			commitDate: parseTestDate("2025-02-03 10:00:00"),
			lastCommitDate: &[]time.Time{
				parseTestDate("2025-02-04 15:00:00"),
			}[0],
			expectValid:     false,
			expectErrorType: "chronology_violation",
		},
		{
			name:       "commit date equal to last commit",
			commitDate: parseTestDate("2025-02-04 15:00:00"),
			lastCommitDate: &[]time.Time{
				parseTestDate("2025-02-04 15:00:00"),
			}[0],
			expectValid:     false,
			expectErrorType: "chronology_violation_equal",
		},
		{
			name:            "first commit in repository",
			commitDate:      parseTestDate("2025-02-05 20:19:19"),
			lastCommitDate:  nil,
			expectValid:     true,
			expectErrorType: "",
		},
		{
			name:       "commit date one second after last",
			commitDate: parseTestDate("2025-02-04 15:00:01"),
			lastCommitDate: &[]time.Time{
				parseTestDate("2025-02-04 15:00:00"),
			}[0],
			expectValid:     true,
			expectErrorType: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This will fail until ValidateChronology is implemented
			valid, errorType := ValidateChronology(tt.commitDate, tt.lastCommitDate)

			if valid != tt.expectValid {
				t.Errorf("ValidateChronology() valid = %v, expected %v", valid, tt.expectValid)
			}

			if errorType != tt.expectErrorType {
				t.Errorf("ValidateChronology() errorType = %q, expected %q", errorType, tt.expectErrorType)
			}
		})
	}
}

// parseTestDate is a helper function to parse test dates.
func parseTestDate(dateStr string) time.Time {
	t, err := time.ParseInLocation(InputDateLayout, dateStr, time.Local)
	if err != nil {
		panic("Invalid test date: " + dateStr)
	}
	return t
}
