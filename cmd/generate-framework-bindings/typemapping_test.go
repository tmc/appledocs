package main

import (
	"testing"

	"github.com/tmc/appledocs/occ2go"
)

// TestLookupTypeMapping tests the type mapping resolution logic
func TestLookupTypeMapping(t *testing.T) {
	tests := []struct {
		name      string
		objcType  string
		framework string
		want      string
		wantFound bool
	}{
		// Foundation geometry types (cross-framework reference to CoreFoundation)
		{
			name:      "NSRect in Foundation",
			objcType:  "NSRect",
			framework: "Foundation",
			want:      "corefoundation.CGRect", // NSRect is CGRect from CoreFoundation
			wantFound: true,
		},
		{
			name:      "NSSize in Foundation",
			objcType:  "NSSize",
			framework: "Foundation",
			want:      "corefoundation.CGSize", // NSSize is CGSize from CoreFoundation
			wantFound: true,
		},
		{
			name:      "NSPoint in Foundation",
			objcType:  "NSPoint",
			framework: "Foundation",
			want:      "corefoundation.CGPoint", // NSPoint is CGPoint from CoreFoundation
			wantFound: true,
		},
		// CoreGraphics types in CoreGraphics (unqualified)
		{
			name:      "CGRect in CoreGraphics",
			objcType:  "CGRect",
			framework: "CoreGraphics",
			want:      "Rect", // Uses local alias in CoreGraphics
			wantFound: true,
		},
		{
			name:      "CGSize in CoreGraphics",
			objcType:  "CGSize",
			framework: "CoreGraphics",
			want:      "Size", // Uses local alias in CoreGraphics
			wantFound: true,
		},
		{
			name:      "CGPoint in CoreGraphics",
			objcType:  "CGPoint",
			framework: "CoreGraphics",
			want:      "Point", // Uses local alias in CoreGraphics
			wantFound: true,
		},
		// AppKit geometry types (cross-framework reference to CoreFoundation)
		{
			name:      "NSRect in AppKit",
			objcType:  "NSRect",
			framework: "AppKit",
			want:      "corefoundation.CGRect", // NSRect is CGRect from CoreFoundation
			wantFound: true,
		},
		{
			name:      "CGRect in AppKit",
			objcType:  "CGRect",
			framework: "AppKit",
			want:      "corefoundation.CGRect", // Cross-framework ref to CoreFoundation
			wantFound: true,
		},
		// ScreenSaver geometry types
		{
			name:      "Rect in ScreenSaver",
			objcType:  "Rect",
			framework: "ScreenSaver",
			want:      "",
			wantFound: false, // "Rect" alone is not in registry
		},
		{
			name:      "NSRect in ScreenSaver",
			objcType:  "NSRect",
			framework: "ScreenSaver",
			want:      "corefoundation.CGRect", // NSRect → corefoundation.CGRect
			wantFound: true,
		},
		{
			name:      "CGRect in ScreenSaver",
			objcType:  "CGRect",
			framework: "ScreenSaver",
			want:      "corefoundation.CGRect", // Cross-framework ref
			wantFound: true,
		},
		// ScreenSaver AppKit types
		{
			name:      "BackingStoreType in ScreenSaver",
			objcType:  "BackingStoreType",
			framework: "ScreenSaver",
			want:      "",
			wantFound: false, // Not in lookupTypeMapping (handled by type registry)
		},
		{
			name:      "NSBackingStoreType in ScreenSaver",
			objcType:  "NSBackingStoreType",
			framework: "ScreenSaver",
			want:      "BackingStoreType", // Stripped
			wantFound: true,
		},
		// AppKit enum types
		{
			name:      "NSWindowStyleMask in AppKit",
			objcType:  "NSWindowStyleMask",
			framework: "AppKit",
			want:      "WindowStyleMask",
			wantFound: true,
		},
		{
			name:      "NSBackingStoreType in AppKit",
			objcType:  "NSBackingStoreType",
			framework: "AppKit",
			want:      "BackingStoreType",
			wantFound: true,
		},
		// Foundation time types
		{
			name:      "NSTimeInterval in Foundation",
			objcType:  "NSTimeInterval",
			framework: "Foundation",
			want:      "float64", // Maps to float64 via static registry
			wantFound: true,
		},
		{
			name:      "NSTimeInterval in AppKit",
			objcType:  "NSTimeInterval",
			framework: "AppKit",
			want:      "float64",
			wantFound: true,
		},
		// Foundation edge enum
		{
			name:      "NSRectEdge in Foundation",
			objcType:  "NSRectEdge",
			framework: "Foundation",
			want:      "RectEdge",
			wantFound: true,
		},
		{
			name:      "NSRectEdge in AppKit",
			objcType:  "NSRectEdge",
			framework: "AppKit",
			want:      "RectEdge", // Stripped to RectEdge
			wantFound: true,
		},
		// CoreGraphics opaque ref types
		{
			name:      "CGColorRef in CoreGraphics",
			objcType:  "CGColorRef",
			framework: "CoreGraphics",
			want:      "ColorRef", // Stripped in CoreGraphics
			wantFound: true,
		},
		{
			name:      "CGImageRef in AppKit",
			objcType:  "CGImageRef",
			framework: "AppKit",
			want:      "ImageRef", // Stripped
			wantFound: true,
		},
		// Generic types (not in lookupTypeMapping - handled by mapObjCTypeToGo)
		{
			name:      "id in any framework",
			objcType:  "id",
			framework: "AppKit",
			want:      "",
			wantFound: false, // Handled by occ2go.MapCTypeToGo, not in static registry
		},
		{
			name:      "NSArray * in any framework",
			objcType:  "NSArray *",
			framework: "Foundation",
			want:      "",
			wantFound: false, // Handled by occ2go.MapCTypeToGo
		},
		// Pointer stripping
		{
			name:      "NSString * pointer stripped",
			objcType:  "NSString *",
			framework: "AppKit",
			want:      "",
			wantFound: false, // Handled by mapObjCTypeToGo (NSString * → string)
		},
		// Type not found
		{
			name:      "unknown type",
			objcType:  "UnknownType",
			framework: "AppKit",
			want:      "",
			wantFound: false,
		},
		// Cross-framework types
		{
			name:      "CGRect in ScreenCaptureKit",
			objcType:  "CGRect",
			framework: "ScreenCaptureKit",
			want:      "corefoundation.CGRect", // Geometry types defined in CoreFoundation
			wantFound: true,
		},
		{
			name:      "CGRect in CoreImage",
			objcType:  "CGRect",
			framework: "CoreImage",
			want:      "corefoundation.CGRect", // Geometry types defined in CoreFoundation
			wantFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := lookupTypeMapping(tt.objcType, tt.framework)
			if found != tt.wantFound {
				t.Errorf("lookupTypeMapping(%q, %q) found = %v, want %v", tt.objcType, tt.framework, found, tt.wantFound)
			}
			if got != tt.want {
				t.Errorf("lookupTypeMapping(%q, %q) = %q, want %q", tt.objcType, tt.framework, got, tt.want)
			}
		})
	}
}

