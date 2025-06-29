package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestValidateCommandLineFlags(t *testing.T) {
	// Store original values
	originals := struct {
		outputDir     string
		cacheDir      string
		baseURL       string
		entryPoint    string
		concurrency   int
		rateLimit     float64
		timeout       time.Duration
		maxTime       time.Duration
		mode          string
		logLevel      string
		exportMetrics string
	}{
		*outputDir, *cacheDir, *baseURL, *entryPoint, *concurrency,
		*rateLimit, *timeout, *maxTime, *mode, *logLevel, *exportMetrics,
	}

	defer func() {
		*outputDir = originals.outputDir
		*cacheDir = originals.cacheDir
		*baseURL = originals.baseURL
		*entryPoint = originals.entryPoint
		*concurrency = originals.concurrency
		*rateLimit = originals.rateLimit
		*timeout = originals.timeout
		*maxTime = originals.maxTime
		*mode = originals.mode
		*logLevel = originals.logLevel
		*exportMetrics = originals.exportMetrics
	}()

	t.Run("valid configuration", func(t *testing.T) {
		*outputDir = "output"
		*cacheDir = ".cache"
		*baseURL = "https://developer.apple.com"
		*entryPoint = "/tutorials/data/documentation/technologies.json"
		*concurrency = 10
		*rateLimit = 10.0
		*timeout = 30 * time.Second
		*maxTime = time.Hour
		*mode = "crawl"
		*logLevel = "info"
		*exportMetrics = ""

		result := ValidateCommandLineFlags()
		if !result.Valid {
			t.Errorf("Expected valid configuration, got errors: %v", result.Errors)
		}
	})

	t.Run("invalid output directory", func(t *testing.T) {
		*outputDir = ""
		result := ValidateCommandLineFlags()
		if result.Valid {
			t.Error("Expected validation to fail for empty output directory")
		}
		found := false
		for _, err := range result.Errors {
			if err.Field == "output" {
				found = true
				break
			}
		}
		if !found {
			t.Error("Expected output directory validation error")
		}
	})

	t.Run("invalid base URL", func(t *testing.T) {
		*baseURL = "not-a-url"
		result := ValidateCommandLineFlags()
		if result.Valid {
			t.Error("Expected validation to fail for invalid base URL")
		}
	})

	t.Run("invalid concurrency", func(t *testing.T) {
		*concurrency = -1
		result := ValidateCommandLineFlags()
		if result.Valid {
			t.Error("Expected validation to fail for negative concurrency")
		}
	})

	t.Run("invalid mode", func(t *testing.T) {
		*mode = "invalid-mode"
		result := ValidateCommandLineFlags()
		if result.Valid {
			t.Error("Expected validation to fail for invalid mode")
		}
	})

	t.Run("high concurrency warning", func(t *testing.T) {
		*concurrency = 150
		result := ValidateCommandLineFlags()
		if len(result.Warnings) == 0 {
			t.Error("Expected warning for high concurrency")
		}
	})
}

func TestValidateAppleDocJSON(t *testing.T) {
	tests := []struct {
		name        string
		json        string
		expectValid bool
		expectError string
	}{
		{
			name: "valid Apple doc JSON",
			json: `{
				"metadata": {
					"title": "SwiftUI",
					"role": "framework",
					"platforms": [
						{"name": "iOS", "introducedAt": "13.0"}
					]
				},
				"abstract": [
					{"type": "text", "text": "A modern UI framework"}
				],
				"references": {
					"ref1": {
						"title": "View",
						"url": "doc://com.apple.documentation/documentation/SwiftUI/View"
					}
				}
			}`,
			expectValid: true,
		},
		{
			name:        "invalid JSON",
			json:        `{invalid json}`,
			expectValid: false,
			expectError: "json",
		},
		{
			name:        "empty JSON",
			json:        `{}`,
			expectValid: true, // Empty docs are technically valid
		},
		{
			name: "invalid platform version",
			json: `{
				"metadata": {
					"platforms": [
						{"name": "iOS", "introducedAt": "invalid-version"}
					]
				}
			}`,
			expectValid: true, // Should produce warnings, not errors
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateAppleDocJSON([]byte(tt.json), "test.json")
			
			if result.Valid != tt.expectValid {
				t.Errorf("Expected valid=%v, got valid=%v", tt.expectValid, result.Valid)
			}

			if tt.expectError != "" {
				found := false
				for _, err := range result.Errors {
					if err.Field == tt.expectError {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected error with field %q, got errors: %v", tt.expectError, result.Errors)
				}
			}
		})
	}
}

