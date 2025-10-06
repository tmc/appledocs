package reader

import (
	"io/fs"
	"path/filepath"
	"testing"
)

const testDataDir = "../output/tutorials/data/documentation"

// TestOpen verifies that we can open the documentation directory.
func TestOpen(t *testing.T) {
	fsys, err := Open(testDataDir)
	if err != nil {
		t.Skipf("skipping test: %v (run from project root with documentation downloaded)", err)
	}

	if fsys == nil {
		t.Fatal("expected non-nil FS")
	}
}

// TestListFrameworks verifies we can list available frameworks.
func TestListFrameworks(t *testing.T) {
	fsys, err := Open(testDataDir)
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	frameworks, err := ListFrameworks(fsys)
	if err != nil {
		t.Fatalf("ListFrameworks failed: %v", err)
	}

	if len(frameworks) == 0 {
		t.Fatal("expected at least one framework")
	}

	// Check for common frameworks
	expectedFrameworks := []string{"Foundation", "AppKit", "UIKit"}
	for _, expected := range expectedFrameworks {
		found := false
		for _, framework := range frameworks {
			if framework == expected {
				found = true
				break
			}
		}
		if !found {
			t.Logf("warning: expected framework %s not found (this may be okay)", expected)
		}
	}

	t.Logf("Found %d frameworks", len(frameworks))
}

// TestGetFramework verifies we can read framework documentation.
func TestGetFramework(t *testing.T) {
	fsys, err := Open(testDataDir)
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	doc, err := GetFramework(fsys, "Foundation")
	if err != nil {
		t.Fatalf("GetFramework failed: %v", err)
	}

	if doc.Metadata.Title != "Foundation" {
		t.Errorf("expected title 'Foundation', got %q", doc.Metadata.Title)
	}

	if doc.Kind != "symbol" {
		t.Errorf("expected kind 'symbol', got %q", doc.Kind)
	}

	if len(doc.Abstract) == 0 {
		t.Error("expected non-empty abstract")
	}

	t.Logf("Foundation abstract: %s", doc.Abstract[0].Text)
}

// TestGetFrameworkInfo verifies we can get framework information.
func TestGetFrameworkInfo(t *testing.T) {
	fsys, err := Open(testDataDir)
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	info, err := GetFrameworkInfo(fsys, "Foundation")
	if err != nil {
		t.Fatalf("GetFrameworkInfo failed: %v", err)
	}

	if info.Name != "Foundation" {
		t.Errorf("expected name 'Foundation', got %q", info.Name)
	}

	if info.Abstract == "" {
		t.Error("expected non-empty abstract")
	}

	if len(info.Platforms) == 0 {
		t.Error("expected at least one platform")
	}

	t.Logf("Foundation: %s", info.Abstract)
	t.Logf("Platforms: %d", len(info.Platforms))
}

// TestListSymbols verifies we can list symbols in a framework.
func TestListSymbols(t *testing.T) {
	fsys, err := Open(testDataDir)
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	symbols, err := ListSymbols(fsys, "Foundation")
	if err != nil {
		t.Fatalf("ListSymbols failed: %v", err)
	}

	if len(symbols) == 0 {
		t.Fatal("expected at least one symbol")
	}

	// Check for a common symbol
	found := false
	for _, symbol := range symbols {
		if symbol == "NSString" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find NSString symbol")
	}

	t.Logf("Found %d symbols in Foundation", len(symbols))
}

// TestGetSymbol verifies we can read symbol documentation.
func TestGetSymbol(t *testing.T) {
	fsys, err := Open(testDataDir)
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	doc, err := GetSymbol(fsys, "Foundation/NSString")
	if err != nil {
		t.Fatalf("GetSymbol failed: %v", err)
	}

	if doc.Metadata.Title != "NSString" {
		t.Errorf("expected title 'NSString', got %q", doc.Metadata.Title)
	}

	t.Logf("NSString kind: %s, role: %s", doc.Kind, doc.Metadata.Role)
}

// TestGetSymbolInfo verifies we can get symbol information.
func TestGetSymbolInfo(t *testing.T) {
	fsys, err := Open(testDataDir)
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	info, err := GetSymbolInfo(fsys, "Foundation/NSString")
	if err != nil {
		t.Fatalf("GetSymbolInfo failed: %v", err)
	}

	if info.Framework != "Foundation" {
		t.Errorf("expected framework 'Foundation', got %q", info.Framework)
	}

	if info.Name != "NSString" {
		t.Errorf("expected name 'NSString', got %q", info.Name)
	}

	if info.URL == "" {
		t.Error("expected non-empty URL")
	}

	t.Logf("NSString URL: %s", info.URL)
	t.Logf("NSString abstract: %s", info.Abstract)
}

// TestSearchSymbols verifies symbol search functionality.
func TestSearchSymbols(t *testing.T) {
	fsys, err := Open(testDataDir)
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	matches, err := SearchSymbols(fsys, "Foundation", "String")
	if err != nil {
		t.Fatalf("SearchSymbols failed: %v", err)
	}

	if len(matches) == 0 {
		t.Fatal("expected at least one match for 'String'")
	}

	// NSString should be in the results
	found := false
	for _, match := range matches {
		if match == "NSString" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find NSString in search results")
	}

	t.Logf("Found %d symbols matching 'String'", len(matches))
}

// TestFrameworkName verifies framework name extraction.
func TestFrameworkName(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"Foundation.json", "Foundation"},
		{"AppKit.json", "AppKit"},
		{"Foundation/NSString.json", "Foundation"},
		{"UIKit/UIView.json", "UIKit"},
		{filepath.Join("Foundation", "NSString.json"), "Foundation"},
	}

	for _, tt := range tests {
		result := FrameworkName(tt.path)
		if result != tt.expected {
			t.Errorf("FrameworkName(%q) = %q, want %q", tt.path, result, tt.expected)
		}
	}
}

// TestIsFramework verifies framework detection.
func TestIsFramework(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{"Foundation.json", true},
		{"AppKit.json", true},
		{"Foundation/NSString.json", false},
		{"UIKit/UIView.json", false},
		{filepath.Join("Foundation", "NSString.json"), false},
		{"notjson.txt", false},
	}

	for _, tt := range tests {
		result := IsFramework(tt.name)
		if result != tt.expected {
			t.Errorf("IsFramework(%q) = %v, want %v", tt.name, result, tt.expected)
		}
	}
}

// TestFSInterface verifies the FS implements fs.FS.
func TestFSInterface(t *testing.T) {
	fsys, err := Open(testDataDir)
	if err != nil {
		t.Skipf("skipping test: %v", err)
	}

	// Verify we can use it as fs.FS
	var _ fs.FS = fsys

	// Test Open
	file, err := fsys.Open("Foundation.json")
	if err != nil {
		t.Fatalf("Open failed: %v", err)
	}
	defer file.Close()

	// Verify we can read from it
	stat, err := file.Stat()
	if err != nil {
		t.Fatalf("Stat failed: %v", err)
	}

	if stat.Size() == 0 {
		t.Error("expected non-zero file size")
	}

	t.Logf("Foundation.json size: %d bytes", stat.Size())
}
