package main

import (
	"testing"

	"github.com/tmc/appledocs/occ2go"
)

// TestIsStringBasedTypedef tests detection of string-based typedefs
func TestIsStringBasedTypedef(t *testing.T) {
	gen := &Generator{
		Framework:   "Foundation",
		PackageName: "foundation",
		Typedefs: []*occ2go.ParsedTypedef{
			{
				Name:     "NSCalendarIdentifier",
				BaseType: "NSString *",
			},
			{
				Name:     "NSErrorDomain",
				BaseType: "NSString *",
			},
			{
				Name:     "unichar",
				BaseType: "unsigned short",
			},
			{
				Name:     "NSInteger",
				BaseType: "long",
			},
		},
	}

	// Build the typedef index
	gen.typedefIndex = make(map[string]*occ2go.ParsedTypedef)
	for _, td := range gen.Typedefs {
		gen.typedefIndex[td.Name] = td
	}

	tests := []struct {
		name     string
		typedef  string
		expected bool
	}{
		{
			name:     "NSCalendarIdentifier is string-based",
			typedef:  "NSCalendarIdentifier",
			expected: true,
		},
		{
			name:     "NSErrorDomain is string-based",
			typedef:  "NSErrorDomain",
			expected: true,
		},
		{
			name:     "unichar is not string-based",
			typedef:  "unichar",
			expected: false,
		},
		{
			name:     "NSInteger is not string-based",
			typedef:  "NSInteger",
			expected: false,
		},
		{
			name:     "unknown typedef",
			typedef:  "UnknownType",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gen.IsStringBasedTypedef(tt.typedef)
			if result != tt.expected {
				t.Errorf("IsStringBasedTypedef(%q) = %v, want %v", tt.typedef, result, tt.expected)
			}
		})
	}
}