func TestValidateDataIntegrity(t *testing.T) {
	tests := []struct {
		name        string
		filePath    string
		data        []byte
		expectValid bool
	}{
		{
			name:        "valid JSON data",
			filePath:    "test.json",
			data:        []byte(`{"valid": "json"}`),
			expectValid: true,
		},
		{
			name:        "empty file",
			filePath:    "test.json",
			data:        []byte{},
			expectValid: false,
		},
		{
			name:        "binary data in JSON file",
			filePath:    "test.json",
			data:        []byte{0, 1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0},
			expectValid: false,
		},
		{
			name:        "valid text file",
			filePath:    "test.txt",
			data:        []byte("This is a text file"),
			expectValid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateDataIntegrity(tt.filePath, tt.data)
			
			if result.Valid != tt.expectValid {
				t.Errorf("Expected valid=%v, got valid=%v. Errors: %v", 
					tt.expectValid, result.Valid, result.Errors)
			}
		})
	}
}

func TestValidateCache(t *testing.T) {
	// Create a temporary cache directory
	tempDir := t.TempDir()
	
	// Create some test files
	validFile := filepath.Join(tempDir, "valid.json")
	if err := os.WriteFile(validFile, []byte(`{"test": "data"}`), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	emptyFile := filepath.Join(tempDir, "empty.json")
	if err := os.WriteFile(emptyFile, []byte{}, 0644); err != nil {
		t.Fatalf("Failed to create empty file: %v", err)
	}

	// Create an old file
	oldFile := filepath.Join(tempDir, "old.json")
	if err := os.WriteFile(oldFile, []byte(`{"old": "data"}`), 0644); err != nil {
		t.Fatalf("Failed to create old file: %v", err)
	}
	// Make it appear old
	oldTime := time.Now().Add(-40 * 24 * time.Hour)
	if err := os.Chtimes(oldFile, oldTime, oldTime); err != nil {
		t.Fatalf("Failed to set old file time: %v", err)
	}

	result := ValidateCache(tempDir)

	// Should find warnings for empty and old files
	if len(result.Warnings) == 0 {
		t.Error("Expected warnings for empty and old files")
	}

	// Check for empty file warning
	foundEmptyWarning := false
	foundOldWarning := false
	for _, warning := range result.Warnings {
		if warning.Path == emptyFile && warning.Field == "cache" {
			foundEmptyWarning = true
		}
		if warning.Path == oldFile && warning.Field == "cache" {
			foundOldWarning = true
		}
	}

	if !foundEmptyWarning {
		t.Error("Expected warning for empty cache file")
	}
	if !foundOldWarning {
		t.Error("Expected warning for old cache file")
	}
}

func TestValidationHelperFunctions(t *testing.T) {
	t.Run("isValidPath", func(t *testing.T) {
		tests := []struct {
			path  string
			valid bool
		}{
			{"valid/path", true},
			{"", false},
			{"path\x00with\x00nulls", false},
			{"path<with>invalid|chars", false},
			{"/absolute/path", true},
			{"./relative/path", true},
		}

		for _, tt := range tests {
			if got := isValidPath(tt.path); got != tt.valid {
				t.Errorf("isValidPath(%q) = %v, want %v", tt.path, got, tt.valid)
			}
		}
	})

	t.Run("containsBinaryData", func(t *testing.T) {
		tests := []struct {
			data   []byte
			binary bool
		}{
			{[]byte("normal text"), false},
			{[]byte{0, 0, 0, 0, 0, 0}, true},
			{[]byte("text with one null\x00"), false},
			{[]byte{0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6}, true},
		}

		for _, tt := range tests {
			if got := containsBinaryData(tt.data); got != tt.binary {
				t.Errorf("containsBinaryData() = %v, want %v", got, tt.binary)
			}
		}
	})

	t.Run("isValidVersionString", func(t *testing.T) {
		tests := []struct {
			version string
			valid   bool
		}{
			{"13.0", true},
			{"10.15.4", true},
			{"1.0", true},
			{"invalid", false},
			{"", false},
			{"13", false},
			{"13.0.0.1", false},
		}

		for _, tt := range tests {
			if got := isValidVersionString(tt.version); got != tt.valid {
				t.Errorf("isValidVersionString(%q) = %v, want %v", tt.version, got, tt.valid)
			}
		}
	})

	t.Run("isValidReferenceURL", func(t *testing.T) {
		tests := []struct {
			url   string
			valid bool
		}{
			{"doc://com.apple.documentation/documentation/SwiftUI", true},
			{"https://developer.apple.com/documentation/swiftui", true},
			{"/documentation/swiftui", true},
			{"invalid-url", false},
			{"", false},
		}

		for _, tt := range tests {
			if got := isValidReferenceURL(tt.url); got != tt.valid {
				t.Errorf("isValidReferenceURL(%q) = %v, want %v", tt.url, got, tt.valid)
			}
		}
	})
}

func TestValidationError(t *testing.T) {
	err := ValidationError{
		Field:   "test_field",
		Value:   "test_value",
		Message: "test message",
		Path:    "test/path",
	}

	expected := "validation error in test/path at test_field: test message (value: test_value)"
	if got := err.Error(); got != expected {
		t.Errorf("ValidationError.Error() = %q, want %q", got, expected)
	}

	// Test without path
	err.Path = ""
	expected = "validation error at test_field: test message (value: test_value)"
	if got := err.Error(); got != expected {
		t.Errorf("ValidationError.Error() without path = %q, want %q", got, expected)
	}
}

func TestValidationResult(t *testing.T) {
	result := ValidationResult{Valid: true}

	// Test adding error
	result.AddError("field1", "error message", "value1", "path1")
	if result.Valid {
		t.Error("Expected Valid to be false after adding error")
	}
	if len(result.Errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(result.Errors))
	}

	// Test adding warning
	result.AddWarning("field2", "warning message", "value2", "path2")
	if len(result.Warnings) != 1 {
		t.Errorf("Expected 1 warning, got %d", len(result.Warnings))
	}
}

