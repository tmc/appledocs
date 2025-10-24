package main

import (
	"testing"

	"github.com/tmc/appledocs/occ2go"
)

func TestBuildTypeRegistryFromParsedData(t *testing.T) {
	// Clear the global registry before test
	crossFrameworkTypeRegistry = make(map[string]string)

	// Create test data
	classes := []*occ2go.ParsedClass{
		{Name: "NSWindow"},
		{Name: "NSButton"},
		{Name: "NSView"},
	}

	enums := []*occ2go.ParsedEnum{
		{Name: "NSImageScaling"},
		{Name: "NSWindowStyleMask"},
		{Name: "NSBackingStoreType"},
	}

	typedefs := []*occ2go.ParsedTypedef{
		{Name: "NSTimeInterval"},
		{Name: "NSInteger"},
	}

	// Build registry from parsed data
	buildTypeRegistryFromParsedData("AppKit", classes, enums, typedefs)

	// Test enum registrations (both ObjC and Go names)
	tests := []struct {
		name     string
		key      string
		expected string
	}{
		// Enum types - ObjC names
		{"NSImageScaling ObjC name", "NSImageScaling", "appkit"},
		{"NSWindowStyleMask ObjC name", "NSWindowStyleMask", "appkit"},
		{"NSBackingStoreType ObjC name", "NSBackingStoreType", "appkit"},

		// Enum types - Go names
		{"ImageScaling Go name", "ImageScaling", "appkit"},
		{"WindowStyleMask Go name", "WindowStyleMask", "appkit"},
		{"BackingStoreType Go name", "BackingStoreType", "appkit"},

		// Class types - ObjC names
		{"NSWindow ObjC name", "NSWindow", "appkit"},
		{"NSButton ObjC name", "NSButton", "appkit"},
		{"NSView ObjC name", "NSView", "appkit"},

		// Class types - Go names
		{"Window Go name", "Window", "appkit"},
		{"Button Go name", "Button", "appkit"},
		{"View Go name", "View", "appkit"},

		// Typedef types - ObjC names
		{"NSTimeInterval ObjC name", "NSTimeInterval", "appkit"},
		{"NSInteger ObjC name", "NSInteger", "appkit"},

		// Typedef types - Go names
		{"TimeInterval Go name", "TimeInterval", "appkit"},
		{"Integer Go name", "Integer", "appkit"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := crossFrameworkTypeRegistry[tt.key]
			if !found {
				t.Errorf("Type %q not found in registry", tt.key)
				return
			}
			if got != tt.expected {
				t.Errorf("crossFrameworkTypeRegistry[%q] = %q, want %q", tt.key, got, tt.expected)
			}
		})
	}

	// Verify total count
	expectedCount := len(classes)*2 + len(enums)*2 + len(typedefs)*2 // Each registered twice (ObjC + Go name)
	if len(crossFrameworkTypeRegistry) != expectedCount {
		t.Errorf("Registry has %d entries, expected %d", len(crossFrameworkTypeRegistry), expectedCount)
	}
}

func TestTypeRegistryMultipleFrameworks(t *testing.T) {
	// Clear the global registry
	crossFrameworkTypeRegistry = make(map[string]string)

	// Register AppKit types
	appkitEnums := []*occ2go.ParsedEnum{
		{Name: "NSImageScaling"},
	}
	buildTypeRegistryFromParsedData("AppKit", nil, appkitEnums, nil)

	// Register Foundation types
	foundationClasses := []*occ2go.ParsedClass{
		{Name: "NSURL"},
		{Name: "NSString"},
	}
	buildTypeRegistryFromParsedData("Foundation", foundationClasses, nil, nil)

	tests := []struct {
		name     string
		key      string
		expected string
	}{
		{"AppKit enum", "ImageScaling", "appkit"},
		{"Foundation class ObjC", "NSURL", "foundation"},
		{"Foundation class Go", "URL", "foundation"},
		{"Foundation string ObjC", "NSString", "foundation"},
		{"Foundation string Go", "String", "foundation"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := crossFrameworkTypeRegistry[tt.key]
			if !found {
				t.Errorf("Type %q not found in registry", tt.key)
				return
			}
			if got != tt.expected {
				t.Errorf("crossFrameworkTypeRegistry[%q] = %q, want %q", tt.key, got, tt.expected)
			}
		})
	}
}

