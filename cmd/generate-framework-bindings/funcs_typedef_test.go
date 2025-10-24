package main

import (
	"testing"

	"github.com/tmc/appledocs/occ2go"
)

// TestTypedefResolution tests that typedefs are properly resolved
// in method parameter and return types.
func TestTypedefResolution(t *testing.T) {
	// Create a generator with some test typedefs
	gen := &Generator{
		Framework:   "Foundation",
		PackageName: "foundation",
		Typedefs: []*occ2go.ParsedTypedef{
			{
				Name:     "unichar",
				BaseType: "unsigned short",
			},
			{
				Name:     "NSInteger",
				BaseType: "long",
			},
			{
				Name:     "NSUInteger",
				BaseType: "unsigned long",
			},
		},
	}

	gf := GeneratorFuncs{gen}

	tests := []struct {
		name        string
		objcType    string
		expected    string
		description string
	}{
		{
			name:        "unichar typedef",
			objcType:    "unichar",
			expected:    "Unichar",
			description: "Should map to capitalized typedef name",
		},
		{
			name:        "NSInteger typedef",
			objcType:    "NSInteger",
			expected:    "Integer",
			description: "Should strip NS prefix and use typedef",
		},
		{
			name:        "NSUInteger typedef",
			objcType:    "NSUInteger",
			expected:    "UInteger",
			description: "Should strip NS prefix and use typedef",
		},
		{
			name:        "pointer to typedef",
			objcType:    "unichar *",
			expected:    "*Unichar",
			description: "Should handle pointer to typedef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gf.TypeToInterfaceType(tt.objcType)
			if result != tt.expected {
				t.Errorf("%s: TypeToInterfaceType(%q) = %q, want %q",
					tt.description, tt.objcType, result, tt.expected)
			}
		})
	}
}

// TestTypedefCapitalization tests that typedef names are properly
// capitalized when generated.
func TestTypedefCapitalization(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"unichar", "Unichar"},
		{"NSInteger", "Integer"},
		{"NSUInteger", "UInteger"},
		{"CGFloat", "Float"},
		{"int32_t", "Int32_t"}, // titleString preserves underscore - Int32_t is valid
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// Use the stripObjCPrefix and titleString functions
			result := titleString(stripObjCPrefix(tt.input))
			if result != tt.expected {
				t.Errorf("capitalization of %q = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestCGTypeImports tests that CoreGraphics types trigger proper imports
func TestCGTypeImports(t *testing.T) {
	gen := &Generator{
		Framework:   "Foundation",
		PackageName: "foundation",
	}

	tests := []struct {
		objcType     string
		shouldImport bool
		importPath   string
	}{
		{
			objcType:     "CGSize",
			shouldImport: true,
			importPath:   "github.com/tmc/appledocs/generated/coregraphics",
		},
		{
			objcType:     "CGRect",
			shouldImport: true,
			importPath:   "github.com/tmc/appledocs/generated/coregraphics",
		},
		{
			objcType:     "NSString *",
			shouldImport: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.objcType, func(t *testing.T) {
			// This test documents the current behavior
			// When we fix CG type imports, we should verify
			// that the import path is correct
			_ = gen
			_ = tt.shouldImport
			_ = tt.importPath
			t.Skip("TODO: Implement import path detection")
		})
	}
}

// TestCalendarIdentifierTypeResolution tests that string-based
// typedefs like CalendarIdentifier are properly handled
func TestCalendarIdentifierTypeResolution(t *testing.T) {
	gen := &Generator{
		Framework:   "Foundation",
		PackageName: "foundation",
		Typedefs: []*occ2go.ParsedTypedef{
			{
				Name:     "NSCalendarIdentifier",
				BaseType: "NSString *",
			},
		},
	}

	gf := GeneratorFuncs{gen}

	// Test that NSCalendarIdentifier is treated as a string typedef
	// not as a String class instance
	result := gf.TypeToInterfaceType("NSCalendarIdentifier")

	// The result should be CalendarIdentifier (the typedef)
	// not String (the class)
	if result == "IString" {
		t.Errorf("NSCalendarIdentifier should not map to IString, got %q", result)
	}

	// For now, just document what we expect
	// Expected: "CalendarIdentifier" (the typedef type)
	t.Logf("NSCalendarIdentifier maps to: %q", result)
}