// TestLookupTypeMappingPrecedence tests that framework-specific mappings take precedence
func TestLookupTypeMappingPrecedence(t *testing.T) {
	// NSRect should map the same in Foundation and AppKit (both to corefoundation.CGRect)
	foundationType, foundFoundation := lookupTypeMapping("NSRect", "Foundation")
	appKitType, foundAppKit := lookupTypeMapping("NSRect", "AppKit")

	if !foundFoundation || !foundAppKit {
		t.Fatal("NSRect should be found in both Foundation and AppKit")
	}

	// Both should map to "corefoundation.CGRect" (cross-framework reference)
	if foundationType != "corefoundation.CGRect" {
		t.Errorf("NSRect in Foundation should be 'corefoundation.CGRect', got %q", foundationType)
	}

	if appKitType != "corefoundation.CGRect" {
		t.Errorf("NSRect in AppKit should be 'corefoundation.CGRect', got %q", appKitType)
	}
}

// TestLookupTypeMappingDetails tests the detailed type mapping lookup
func TestLookupTypeMappingDetails(t *testing.T) {
	tests := []struct {
		name       string
		objcType   string
		framework  string
		wantGoType string
		wantNil    bool
	}{
		{
			name:       "NSRect in Foundation returns details",
			objcType:   "NSRect",
			framework:  "Foundation",
			wantGoType: "",
			wantNil:    true, // NSRect is handled by stripping logic, not in registry
		},
		{
			name:       "NSRect in AppKit returns details",
			objcType:   "NSRect",
			framework:  "AppKit",
			wantGoType: "",
			wantNil:    true, // NSRect is handled by stripping logic, not in registry
		},
		{
			name:       "Unknown type returns nil",
			objcType:   "UnknownType",
			framework:  "AppKit",
			wantGoType: "",
			wantNil:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lookupTypeMappingDetails(tt.objcType, tt.framework)
			if tt.wantNil {
				if got != nil {
					t.Errorf("lookupTypeMappingDetails(%q, %q) = %+v, want nil", tt.objcType, tt.framework, got)
				}
				return
			}
			if got == nil {
				t.Fatalf("lookupTypeMappingDetails(%q, %q) = nil, want non-nil", tt.objcType, tt.framework)
			}
			if got.GoType != tt.wantGoType {
				t.Errorf("lookupTypeMappingDetails(%q, %q).GoType = %q, want %q", tt.objcType, tt.framework, got.GoType, tt.wantGoType)
			}
		})
	}
}

