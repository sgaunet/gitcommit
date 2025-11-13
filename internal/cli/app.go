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
	var lastCommitDate *time.Time
	if git.HasCommits() {
		lastDate, err := git.GetLastCommitDate()
		if err != nil {
			slog.Warn("Could not retrieve last commit date", "error", err)
		} else {
			lastCommitDate = &lastDate
			slog.Debug("Last commit date retrieved", "date", lastDate)
		}
	} else {
		slog.Debug("No previous commits in repository")
	}

	// Step 3: Parse the date
	parsedDate, err := datetime.ParseDate(request.InputDate)
	if err != nil {
		slog.Error("Date parsing failed", "error", err)
		return NewInvalidDateFormatError(request.InputDate)
	}
	request.ParsedDate = parsedDate

	slog.Debug("Date parsed successfully", "parsed", parsedDate)

	// Step 4: Validate chronology
	if lastCommitDate != nil {
		valid, errorType := datetime.ValidateChronology(parsedDate, lastCommitDate)
		if !valid {
			slog.Error("Chronology validation failed",
				"provided", parsedDate,
				"lastCommit", lastCommitDate,
				"errorType", errorType)

			equal := errorType == "chronology_violation_equal"
			return NewChronologyViolationError(
				datetime.FormatForGit(parsedDate),
				datetime.FormatForGit(*lastCommitDate),
				equal,
			)
		}
	}

	slog.Debug("Chronology validation passed")

	// Step 5: Format the date for Git
	gitFormattedDate := datetime.FormatForGit(parsedDate)
	request.GitFormattedDate = gitFormattedDate

	slog.Debug("Date formatted for Git", "formatted", gitFormattedDate)

	// Step 6: Execute the commit
	if err := git.ExecuteCommit(gitFormattedDate, request.CommitMessage); err != nil {
		slog.Error("Git commit failed", "error", err)
		return NewGitCommandError(err.Error())
	}

	// Step 7: Display success message
	fmt.Println(FormatSuccessMessage(gitFormattedDate))

	slog.Info("Commit created successfully")
	return nil
}
