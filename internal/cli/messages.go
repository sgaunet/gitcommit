package cli

import "fmt"

// FormatSuccessMessage formats a success message with a checkmark.
func FormatSuccessMessage(gitFormattedDate string) string {
	return fmt.Sprintf("✓ Commit created with date: %s", gitFormattedDate)
}