// TestGetTypeImportPath tests import path resolution
func TestGetTypeImportPath(t *testing.T) {
	tests := []struct {
		name      string
		objcType  string
		framework string
		want      string
	}{
		{
			name:      "simple type no import",
			objcType:  "NSRect",
			framework: "Foundation",
			want:      "",
		},
		{
			name:      "unknown type no import",
			objcType:  "UnknownType",
			framework: "AppKit",
			want:      "",
		},
		// Most types don't require imports in our current setup
		// as we use qualified package names directly
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTypeImportPath(tt.objcType, tt.framework)
			if got != tt.want {
				t.Errorf("getTypeImportPath(%q, %q) = %q, want %q", tt.objcType, tt.framework, got, tt.want)
			}
		})
	}
}

// TestGetAllAppKitEnumTypes tests that we can retrieve all AppKit enum types
func TestGetAllAppKitEnumTypes(t *testing.T) {
	// Save and restore original registries
	oldRegistry := typeRegistry
	oldCrossFrameworkRegistry := make(map[string]string)
	for k, v := range crossFrameworkTypeRegistry {
		oldCrossFrameworkRegistry[k] = v
	}
	defer func() {
		typeRegistry = oldRegistry
		crossFrameworkTypeRegistry = oldCrossFrameworkRegistry
	}()

	// Clear the cross-framework registry for a clean test
	crossFrameworkTypeRegistry = make(map[string]string)

	// Populate test data - simulate what buildTypeRegistryFromParsedData does
	appkitEnums := []*occ2go.ParsedEnum{
		{Name: "NSWindowStyleMask"},
		{Name: "NSBackingStoreType"},
		{Name: "NSWindowOrderingMode"},
		{Name: "NSWindowLevel"},
		{Name: "NSEventType"},
		{Name: "NSEventModifierFlags"},
	}
	buildTypeRegistryFromParsedData("AppKit", nil, appkitEnums, nil)

	enums := getAllAppKitEnumTypes()

	if len(enums) == 0 {
		t.Error("getAllAppKitEnumTypes() returned empty list")
	}

	// Check for known AppKit enum types
	expectedEnums := []string{
		"WindowStyleMask",
		"BackingStoreType",
		"WindowOrderingMode",
		"WindowLevel",
		"EventType",
		"EventModifierFlags",
	}

	found := make(map[string]bool)
	for _, enum := range enums {
		found[enum] = true
	}

	for _, expected := range expectedEnums {
		if !found[expected] {
			t.Errorf("getAllAppKitEnumTypes() missing expected enum: %q", expected)
		}
	}

	// Ensure no duplicates
	if len(enums) != len(found) {
		t.Errorf("getAllAppKitEnumTypes() returned duplicates: got %d enums, %d unique", len(enums), len(found))
	}
}

// TestGetAllMappedTypes tests that we can retrieve all type mappings
func TestGetAllMappedTypes(t *testing.T) {
	// Save and restore original registries
	oldRegistry := typeRegistry
	oldCrossFrameworkRegistry := make(map[string]string)
	for k, v := range crossFrameworkTypeRegistry {
		oldCrossFrameworkRegistry[k] = v
	}
	defer func() {
		typeRegistry = oldRegistry
		crossFrameworkTypeRegistry = oldCrossFrameworkRegistry
	}()

	// Clear the cross-framework registry for a clean test
	crossFrameworkTypeRegistry = make(map[string]string)

	// Populate test data - simulate what buildTypeRegistryFromParsedData does
	foundationClasses := []*occ2go.ParsedClass{
		{Name: "NSRect"}, // This will create NSRect → Rect mapping for Foundation
	}
	screenSaverClasses := []*occ2go.ParsedClass{
		{Name: "ScreenSaverView"},
	}

	buildTypeRegistryFromParsedData("Foundation", foundationClasses, nil, nil)
	buildTypeRegistryFromParsedData("ScreenSaver", screenSaverClasses, nil, nil)

	mappings := getAllMappedTypes()

	if len(mappings) == 0 {
		t.Error("getAllMappedTypes() returned empty list")
	}

	// Check for some known mappings
	// Note: NSRect gets stripped to "Rect" in Foundation framework
	// Framework names are stored in lowercase in crossFrameworkTypeRegistry
	found := false
	for _, mapping := range mappings {
		if mapping.ObjCType == "NSRect" && mapping.Framework == "foundation" && mapping.GoType == "Rect" {
			found = true
			break
		}
	}

	if !found {
		t.Error("getAllMappedTypes() did not include Foundation NSRect mapping")
	}

	// Check for ScreenSaver mappings (recently added)
	// Framework names are stored in lowercase
	screenSaverFound := false
	for _, mapping := range mappings {
		if mapping.Framework == "screensaver" {
			screenSaverFound = true
			break
		}
	}

	if !screenSaverFound {
		t.Error("getAllMappedTypes() did not include any ScreenSaver framework mappings")
	}
}

