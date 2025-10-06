package appledocs

import (
	"testing"
)

func TestGetFramework(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	doc, err := GetFramework(fsys, "Foundation")
	if err != nil {
		t.Fatalf("GetFramework failed: %v", err)
	}

	if doc.Kind != "symbol" {
		t.Errorf("Kind = %q, want 'symbol'", doc.Kind)
	}
	if doc.Metadata.Title != "Foundation" {
		t.Errorf("Title = %q, want 'Foundation'", doc.Metadata.Title)
	}
}

func TestListFrameworks(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	frameworks, err := ListFrameworks(fsys)
	if err != nil {
		t.Fatalf("ListFrameworks failed: %v", err)
	}

	if len(frameworks) == 0 {
		t.Error("expected at least one framework")
	}

	// Check that frameworks are sorted
	for i := 1; i < len(frameworks); i++ {
		if frameworks[i-1] > frameworks[i] {
			t.Errorf("frameworks not sorted: %q > %q", frameworks[i-1], frameworks[i])
		}
	}

	t.Logf("Found %d frameworks", len(frameworks))
}

func TestGetSymbol(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	// Test with .json extension
	doc1, err := GetSymbol(fsys, "Foundation/NSString.json")
	if err != nil {
		t.Fatalf("GetSymbol failed: %v", err)
	}

	// Test without .json extension
	doc2, err := GetSymbol(fsys, "Foundation/NSString")
	if err != nil {
		t.Fatalf("GetSymbol failed: %v", err)
	}

	// Both should load the same document
	if doc1.Metadata.Title != doc2.Metadata.Title {
		t.Errorf("inconsistent results: %q vs %q", doc1.Metadata.Title, doc2.Metadata.Title)
	}

	if doc1.Metadata.Title != "NSString" {
		t.Errorf("Title = %q, want 'NSString'", doc1.Metadata.Title)
	}
	if doc1.Metadata.SymbolKind != "class" {
		t.Errorf("SymbolKind = %q, want 'class'", doc1.Metadata.SymbolKind)
	}
	if doc1.Metadata.ExternalID == "" {
		t.Error("ExternalID is empty")
	}

	t.Logf("NSString externalID: %s", doc1.Metadata.ExternalID)
}

func TestGetSymbolByURL(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	url := "doc://com.apple.foundation/documentation/Foundation/NSString"
	doc, err := GetSymbolByURL(fsys, url)
	if err != nil {
		t.Fatalf("GetSymbolByURL failed: %v", err)
	}

	if doc.Metadata.Title != "NSString" {
		t.Errorf("Title = %q, want 'NSString'", doc.Metadata.Title)
	}
}

func TestListSymbols(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	symbols, err := ListSymbols(fsys, "Foundation")
	if err != nil {
		t.Fatalf("ListSymbols failed: %v", err)
	}

	if len(symbols) == 0 {
		t.Error("expected at least one symbol")
	}

	// Check that symbols are sorted
	for i := 1; i < len(symbols); i++ {
		if symbols[i-1] > symbols[i] {
			t.Errorf("symbols not sorted: %q > %q", symbols[i-1], symbols[i])
		}
	}

	// Check that NSString is in the list
	found := false
	for _, s := range symbols {
		if s == "NSString" {
			found = true
			break
		}
	}
	if !found {
		t.Error("NSString not found in Foundation symbols")
	}

	t.Logf("Found %d symbols in Foundation", len(symbols))
}

func TestGetFrameworkInfo(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	info, err := GetFrameworkInfo(fsys, "Foundation")
	if err != nil {
		t.Fatalf("GetFrameworkInfo failed: %v", err)
	}

	if info.Name != "Foundation" {
		t.Errorf("Name = %q, want 'Foundation'", info.Name)
	}
	if info.Title == "" {
		t.Error("Title is empty")
	}
	if len(info.Platforms) == 0 {
		t.Error("expected at least one platform")
	}

	t.Logf("Foundation: %s", info.Abstract)
	t.Logf("Platforms: %d", len(info.Platforms))
}

func TestGetSymbolInfo(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	info, err := GetSymbolInfo(fsys, "Foundation/NSString")
	if err != nil {
		t.Fatalf("GetSymbolInfo failed: %v", err)
	}

	if info.Framework != "Foundation" {
		t.Errorf("Framework = %q, want 'Foundation'", info.Framework)
	}
	if info.Name != "NSString" {
		t.Errorf("Name = %q, want 'NSString'", info.Name)
	}
	if info.Title != "NSString" {
		t.Errorf("Title = %q, want 'NSString'", info.Title)
	}
	if info.SymbolKind != "class" {
		t.Errorf("SymbolKind = %q, want 'class'", info.SymbolKind)
	}
	if len(info.Platforms) == 0 {
		t.Error("expected at least one platform")
	}

	t.Logf("NSString: %s", info.Abstract)
}

func TestSearchSymbols(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	// Search for "string" in Foundation
	matches, err := SearchSymbols(fsys, "Foundation", "string")
	if err != nil {
		t.Fatalf("SearchSymbols failed: %v", err)
	}

	// Should find at least NSString
	found := false
	for _, match := range matches {
		if match == "NSString" {
			found = true
			break
		}
	}
	if !found {
		t.Error("NSString not found in search results for 'string'")
	}

	t.Logf("Found %d matches for 'string'", len(matches))
}

func TestTypedAPIReferences(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	doc, err := GetSymbol(fsys, "Foundation/NSString")
	if err != nil {
		t.Fatalf("GetSymbol failed: %v", err)
	}

	if len(doc.References) == 0 {
		t.Error("expected at least one reference")
	}

	// Count different types of references
	symbolCount := 0
	methodCount := 0
	for _, ref := range doc.References {
		if ref.Role == "symbol" {
			symbolCount++
			if ref.SymbolKind == "method" {
				methodCount++
			}
		}
	}

	t.Logf("NSString has %d references, %d symbols, %d methods",
		len(doc.References), symbolCount, methodCount)

	if symbolCount == 0 {
		t.Error("expected at least one symbol reference")
	}
}

func TestTypedAPIPlatforms(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	doc, err := GetSymbol(fsys, "Foundation/NSString")
	if err != nil {
		t.Fatalf("GetSymbol failed: %v", err)
	}

	if len(doc.Metadata.Platforms) == 0 {
		t.Error("expected at least one platform")
	}

	// Check platform structure
	for _, p := range doc.Metadata.Platforms {
		if p.Name == "" {
			t.Error("platform has empty name")
		}
		t.Logf("Platform: %s (introduced: %s, deprecated: %v)",
			p.Name, p.IntroducedAt, p.Deprecated)
	}
}

func TestTypedAPIFragments(t *testing.T) {
	fsys, err := Open(testDocsDir)
	if err != nil {
		t.Skipf("skipping: %v", err)
	}

	doc, err := GetSymbol(fsys, "Foundation/NSString")
	if err != nil {
		t.Fatalf("GetSymbol failed: %v", err)
	}

	// Find a reference with fragments
	foundFragments := false
	for _, ref := range doc.References {
		if len(ref.Fragments) > 0 {
			foundFragments = true
			for _, frag := range ref.Fragments {
				if frag.Kind == "" {
					t.Error("fragment has empty kind")
				}
				if frag.Text == "" {
					t.Error("fragment has empty text")
				}
			}
			break
		}
	}

	if !foundFragments {
		t.Error("expected to find at least one reference with fragments")
	}
}
