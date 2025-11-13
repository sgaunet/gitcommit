package datetime

import "time"

// FormatForGit formats a time.Time into the format Git expects for
// GIT_AUTHOR_DATE and GIT_COMMITTER_DATE environment variables.
//
// Format: "Dow DD Mon YYYY HH:MM:SS TZ"
// Example: "Wed 5 Feb 2025 20:19:19 CEST"
//
// The timezone used is the local timezone of the system.
func FormatForGit(t time.Time) string {
	return t.Format(GitDateLayout)
}
