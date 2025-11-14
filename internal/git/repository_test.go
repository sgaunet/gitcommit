package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestIsGitRepository tests the Git repository detection logic.
func TestIsGitRepository(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(t *testing.T) string // Returns path to test directory
		expected bool
	}{
		{
			name: "valid git repository",
			setup: func(t *testing.T) string {
				// Create temporary directory
				tmpDir := t.TempDir()

				// Initialize git repository using git init
				cmd := exec.Command("git", "init")
				cmd.Dir = tmpDir
				if err := cmd.Run(); err != nil {
					t.Fatalf("Failed to initialize git repository: %v", err)
				}

				return tmpDir
			},
			expected: true,
		},
		{
			name: "not a git repository",
			setup: func(t *testing.T) string {
				// Create temporary directory without .git
				return t.TempDir()
			},
			expected: false,
		},
		{
			name: "subdirectory of git repository",
			setup: func(t *testing.T) string {
				// Create temporary directory and initialize git
				tmpDir := t.TempDir()

				// Initialize git repository using git init
				cmd := exec.Command("git", "init")
				cmd.Dir = tmpDir
				if err := cmd.Run(); err != nil {
					t.Fatalf("Failed to initialize git repository: %v", err)
				}

				// Create subdirectory
				subDir := filepath.Join(tmpDir, "subdir", "nested")
				if err := os.MkdirAll(subDir, 0755); err != nil {
					t.Fatalf("Failed to create subdirectory: %v", err)
				}

				return subDir
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testPath := tt.setup(t)

			// Change to test directory
			originalDir, err := os.Getwd()
			if err != nil {
				t.Fatalf("Failed to get current directory: %v", err)
			}
			defer os.Chdir(originalDir)

			if err := os.Chdir(testPath); err != nil {
				t.Fatalf("Failed to change to test directory: %v", err)
			}

			// Test repository detection
			result := IsGitRepository()

			if result != tt.expected {
				t.Errorf("IsGitRepository() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestGetLastCommitDate tests retrieving the last commit date.
func TestGetLastCommitDate(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(t *testing.T) string // Returns path to test directory
		expectError bool
		errorMsg    string
	}{
		{
			name: "repository with commits",
			setup: func(t *testing.T) string {
				// This test will be implemented after Repository struct is created
				// For TDD, we expect this to fail initially
				return t.TempDir()
			},
			expectError: true,
			errorMsg:    "not implemented yet",
		},
		{
			name: "empty repository",
			setup: func(t *testing.T) string {
				return t.TempDir()
			},
			expectError: true,
			errorMsg:    "no commits",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testPath := tt.setup(t)

			// Change to test directory
			originalDir, err := os.Getwd()
			if err != nil {
				t.Fatalf("Failed to get current directory: %v", err)
			}
			defer os.Chdir(originalDir)

			if err := os.Chdir(testPath); err != nil {
				t.Fatalf("Failed to change to test directory: %v", err)
			}

			// Test last commit date retrieval
			_, err = GetLastCommitDate()

			if tt.expectError && err == nil {
				t.Errorf("Expected error containing %q, but got nil", tt.errorMsg)
			}

			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
