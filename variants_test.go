package appledocs

import (
	"os"
	"path/filepath"
	"testing"
)

// TestGetCrossReference tests cross-referencing between Swift and Objective-C APIs.
func TestGetCrossReference(t *testing.T) {
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
	fsys, err := Open(cacheDir)
	if err != nil {
		t.Fatal("Failed to open cache:", err)
	}

	// Test with NSString from Foundation - a class that exists in both Swift and ObjC
	doc, err := GetSymbol(fsys, "Foundation/NSString")
	if err != nil {
		t.Skip("NSString not available:", err)
	}

	// Get cross-reference information
	ref := GetCrossReference(doc)

	t.Logf("Symbol: %s", ref.Title)
	t.Logf("Kind: %s", ref.SymbolKind)
	t.Logf("Swift available: %v", ref.Available.Swift)
	t.Logf("Objective-C available: %v", ref.Available.ObjectiveC)

	// NSString should be available in both languages
	if !ref.Available.Swift {
		t.Error("Expected Swift variant to be available for NSString")
	}
	if !ref.Available.ObjectiveC {
		t.Error("Expected Objective-C variant to be available for NSString")
	}

	if ref.SwiftDeclaration != "" {
		t.Logf("Swift declaration: %s", ref.SwiftDeclaration)
	}
	if ref.ObjCDeclaration != "" {
		t.Logf("Objective-C declaration: %s", ref.ObjCDeclaration)
	}
}

// TestHasVariants tests the variant detection functions.
func TestHasVariants(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get home directory:", err)
	}
	cacheDir := filepath.Join(homeDir, ".appledocs/cache/developer.apple.com/tutorials/data/documentation")

	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		t.Skip("Cache directory not found:", cacheDir)
	}

	fsys, err := Open(cacheDir)
	if err != nil {
		t.Fatal("Failed to open cache:", err)
	}

	// Test with a Foundation class
	doc, err := GetSymbol(fsys, "Foundation/NSArray")
	if err != nil {
		t.Skip("NSArray not available:", err)
	}

	hasSwift := HasSwiftVariant(doc)
	hasObjC := HasObjectiveCVariant(doc)

	t.Logf("NSArray - Swift variant: %v, Objective-C variant: %v", hasSwift, hasObjC)

	// At least one should be true
	if !hasSwift && !hasObjC {
		t.Error("Expected at least one language variant to be available")
	}
}

// TestGetVariant tests getting specific language variants.
func TestGetVariant(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get home directory:", err)
	}
	cacheDir := filepath.Join(homeDir, ".appledocs/cache/developer.apple.com/tutorials/data/documentation")

	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		t.Skip("Cache directory not found:", cacheDir)
	}

	fsys, err := Open(cacheDir)
	if err != nil {
		t.Fatal("Failed to open cache:", err)
	}

	// Test with NSString
	doc, err := GetSymbol(fsys, "Foundation/NSString")
	if err != nil {
		t.Skip("NSString not available:", err)
	}

	t.Logf("Original language: %s", doc.Identifier.InterfaceLanguage)

	// Try to get Swift variant
	swiftDoc, err := GetSwiftVariant(doc)
	if err != nil {
		t.Logf("Swift variant not available: %v", err)
	} else {
		t.Logf("Swift variant language: %s", swiftDoc.Identifier.InterfaceLanguage)
		t.Logf("Swift declaration: %s", GetDeclarationText(swiftDoc, LanguageSwift))
	}

	// Try to get Objective-C variant
	objcDoc, err := GetObjectiveCVariant(doc)
	if err != nil {
		t.Logf("Objective-C variant not available: %v", err)
	} else {
		t.Logf("Objective-C variant language: %s", objcDoc.Identifier.InterfaceLanguage)
		t.Logf("Objective-C declaration: %s", GetDeclarationText(objcDoc, LanguageObjectiveC))
	}
}
