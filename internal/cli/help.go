package cli

import "fmt"

// HelpText returns the formatted help text for the gitcommit tool.
func HelpText() string {
	return fmt.Sprintf(`gitcommit - Create Git commits with custom dates

Usage:
  gitcommit <date> <message>
  gitcommit --help
  gitcommit --version

Arguments:
  <date>     Date and time in format: YYYY-MM-DD HH:MM:SS
             Example: 2025-02-05 20:19:19

  <message>  Commit message (quote if contains spaces)

Flags:
  --help, -h       Show this help message
  --version, -v    Show version information

Description:
  gitcommit allows you to create Git commits with custom author and
  committer dates. This is useful for:
  - Backdating commits for historical documentation
  - Maintaining chronological commit history
  - Testing date-dependent Git workflows

Requirements:
  - Must be run inside a Git repository
  - Dates must be after the last commit (chronological order)
  - Changes must be staged before committing (git add)

Examples:
  # Create a commit with a specific date
  gitcommit "2025-02-05 20:19:19" "Add new feature"

  # Commit at midnight on New Year's Day
  gitcommit "2025-01-01 00:00:00" "Happy New Year!"

  # Show version
  gitcommit --version

Date Format:
  The date must be in the exact format: YYYY-MM-DD HH:MM:SS
  - Year: 4 digits (2025)
  - Month: 2 digits (01-12)
  - Day: 2 digits (01-31, valid for the month)
  - Hour: 2 digits (00-23)
  - Minute: 2 digits (00-59)
  - Second: 2 digits (00-59)

Error Messages:
  If you encounter errors, the tool provides helpful messages with:
  - Clear description of what went wrong
  - Specific details about the issue
  - Suggestions for how to fix it

For more information, visit:
  https://github.com/sgaunet/gitcommit
`)
}
