// Package git provides Git repository operations and commit execution for gitcommit.
package git

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ExecuteCommit executes a git commit with the provided date and message.
// It sets the GIT_AUTHOR_DATE and GIT_COMMITTER_DATE environment variables
// to the provided gitFormattedDate before executing the commit.
//
// Parameters:
//   - gitFormattedDate: Date in Git format (e.g., "Wed 5 Feb 2025 20:19:19 CEST")
//   - message: The commit message
//
// Returns an error if the git commit command fails.
func ExecuteCommit(gitFormattedDate, message string) error {
	// Prepare git commit command
	cmd := exec.CommandContext(context.Background(), "git", "commit", "-m", message)

	// Set environment variables for commit dates
	cmd.Env = append(os.Environ(),
		"GIT_AUTHOR_DATE="+gitFormattedDate,
		"GIT_COMMITTER_DATE="+gitFormattedDate,
	)

	// Capture output for error reporting
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the commit
	if err := cmd.Run(); err != nil {
		// Try to get more details about the error
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			// Git command failed, return a more informative error
			return fmt.Errorf("git commit failed with exit code %d: %w", exitErr.ExitCode(), err)
		}
		return fmt.Errorf("git commit failed: %w", err)
	}

	return nil
}

// GetGitError extracts a user-friendly error message from git command output.
func GetGitError(output []byte) string {
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "On branch") {
			return line
		}
	}
	return "unknown git error"
}
