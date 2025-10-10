// Package appledocs provides access to Apple's developer documentation.
//
// This file contains functionality for cross-referencing between Swift
// and Objective-C language variants of the same API.
//
// Many Apple APIs are available in both Swift and Objective-C, but the
// documentation may be stored with one as the primary language. The variant
// system allows you to transform between these representations.
//
// Example usage:
//
//	// Get a symbol's documentation
//	doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
//
//	// Get cross-reference information
//	ref := appledocs.GetCrossReference(doc)
//	fmt.Printf("Swift: %s\n", ref.SwiftDeclaration)
//	fmt.Printf("Objective-C: %s\n", ref.ObjCDeclaration)
//
//	// Get specific language variant
//	swiftDoc, _ := appledocs.GetSwiftVariant(doc)
//	objcDoc, _ := appledocs.GetObjectiveCVariant(doc)
package appledocs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// LanguageVariant represents the interface language for documentation.
type LanguageVariant string

const (
	LanguageSwift      LanguageVariant = "swift"
	LanguageObjectiveC LanguageVariant = "occ"
)

// GetSwiftVariant returns the Swift variant of a documentation symbol.
// If the document is already in Swift, it returns the original document.
// If an Objective-C variant exists, it applies the necessary transformations.
func GetSwiftVariant(doc *Document) (*Document, error) {
	if doc.Identifier.InterfaceLanguage == string(LanguageSwift) {
		return doc, nil
	}

	// Apply variant overrides to get Swift version
	return applyVariantOverrides(doc, LanguageSwift)
}

// GetObjectiveCVariant returns the Objective-C variant of a documentation symbol.
// If the document is already in Objective-C, it returns the original document.
// If a Swift variant exists, it applies the necessary transformations.
func GetObjectiveCVariant(doc *Document) (*Document, error) {
	if doc.Identifier.InterfaceLanguage == string(LanguageObjectiveC) {
		return doc, nil
	}

	// Apply variant overrides to get Objective-C version
	return applyVariantOverrides(doc, LanguageObjectiveC)
}

// HasSwiftVariant returns true if the document has a Swift language variant available.
func HasSwiftVariant(doc *Document) bool {
	return hasVariant(doc, LanguageSwift)
}

// HasObjectiveCVariant returns true if the document has an Objective-C language variant available.
func HasObjectiveCVariant(doc *Document) bool {
	return hasVariant(doc, LanguageObjectiveC)
}

// hasVariant checks if a document has a specific language variant.
func hasVariant(doc *Document, lang LanguageVariant) bool {
	// Check current language
	if doc.Identifier.InterfaceLanguage == string(lang) {
		return true
	}

	// Check variant overrides
	for _, variant := range doc.VariantOverrides {
		for _, trait := range variant.Traits {
			if trait.InterfaceLanguage == string(lang) {
				return true
			}
		}
	}

	// Check variants list
	for _, variant := range doc.Variants {
		for _, trait := range variant.Traits {
			if trait.InterfaceLanguage == string(lang) {
				return true
			}
		}
	}

	return false
}

// applyVariantOverrides applies JSON Patch operations to transform the document
// to the specified language variant.
func applyVariantOverrides(doc *Document, targetLang LanguageVariant) (*Document, error) {
	// Find the variant override for the target language
	var override *VariantOverride
	for i := range doc.VariantOverrides {
		for _, trait := range doc.VariantOverrides[i].Traits {
			if trait.InterfaceLanguage == string(targetLang) {
				override = &doc.VariantOverrides[i]
				break
			}
		}
		if override != nil {
			break
		}
	}

	if override == nil {
		return nil, fmt.Errorf("no %s variant available", targetLang)
	}

	// Marshal the document to JSON
	docJSON, err := json.Marshal(doc)
	if err != nil {
		return nil, fmt.Errorf("marshal document: %w", err)
	}

	// Unmarshal to generic map for patch application
	var docMap map[string]interface{}
	if err := json.Unmarshal(docJSON, &docMap); err != nil {
		return nil, fmt.Errorf("unmarshal to map: %w", err)
	}

	// Apply each patch operation
	// Note: We skip patches that reference non-existent paths since
	// variant patches may reference optional fields
	for _, op := range override.Patch {
		if err := applyPatch(docMap, op); err != nil {
			// Log but don't fail on path not found errors
			if !isPathNotFoundError(err) {
				return nil, fmt.Errorf("apply patch: %w", err)
			}
			// Skip patches for paths that don't exist
		}
	}

	// Marshal back to JSON
	patchedJSON, err := json.Marshal(docMap)
	if err != nil {
		return nil, fmt.Errorf("marshal patched document: %w", err)
	}

	// Unmarshal to Document struct
	var patchedDoc Document
	if err := json.Unmarshal(patchedJSON, &patchedDoc); err != nil {
		return nil, fmt.Errorf("unmarshal patched document: %w", err)
	}

	return &patchedDoc, nil
}

// applyPatch applies a single JSON Patch operation to a document map.
// Implements RFC 6902 JSON Patch operations.
func applyPatch(doc map[string]interface{}, op PatchOperation) error {
	switch op.Op {
	case "replace":
		return applyReplace(doc, op.Path, op.Value)
	case "add":
		return applyAdd(doc, op.Path, op.Value)
	case "remove":
		return applyRemove(doc, op.Path)
	default:
		return fmt.Errorf("unsupported patch operation: %s", op.Op)
	}
}

// isPathNotFoundError checks if an error is a "path not found" error.
func isPathNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "path not found")
}

// applyReplace replaces a value at the given JSON Pointer path.
func applyReplace(doc map[string]interface{}, path string, value interface{}) error {
	parent, key, err := navigateToParent(doc, path)
	if err != nil {
		return err
	}
	return setValue(parent, key, value)
}

