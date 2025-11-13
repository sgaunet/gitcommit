// Package cli provides the command-line interface logic and orchestration for gitcommit.
package cli

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/sgaunet/gitcommit/internal/datetime"
	"github.com/sgaunet/gitcommit/internal/git"
)

// App represents the main application logic.
type App struct {
	config *Config
}

// NewApp creates a new App instance.
func NewApp(config *Config) *App {
	return &App{
		config: config,
	}
}

// Run executes the main application logic.
func (a *App) Run() error {
	// Create commit request
	request := NewCommitRequest(a.config.GetDate(), a.config.GetMessage())

	slog.Info("Processing commit request",
		"date", request.InputDate,
		"message", request.CommitMessage)

	// Step 1: Validate Git repository
	if !git.IsGitRepository() {
		slog.Error("Not in a Git repository")
		return NewNoRepositoryError()
	}

	slog.Debug("Git repository detected")

	// Step 2: Get last commit date (if any)
	lastCommitDate := a.getLastCommitDate()

	// Step 3: Parse and validate the date
	parsedDate, err := a.parseAndValidateDate(request.InputDate, lastCommitDate)
	if err != nil {
		return err
	}
	request.ParsedDate = parsedDate

	// Step 4: Format the date for Git
	gitFormattedDate := datetime.FormatForGit(parsedDate)
	request.GitFormattedDate = gitFormattedDate
	slog.Debug("Date formatted for Git", "formatted", gitFormattedDate)

	// Step 5: Execute the commit
	if err := git.ExecuteCommit(gitFormattedDate, request.CommitMessage); err != nil {
		slog.Error("Git commit failed", "error", err)
		return NewGitCommandError(err.Error())
	}

	// Step 6: Display success message
	fmt.Println(FormatSuccessMessage(gitFormattedDate))
	slog.Info("Commit created successfully")
	return nil
}

// getLastCommitDate retrieves the last commit date from the repository.
func (a *App) getLastCommitDate() *time.Time {
	if !git.HasCommits() {
		slog.Debug("No previous commits in repository")
		return nil
	}

	lastDate, err := git.GetLastCommitDate()
	if err != nil {
		slog.Warn("Could not retrieve last commit date", "error", err)
		return nil
	}

	slog.Debug("Last commit date retrieved", "date", lastDate)
	return &lastDate
}

// parseAndValidateDate parses and validates the commit date.
func (a *App) parseAndValidateDate(dateStr string, lastCommitDate *time.Time) (time.Time, error) {
	// Parse the date
	parsedDate, err := datetime.ParseDate(dateStr)
	if err != nil {
		slog.Error("Date parsing failed", "error", err)
		return time.Time{}, NewInvalidDateFormatError(dateStr)
	}

	slog.Debug("Date parsed successfully", "parsed", parsedDate)

	// Validate chronology
	if lastCommitDate != nil {
		valid, errorType := datetime.ValidateChronology(parsedDate, lastCommitDate)
		if !valid {
			slog.Error("Chronology validation failed",
				"provided", parsedDate,
				"lastCommit", lastCommitDate,
				"errorType", errorType)

			equal := errorType == "chronology_violation_equal"
			return time.Time{}, NewChronologyViolationError(
				datetime.FormatForGit(parsedDate),
				datetime.FormatForGit(*lastCommitDate),
				equal,
			)
		}
	}

	slog.Debug("Chronology validation passed")
	return parsedDate, nil
}
