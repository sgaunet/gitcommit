package integration

import (
	"os/exec"
	"strings"
	"testing"
)

// TestHelpFlag tests that the --help flag displays proper help text.
func TestHelpFlag(t *testing.T) {
	tests := []struct {
		name     string
		flags    []string
		expected []string // strings that should be in the output
	}{
		{
			name:  "long form --help",
			flags: []string{"--help"},
			expected: []string{
				"gitcommit",
				"Usage:",
				"date",
				"message",
				"Examples:",
				"--help",
				"--version",
			},
		},
		{
			name:  "short form -h",
			flags: []string{"-h"},
			expected: []string{
				"gitcommit",
				"Usage:",
				"Examples:",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get path to gitcommit binary
			binaryPath := "../../build/gitcommit"

			// Run with help flag
			cmd := exec.Command(binaryPath, tt.flags...)
			output, err := cmd.CombinedOutput()

			// Help should exit with code 0
			if err != nil {
				t.Errorf("Help command should exit with 0, got error: %v", err)
			}

			outputStr := string(output)

			// Check for expected strings
			for _, expected := range tt.expected {
				if !strings.Contains(outputStr, expected) {
					t.Errorf("Help output should contain %q, got:\n%s", expected, outputStr)
				}
			}
		})
	}
}
