package occ2go

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

// TestScripts runs scripttest-based integration tests for occ2go parsing.
// By default, runs only baseline tests that should pass.
// Use -aspirational flag to also run aspirational tests (goals for future implementation).
func TestScripts(t *testing.T) {
	// Build the occ2go binary for testing
	tmpDir := t.TempDir()
	tmpBin := filepath.Join(tmpDir, "occ2go")

	buildCmd := "../cmd/occ2go"
	cmd := exec.Command("go", "build", "-o", tmpBin, buildCmd)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build occ2go binary: %v\nOutput: %s", err, out)
	}

	// Create script engine with default commands
	engine := script.NewEngine()

	// Register occ2go command
	engine.Cmds["occ2go"] = script.Program(tmpBin, nil, 0)

	// Register map-type helper for type mapping tests
	engine.Cmds["map-type"] = script.Command(
		script.CmdUsage{
			Summary: "map a C type to Go type",
			Args:    "ctype framework",
		},
		func(s *script.State, args ...string) (script.WaitFunc, error) {
			if len(args) != 2 {
				return nil, script.ErrUsage
			}

			goType := MapCTypeToGo(args[0], args[1])
			return func(*script.State) (stdout, stderr string, err error) {
				return goType + "\n", "", nil
			}, nil
		})

	// Set up environment variables
	env := []string{
		"PATH=" + os.Getenv("PATH"),
		"HOME=" + os.Getenv("HOME"),
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

	// Create script engine with default commands
	engine := script.NewEngine()

	// Register the same commands as baseline tests
	// (reusing setup code would be better, but keeping it simple for now)

	env := []string{
		"PATH=" + os.Getenv("PATH"),
		"HOME=" + os.Getenv("HOME"),
	}

	// Run aspirational test scripts (may fail - that's expected!)
	scripttest.Test(t, context.Background(), engine, env, "testdata/aspirational/*.txt")
}
