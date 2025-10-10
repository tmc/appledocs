package occ2go

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tmc/appledocs"
)

// TestParseDocumentComprehensive tests that ParseDocument can handle all documents
// in the appledocs cache without panicking or crashing.
func TestParseDocumentComprehensive(t *testing.T) {
	// Get the default cache directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get home directory, skipping comprehensive test")
	}

	cacheDir := filepath.Join(homeDir, ".appledocs/cache/developer.apple.com/tutorials/data/documentation")
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		t.Skipf("Cache directory %s does not exist, skipping comprehensive test", cacheDir)
	}

	// Open the appledocs filesystem
	fsys, err := appledocs.Open(cacheDir)
	if err != nil {
		t.Fatalf("Failed to open appledocs filesystem: %v", err)
	}

	// Discover all available frameworks
	frameworks, err := appledocs.ListFrameworks(fsys)
	if err != nil {
		t.Fatalf("Failed to list frameworks: %v", err)
	}

	if len(frameworks) == 0 {
		t.Skip("No frameworks found in cache")
	}

	t.Logf("Testing %d frameworks", len(frameworks))

	stats := make(map[string]*FrameworkStats)

	for _, framework := range frameworks {
		t.Run(framework, func(t *testing.T) {
			t.Parallel()
			fwStats := &FrameworkStats{
				Framework: framework,
				Results:   make(map[string]int),
			}
			stats[framework] = fwStats

			// Process all symbols in the framework
			for path, doc := range appledocs.Symbols(fsys, framework) {
				fwStats.Total++

				// Create subtest for each symbol (only in verbose mode to avoid spam)
				symbolName := filepath.Base(path)
				testFunc := func(t *testing.T) {
					// The main test: ParseDocument should not panic
					fn, cls, proto, err := ParseDocument(doc)

					// Track results
					if err != nil {
						fwStats.Errors++
						// Categorize error types
						errStr := err.Error()
						switch {
						case strings.Contains(errStr, "unsupported symbol type"):
							fwStats.Results["unsupported"]++
						case strings.Contains(errStr, "no declaration found"):
							fwStats.Results["no_declaration"]++
						case strings.Contains(errStr, "failed to parse function name"):
							fwStats.Results["parse_error_function"]++
						case strings.Contains(errStr, "failed to parse class"):
							fwStats.Results["parse_error_class"]++
						case strings.Contains(errStr, "failed to parse protocol"):
							fwStats.Results["parse_error_protocol"]++
						case strings.Contains(errStr, "skipping"):
							fwStats.Results["skipped"]++
						default:
							fwStats.Results["other_error"]++
							// Log unexpected errors
							if testing.Verbose() {
								t.Logf("Unexpected error parsing %s: %v", path, err)
							}
						}
						return
					}

					// Count successes by type
					if fn != nil {
						fwStats.Functions++
						// Basic validation
						if fn.Name == "" {
							t.Errorf("Function from %s has empty name", path)
						}
					}
					if cls != nil {
						fwStats.Classes++
						// Basic validation
						if cls.Name == "" {
							t.Errorf("Class from %s has empty name", path)
						}
					}
					if proto != nil {
						fwStats.Protocols++
						// Basic validation
						if proto.Name == "" {
							t.Errorf("Protocol from %s has empty name", path)
						}
					}

					// Exactly one should be non-nil
					count := 0
					if fn != nil {
						count++
					}
					if cls != nil {
						count++
					}
					if proto != nil {
						count++
					}
					if count > 1 {
						t.Errorf("ParseDocument returned multiple non-nil values for %s", path)
					}
				}

				// Run as subtest in verbose mode, otherwise run directly
				if testing.Verbose() {
					t.Run(symbolName, testFunc)
				} else {
					testFunc(t)
				}
			}

			// Report statistics
			t.Logf("\n%s Statistics:", framework)
			t.Logf("  Total symbols: %d", fwStats.Total)
			t.Logf("  Functions: %d", fwStats.Functions)
			t.Logf("  Classes: %d", fwStats.Classes)
			t.Logf("  Protocols: %d", fwStats.Protocols)
			t.Logf("  Errors: %d (%.1f%%)", fwStats.Errors,
				float64(fwStats.Errors)/float64(fwStats.Total)*100)

			if len(fwStats.Results) > 0 {
				t.Logf("  Error breakdown:")
				for errType, count := range fwStats.Results {
					t.Logf("    %s: %d", errType, count)
				}
			}

			// Verify we got at least some successful parses
			if fwStats.Functions == 0 && fwStats.Classes == 0 && fwStats.Protocols == 0 {
				t.Errorf("Failed to parse any symbols successfully for %s", framework)
			}
		})
	}
}

// FrameworkStats tracks parsing statistics for a framework
type FrameworkStats struct {
	Framework string
	Total     int
	Functions int
	Classes   int
	Protocols int
	Errors    int
	Results   map[string]int // Error categorization
}

// TestParseDocumentSanity is a quick sanity check that doesn't require the cache
func TestParseDocumentSanity(t *testing.T) {
	// Test that ParseDocument handles nil document gracefully
	_, _, _, err := ParseDocument(nil)
	if err == nil {
		t.Error("Expected error for nil document, got nil")
	}
	if !strings.Contains(err.Error(), "nil") {
		t.Errorf("Expected error about nil document, got: %v", err)
	}
}

// TestParseDocumentErrorCategories verifies all error paths return proper errors
func TestParseDocumentErrorCategories(t *testing.T) {
	testCases := []struct {
		name        string
		doc         *appledocs.Document
		wantErr     bool
		errContains string
	}{
		{
			name: "no_declaration",
			doc: &appledocs.Document{
				Metadata: appledocs.Metadata{
					ExternalID: "c:@F@SomeFunction",
				},
				PrimaryContentSections: []appledocs.ContentSection{},
			},
			wantErr:     true,
			errContains: "no declaration found",
		},
		{
			name: "unsupported_symbol_type",
			doc: &appledocs.Document{
				Metadata: appledocs.Metadata{
					ExternalID: "c:@E@SomeEnum",
				},
				PrimaryContentSections: []appledocs.ContentSection{
					{
						Declarations: []appledocs.Declaration{
							{
								Tokens: []appledocs.Token{
									{Kind: "keyword", Text: "enum"},
								},
							},
						},
					},
				},
			},
			wantErr:     true,
			errContains: "unsupported symbol type",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, _, _, err := ParseDocument(tc.doc)
			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected error containing %q, got nil", tc.errContains)
				} else if !strings.Contains(err.Error(), tc.errContains) {
					t.Errorf("Expected error containing %q, got: %v", tc.errContains, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got: %v", err)
				}
			}
		})
	}
}
