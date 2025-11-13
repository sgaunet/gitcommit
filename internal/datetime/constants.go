package datetime

import "time"

const (
	// InputDateLayout is the format users provide for commit dates.
	// Example: "2025-02-05 20:19:19"
	InputDateLayout = "2006-01-02 15:04:05"

	// GitDateLayout is the format Git expects for GIT_AUTHOR_DATE and GIT_COMMITTER_DATE environment variables.
	// Example: "Wed 5 Feb 2025 20:19:19 CEST"
	GitDateLayout = "Mon 2 Jan 2006 15:04:05 MST"

	// GitLogDateLayout is the format git log returns when using --format=%cI (ISO 8601).
	// This is used to parse the last commit date from git log output.
	GitLogDateLayout = time.RFC3339
)
