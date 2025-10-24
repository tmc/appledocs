package main

import (
	"testing"

	"github.com/tmc/appledocs/occ2go"
)

// TestGetStructEmbeddedField tests that struct embedding resolves parent classes correctly
func TestGetStructEmbeddedField(t *testing.T) {
	tests := []struct {
		name      string
		class     *occ2go.ParsedClass
		framework string
		expected  string
	}{
		{
			name: "SimpleCString inherits from NSString in Foundation",
			class: &occ2go.ParsedClass{
				Name:       "NSSimpleCString",
				SuperClass: "NSString",
			},
			framework: "Foundation",
			expected:  "String", // Should be "String" (the Foundation class), not "string" (the primitive)
		},
		{
			name: "ConstantString inherits from NSSimpleCString in Foundation",
			class: &occ2go.ParsedClass{
				Name:       "NSConstantString",
				SuperClass: "NSSimpleCString",
			},
			framework: "Foundation",
			expected:  "SimpleCString",
		},
		{
			name: "MutableString inherits from NSString in Foundation",
			class: &occ2go.ParsedClass{
				Name:       "NSMutableString",
				SuperClass: "NSString",
			},
			framework: "Foundation",
			expected:  "String",
		},
		{
			name: "NSObject in Foundation",
			class: &occ2go.ParsedClass{
				Name:       "NSObject",
				SuperClass: "",
			},
			framework: "Foundation",
			expected:  "objectivec.Object",
		},
		{
			name: "NSObject in ObjectiveC framework",
			class: &occ2go.ParsedClass{
				Name:       "NSObject",
				SuperClass: "",
			},
			framework: "ObjectiveC",
			expected:  "objc.ID",
		},
		{
			name: "Class with NSObject superclass",
			class: &occ2go.ParsedClass{
				Name:       "NSData",
				SuperClass: "NSObject",
			},
			framework: "Foundation",
			expected:  "objectivec.Object",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getStructEmbeddedField(tt.class, tt.framework)
			if result != tt.expected {
				t.Errorf("getStructEmbeddedField(%+v, %q) = %q, want %q",
					tt.class, tt.framework, result, tt.expected)
			}
		})
	}
}

// TestResolveTypeWithStringClass tests that resolveType handles "String" class correctly
func TestResolveTypeWithStringClass(t *testing.T) {
	tests := []struct {
		name      string
		framework string
		typeName  string
		expected  string
	}{
		{
			name:      "String class in Foundation should not become primitive string",
			framework: "Foundation",
			typeName:  "String",
			expected:  "String", // Class name, not primitive
		},
		{
			name:      "Primitive string should stay unqualified",
			framework: "Foundation",
			typeName:  "string",
			expected:  "string",
		},
		{
			name:      "Data class stays unqualified in Foundation",
			framework: "Foundation",
			typeName:  "Data",
			expected:  "Data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resolveType(tt.framework, tt.typeName)
			if result != tt.expected {
				t.Errorf("resolveType(%q, %q) = %q, want %q",
					tt.framework, tt.typeName, result, tt.expected)
			}
		})
	}
}
