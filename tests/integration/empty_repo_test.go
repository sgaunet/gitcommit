package integration

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestEmptyRepository tests that commits work in an empty repository (no previous commits).
func TestEmptyRepository(t *testing.T) {
	// Create temporary directory
	tmpDir := t.TempDir()

	// Initialize empty git repository
	cmd := exec.Command("git", "init")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init repository: %v", err)
	}

	// Configure git user
	configEmailCmd := exec.Command("git", "config", "user.email", "test@test.com")
	configEmailCmd.Dir = tmpDir
	if err := configEmailCmd.Run(); err != nil {
		t.Fatalf("Failed to configure git user.email: %v", err)
	}

	configNameCmd := exec.Command("git", "config", "user.name", "Test User")
	configNameCmd.Dir = tmpDir
	if err := configNameCmd.Run(); err != nil {
		t.Fatalf("Failed to configure git user.name: %v", err)
	}

	// Create a test file
	testFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Stage the file
	cmd = exec.Command("git", "add", "test.txt")
	cmd.Dir = tmpDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to stage file: %v", err)
	}

	// Run gitcommit with any date (should work since repo is empty)
	binaryPath := getBinaryPath(t)
	cmd = exec.Command(binaryPath, "2020-01-01 00:00:00", "Initial commit")
	cmd.Dir = tmpDir
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Command should succeed in empty repository: %v\nOutput: %s", err, output)
	}

	outputStr := string(output)

	// Check for success message
	if !strings.Contains(outputStr, "✓") || !strings.Contains(outputStr, "Commit created") {
		t.Errorf("Expected success message, got: %s", outputStr)
	}

	// Verify commit was created with correct date
	cmd = exec.Command("git", "log", "-1", "--format=%ai")
	cmd.Dir = tmpDir
	output, err = cmd.Output()
	if err != nil {
		t.Fatalf("Failed to check commit date: %v", err)
	}

	dateStr := strings.TrimSpace(string(output))
	if !strings.Contains(dateStr, "2020-01-01 00:00:00") {
		t.Errorf("Expected commit date to contain '2020-01-01 00:00:00', got: %s", dateStr)
	}
}
