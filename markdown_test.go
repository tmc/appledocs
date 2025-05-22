package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestChangeExtension(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		newExt   string
		expected string
	}{
		{
			name:     "change json to md",
			path:     "test.json",
			newExt:   ".md",
			expected: "test.md",
		},
		{
			name:     "change with path",
			path:     "path/to/file.json",
			newExt:   ".md",
			expected: "path/to/file.md",
		},
		{
			name:     "no extension",
			path:     "test",
			newExt:   ".md",
			expected: "test.md",
		},
		{
			name:     "different extension",
			path:     "test.html",
			newExt:   ".md",
			expected: "test.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := changeExtension(tt.path, tt.newExt)
			if result != tt.expected {
				t.Errorf("changeExtension() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatURL(t *testing.T) {
	tests := []struct {
		name     string
		urlStr   string
		expected string
	}{
		{
			name:     "doc URL",
			urlStr:   "doc://com.apple.documentation/documentation/SwiftUI",
			expected: "/documentation/documentation/SwiftUI.md",
		},
		{
			name:     "doc URL with nested path",
			urlStr:   "doc://com.apple.documentation/documentation/UIKit/UIView",
			expected: "/documentation/documentation/UIKit/UIView.md",
		},
		{
			name:     "regular URL",
			urlStr:   "https://developer.apple.com/documentation/SwiftUI",
			expected: "https://developer.apple.com/documentation/SwiftUI",
		},
		{
			name:     "empty URL",
			urlStr:   "",
			expected: ".md",
		},
		{
			name:     "doc URL without path",
			urlStr:   "doc://com.apple.documentation",
			expected: "#",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatURL(tt.urlStr)
			if result != tt.expected {
				t.Errorf("formatURL() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestScanDirectoryForJSON(t *testing.T) {
	// Create a temporary directory with test files
	tempDir := t.TempDir()

	// Create test JSON files
	testFiles := []string{
		"test1.json",
		"nested/test2.json",
		"nested/deep/test3.json",
		"notjson.txt",
	}

	for _, file := range testFiles {
		fullPath := filepath.Join(tempDir, file)
		dir := filepath.Dir(fullPath)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}

		content := ""
		if strings.HasSuffix(file, ".json") {
			content = `{"test": "data"}`
		} else {
			content = "not json content"
		}

		err = os.WriteFile(fullPath, []byte(content), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", file, err)
		}
	}

	// Test scanning
	jsonFiles, err := scanDirectoryForJSON(tempDir)
	if err != nil {
		t.Errorf("scanDirectoryForJSON() error = %v", err)
	}

	// Should find 3 JSON files
	expectedCount := 3
	if len(jsonFiles) != expectedCount {
		t.Errorf("scanDirectoryForJSON() found %d files, want %d", len(jsonFiles), expectedCount)
	}

	// Verify all found files are JSON files
	for _, file := range jsonFiles {
		if !strings.HasSuffix(file, ".json") {
			t.Errorf("scanDirectoryForJSON() returned non-JSON file: %s", file)
		}
	}
}

func TestConvertJSONToMarkdown(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()

	// Create a test JSON file
	testDoc := DocJSONData{
		Metadata: Metadata{
			Title: "Test Document",
			Role:  "article",
		},
		Abstract: []TextContent{
			{
				Type: "paragraph",
				Text: "This is a test document.",
			},
		},
	}

	testDocBytes, err := json.Marshal(testDoc)
	if err != nil {
		t.Fatalf("Failed to marshal test document: %v", err)
	}

	jsonPath := filepath.Join(tempDir, "test.json")
	err = os.WriteFile(jsonPath, testDocBytes, 0644)
	if err != nil {
		t.Fatalf("Failed to create test JSON file: %v", err)
	}

	// Test conversion
	mdPath := filepath.Join(tempDir, "test.md")
	err = convertJSONToMarkdown(jsonPath, mdPath)
	if err != nil {
		t.Errorf("convertJSONToMarkdown() error = %v", err)
	}

	// Verify markdown file was created
	if _, err := os.Stat(mdPath); os.IsNotExist(err) {
		t.Errorf("convertJSONToMarkdown() did not create markdown file")
	}

	// Read and verify content
	content, err := os.ReadFile(mdPath)
	if err != nil {
		t.Errorf("Failed to read markdown file: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "# Test Document") {
		t.Errorf("Markdown content should contain title")
	}
}

func TestConvertJSONToMarkdownWithInvalidJSON(t *testing.T) {
	tempDir := t.TempDir()

	// Create invalid JSON file
	jsonPath := filepath.Join(tempDir, "invalid.json")
	err := os.WriteFile(jsonPath, []byte(`{invalid json`), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid JSON file: %v", err)
	}

	// Test conversion should handle error gracefully
	mdPath := filepath.Join(tempDir, "invalid.md")
	err = convertJSONToMarkdown(jsonPath, mdPath)
	if err == nil {
		t.Errorf("convertJSONToMarkdown() should return error for invalid JSON")
	}
}

func TestGenerateMarkdown(t *testing.T) {
	// Create temporary directories
	tempInputDir := t.TempDir()
	tempOutputDir := t.TempDir()

	// Create test JSON files
	testDocs := []struct {
		path    string
		content DocJSONData
	}{
		{
			path: "tutorials/data/documentation/SwiftUI.json",
			content: DocJSONData{
				Metadata: Metadata{
					Title: "SwiftUI",
					Role:  "framework",
				},
				Abstract: []TextContent{
					{
						Type: "paragraph",
						Text: "A declarative UI framework.",
					},
				},
			},
		},
		{
			path: "tutorials/data/documentation/UIKit.json",
			content: DocJSONData{
				Metadata: Metadata{
					Title: "UIKit",
					Role:  "framework",
				},
				Abstract: []TextContent{
					{
						Type: "paragraph",
						Text: "An imperative UI framework.",
					},
				},
			},
		},
	}

	// Create the directory structure and files
	for _, doc := range testDocs {
		fullPath := filepath.Join(tempInputDir, doc.path)
		dir := filepath.Dir(fullPath)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}

		docBytes, err := json.Marshal(doc.content)
		if err != nil {
			t.Fatalf("Failed to marshal document: %v", err)
		}

		err = os.WriteFile(fullPath, docBytes, 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", doc.path, err)
		}
	}

	// Create technologies.json for framework index
	technologiesDoc := map[string]interface{}{
		"technologies": []interface{}{
			map[string]interface{}{
				"title": "SwiftUI",
				"path":  "/documentation/swiftui",
			},
			map[string]interface{}{
				"title": "UIKit",
				"path":  "/documentation/uikit",
			},
		},
	}
	techBytes, _ := json.Marshal(technologiesDoc)
	techPath := filepath.Join(tempInputDir, "tutorials/data/documentation/technologies.json")
	err := os.WriteFile(techPath, techBytes, 0644)
	if err != nil {
		t.Fatalf("Failed to create technologies.json: %v", err)
	}

	// Test markdown generation
	err = generateMarkdown(tempInputDir, tempOutputDir)
	if err != nil {
		t.Errorf("generateMarkdown() error = %v", err)
	}

	// Verify markdown files were created
	expectedFiles := []string{
		"tutorials/data/documentation/SwiftUI.md",
		"tutorials/data/documentation/UIKit.md",
		"index.md",
	}

	for _, expectedFile := range expectedFiles {
		mdPath := filepath.Join(tempOutputDir, expectedFile)
		if _, err := os.Stat(mdPath); os.IsNotExist(err) {
			t.Errorf("generateMarkdown() should create %s", expectedFile)
		}
	}

	// Verify index.md content
	indexPath := filepath.Join(tempOutputDir, "index.md")
	indexContent, err := os.ReadFile(indexPath)
	if err != nil {
		t.Errorf("Failed to read index.md: %v", err)
	}

	indexStr := string(indexContent)
	if !strings.Contains(indexStr, "Apple Documentation") {
		t.Errorf("index.md should contain main title")
	}
}

func TestGenerateMarkdownWithNonExistentInput(t *testing.T) {
	tempOutputDir := t.TempDir()

	// Test with non-existent input directory
	err := generateMarkdown("/nonexistent/input", tempOutputDir)
	if err == nil {
		t.Errorf("generateMarkdown() should return error for non-existent input directory")
	}
}

func TestWriteMarkdownContent(t *testing.T) {
	// Create test document
	testDoc := &DocJSONData{
		Metadata: Metadata{
			Title: "Test Document",
			Role:  "article",
		},
		Abstract: []TextContent{
			{
				Type: "paragraph",
				Text: "This is the abstract.",
			},
		},
		PrimaryContentSections: []ContentSection{
			{
				Kind: "content",
				Content: []ContentBlock{
					{
						Type:  "heading",
						Level: 2,
						Text:  "Overview",
					},
					{
						Type: "paragraph",
						InlineContent: []InlineContent{
							{
								Type: "text",
								Text: "This is the overview content.",
							},
						},
					},
				},
			},
		},
	}

	// Test writing content
	var buf strings.Builder
	err := writeMarkdownContent(&buf, testDoc)
	if err != nil {
		t.Errorf("writeMarkdownContent() error = %v", err)
	}

	result := buf.String()

	// Check expected content
	if !strings.Contains(result, "# Test Document") {
		t.Errorf("writeMarkdownContent() should contain title")
	}
	if !strings.Contains(result, "This is the abstract.") {
		t.Errorf("writeMarkdownContent() should contain abstract")
	}
	if !strings.Contains(result, "## Overview") {
		t.Errorf("writeMarkdownContent() should contain heading")
	}
	if !strings.Contains(result, "This is the overview content.") {
		t.Errorf("writeMarkdownContent() should contain paragraph content")
	}
}

func TestCreateFrameworkIndex(t *testing.T) {
	tempDir := t.TempDir()

	// Create some test JSON files
	jsonFiles := []string{
		"tutorials/data/documentation/SwiftUI.json",
		"tutorials/data/documentation/UIKit.json",
		"tutorials/data/documentation/Foundation.json",
	}

	// Create test files
	for _, file := range jsonFiles {
		fullPath := filepath.Join(tempDir, file)
		dir := filepath.Dir(fullPath)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			t.Fatalf("Failed to create directory: %v", err)
		}

		testDoc := map[string]interface{}{
			"metadata": map[string]interface{}{
				"title": strings.TrimSuffix(filepath.Base(file), ".json"),
			},
		}
		docBytes, _ := json.Marshal(testDoc)
		err = os.WriteFile(fullPath, docBytes, 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	// Test framework index creation
	err := createFrameworkIndex(tempDir, jsonFiles)
	if err != nil {
		t.Errorf("createFrameworkIndex() error = %v", err)
	}

	// Verify index.md was created
	indexPath := filepath.Join(tempDir, "index.md")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Errorf("createFrameworkIndex() should create index.md")
	}

	// Read and verify content
	content, err := os.ReadFile(indexPath)
	if err != nil {
		t.Errorf("Failed to read index.md: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "Apple Documentation") {
		t.Errorf("Index should contain main title")
	}
}

// Benchmark for markdown generation performance
func BenchmarkConvertJSONToMarkdown(b *testing.B) {
	tempDir := b.TempDir()

	// Create a test document
	testDoc := DocJSONData{
		Metadata: Metadata{
			Title: "Benchmark Document",
			Role:  "article",
		},
		Abstract: []TextContent{
			{
				Type: "paragraph",
				Text: "This is a benchmark document.",
			},
		},
		PrimaryContentSections: make([]ContentSection, 5),
	}

	// Populate with some content
	for i := 0; i < 5; i++ {
		testDoc.PrimaryContentSections[i] = ContentSection{
			Kind: "content",
			Content: []ContentBlock{
				{
					Type:  "heading",
					Level: 2,
					Text:  fmt.Sprintf("Section %d", i),
				},
				{
					Type: "paragraph",
					InlineContent: []InlineContent{
						{
							Type: "text",
							Text: fmt.Sprintf("Content for section %d", i),
						},
					},
				},
			},
		}
	}

	testDocBytes, _ := json.Marshal(testDoc)
	jsonPath := filepath.Join(tempDir, "benchmark.json")
	err := os.WriteFile(jsonPath, testDocBytes, 0644)
	if err != nil {
		b.Fatalf("Failed to create benchmark file: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mdPath := filepath.Join(tempDir, fmt.Sprintf("benchmark_%d.md", i))
		_ = convertJSONToMarkdown(jsonPath, mdPath)
	}
}
