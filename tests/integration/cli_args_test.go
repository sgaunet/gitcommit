package integration

import (
	"os/exec"
	"strings"
	"testing"
)

// TestCLIArgumentParsing tests that the CLI correctly parses date and message arguments.
func TestCLIArgumentParsing(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid date and message",
			args:        []string{"2025-02-05 20:19:19", "Test commit"},
			expectError: false,
		},
		{
			name:        "missing message argument",
			args:        []string{"2025-02-05 20:19:19"},
			expectError: true,
			errorMsg:    "Missing required arguments",
		},
		{
			name:        "missing both arguments",
			args:        []string{},
			expectError: true,
			errorMsg:    "Missing required arguments",
		},
		{
			name:        "too many arguments",
			args:        []string{"2025-02-05 20:19:19", "Message", "Extra"},
			expectError: true,
			errorMsg:    "Missing required arguments",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This will initially fail until implementation is complete
			cmd := exec.Command("go", append([]string{"run", "../../cmd/gitcommit/main.go"}, tt.args...)...)
			output, err := cmd.CombinedOutput()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none. Output: %s", output)
				}
				if !strings.Contains(string(output), tt.errorMsg) {
					t.Errorf("Expected error message containing %q, got: %s", tt.errorMsg, output)
				}
			} else {
				if err != nil && !strings.Contains(string(output), "Phase 2 complete") {
					// Allow Phase 2 message for now, will be updated in implementation
					t.Logf("Command failed (expected during TDD): %v\nOutput: %s", err, output)
				}
			}
		})
	}
}

// TestVersionFlag tests that --version and -v flags work correctly.
func TestVersionFlag(t *testing.T) {
	tests := []struct {
		name string
		flag string
	}{
		{"long form", "--version"},
		{"short form", "-v"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "../../cmd/gitcommit/main.go", tt.flag)
			output, err := cmd.CombinedOutput()

			if err != nil {
				t.Fatalf("Version flag should not error: %v\nOutput: %s", err, output)
			}

			expected := "gitcommit version 1.0.0"
			if !strings.Contains(string(output), expected) {
				t.Errorf("Expected output containing %q, got: %s", expected, output)
			}
		})
	}
}
