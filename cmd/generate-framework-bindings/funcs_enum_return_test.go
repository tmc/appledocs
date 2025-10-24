package main

import (
	"testing"

	"github.com/tmc/appledocs/occ2go"
)

// TestConcreteReturnTypeWithTypedef tests that typedef return types are preserved
// Key insight: objc.Send[T] needs the actual typedef type, not the underlying primitive
// For example: objc.Send[TextCheckingTypes] not objc.Send[uint64]
func TestConcreteReturnTypeWithTypedef(t *testing.T) {
	tests := []struct {
		name      string
		goType    string // The mapped Go type
		framework string
		typedefs  []*occ2go.ParsedTypedef
		expected  string
	}{
		{
			name:      "Typedef should be preserved for objc.Send",
			goType:    "TextCheckingTypes",
			framework: "Foundation",
			typedefs: []*occ2go.ParsedTypedef{
				{Name: "NSTextCheckingTypes", BaseType: "uint64_t"},
			},
			expected: "TextCheckingTypes", // Preserve typedef, don't resolve to uint64
		},
		{
			name:      "TimeInterval typedef should be preserved",
			goType:    "TimeInterval",
			framework: "Foundation",
			typedefs: []*occ2go.ParsedTypedef{
				{Name: "NSTimeInterval", BaseType: "double"},
			},
			expected: "TimeInterval", // Preserve typedef, don't resolve to float64
		},
		{
			name:      "Non-typedef primitive should pass through",
			goType:    "uint",
			framework: "Foundation",
			typedefs:  nil,
			expected:  "uint",
		},
		{
			name:      "Interface type IObject becomes objc.ID (cross-framework)",
			goType:    "IObject",
			framework: "Foundation",
			typedefs:  nil,
			expected:  "objc.ID", // IObject without package resolves to objc.ID
		},
		{
			name:      "objc.IObject should become objc.ID",
			goType:    "objc.IObject",
			framework: "Foundation",
			typedefs:  nil,
			expected:  "objc.ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a generator with typedef index
			gen := &Generator{
				Framework: tt.framework,
				Typedefs:  tt.typedefs,
			}
			gen.prepare() // Build indexes

			gf := GeneratorFuncs{gen}
			result := gf.concreteReturnType(tt.goType)

			if result != tt.expected {
				t.Errorf("concreteReturnType(%q) = %q, want %q",
					tt.goType, result, tt.expected)
			}
		})
	}
}
