package cli

import "fmt"

// ErrorType represents different categories of errors that can occur.
type ErrorType int

const (
	// ExitSuccess indicates successful execution.
	ExitSuccess = 0
	// ExitError indicates an error occurred.
	ExitError = 1
)

// UserError represents an error with a user-friendly message.
type UserError struct {
	Type    string
	Message string
	Details string
	Hint    string
}

// NewInvalidDateFormatError creates an error for invalid date format.
func NewInvalidDateFormatError(provided string) *UserError {
	return &UserError{
		Type:    "InvalidDateFormat",
		Message: "Invalid date format",
		Details: fmt.Sprintf(
			"Expected format: YYYY-MM-DD HH:MM:SS\nExample:         2025-02-05 20:19:19\n\n"+
				"You provided:    %s",
			provided,
		),
	}
}

// NewInvalidDateValueError creates an error for invalid calendar dates.
func NewInvalidDateValueError(date string, reason string) *UserError {
	return &UserError{
		Type:    "InvalidDateValue",
		Message: "Invalid calendar date",
		Details: fmt.Sprintf("The date \"%s\" does not exist.\n%s", date, reason),
	}
}

// NewNoRepositoryError creates an error when not in a Git repository.
func NewNoRepositoryError() *UserError {
	return &UserError{
		Type:    "NoRepository",
		Message: "Not a Git repository",
		Details: "The current directory is not inside a Git repository.",
		Hint: "To fix this:\n  - Navigate to a Git repository: cd /path/to/repo\n" +
			"  - Or initialize a new repository: git init",
	}
}

// NewChronologyViolationError creates an error when date is before last commit.
func NewChronologyViolationError(providedDate, lastCommitDate string, equal bool) *UserError {
	details := fmt.Sprintf("Your date:        %s\nLast commit date: %s", providedDate, lastCommitDate)
	hint := "Commits must be dated after the last commit to maintain chronological order."

	if equal {
		hint = "Commits must be dated AFTER the last commit (not equal).\n" +
			"Suggestion: Try adding 1 second to your date."
	}

	return &UserError{
		Type:    "ChronologyViolation",
		Message: "Chronology violation",
		Details: details,
		Hint:    hint,
	}
}

// NewGitCommandError creates an error when a Git command fails.
func NewGitCommandError(gitError string) *UserError {
	return &UserError{
		Type:    "GitCommandError",
		Message: "Git commit failed",
		Details: "Git error: " + gitError,
		Hint: "Possible solutions:\n  - Stage your changes: git add <files>\n" +
			"  - Check if changes exist: git status",
	}
}

// NewMissingArgumentsError creates an error when arguments are missing.
func NewMissingArgumentsError(expected, received int) *UserError {
	return &UserError{
		Type:    "MissingArguments",
		Message: "Missing required arguments",
		Details: fmt.Sprintf(
			"Usage: gitcommit <date> <message>\n\nExpected: %d arguments\nReceived: %d argument(s)",
			expected,
			received,
		),
		Hint: "Examples:\n" +
			"  gitcommit \"2025-02-05 20:19:19\" \"Add new feature\"\n" +
			"  gitcommit \"2025-12-31 23:59:59\" \"End of year commit\"\n\n" +
			"Run 'gitcommit --help' for more information.",
	}
}

// Error implements the error interface.
func (e *UserError) Error() string {
	if e.Details != "" && e.Hint != "" {
		return "Error: " + e.Message + "\n\n" + e.Details + "\n\n" + e.Hint
	}
	if e.Details != "" {
		return "Error: " + e.Message + "\n\n" + e.Details
	}
	if e.Hint != "" {
		return "Error: " + e.Message + "\n\n" + e.Hint
	}
	return "Error: " + e.Message
}
