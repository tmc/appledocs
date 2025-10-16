package main

import (
	"context"
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"rsc.io/script"
	"rsc.io/script/scripttest"
)

var runAspirational = flag.Bool("aspirational", false, "run aspirational (not-yet-passing) tests")

// TestScripts runs scripttest-based integration tests for the framework generator.
// By default, runs only baseline tests that should pass.
// Use -aspirational flag to also run aspirational tests (goals for future implementation).
func TestScripts(t *testing.T) {
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

	// Run baseline test scripts (should always pass)
	scripttest.Test(t, context.Background(), engine, env, "testdata/*.txt")
}

// TestScriptsAspirational runs aspirational tests that document future goals.
// These tests define the target API but are expected to fail until implemented.
// Skipped by default. Run with: go test -aspirational
func TestScriptsAspirational(t *testing.T) {
	if !*runAspirational {
		t.Skip("Skipping aspirational tests (expected to fail). Use -aspirational to run anyway.")
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

	// Run aspirational test scripts (may fail - that's expected!)
	scripttest.Test(t, context.Background(), engine, env, "testdata/aspirational/*.txt")
}
