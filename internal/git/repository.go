package git

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	// GitExitCodeNoCommits is the exit code returned by git log when there are no commits.
	GitExitCodeNoCommits = 128
)

var (
	// ErrNoCommits is returned when the repository has no commits.
	ErrNoCommits = errors.New("no commits in repository")
	// ErrNotGitRepository is returned when the directory is not a Git repository.
	ErrNotGitRepository = errors.New("not a git repository")
	// ErrGitDirNotFound is returned when the .git directory cannot be found.
	ErrGitDirNotFound = errors.New(".git directory not found")
)

// Repository represents a Git repository and its properties.
type Repository struct {
	// Path is the root path of the Git repository.
	Path string

	// IsValid indicates whether this is a valid Git repository.
	IsValid bool
}

// IsGitRepository checks if the current directory is inside a Git repository.
// It searches up the directory tree for a .git directory.
func IsGitRepository() bool {
	// Try using git rev-parse to check if we're in a git repository
	cmd := exec.CommandContext(context.Background(), "git", "rev-parse", "--git-dir")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// GetLastCommitDate retrieves the date of the last commit in the repository.
// Returns an error if there are no commits or if not in a Git repository.
func GetLastCommitDate() (time.Time, error) {
	// Use git log to get the last commit date in RFC3339 format
	cmd := exec.CommandContext(context.Background(), "git", "log", "-1", "--format=%aI")
	output, err := cmd.Output()
	if err != nil {
		// Check if it's because there are no commits
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() == GitExitCodeNoCommits {
			return time.Time{}, ErrNoCommits
		}
		return time.Time{}, fmt.Errorf("failed to get last commit date: %w", err)
	}

	// Parse the date
	dateStr := strings.TrimSpace(string(output))
	if dateStr == "" {
		return time.Time{}, ErrNoCommits
	}

	commitDate, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse commit date %q: %w", dateStr, err)
	}

	return commitDate, nil
}

// NewRepository creates a Repository instance for the current directory.
func NewRepository() *Repository {
	isValid := IsGitRepository()

	repo := &Repository{
		IsValid: isValid,
	}

	if isValid {
		// Get the repository root path
		cmd := exec.CommandContext(context.Background(), "git", "rev-parse", "--show-toplevel")
		if output, err := cmd.Output(); err == nil {
			repo.Path = strings.TrimSpace(string(output))
		}
	}

	return repo
}

// GetRepositoryRoot returns the root path of the Git repository.
// Returns an error if not in a Git repository.
func GetRepositoryRoot() (string, error) {
	cmd := exec.CommandContext(context.Background(), "git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", ErrNotGitRepository
	}

	return strings.TrimSpace(string(output)), nil
}

// HasCommits checks if the repository has any commits.
func HasCommits() bool {
	cmd := exec.CommandContext(context.Background(), "git", "rev-parse", "HEAD")
	return cmd.Run() == nil
}

// ValidateRepository checks if the current directory is a valid Git repository.
// Returns a Repository instance and an error if validation fails.
func ValidateRepository() (*Repository, error) {
	if !IsGitRepository() {
		return nil, ErrNotGitRepository
	}

	repo := NewRepository()
	return repo, nil
}

// GetGitDirectory returns the path to the .git directory.
func GetGitDirectory() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	// Search up the directory tree for .git
	dir := currentDir
	for {
		gitDir := filepath.Join(dir, ".git")
		if info, err := os.Stat(gitDir); err == nil && info.IsDir() {
			return gitDir, nil
		}

		// Move to parent directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root without finding .git
			return "", ErrGitDirNotFound
		}
		dir = parent
	}
}
