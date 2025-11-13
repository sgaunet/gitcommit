package cli

import (
	"time"
)

// CommitRequest represents a user's request to create a commit with a specific date.
type CommitRequest struct {
	// InputDate is the raw date string provided by the user.
	InputDate string

	// CommitMessage is the commit message text.
	CommitMessage string

	// ParsedDate is the parsed and validated date with timezone.
	ParsedDate time.Time

	// GitFormattedDate is the date formatted for Git environment variables.
	GitFormattedDate string
}

// NewCommitRequest creates a new CommitRequest from user input.
func NewCommitRequest(date, message string) *CommitRequest {
	return &CommitRequest{
		InputDate:     date,
		CommitMessage: message,
	}
}