func TestTypeRegistryCoreGraphicsTypes(t *testing.T) {
	// Clear the global registry
	crossFrameworkTypeRegistry = make(map[string]string)

	// CoreGraphics types should NOT strip CG prefix
	cgClasses := []*occ2go.ParsedClass{
		{Name: "CGContext"},
		{Name: "CGImage"},
	}
	cgTypedefs := []*occ2go.ParsedTypedef{
		{Name: "CGFloat"},
	}
	buildTypeRegistryFromParsedData("CoreGraphics", cgClasses, nil, cgTypedefs)

	tests := []struct {
		name     string
		key      string
		expected string
	}{
		{"CGContext keeps CG prefix", "CGContext", "coregraphics"},
		{"CGImage keeps CG prefix", "CGImage", "coregraphics"},
		{"CGFloat typedef", "CGFloat", "coregraphics"},
		// Note: StripObjCPrefix should keep CG prefix for CoreGraphics types
		{"CGContext without CG", "Context", "coregraphics"},
		{"CGImage without CG", "Image", "coregraphics"},
		{"CGFloat without CG", "Float", "coregraphics"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := crossFrameworkTypeRegistry[tt.key]
			if !found {
				t.Errorf("Type %q not found in registry", tt.key)
				return
			}
			if got != tt.expected {
				t.Errorf("crossFrameworkTypeRegistry[%q] = %q, want %q", tt.key, got, tt.expected)
			}
		})
	}
}

func TestResolveTypeWithRegistry(t *testing.T) {
	// Clear and populate registry
	crossFrameworkTypeRegistry = make(map[string]string)

	appkitClasses := []*occ2go.ParsedClass{
		{Name: "NSWindow"},
		{Name: "NSColor"},
	}
	appkitEnums := []*occ2go.ParsedEnum{
		{Name: "NSImageScaling"},
	}
	foundationClasses := []*occ2go.ParsedClass{
		{Name: "NSURL"},
	}

	buildTypeRegistryFromParsedData("AppKit", appkitClasses, appkitEnums, nil)
	buildTypeRegistryFromParsedData("Foundation", foundationClasses, nil, nil)

	tests := []struct {
		name       string
		framework  string
		typeName   string
		want       string
		wantReason string
	}{
		{
			name:       "Same framework - unqualified",
			framework:  "AppKit",
			typeName:   "Window",
			want:       "Window",
			wantReason: "Window is in AppKit, so return unqualified",
		},
		{
			name:       "Cross framework - resolves with qualification",
			framework:  "AppKit",
			typeName:   "URL",
			want:       "foundation.URL", // URL is from Foundation, qualified in AppKit
			wantReason: "URL from Foundation is qualified when used in AppKit",
		},
		{
			name:       "Enum same framework",
			framework:  "AppKit",
			typeName:   "ImageScaling",
			want:       "ImageScaling",
			wantReason: "ImageScaling is in AppKit, return unqualified",
		},
		{
			name:       "Foundation type in Foundation",
			framework:  "Foundation",
			typeName:   "URL",
			want:       "URL",
			wantReason: "URL in its own framework",
		},
		{
			name:       "Registry lookup for non-hardcoded type",
			framework:  "ScreenSaver",
			typeName:   "Window",
			want:       "appkit.Window",
			wantReason: "Window not in hardcoded maps, uses registry",
		},
		{
			name:       "TextField.TextColor should resolve to Color (issue: appledocs-437)",
			framework:  "AppKit",
			typeName:   "Color",
			want:       "Color",
			wantReason: "Color is in AppKit, return unqualified",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveType(tt.framework, tt.typeName)
			if got != tt.want {
				t.Errorf("resolveType(%q, %q) = %q, want %q (%s)",
					tt.framework, tt.typeName, got, tt.want, tt.wantReason)
			}
		})
	}
}
