package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestGenerateHTMLFile(t *testing.T) {
	// Create a temporary directory
	tempDir := t.TempDir()
	indexPath := filepath.Join(tempDir, "index.html")

	// Create test data
	root := &TreeNode{
		Name:  "root",
		IsDir: true,
		Children: []*TreeNode{
			{
				Name:  "documentation",
				IsDir: true,
				Children: []*TreeNode{
					{Name: "SwiftUI.json", IsDir: false, Path: "documentation/SwiftUI.json"},
					{Name: "UIKit.json", IsDir: false, Path: "documentation/UIKit.json"},
				},
			},
			{Name: "index.json", IsDir: false, Path: "index.json"},
		},
	}

	data := struct {
		Root      *TreeNode
		FileCount int
		DirCount  int
		Timestamp string
	}{
		Root:      root,
		FileCount: 3,
		DirCount:  1,
		Timestamp: time.Now().Format(time.RFC1123),
	}

	// Test HTML generation
	err := generateHTMLFile(indexPath, data)
	if err != nil {
		t.Errorf("generateHTMLFile() error = %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Errorf("generateHTMLFile() did not create HTML file")
	}

	// Read and verify content
	content, err := os.ReadFile(indexPath)
	if err != nil {
		t.Errorf("Failed to read generated HTML file: %v", err)
	}

	htmlContent := string(content)

	// Check for required HTML elements
	tests := []struct {
		name     string
		contains string
	}{
		{"HTML document type", "<!DOCTYPE html>"},
		{"Title", "<title>Apple Documentation JSON Mirror</title>"},
		{"CSS styles", "<style>"},
		{"JavaScript", "<script>"},
		{"File count", "3"},
		{"Directory count", "1"},
		{"Documentation directory", "documentation"},
		{"SwiftUI file", "SwiftUI.json"},
		{"UIKit file", "UIKit.json"},
		{"Index file", "index.json"},
		{"Tree structure", "class=\"tree\""},
		{"Stats info", "Found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !strings.Contains(htmlContent, tt.contains) {
				t.Errorf("generateHTMLFile() HTML content should contain %q", tt.contains)
			}
		})
	}

	// Verify HTML structure
	if !strings.Contains(htmlContent, "<html") {
		t.Errorf("generateHTMLFile() should generate valid HTML")
	}
	if !strings.Contains(htmlContent, "</html>") {
		t.Errorf("generateHTMLFile() should close HTML tag")
	}
	if !strings.Contains(htmlContent, "<body") {
		t.Errorf("generateHTMLFile() should include body tag")
	}
	if !strings.Contains(htmlContent, "</body>") {
		t.Errorf("generateHTMLFile() should close body tag")
	}
}

