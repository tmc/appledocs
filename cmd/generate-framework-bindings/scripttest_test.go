package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"rsc.io/script"
	"rsc.io/script/scripttest"
)

// TestScripts runs scripttest-based integration tests for the framework generator.
// These tests are skipped by default. Set RUN_SCRIPTTEST=1 to run them.
func TestScripts(t *testing.T) {
	if os.Getenv("RUN_SCRIPTTEST") == "" {
		t.Skip("Skipping scripttest integration tests. Set RUN_SCRIPTTEST=1 to run.")
	}

	// Build the binary for testing in a temporary location
	tmpDir := t.TempDir()
	tmpBin := filepath.Join(tmpDir, "generate-framework-bindings")

	// Build from current package directory
	buildCmd := exec.Command("go", "build", "-o", tmpBin, ".")
	if out, err := buildCmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build binary: %v\nOutput: %s", err, out)
	}

	// Create script engine with default commands
	engine := script.NewEngine()

	// Register our command as "generate-framework-bindings"
	engine.Cmds["generate-framework-bindings"] = script.Program(tmpBin, nil, 0)

	// Set up environment variables
	env := []string{
		"PATH=" + os.Getenv("PATH"),
		"HOME=" + os.Getenv("HOME"),
	}

	// Propagate coverage directory if set
	if gcd := os.Getenv("GOCOVERDIR"); gcd != "" {
		env = append(env, "GOCOVERDIR="+gcd)
	}

	// Run all test scripts in testdata directory
	scripttest.Test(t, context.Background(), engine, env, "testdata/*.txt")
}
