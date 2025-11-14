package cli

// FormatSuccessMessage formats a success message with a checkmark.
func FormatSuccessMessage(gitFormattedDate string) string {
	return "✓ Commit created with date: " + gitFormattedDate
}