// Integration test with real Apple doc structure
func TestValidateRealAppleDocStructure(t *testing.T) {
	// Create a realistic Apple documentation JSON
	doc := DocJSONData{
		Metadata: Metadata{
			Title: "SwiftUI",
			Role:  "framework",
			Platforms: []Platform{
				{Name: "iOS", IntroducedAt: "13.0"},
				{Name: "macOS", IntroducedAt: "10.15"},
			},
		},
		Abstract: []TextContent{
			{Type: "text", Text: "A modern UI framework"},
		},
		References: map[string]Reference{
			"ref1": {
				Title: "View",
				URL:   "doc://com.apple.documentation/documentation/SwiftUI/View",
			},
		},
		PrimaryContentSections: []ContentSection{
			{Kind: "content"},
			{Kind: "declarations"},
		},
		TopicSections: []TopicSection{
			{Title: "Essentials", Identifiers: []string{"ref1"}},
		},
	}

	jsonData, err := json.Marshal(doc)
	if err != nil {
		t.Fatalf("Failed to marshal test doc: %v", err)
	}

	result := ValidateAppleDocJSON(jsonData, "test.json")
	if !result.Valid {
		t.Errorf("Valid Apple doc failed validation: %v", result.Errors)
	}

	// Should have no errors for a well-formed document
	if len(result.Errors) > 0 {
		t.Errorf("Expected no errors for valid doc, got: %v", result.Errors)
	}
}