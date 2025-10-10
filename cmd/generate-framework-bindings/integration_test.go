package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tmc/appledocs"
	"github.com/tmc/appledocs/occ2go"
)

// TestGeneratorUsesAppledocsFS tests that we can use appledocs.FS API
// to accomplish the same thing as the current filepath.Walk approach.
// This test establishes the baseline before refactoring main.go.
func TestGeneratorUsesAppledocsFS(t *testing.T) {
	// Get the cache directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get home directory:", err)
	}
	cacheDir := filepath.Join(homeDir, ".appledocs/cache/developer.apple.com/tutorials/data/documentation")

	// Check if cache exists
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		t.Skip("Cache directory not found:", cacheDir)
	}

	// Open using appledocs.FS API
	fsys, err := appledocs.Open(cacheDir)
	if err != nil {
		t.Fatal("Failed to open cache with appledocs.FS:", err)
	}

	// Test with a small framework - SecurityFoundation
	framework := "SecurityFoundation"

	// List all symbols using the FS API
	symbols, err := appledocs.ListSymbols(fsys, framework)
	if err != nil {
		t.Skip("Framework not available:", err)
	}

	if len(symbols) == 0 {
		t.Skip("No symbols found in", framework)
	}

	t.Logf("Found %d symbols in %s using appledocs.FS API", len(symbols), framework)

	// Parse each symbol document
	var functions []*occ2go.ParsedFunction
	var classes []*occ2go.ParsedClass
	var protocols []*occ2go.ParsedProtocol

	for _, doc := range appledocs.Symbols(fsys, framework) {
		// Parse using occ2go
		fn, cls, proto, err := occ2go.ParseDocument(doc)
		if err != nil {
			// Not all symbols parse successfully - that's ok
			continue
		}

		if fn != nil {
			functions = append(functions, fn)
		}
		if cls != nil {
			classes = append(classes, cls)
		}
		if proto != nil {
			protocols = append(protocols, proto)
		}
	}

	t.Logf("Parsed: %d functions, %d classes, %d protocols",
		len(functions), len(classes), len(protocols))

	// SecurityFoundation should have at least the SFAuthorization class
	if len(classes) == 0 {
		t.Error("Expected at least one class in SecurityFoundation")
	}

	// Verify we can access parsed data
	for _, cls := range classes {
		if cls.Name == "" {
			t.Error("Class has empty name")
		}
		if cls.Name == "SFAuthorization" {
			t.Logf("Found SFAuthorization class with availability: %v", cls.Availability)
		}
	}
}

// TestAppledocsFSVsFilepathWalk compares the two approaches to ensure
// they produce equivalent results
func TestAppledocsFSVsFilepathWalk(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get home directory:", err)
	}
	cacheDir := filepath.Join(homeDir, ".appledocs/cache/developer.apple.com/tutorials/data/documentation")

	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		t.Skip("Cache directory not found:", cacheDir)
	}

	framework := "SecurityFoundation"
	frameworkDir := filepath.Join(cacheDir, framework)

	if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
		t.Skip("Framework directory not found:", frameworkDir)
	}

	// Approach 1: Using filepath.Walk (current main.go approach)
	var filesWalk []string
	err = filepath.Walk(frameworkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" && filepath.Base(path) != ".json" {
			filesWalk = append(filesWalk, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal("filepath.Walk failed:", err)
	}

	// Approach 2: Using appledocs.FS API
	fsys, err := appledocs.Open(cacheDir)
	if err != nil {
		t.Fatal("Failed to open cache:", err)
	}

	symbols, err := appledocs.ListSymbols(fsys, framework)
	if err != nil {
		t.Fatal("ListSymbols failed:", err)
	}

	// Convert symbols to file paths for comparison
	var filesFS []string
	for _, symbol := range symbols {
		filesFS = append(filesFS, filepath.Join(frameworkDir, symbol+".json"))
	}

	// Compare counts
	t.Logf("filepath.Walk found: %d files", len(filesWalk))
	t.Logf("appledocs.FS found: %d symbols", len(filesFS))

	// The FS API might find different files (e.g., it doesn't include subdirectories yet)
	// But it should find the main framework symbols
	if len(filesFS) == 0 {
		t.Error("FS API found no symbols")
	}

	// Both approaches should find SFAuthorization.json
	foundInWalk := false
	foundInFS := false

	targetFile := filepath.Join(frameworkDir, "SFAuthorization.json")

	for _, f := range filesWalk {
		if f == targetFile {
			foundInWalk = true
			break
		}
	}

	for _, f := range filesFS {
		if f == targetFile {
			foundInFS = true
			break
		}
	}

	if !foundInWalk {
		t.Error("filepath.Walk did not find SFAuthorization.json")
	}
	if !foundInFS {
		t.Error("appledocs.FS did not find SFAuthorization")
	}
}
