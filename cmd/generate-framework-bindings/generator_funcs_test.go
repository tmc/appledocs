package main

import (
	"testing"

	"github.com/tmc/appledocs/occ2go"
)

func TestConcreteReturnType(t *testing.T) {
	tests := []struct {
		name       string
		framework  string
		goType     string
		typedefs   []*occ2go.ParsedTypedef
		want       string
	}{
		{
			name:      "Basic type float64",
			framework: "Foundation",
			goType:    "float64",
			want:      "float64",
		},
		{
			name:      "Function type",
			framework: "Foundation",
			goType:    "func()",
			want:      "func()",
		},
		{
			name:      "Function type with signature",
			framework: "Foundation",
			goType:    "func(int) string",
			want:      "func(int) string",
		},
		{
			name:      "IMP special case",
			framework: "ObjectiveC",
			goType:    "IMP",
			want:      "IMP", // IMP is preserved as special cross-framework typedef
		},
		{
			name:      "TimeInterval typedef to float64",
			framework: "Foundation",
			goType:    "TimeInterval",
			typedefs: []*occ2go.ParsedTypedef{
				{Name: "TimeInterval", BaseType: "double"},
			},
			want: "float64", // double maps to float64
		},
		{
			name:      "IMP typedef to func()",
			framework: "ObjectiveC",
			goType:    "IMP",
			typedefs: []*occ2go.ParsedTypedef{
				{Name: "IMP", BaseType: "void (*)(void)"},
			},
			want: "IMP", // IMP special case takes precedence over typedef resolution
		},
		{
			name:      "objc.IObject to objc.ID",
			framework: "Foundation",
			goType:    "objc.IObject",
			want:      "objc.ID",
		},
		{
			name:      "TimeInterval without typedef index",
			framework: "Foundation",
			goType:    "TimeInterval",
			want:      "objc.ID", // Falls through to default without typedef
		},
		{
			name:      "NSTimeInterval via static registry",
			framework: "Foundation",
			goType:    "NSTimeInterval",
			want:      "objc.ID", // Will be mapped to float64 by mapObjCTypeToGo, but this test is post-mapping
		},
		{
			name:      "Slice of foundation.Number in Foundation",
			framework: "Foundation",
			goType:    "[]foundation.Number",
			want:      "[]foundation.Number", // concreteReturnType preserves slices in same framework
		},
		{
			name:      "Slice of foundation.Number in ObjectiveC (hierarchy violation)",
			framework: "ObjectiveC",
			goType:    "[]foundation.Number",
			want:      "[]objc.ID", // Should convert to []objc.ID due to hierarchy violation
		},
		{
			name:      "Slice of coredata.AttributeDescription in ObjectiveC (hierarchy violation)",
			framework: "ObjectiveC",
			goType:    "[]coredata.AttributeDescription",
			want:      "[]objc.ID", // Should convert to []objc.ID due to hierarchy violation
		},
		{
			name:      "Slice of objc.ID",
			framework: "Foundation",
			goType:    "[]objc.ID",
			want:      "[]objc.ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			framework := tt.framework
			if framework == "" {
				framework = "Foundation"
			}
			// Create a minimal generator with typedef index
			gen := &Generator{
				Framework:    framework,
				Typedefs:     tt.typedefs,
				typedefIndex: make(map[string]*occ2go.ParsedTypedef),
			}
			// Build typedef index
			for _, td := range tt.typedefs {
				gen.typedefIndex[td.Name] = td
			}

			gf := GeneratorFuncs{gen}
			got := gf.concreteReturnType(tt.goType)
			if got != tt.want {
				t.Errorf("concreteReturnType(%q) in framework %q = %q, want %q", tt.goType, framework, got, tt.want)
			}
		})
	}
}
