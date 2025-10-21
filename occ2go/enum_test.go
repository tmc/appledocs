package occ2go

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/tmc/appledocs"
)

// TestParseEnumCase tests parsing of individual enum constants like NSTitledWindowMask
func TestParseEnumCase(t *testing.T) {
	// Load the NSTitledWindowMask document from cache
	cacheFile := filepath.Join(os.Getenv("HOME"), ".appledocs/cache/developer.apple.com/tutorials/data/documentation/AppKit/NSTitledWindowMask.json")

	data, err := os.ReadFile(cacheFile)
	if err != nil {
		t.Skipf("Skipping test: cache file not found: %v", err)
	}

	var doc appledocs.Document
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("Failed to unmarshal document: %v", err)
	}

	// Parse the enum case
	enumCase, err := ParseEnumCase(&doc)
	if err != nil {
		t.Fatalf("ParseEnumCase failed: %v", err)
	}

	// Verify the parsed data
	if enumCase.Name != "NSTitledWindowMask" {
		t.Errorf("Expected name 'NSTitledWindowMask', got '%s'", enumCase.Name)
	}

	if enumCase.Abstract == "" {
		t.Errorf("Expected non-empty abstract")
	}

	if enumCase.DocURL == "" {
		t.Errorf("Expected non-empty DocURL")
	}

	t.Logf("Successfully parsed enum case: %s", enumCase.Name)
	t.Logf("  Abstract: %s", enumCase.Abstract)
	t.Logf("  DocURL: %s", enumCase.DocURL)
}

// TestParseEnumCaseDeclaration tests the token parsing logic
func TestParseEnumCaseDeclaration(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []appledocs.Token
		expected string
		wantErr  bool
	}{
		{
			name: "Objective-C enum constant",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "static"},
				{Kind: "text", Text: " "},
				{Kind: "keyword", Text: "const"},
				{Kind: "text", Text: " "},
				{Kind: "typeIdentifier", Text: "NSWindowStyleMask"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "NSTitledWindowMask"},
				{Kind: "text", Text: ";"},
			},
			expected: "NSTitledWindowMask",
			wantErr:  false,
		},
		{
			name: "Swift enum case",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "static"},
				{Kind: "text", Text: " "},
				{Kind: "keyword", Text: "var"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "titled"},
				{Kind: "text", Text: ": "},
				{Kind: "typeIdentifier", Text: "NSWindow"},
				{Kind: "text", Text: "."},
				{Kind: "typeIdentifier", Text: "StyleMask"},
			},
			expected: "titled",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enumCase, err := ParseEnumCaseDeclaration(tt.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseEnumCaseDeclaration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && enumCase.Name != tt.expected {
				t.Errorf("ParseEnumCaseDeclaration() name = %v, want %v", enumCase.Name, tt.expected)
			}
		})
	}
}