// applyAdd adds a value at the given JSON Pointer path.
func applyAdd(doc map[string]interface{}, path string, value interface{}) error {
	parent, key, err := navigateToParent(doc, path)
	if err != nil {
		return err
	}
	return setValue(parent, key, value)
}

// applyRemove removes a value at the given JSON Pointer path.
func applyRemove(doc map[string]interface{}, path string) error {
	parent, key, err := navigateToParent(doc, path)
	if err != nil {
		return err
	}
	return removeValue(parent, key)
}

// setValue sets a value in either a map or an array.
func setValue(parent interface{}, key string, value interface{}) error {
	if m, ok := parent.(map[string]interface{}); ok {
		m[key] = value
		return nil
	}
	if arr, ok := parent.([]interface{}); ok {
		var idx int
		if _, err := fmt.Sscanf(key, "%d", &idx); err != nil {
			return fmt.Errorf("invalid array index %s: %w", key, err)
		}
		if idx < 0 || idx >= len(arr) {
			return fmt.Errorf("array index %d out of bounds", idx)
		}
		arr[idx] = value
		return nil
	}
	return fmt.Errorf("parent is neither map nor array: %T", parent)
}

// removeValue removes a value from either a map or an array.
func removeValue(parent interface{}, key string) error {
	if m, ok := parent.(map[string]interface{}); ok {
		delete(m, key)
		return nil
	}
	// Array removal is not commonly used in variant patches, but we could implement it if needed
	return fmt.Errorf("remove from array not implemented")
}

// navigateToParent navigates to the parent of the target location and returns
// the parent (map or array) and the final key.
func navigateToParent(doc map[string]interface{}, path string) (interface{}, string, error) {
	// JSON Pointer paths start with /
	if len(path) == 0 || path[0] != '/' {
		return nil, "", fmt.Errorf("invalid JSON Pointer: %s", path)
	}

	// Split path into components
	components := splitJSONPointer(path[1:])
	if len(components) == 0 {
		return nil, "", fmt.Errorf("empty path")
	}

	// Navigate to parent
	var current interface{} = doc
	for i := 0; i < len(components)-1; i++ {
		key := components[i]

		// Handle map access
		if m, ok := current.(map[string]interface{}); ok {
			next, exists := m[key]
			if !exists {
				return nil, "", fmt.Errorf("path not found: %s", path)
			}
			current = next
			continue
		}

		// Handle array access
		if arr, ok := current.([]interface{}); ok {
			// Parse array index
			var idx int
			if _, err := fmt.Sscanf(key, "%d", &idx); err != nil {
				return nil, "", fmt.Errorf("invalid array index %s: %w", key, err)
			}
			if idx < 0 || idx >= len(arr) {
				return nil, "", fmt.Errorf("array index %d out of bounds", idx)
			}
			current = arr[idx]
			continue
		}

		return nil, "", fmt.Errorf("path component %s: cannot navigate through %T", key, current)
	}

	finalKey := components[len(components)-1]
	return current, finalKey, nil
}

// splitJSONPointer splits a JSON Pointer path into components.
func splitJSONPointer(path string) []string {
	if path == "" {
		return nil
	}
	var components []string
	start := 0
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			components = append(components, path[start:i])
			start = i + 1
		}
	}
	if start < len(path) {
		components = append(components, path[start:])
	}
	return components
}

// GetDeclarationText extracts the declaration text from a document for the specified language.
func GetDeclarationText(doc *Document, lang LanguageVariant) string {
	// Try to get the variant-specific document
	variantDoc := doc
	if doc.Identifier.InterfaceLanguage != string(lang) {
		var err error
		if lang == LanguageSwift {
			variantDoc, err = GetSwiftVariant(doc)
		} else {
			variantDoc, err = GetObjectiveCVariant(doc)
		}
		if err != nil {
			return ""
		}
	}

	// Extract declaration from primary content sections
	for _, section := range variantDoc.PrimaryContentSections {
		if section.Kind == "declarations" {
			for _, decl := range section.Declarations {
				// Check if this declaration is for the target language
				for _, declLang := range decl.Languages {
					if (lang == LanguageSwift && declLang == "swift") ||
						(lang == LanguageObjectiveC && (declLang == "occ" || declLang == "objective-c")) {
						return renderDeclaration(decl.Tokens)
					}
				}
			}
		}
	}

	return ""
}

// renderDeclaration renders declaration tokens as a string.
func renderDeclaration(tokens []Token) string {
	var result string
	for _, token := range tokens {
		result += token.Text
	}
	return result
}

// CrossReference represents a cross-reference between Swift and Objective-C APIs.
type CrossReference struct {
	SwiftDeclaration string
	ObjCDeclaration  string
	SymbolKind       string
	Title            string
	Available        struct {
		Swift      bool
		ObjectiveC bool
	}
}

// GetCrossReference returns cross-reference information between Swift and Objective-C
// for the given symbol document.
func GetCrossReference(doc *Document) CrossReference {
	ref := CrossReference{
		SymbolKind: doc.Metadata.SymbolKind,
		Title:      doc.Metadata.Title,
	}

	ref.Available.Swift = HasSwiftVariant(doc)
	ref.Available.ObjectiveC = HasObjectiveCVariant(doc)

	if ref.Available.Swift {
		ref.SwiftDeclaration = GetDeclarationText(doc, LanguageSwift)
	}

	if ref.Available.ObjectiveC {
		ref.ObjCDeclaration = GetDeclarationText(doc, LanguageObjectiveC)
	}

	return ref
}
