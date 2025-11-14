package datetime

import (
	"time"
)

// ValidationResult represents the result of a date validation operation.
type ValidationResult struct {
	// Valid indicates whether the date is valid.
	Valid bool

	// ErrorType describes the type of validation error, if any.
	// Possible values: "chronology_violation", "chronology_violation_equal", "invalid_format", "invalid_value"
	ErrorType string

	// ErrorMessage provides a human-readable error message.
	ErrorMessage string

	// ProvidedDate is the date that was validated.
	ProvidedDate time.Time

	// LastCommitDate is the date of the last commit, if any.
	LastCommitDate *time.Time
}

// ValidateChronology validates that a commit date is chronologically after the last commit.
// Returns true if valid, and an error type string if invalid.
//
// Parameters:
//   - commitDate: The proposed commit date
//   - lastCommitDate: The date of the last commit (nil if no previous commits)
//
// Returns:
//   - bool: true if the date is valid, false otherwise
//   - string: empty string if valid, or one of:
//   - "chronology_violation": date is before last commit
//   - "chronology_violation_equal": date is equal to last commit
func ValidateChronology(commitDate time.Time, lastCommitDate *time.Time) (bool, string) {
	// If there's no last commit, any date is valid
	if lastCommitDate == nil {
		return true, ""
	}

	// Check if dates are equal
	if commitDate.Equal(*lastCommitDate) {
		return false, "chronology_violation_equal"
	}

	// Check if commit date is before last commit
	if commitDate.Before(*lastCommitDate) {
		return false, "chronology_violation"
	}

	// Date is after last commit - valid
	return true, ""
}

// ValidateDate performs comprehensive validation of a commit date.
// It checks format, value, and chronology against the last commit.
//
// Parameters:
//   - dateStr: The date string to validate
//   - lastCommitDate: The date of the last commit (nil if no previous commits)
//
// Returns a ValidationResult with details about the validation.
func ValidateDate(dateStr string, lastCommitDate *time.Time) *ValidationResult {
	result := &ValidationResult{
		LastCommitDate: lastCommitDate,
	}

	// Step 1: Parse the date
	parsedDate, err := ParseDate(dateStr)
	if err != nil {
		result.Valid = false
		result.ErrorType = "invalid_format"
		result.ErrorMessage = "Invalid date format. Expected: YYYY-MM-DD HH:MM:SS"
		return result
	}

	result.ProvidedDate = parsedDate

	// Step 2: Validate chronology
	valid, errorType := ValidateChronology(parsedDate, lastCommitDate)
	if !valid {
		result.Valid = false
		result.ErrorType = errorType

		if errorType == "chronology_violation_equal" {
			result.ErrorMessage = "Commit date must be after the last commit (dates are equal)"
		} else {
			result.ErrorMessage = "Commit date must be after the last commit"
		}
		return result
	}

	// Validation passed
	result.Valid = true
	return result
}

// ValidateDateValue checks if a parsed date is a valid calendar date.
// This is a secondary check after parsing to catch edge cases.
func ValidateDateValue(t time.Time, originalStr string) bool {
	// Re-format the time and compare with original
	// This catches cases like Feb 30 that might parse incorrectly
	reformatted := t.Format(InputDateLayout)
	return reformatted == originalStr
}
