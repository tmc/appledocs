package main

import (
	"testing"

	"github.com/tmc/appledocs/occ2go"
)

// TestEnumsNotInUndefinedTypes verifies that enum types are not
// incorrectly added to the undefined types list, which would cause
// duplicate type declarations (once in enums.gen.go and once in types.gen.go)
func TestEnumsNotInUndefinedTypes(t *testing.T) {
	// Create a generator with an enum type
	g := &Generator{
		Framework:   "ServiceManagement",
		PackageName: "servicemanagement",
		Enums: []*occ2go.ParsedEnum{
			{
				Name: "SMAppServiceStatus",
				Cases: []*occ2go.ParsedEnumCase{
					{Name: "SMAppServiceStatusEnabled", IntValue: 0},
					{Name: "SMAppServiceStatusNotFound", IntValue: 1},
				},
			},
		},
		Classes: []*occ2go.ParsedClass{
			{
				Name: "SMAppService",
				Methods: []*occ2go.ParsedMethod{
					{
						Name:       "status",
						ReturnType: "SMAppServiceStatus",
					},
				},
			},
		},
	}

	// Get defined types
	defined := g.getDefinedTypes()

	// Verify that both the original enum name and stripped version are in defined types
	if !defined["SMAppServiceStatus"] {
		t.Error("SMAppServiceStatus should be in defined types")
	}
	if !defined["AppServiceStatus"] {
		t.Error("AppServiceStatus (stripped) should be in defined types")
	}

	// Collect undefined types - should NOT include the enum
	undefined := g.CollectUndefinedTypes()

	if _, found := undefined["SMAppServiceStatus"]; found {
		t.Error("SMAppServiceStatus should NOT be in undefined types (it's an enum)")
	}
	if _, found := undefined["AppServiceStatus"]; found {
		t.Error("AppServiceStatus should NOT be in undefined types (it's an enum)")
	}
}

// TestEnumsWithMultipleFrameworks verifies enum handling across frameworks
func TestEnumsWithMultipleFrameworks(t *testing.T) {
	tests := []struct {
		framework string
		enumName  string
		stripped  string
	}{
		{"AppKit", "NSWindowStyleMask", "WindowStyleMask"},
		{"Foundation", "NSComparisonResult", "ComparisonResult"},
		{"ServiceManagement", "SMAppServiceStatus", "AppServiceStatus"},
	}

	for _, tt := range tests {
		t.Run(tt.framework+"/"+tt.enumName, func(t *testing.T) {
			g := &Generator{
				Framework:   tt.framework,
				PackageName: tt.framework,
				Enums: []*occ2go.ParsedEnum{
					{Name: tt.enumName},
				},
			}

			defined := g.getDefinedTypes()

			if !defined[tt.enumName] {
				t.Errorf("Enum %s should be in defined types", tt.enumName)
			}
			if !defined[tt.stripped] {
				t.Errorf("Stripped enum name %s should be in defined types", tt.stripped)
			}

			// Add a class that uses this enum
			g.Classes = []*occ2go.ParsedClass{
				{
					Name: "TestClass",
					Methods: []*occ2go.ParsedMethod{
						{
							Name:       "testMethod",
							ReturnType: tt.enumName,
						},
					},
				},
			}

			undefined := g.CollectUndefinedTypes()
			if _, found := undefined[tt.enumName]; found {
				t.Errorf("Enum %s should NOT be in undefined types", tt.enumName)
			}
		})
	}
}
