package main

import (
	"testing"
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
		// Foundation geometry types (unqualified)
		{
			name:      "NSRect in Foundation",
			objcType:  "NSRect",
			framework: "Foundation",
			want:      "Rect",
			wantFound: true,
		},
		{
			name:      "NSSize in Foundation",
			objcType:  "NSSize",
			framework: "Foundation",
			want:      "Size",
			wantFound: true,
		},
		{
			name:      "NSPoint in Foundation",
			objcType:  "NSPoint",
			framework: "Foundation",
			want:      "Point",
			wantFound: true,
		},
		// CoreGraphics types in CoreGraphics (unqualified)
		{
			name:      "CGRect in CoreGraphics",
			objcType:  "CGRect",
			framework: "CoreGraphics",
			want:      "CGRect",
			wantFound: true,
		},
		{
			name:      "CGSize in CoreGraphics",
			objcType:  "CGSize",
			framework: "CoreGraphics",
			want:      "CGSize",
			wantFound: true,
		},
		{
			name:      "CGPoint in CoreGraphics",
			objcType:  "CGPoint",
			framework: "CoreGraphics",
			want:      "CGPoint",
			wantFound: true,
		},
		// AppKit geometry types (qualified with coregraphics)
		{
			name:      "NSRect in AppKit",
			objcType:  "NSRect",
			framework: "AppKit",
			want:      "coregraphics.CGRect",
			wantFound: true,
		},
		{
			name:      "CGRect in AppKit",
			objcType:  "CGRect",
			framework: "AppKit",
			want:      "coregraphics.CGRect",
			wantFound: true,
		},
		// ScreenSaver geometry types (qualified with coregraphics)
		{
			name:      "Rect in ScreenSaver",
			objcType:  "Rect",
			framework: "ScreenSaver",
			want:      "coregraphics.CGRect",
			wantFound: true,
		},
		{
			name:      "NSRect in ScreenSaver",
			objcType:  "NSRect",
			framework: "ScreenSaver",
			want:      "coregraphics.CGRect",
			wantFound: true,
		},
		{
			name:      "CGRect in ScreenSaver",
			objcType:  "CGRect",
			framework: "ScreenSaver",
			want:      "coregraphics.CGRect",
			wantFound: true,
		},
		// ScreenSaver AppKit types
		{
			name:      "BackingStoreType in ScreenSaver",
			objcType:  "BackingStoreType",
			framework: "ScreenSaver",
			want:      "appkit.BackingStoreType",
			wantFound: true,
		},
		{
			name:      "NSBackingStoreType in ScreenSaver",
			objcType:  "NSBackingStoreType",
			framework: "ScreenSaver",
			want:      "appkit.BackingStoreType",
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
			want:      "TimeInterval",
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
			want:      "int",
			wantFound: true,
		},
		// CoreGraphics opaque ref types
		{
			name:      "CGColorRef in CoreGraphics",
			objcType:  "CGColorRef",
			framework: "CoreGraphics",
			want:      "CGColorRef",
			wantFound: true,
		},
		{
			name:      "CGImageRef in AppKit",
			objcType:  "CGImageRef",
			framework: "AppKit",
			want:      "coregraphics.CGImageRef",
			wantFound: true,
		},
		// Generic types (framework-agnostic)
		{
			name:      "id in any framework",
			objcType:  "id",
			framework: "AppKit",
			want:      "objc.ID",
			wantFound: true,
		},
		{
			name:      "NSArray * in any framework",
			objcType:  "NSArray *",
			framework: "Foundation",
			want:      "objc.ID",
			wantFound: true,
		},
		// Pointer stripping
		{
			name:      "NSString * pointer stripped",
			objcType:  "NSString *",
			framework: "AppKit",
			want:      "string",
			wantFound: true,
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
			want:      "coregraphics.CGRect",
			wantFound: true,
		},
		{
			name:      "CGRect in CoreImage",
			objcType:  "CGRect",
			framework: "CoreImage",
			want:      "coregraphics.CGRect",
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
	// NSRect should map differently in Foundation vs AppKit
	foundationType, foundFoundation := lookupTypeMapping("NSRect", "Foundation")
	appKitType, foundAppKit := lookupTypeMapping("NSRect", "AppKit")

	if !foundFoundation || !foundAppKit {
		t.Fatal("NSRect should be found in both Foundation and AppKit")
	}

	if foundationType == appKitType {
		t.Errorf("NSRect should have different types in Foundation (%q) vs AppKit (%q)", foundationType, appKitType)
	}

	// Foundation should have unqualified type
	if foundationType != "Rect" {
		t.Errorf("NSRect in Foundation should be 'Rect', got %q", foundationType)
	}

	// AppKit should have qualified type
	if appKitType != "coregraphics.CGRect" {
		t.Errorf("NSRect in AppKit should be 'coregraphics.CGRect', got %q", appKitType)
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
			wantGoType: "Rect",
			wantNil:    false,
		},
		{
			name:       "NSRect in AppKit returns details",
			objcType:   "NSRect",
			framework:  "AppKit",
			wantGoType: "coregraphics.CGRect",
			wantNil:    false,
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
	mappings := getAllMappedTypes()

	if len(mappings) == 0 {
		t.Error("getAllMappedTypes() returned empty list")
	}

	// Check for some known mappings
	found := false
	for _, mapping := range mappings {
		if mapping.ObjCType == "NSRect" && mapping.Framework == "Foundation" && mapping.GoType == "Rect" {
			found = true
			break
		}
	}

	if !found {
		t.Error("getAllMappedTypes() did not include Foundation NSRect mapping")
	}

	// Check for ScreenSaver mappings (recently added)
	screenSaverFound := false
	for _, mapping := range mappings {
		if mapping.Framework == "ScreenSaver" {
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

		// All should use the same prefix (coregraphics.)
		if rect != "coregraphics.CGRect" {
			t.Errorf("CGRect in %s = %q, expected 'coregraphics.CGRect'", framework, rect)
		}
		if size != "coregraphics.CGSize" {
			t.Errorf("CGSize in %s = %q, expected 'coregraphics.CGSize'", framework, size)
		}
		if point != "coregraphics.CGPoint" {
			t.Errorf("CGPoint in %s = %q, expected 'coregraphics.CGPoint'", framework, point)
		}
	}
}

// TestCrossFrameworkTypeConsistency tests that cross-framework references are consistent
func TestCrossFrameworkTypeConsistency(t *testing.T) {
	// All non-CoreGraphics frameworks should reference CGRect as coregraphics.CGRect
	frameworks := []string{"AppKit", "ScreenSaver", "CoreImage", "ScreenCaptureKit", "CoreText", "CoreVideo"}

	for _, framework := range frameworks {
		got, found := lookupTypeMapping("CGRect", framework)
		if !found {
			t.Errorf("CGRect not found in %s", framework)
			continue
		}
		if got != "coregraphics.CGRect" {
			t.Errorf("CGRect in %s = %q, want 'coregraphics.CGRect'", framework, got)
		}
	}

	// CoreGraphics itself should use unqualified CGRect
	got, found := lookupTypeMapping("CGRect", "CoreGraphics")
	if !found {
		t.Fatal("CGRect not found in CoreGraphics")
	}
	if got != "CGRect" {
		t.Errorf("CGRect in CoreGraphics = %q, want 'CGRect'", got)
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
			want:      "func(error objc.ID)",
		},
		{
			name:      "bool completion block",
			objcType:  "void (^)(BOOL)",
			framework: "AppKit",
			want:      "func(success bool)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := lookupTypeMapping(tt.objcType, tt.framework)
			if !found {
				t.Errorf("lookupTypeMapping(%q, %q) not found", tt.objcType, tt.framework)
				return
			}
			if got != tt.want {
				t.Errorf("lookupTypeMapping(%q, %q) = %q, want %q", tt.objcType, tt.framework, got, tt.want)
			}
		})
	}
}