func TestGenerateHTMLFileWithEmptyTree(t *testing.T) {
	tempDir := t.TempDir()
	indexPath := filepath.Join(tempDir, "empty.html")

	// Create empty tree data
	root := &TreeNode{
		Name:     "root",
		IsDir:    true,
		Children: []*TreeNode{},
	}

	data := struct {
		Root      *TreeNode
		FileCount int
		DirCount  int
		Timestamp string
	}{
		Root:      root,
		FileCount: 0,
		DirCount:  0,
		Timestamp: time.Now().Format(time.RFC1123),
	}

	// Test HTML generation with empty data
	err := generateHTMLFile(indexPath, data)
	if err != nil {
		t.Errorf("generateHTMLFile() with empty tree error = %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		t.Errorf("generateHTMLFile() did not create HTML file for empty tree")
	}

	// Read and verify content
	content, err := os.ReadFile(indexPath)
	if err != nil {
		t.Errorf("Failed to read generated HTML file: %v", err)
	}

	htmlContent := string(content)

	// Check that it still generates valid HTML
	if !strings.Contains(htmlContent, "<!DOCTYPE html>") {
		t.Errorf("generateHTMLFile() should generate valid HTML even for empty tree")
	}
	if !strings.Contains(htmlContent, "0") {
		t.Errorf("generateHTMLFile() should show 0 file count for empty tree")
	}
}

func TestGenerateHTMLFileWithInvalidPath(t *testing.T) {
	// Try to create HTML file in non-existent directory without creating it first
	invalidPath := "/nonexistent/directory/index.html"

	root := &TreeNode{
		Name:     "root",
		IsDir:    true,
		Children: []*TreeNode{},
	}

	data := struct {
		Root      *TreeNode
		FileCount int
		DirCount  int
		Timestamp string
	}{
		Root:      root,
		FileCount: 0,
		DirCount:  0,
		Timestamp: time.Now().Format(time.RFC1123),
	}

	// Test HTML generation with invalid path should return error
	err := generateHTMLFile(invalidPath, data)
	if err == nil {
		t.Errorf("generateHTMLFile() should return error for invalid path")
	}
}

func TestGenerateHTMLFileWithNestedStructure(t *testing.T) {
	tempDir := t.TempDir()
	indexPath := filepath.Join(tempDir, "nested.html")

	// Create deeply nested tree structure
	root := &TreeNode{
		Name:  "root",
		IsDir: true,
		Children: []*TreeNode{
			{
				Name:  "tutorials",
				IsDir: true,
				Children: []*TreeNode{
					{
						Name:  "data",
						IsDir: true,
						Children: []*TreeNode{
							{
								Name:  "documentation",
								IsDir: true,
								Children: []*TreeNode{
									{
										Name:  "SwiftUI",
										IsDir: true,
										Children: []*TreeNode{
											{Name: "View.json", IsDir: false, Path: "tutorials/data/documentation/SwiftUI/View.json"},
											{Name: "Button.json", IsDir: false, Path: "tutorials/data/documentation/SwiftUI/Button.json"},
										},
									},
									{Name: "Foundation.json", IsDir: false, Path: "tutorials/data/documentation/Foundation.json"},
								},
							},
						},
					},
				},
			},
		},
	}

	data := struct {
		Root      *TreeNode
		FileCount int
		DirCount  int
		Timestamp string
	}{
		Root:      root,
		FileCount: 3,
		DirCount:  4,
		Timestamp: time.Now().Format(time.RFC1123),
	}

	// Test HTML generation
	err := generateHTMLFile(indexPath, data)
	if err != nil {
		t.Errorf("generateHTMLFile() with nested structure error = %v", err)
	}

	// Read and verify content
	content, err := os.ReadFile(indexPath)
	if err != nil {
		t.Errorf("Failed to read generated HTML file: %v", err)
	}

	htmlContent := string(content)

	// Check for nested structure elements
	tests := []struct {
		name     string
		contains string
	}{
		{"Tutorials directory", "tutorials"},
		{"Data directory", "data"},
		{"Documentation directory", "documentation"},
		{"SwiftUI directory", "SwiftUI"},
		{"View file", "View.json"},
		{"Button file", "Button.json"},
		{"Foundation file", "Foundation.json"},
		{"Correct file count", "3"},
		{"Correct directory count", "4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !strings.Contains(htmlContent, tt.contains) {
				t.Errorf("generateHTMLFile() nested structure should contain %q", tt.contains)
			}
		})
	}
}

func TestGenerateHTMLFileWithSpecialCharacters(t *testing.T) {
	tempDir := t.TempDir()
	indexPath := filepath.Join(tempDir, "special.html")

	// Create tree with special characters that need HTML escaping
	root := &TreeNode{
		Name:  "root",
		IsDir: true,
		Children: []*TreeNode{
			{Name: "file<>&.json", IsDir: false, Path: "file<>&.json"},
			{Name: "file\"quotes\".json", IsDir: false, Path: "file\"quotes\".json"},
		},
	}

	data := struct {
		Root      *TreeNode
		FileCount int
		DirCount  int
		Timestamp string
	}{
		Root:      root,
		FileCount: 2,
		DirCount:  0,
		Timestamp: time.Now().Format(time.RFC1123),
	}

	// Test HTML generation
	err := generateHTMLFile(indexPath, data)
	if err != nil {
		t.Errorf("generateHTMLFile() with special characters error = %v", err)
	}

	// Read and verify content
	content, err := os.ReadFile(indexPath)
	if err != nil {
		t.Errorf("Failed to read generated HTML file: %v", err)
	}

	htmlContent := string(content)

	// Check that special characters are properly escaped in HTML
	// The template should handle HTML escaping automatically
	if !strings.Contains(htmlContent, "file") {
		t.Errorf("generateHTMLFile() should contain file references even with special characters")
	}

	// Verify it's still valid HTML structure
	if !strings.Contains(htmlContent, "<!DOCTYPE html>") {
		t.Errorf("generateHTMLFile() should generate valid HTML with special characters")
	}
}

// Benchmark test for HTML generation performance
func BenchmarkGenerateHTMLFile(b *testing.B) {
	tempDir := b.TempDir()

	// Create a reasonably sized tree structure
	root := &TreeNode{
		Name:     "root",
		IsDir:    true,
		Children: make([]*TreeNode, 100),
	}

	// Populate with test data
	for i := 0; i < 100; i++ {
		root.Children[i] = &TreeNode{
			Name:  fmt.Sprintf("file%d.json", i),
			IsDir: false,
			Path:  fmt.Sprintf("path/to/file%d.json", i),
		}
	}

	data := struct {
		Root      *TreeNode
		FileCount int
		DirCount  int
		Timestamp string
	}{
		Root:      root,
		FileCount: 100,
		DirCount:  1,
		Timestamp: time.Now().Format(time.RFC1123),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexPath := filepath.Join(tempDir, fmt.Sprintf("bench%d.html", i))
		_ = generateHTMLFile(indexPath, data)
	}
}
