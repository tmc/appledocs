package main

import (
	"strings"
	"testing"

	"github.com/tmc/appledocs/occ2go"
)

// TestDetermineRequiredImports tests that we correctly identify required imports
func TestDetermineRequiredImports(t *testing.T) {
	tests := []struct {
		name            string
		methods         []MethodInfo
		framework       string
		wantImports     []string
		dontWantImports []string
	}{
		{
			name: "ScreenSaver with CGRect parameters",
			methods: []MethodInfo{
				{
					Selector:   "initWithFrame:isPreview:",
					ReturnType: "instancetype",
					Parameters: []occ2go.Parameter{
						{Type: "Rect"},
						{Type: "bool"},
					},
				},
			},
			framework: "ScreenSaver",
			wantImports: []string{
				"github.com/tmc/appledocs/generated/coregraphics",
				"github.com/tmc/appledocs/generated/objc",
			},
		},
		{
			name: "ScreenSaver with BackingStoreType",
			methods: []MethodInfo{
				{
					Selector:   "backingStoreType",
					ReturnType: "BackingStoreType",
					Parameters: []occ2go.Parameter{},
				},
			},
			framework: "ScreenSaver",
			wantImports: []string{
				"github.com/tmc/appledocs/generated/appkit",
				"github.com/tmc/appledocs/generated/objc",
			},
		},
		{
			name: "AppKit with NSRect - no qualified import needed",
			methods: []MethodInfo{
				{
					Selector:   "initWithFrame:",
					ReturnType: "instancetype",
					Parameters: []occ2go.Parameter{
						{Type: "NSRect"},
					},
				},
			},
			framework: "AppKit",
			wantImports: []string{
				"github.com/tmc/appledocs/generated/coregraphics",
				"github.com/tmc/appledocs/generated/objc",
			},
			// AppKit shouldn't import itself
			dontWantImports: []string{
				"github.com/tmc/appledocs/generated/appkit",
			},
		},
		{
			name: "Foundation with basic types only",
			methods: []MethodInfo{
				{
					Selector:   "count",
					ReturnType: "int",
					Parameters: []occ2go.Parameter{},
				},
			},
			framework: "Foundation",
			wantImports: []string{
				"github.com/tmc/appledocs/generated/objc",
			},
			dontWantImports: []string{
				"github.com/tmc/appledocs/generated/coregraphics",
				"github.com/tmc/appledocs/generated/appkit",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a class with the methods
			classInfo := ClassInfo{
				Name:    "TestClass",
				Methods: tt.methods,
			}

			// Use our import determination logic
			imports := determineImports(classInfo, tt.framework)

			// Check for expected imports
			for _, want := range tt.wantImports {
				found := false
				for _, imp := range imports {
					if imp == want {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected import %q not found in %v", want, imports)
				}
			}

			// Check that unwanted imports are not present
			for _, dontWant := range tt.dontWantImports {
				for _, imp := range imports {
					if imp == dontWant {
						t.Errorf("Unexpected import %q found in %v", dontWant, imports)
					}
				}
			}
		})
	}
}

// determineImports is a helper that extracts imports from a ClassInfo
// This mirrors the logic in the actual generator
func determineImports(class ClassInfo, framework string) []string {
	imports := make(map[string]bool)

	// Always need objc
	imports["github.com/tmc/appledocs/generated/objc"] = true

	// Check all method parameters and return types
	for _, method := range class.Methods {
		// Check return type
		if goType, found := lookupTypeMapping(method.ReturnType, framework); found {
			if strings.Contains(goType, ".") {
				// Qualified type - need import
				parts := strings.Split(goType, ".")
				if len(parts) == 2 {
					pkg := parts[0]
					// Don't import self
					if !strings.EqualFold(pkg, framework) {
						imports["github.com/tmc/appledocs/generated/"+pkg] = true
					}
				}
			}
		}

		// Check parameters
		for _, param := range method.Parameters {
			if goType, found := lookupTypeMapping(param.Type, framework); found {
				if strings.Contains(goType, ".") {
					parts := strings.Split(goType, ".")
					if len(parts) == 2 {
						pkg := parts[0]
						if !strings.EqualFold(pkg, framework) {
							imports["github.com/tmc/appledocs/generated/"+pkg] = true
						}
					}
				}
			}
		}
	}

	result := make([]string, 0, len(imports))
	for imp := range imports {
		result = append(result, imp)
	}
	return result
}

// TestImportCycles tests that we don't create import cycles
func TestImportCycles(t *testing.T) {
	// A framework should never import itself
	tests := []struct {
		framework string
		methods   []MethodInfo
	}{
		{
			framework: "AppKit",
			methods: []MethodInfo{
				{
					Selector:   "backingStoreType",
					ReturnType: "NSBackingStoreType",
				},
			},
		},
		{
			framework: "Foundation",
			methods: []MethodInfo{
				{
					Selector:   "frame",
					ReturnType: "NSRect",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.framework, func(t *testing.T) {
			classInfo := ClassInfo{
				Name:    "TestClass",
				Methods: tt.methods,
			}

			imports := determineImports(classInfo, tt.framework)

			// Check that we don't import ourselves
			selfImport := "github.com/tmc/appledocs/generated/" + strings.ToLower(tt.framework)
			for _, imp := range imports {
				if strings.ToLower(imp) == selfImport {
					t.Errorf("Framework %s imports itself: %s", tt.framework, imp)
				}
			}
		})
	}
}

// TestQualifiedTypeGeneration tests that qualified types are generated correctly
func TestQualifiedTypeGeneration(t *testing.T) {
	tests := []struct {
		name      string
		objcType  string
		framework string
		wantType  string
		wantPkg   string
	}{
		{
			name:      "ScreenSaver CGRect",
			objcType:  "CGRect",
			framework: "ScreenSaver",
			wantType:  "coregraphics.CGRect",
			wantPkg:   "coregraphics",
		},
		{
			name:      "ScreenSaver BackingStoreType",
			objcType:  "BackingStoreType",
			framework: "ScreenSaver",
			wantType:  "appkit.BackingStoreType",
			wantPkg:   "appkit",
		},
		{
			name:      "AppKit CGRect",
			objcType:  "CGRect",
			framework: "AppKit",
			wantType:  "coregraphics.CGRect",
			wantPkg:   "coregraphics",
		},
		{
			name:      "Foundation unqualified",
			objcType:  "NSRect",
			framework: "Foundation",
			wantType:  "Rect",
			wantPkg:   "",
		},
		{
			name:      "CoreGraphics unqualified",
			objcType:  "CGRect",
			framework: "CoreGraphics",
			wantType:  "CGRect",
			wantPkg:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := lookupTypeMapping(tt.objcType, tt.framework)
			if !found {
				t.Fatalf("Type mapping not found for %s in %s", tt.objcType, tt.framework)
			}

			if got != tt.wantType {
				t.Errorf("Type = %q, want %q", got, tt.wantType)
			}

			// Check package prefix
			if tt.wantPkg != "" {
				if !strings.HasPrefix(got, tt.wantPkg+".") {
					t.Errorf("Type %q doesn't have expected package prefix %q", got, tt.wantPkg)
				}
			} else {
				if strings.Contains(got, ".") {
					t.Errorf("Type %q should be unqualified but contains '.'", got)
				}
			}
		})
	}
}

// TestCrossFrameworkReferences tests that cross-framework references work correctly
func TestCrossFrameworkReferences(t *testing.T) {
	// Test that ScreenSaver can reference both AppKit and CoreGraphics types
	screenSaverClass := ClassInfo{
		Name: "ScreenSaverView",
		Methods: []MethodInfo{
			{
				Selector:   "initWithFrame:isPreview:",
				ReturnType: "instancetype",
				Parameters: []occ2go.Parameter{
					{Type: "Rect"}, // Should map to coregraphics.CGRect
					{Type: "bool"},
				},
			},
			{
				Selector:   "backingStoreType",
				ReturnType: "BackingStoreType", // Should map to appkit.BackingStoreType
			},
		},
	}

	imports := determineImports(screenSaverClass, "ScreenSaver")

	// Should import both coregraphics and appkit
	hasCoregraphics := false
	hasAppKit := false

	for _, imp := range imports {
		if strings.Contains(imp, "coregraphics") {
			hasCoregraphics = true
		}
		if strings.Contains(imp, "appkit") {
			hasAppKit = true
		}
	}

	if !hasCoregraphics {
		t.Error("ScreenSaver should import coregraphics for Rect type")
	}
	if !hasAppKit {
		t.Error("ScreenSaver should import appkit for BackingStoreType")
	}
}

// TestImportDeduplication tests that we don't have duplicate imports
func TestImportDeduplication(t *testing.T) {
	// Multiple methods using the same external type should only result in one import
	classInfo := ClassInfo{
		Name: "TestClass",
		Methods: []MethodInfo{
			{
				Selector:   "initWithFrame:",
				ReturnType: "instancetype",
				Parameters: []occ2go.Parameter{{Type: "CGRect"}},
			},
			{
				Selector:   "setFrame:",
				ReturnType: "void",
				Parameters: []occ2go.Parameter{{Type: "CGRect"}},
			},
			{
				Selector:   "bounds",
				ReturnType: "CGRect",
			},
		},
	}

	imports := determineImports(classInfo, "ScreenSaver")

	// Count coregraphics imports
	count := 0
	for _, imp := range imports {
		if strings.Contains(imp, "coregraphics") {
			count++
		}
	}

	if count != 1 {
		t.Errorf("Expected 1 coregraphics import, got %d: %v", count, imports)
	}
}