// TestIsFrameworkLocalType tests framework-local type detection
func TestIsFrameworkLocalType(t *testing.T) {
	tests := []struct {
		name       string
		importPath string
		framework  string
		want       bool
	}{
		{
			name:       "foundation import for Foundation framework",
			importPath: "github.com/progrium/darwinkit/macos/foundation",
			framework:  "Foundation",
			want:       true,
		},
		{
			name:       "appkit import for AppKit framework",
			importPath: "github.com/progrium/darwinkit/macos/appkit",
			framework:  "AppKit",
			want:       true,
		},
		{
			name:       "foundation import for AppKit framework",
			importPath: "github.com/progrium/darwinkit/macos/foundation",
			framework:  "AppKit",
			want:       false,
		},
		{
			name:       "empty import path",
			importPath: "",
			framework:  "Foundation",
			want:       false,
		},
		{
			name:       "case insensitive matching",
			importPath: "github.com/progrium/darwinkit/macos/Foundation",
			framework:  "foundation",
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isFrameworkLocalType(tt.importPath, tt.framework)
			if got != tt.want {
				t.Errorf("isFrameworkLocalType(%q, %q) = %v, want %v", tt.importPath, tt.framework, got, tt.want)
			}
		})
	}
}

// TestTypeMappingConsistency tests that related types are mapped consistently
func TestTypeMappingConsistency(t *testing.T) {
	// Test that all geometry types in the same framework use consistent prefixes
	frameworks := []string{"AppKit", "ScreenSaver", "CoreImage", "ScreenCaptureKit"}

	for _, framework := range frameworks {
		rect, foundRect := lookupTypeMapping("CGRect", framework)
		size, foundSize := lookupTypeMapping("CGSize", framework)
		point, foundPoint := lookupTypeMapping("CGPoint", framework)

		if !foundRect || !foundSize || !foundPoint {
			t.Errorf("Framework %s missing geometry type mappings", framework)
			continue
		}

		// All should use the same prefix (corefoundation. where geometry types are defined)
		if rect != "corefoundation.CGRect" {
			t.Errorf("CGRect in %s = %q, expected 'corefoundation.CGRect'", framework, rect)
		}
		if size != "corefoundation.CGSize" {
			t.Errorf("CGSize in %s = %q, expected 'corefoundation.CGSize'", framework, size)
		}
		if point != "corefoundation.CGPoint" {
			t.Errorf("CGPoint in %s = %q, expected 'corefoundation.CGPoint'", framework, point)
		}
	}
}

// TestCrossFrameworkTypeConsistency tests that cross-framework references are consistent
func TestCrossFrameworkTypeConsistency(t *testing.T) {
	// All non-CoreGraphics, non-CoreFoundation frameworks should reference CGRect as corefoundation.CGRect
	frameworks := []string{"AppKit", "ScreenSaver", "CoreImage", "ScreenCaptureKit", "CoreText", "CoreVideo"}

	for _, framework := range frameworks {
		got, found := lookupTypeMapping("CGRect", framework)
		if !found {
			t.Errorf("CGRect not found in %s", framework)
			continue
		}
		if got != "corefoundation.CGRect" {
			t.Errorf("CGRect in %s = %q, want 'corefoundation.CGRect'", framework, got)
		}
	}

	// CoreGraphics itself should use local alias "Rect" (stripped prefix)
	got, found := lookupTypeMapping("CGRect", "CoreGraphics")
	if !found {
		t.Fatal("CGRect not found in CoreGraphics")
	}
	if got != "Rect" {
		t.Errorf("CGRect in CoreGraphics = %q, want 'Rect'", got)
	}
}

// TestBlockTypeMappings tests that block/closure types are mapped correctly
func TestBlockTypeMappings(t *testing.T) {
	tests := []struct {
		name      string
		objcType  string
		framework string
		want      string
	}{
		{
			name:      "void completion block",
			objcType:  "void (^)(void)",
			framework: "AppKit",
			want:      "func()",
		},
		{
			name:      "error completion block",
			objcType:  "void (^)(NSError *)",
			framework: "Foundation",
			want:      "func(unsafe.Pointer)", // NSError * → unsafe.Pointer (no parameter names)
		},
		{
			name:      "bool completion block",
			objcType:  "void (^)(BOOL)",
			framework: "AppKit",
			want:      "func(bool)", // No parameter names in type signature
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test via mapObjCTypeToGo which calls occ2go.MapCTypeToGo for block types
			got := mapObjCTypeToGo(tt.objcType, tt.framework)
			if got != tt.want {
				t.Errorf("mapObjCTypeToGo(%q, %q) = %q, want %q", tt.objcType, tt.framework, got, tt.want)
			}
		})
	}
}
