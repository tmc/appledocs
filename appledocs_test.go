package appledocs

import (
	"testing"
)

const testDocsDir = "output/tutorials/data/documentation"

func TestOpen(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v (documentation not downloaded)", err)
	}
	if fsys == nil {
		t.Fatal("expected non-nil fs.FS")
	}
}

func TestLoadMap(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	doc, err := LoadMap(fsys, "Foundation.json")
	if err != nil {
		t.Fatalf("LoadMap failed: %v", err)
	}

	if doc == nil {
		t.Fatal("expected non-nil map")
	}

	// All docs should have these fields
	if doc["kind"] == nil {
		t.Error("missing 'kind' field")
	}
	if doc["metadata"] == nil {
		t.Error("missing 'metadata' field")
	}
	if doc["identifier"] == nil {
		t.Error("missing 'identifier' field")
	}
}

func TestGetString(t *testing.T) {
	m := map[string]interface{}{
		"kind": "symbol",
		"metadata": map[string]interface{}{
			"title": "Foundation",
			"role":  "collection",
		},
	}

	tests := []struct {
		path     []string
		expected string
	}{
		{[]string{"kind"}, "symbol"},
		{[]string{"metadata", "title"}, "Foundation"},
		{[]string{"metadata", "role"}, "collection"},
		{[]string{"nonexistent"}, ""},
		{[]string{"metadata", "nonexistent"}, ""},
	}

	for _, tt := range tests {
		result := GetString(m, tt.path...)
		if result != tt.expected {
			t.Errorf("GetString(%v) = %q, want %q", tt.path, result, tt.expected)
		}
	}
}

func TestGetInt(t *testing.T) {
	m := map[string]interface{}{
		"schemaVersion": map[string]interface{}{
			"major": 0.0, // JSON numbers are float64
			"minor": 3.0,
		},
	}

	major := GetInt(m, "schemaVersion", "major")
	if major != 0 {
		t.Errorf("GetInt(major) = %d, want 0", major)
	}

	minor := GetInt(m, "schemaVersion", "minor")
	if minor != 3 {
		t.Errorf("GetInt(minor) = %d, want 3", minor)
	}

	missing := GetInt(m, "nonexistent")
	if missing != 0 {
		t.Errorf("GetInt(nonexistent) = %d, want 0", missing)
	}
}

func TestHelpers(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	doc, err := LoadMap(fsys, "Foundation.json")
	if err != nil {
		t.Fatalf("LoadMap failed: %v", err)
	}

	kind := Kind(doc)
	if kind != "symbol" {
		t.Errorf("Kind() = %q, want 'symbol'", kind)
	}

	title := Title(doc)
	if title != "Foundation" {
		t.Errorf("Title() = %q, want 'Foundation'", title)
	}

	role := Role(doc)
	if role == "" {
		t.Error("Role() returned empty string")
	}

	url := URL(doc)
	if url == "" {
		t.Error("URL() returned empty string")
	}

	refs := References(doc)
	if refs == nil {
		t.Error("References() returned nil")
	}

	meta := MetadataMap(doc)
	if meta == nil {
		t.Error("MetadataMap() returned nil")
	}
}

func TestSymbolDocument(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	doc, err := LoadMap(fsys, "Foundation/NSString.json")
	if err != nil {
		t.Fatalf("LoadMap failed: %v", err)
	}

	kind := Kind(doc)
	if kind != "symbol" {
		t.Errorf("Kind() = %q, want 'symbol'", kind)
	}

	title := Title(doc)
	if title != "NSString" {
		t.Errorf("Title() = %q, want 'NSString'", title)
	}

	symbolKind := SymbolKind(doc)
	if symbolKind != "class" {
		t.Errorf("SymbolKind() = %q, want 'class'", symbolKind)
	}

	extID := ExternalID(doc)
	if extID == "" {
		t.Error("ExternalID() returned empty string")
	}
	t.Logf("NSString externalID: %s", extID)

	platforms := Platforms(doc)
	if len(platforms) == 0 {
		t.Error("Platforms() returned empty array")
	}
	t.Logf("NSString available on %d platforms", len(platforms))
}

func TestIsFramework(t *testing.T) {
	tests := []struct {
		path     string
		expected bool
	}{
		{"Foundation.json", true},
		{"AppKit.json", true},
		{"Foundation/NSString.json", false},
		{"UIKit/UIView.json", false},
		{"notjson.txt", false},
	}

	for _, tt := range tests {
		result := IsFramework(tt.path)
		if result != tt.expected {
			t.Errorf("IsFramework(%q) = %v, want %v", tt.path, result, tt.expected)
		}
	}
}

func TestFrameworkName(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"Foundation.json", "Foundation"},
		{"AppKit.json", "AppKit"},
		{"Foundation/NSString.json", "Foundation"},
		{"UIKit/UIView.json", "UIKit"},
	}

	for _, tt := range tests {
		result := FrameworkName(tt.path)
		if result != tt.expected {
			t.Errorf("FrameworkName(%q) = %q, want %q", tt.path, result, tt.expected)
		}
	}
}

func TestSymbolPath(t *testing.T) {
	tests := []struct {
		framework string
		symbol    string
		expected  string
	}{
		{"Foundation", "NSString", "Foundation/NSString.json"},
		{"AppKit", "NSView", "AppKit/NSView.json"},
	}

	for _, tt := range tests {
		result := SymbolPath(tt.framework, tt.symbol)
		if result != tt.expected {
			t.Errorf("SymbolPath(%q, %q) = %q, want %q",
				tt.framework, tt.symbol, result, tt.expected)
		}
	}
}
