package integration

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestInvalidDateFormatRejection tests that invalid date formats are properly rejected.
func TestInvalidDateFormatRejection(t *testing.T) {
	tests := []struct {
		name          string
		dateInput     string
		message       string
		expectedError string
	}{
		{
			name:          "wrong date separator",
			dateInput:     "2025/02/05 20:19:19",
			message:       "Test commit",
			expectedError: "Invalid date format",
		},
		{
			name:          "missing time component",
			dateInput:     "2025-02-05",
			message:       "Test commit",
			expectedError: "Invalid date format",
		},
		{
			name:          "invalid month",
			dateInput:     "2025-13-05 20:19:19",
			message:       "Test commit",
			expectedError: "Invalid date format",
		},
		{
			name:          "invalid day",
			dateInput:     "2025-02-30 20:19:19",
			message:       "Test commit",
			expectedError: "Invalid date format",
		},
		{
			name:          "completely invalid format",
			dateInput:     "not-a-date",
			message:       "Test commit",
			expectedError: "Invalid date format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup a temporary git repository
			repoDir := setupTestRepo(t)
			defer os.RemoveAll(repoDir)

			// Run the command
			binaryPath := getBinaryPath(t)
			cmd := exec.Command(binaryPath, tt.dateInput, tt.message)
			cmd.Dir = repoDir
			output, err := cmd.CombinedOutput()

			// Expect error
			if err == nil {
				t.Errorf("Expected error for invalid date format, but command succeeded")
			}

			// Check error message
			outputStr := string(output)
			if !strings.Contains(outputStr, tt.expectedError) {
				t.Errorf("Expected error message to contain %q, got: %s", tt.expectedError, outputStr)
			}
		})
	}
}

// TestChronologyViolationErrors tests that chronology violations are properly detected.
func TestChronologyViolationErrors(t *testing.T) {
	tests := []struct {
		name             string
		firstCommitDate  string
		secondCommitDate string
		expectedError    string
	}{
		{
			name:             "commit before last commit",
			firstCommitDate:  "2025-02-05 20:00:00",
			secondCommitDate: "2025-02-04 10:00:00",
			expectedError:    "Chronology violation",
		},
		{
			name:             "commit equal to last commit",
			firstCommitDate:  "2025-02-05 20:00:00",
			secondCommitDate: "2025-02-05 20:00:00",
			expectedError:    "Chronology violation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup a temporary git repository
			repoDir := setupTestRepo(t)
			defer os.RemoveAll(repoDir)

			binaryPath := getBinaryPath(t)

			// Create first commit with first date
			testFile := filepath.Join(repoDir, "test.txt")
			os.WriteFile(testFile, []byte("first commit"), 0644)
			gitAddCmd1 := exec.Command("git", "add", "test.txt")
			gitAddCmd1.Dir = repoDir
			gitAddCmd1.Run()

			cmd1 := exec.Command(binaryPath, tt.firstCommitDate, "First commit")
			cmd1.Dir = repoDir
			if err := cmd1.Run(); err != nil {
				t.Fatalf("First commit failed: %v", err)
			}

			// Try to create second commit with earlier/equal date
			os.WriteFile(testFile, []byte("second commit"), 0644)
			gitAddCmd := exec.Command("git", "add", "test.txt")
			gitAddCmd.Dir = repoDir
			gitAddCmd.Run()

			cmd2 := exec.Command(binaryPath, tt.secondCommitDate, "Second commit")
			cmd2.Dir = repoDir
			output, err := cmd2.CombinedOutput()

			// Expect error
			if err == nil {
				t.Errorf("Expected chronology violation error, but command succeeded")
			}

			// Check error message
			outputStr := string(output)
			if !strings.Contains(outputStr, tt.expectedError) {
				t.Errorf("Expected error message to contain %q, got: %s", tt.expectedError, outputStr)
			}
		})
	}
}

// TestNoRepositoryError tests that running outside a Git repository shows proper error.
func TestNoRepositoryError(t *testing.T) {
	// Create temporary directory (NOT a git repository)
	tmpDir := t.TempDir()

	// Run the command
	binaryPath := getBinaryPath(t)
	cmd := exec.Command(binaryPath, "2025-02-05 20:19:19", "Test commit")
	cmd.Dir = tmpDir
	output, err := cmd.CombinedOutput()

	// Expect error
	if err == nil {
		t.Errorf("Expected error for no repository, but command succeeded")
	}

	// Check error message
	outputStr := string(output)
	expectedErrors := []string{"Not a Git repository", "not a git repository"}

	foundExpected := false
	for _, expectedErr := range expectedErrors {
		if strings.Contains(outputStr, expectedErr) {
			foundExpected = true
			break
		}
	}

	if !foundExpected {
		t.Errorf("Expected error message to contain repository error, got: %s", outputStr)
	}
}
