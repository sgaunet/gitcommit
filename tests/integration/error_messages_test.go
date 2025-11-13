package integration

import (
	"os/exec"
	"strings"
	"testing"
)

// TestErrorMessageFormats tests that error messages follow the consistent format.
func TestErrorMessageFormats(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		setupRepo       bool // whether to setup a git repository
		expectedStrings []string
	}{
		{
			name:      "invalid date format error",
			args:      []string{"invalid-date", "Test message"},
			setupRepo: true,
			expectedStrings: []string{
				"Error:",
				"Invalid date format",
				"Expected format: YYYY-MM-DD HH:MM:SS",
			},
		},
		{
			name:      "missing arguments error",
			args:      []string{},
			setupRepo: true,
			expectedStrings: []string{
				"Error:",
				"Missing required arguments",
				"Usage:",
			},
		},
		{
			name:      "not a repository error",
			args:      []string{"2025-02-05 20:19:19", "Test"},
			setupRepo: false, // Don't setup git repo
			expectedStrings: []string{
				"Error:",
				"Not a Git repository",
				"git init",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			binaryPath := getBinaryPath(t)
			var cmd *exec.Cmd

			if tt.setupRepo {
				// Setup temporary git repository
				repoDir := setupTestRepo(t)
				defer cleanup(t, repoDir)

				// Run in repository
				cmd = exec.Command(binaryPath, tt.args...)
				cmd.Dir = repoDir
			} else {
				// Run in temp directory (not a git repo)
				tmpDir := t.TempDir()
				cmd = exec.Command(binaryPath, tt.args...)
				cmd.Dir = tmpDir
			}

			output, err := cmd.CombinedOutput()
			outputStr := string(output)

			// Error expected
			if err == nil {
				t.Errorf("Expected error, but command succeeded")
			}

			// Check for expected strings in error message
			for _, expected := range tt.expectedStrings {
				if !strings.Contains(outputStr, expected) {
					t.Errorf("Error message should contain %q, got:\n%s", expected, outputStr)
				}
			}
		})
	}
}

// TestSuccessMessageFormat tests that success messages have the correct format.
func TestSuccessMessageFormat(t *testing.T) {
	// Setup temporary git repository
	repoDir := setupTestRepo(t)
	defer cleanup(t, repoDir)

	// Create a file to commit
	testFile := repoDir + "/test.txt"
	if err := exec.Command("bash", "-c", "echo 'test' > "+testFile).Run(); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Stage the file
	cmd := exec.Command("git", "add", "test.txt")
	cmd.Dir = repoDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to stage file: %v", err)
	}

	// Run gitcommit
	binaryPath := getBinaryPath(t)
	cmd = exec.Command(binaryPath, "2025-02-05 20:19:19", "Test commit")
	cmd.Dir = repoDir
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, output)
	}

	outputStr := string(output)

	// Check for checkmark and formatted date
	expectedStrings := []string{
		"✓",
		"Commit created",
		"Wed 5 Feb 2025",
		"20:19:19",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(outputStr, expected) {
			t.Errorf("Success message should contain %q, got:\n%s", expected, outputStr)
		}
	}
}

// cleanup removes a temporary directory.
func cleanup(t *testing.T, dir string) {
	t.Helper()
	// t.TempDir() will handle cleanup automatically
}
