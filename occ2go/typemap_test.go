package occ2go

import (
	"testing"
)

func TestParseBlockType(t *testing.T) {
	tests := []struct {
		name      string
		blockType string
		framework string
		want      string
	}{
		{
			name:      "simple void block",
			blockType: "void (^)(void)",
			framework: "Foundation",
			want:      "func()",
		},
		{
			name:      "block with id return",
			blockType: "id (^)(void)",
			framework: "Foundation",
			want:      "func() unsafe.Pointer",
		},
		{
			name:      "block with single param",
			blockType: "void (^)(id)",
			framework: "Foundation",
			want:      "func(unsafe.Pointer)",
		},
		{
			name:      "block with multiple params",
			blockType: "BOOL (^)(id, NSError *)",
			framework: "Foundation",
			want:      "func(unsafe.Pointer, unsafe.Pointer) bool",
		},
		{
			name:      "block with NSString param",
			blockType: "void (^)(NSString *)",
			framework: "Foundation",
			want:      "func(unsafe.Pointer)",
		},
		{
			name:      "not a block type",
			blockType: "NSString *",
			framework: "Foundation",
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseBlockType(tt.blockType, tt.framework)
			if got != tt.want {
				t.Errorf("parseBlockType(%q, %q) = %q, want %q", tt.blockType, tt.framework, got, tt.want)
			}
		})
	}
}

func TestMapCTypeToGo_Blocks(t *testing.T) {
	tests := []struct {
		name      string
		cType     string
		framework string
		want      string
	}{
		{
			name:      "void block",
			cType:     "void (^)(void)",
			framework: "Foundation",
			want:      "func()",
		},
		{
			name:      "BOOL block with params",
			cType:     "BOOL (^)(id, NSError *)",
			framework: "Foundation",
			want:      "func(unsafe.Pointer, unsafe.Pointer) bool",
		},
		{
			name:      "regular type",
			cType:     "NSString *",
			framework: "Foundation",
			want:      "unsafe.Pointer",
		},
		{
			name:      "NSArray of blocks",
			cType:     "NSArray<void (^)(void)>",
			framework: "Foundation",
			want:      "[]func()",
		},
		{
			name:      "NSArray of BOOL blocks with params",
			cType:     "NSArray<BOOL (^)(id, NSError *)>",
			framework: "Foundation",
			want:      "[]func(unsafe.Pointer, unsafe.Pointer) bool",
		},
		{
			name:      "NSArray of blocks with pointer",
			cType:     "NSArray<void (^)(void)> *",
			framework: "Foundation",
			want:      "[]func()",  // NSArray<T> * maps to []T in Go
		},
		{
			name:      "exact executionBlocks type from JSON",
			cType:     "NSArray<void (^)(void)>",
			framework: "Foundation",
			want:      "[]func()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapCTypeToGo(tt.cType, tt.framework)
			if got != tt.want {
				t.Errorf("MapCTypeToGo(%q, %q) = %q, want %q", tt.cType, tt.framework, got, tt.want)
			}
		})
	}
}

func TestMapCTypeToGo_DebugExecutionBlocks(t *testing.T) {
	// This test replicates exactly what the executionBlocks property type should be
	input := "NSArray<void (^)(void)>"
	expected := "[]func()"
	result := MapCTypeToGo(input, "Foundation")

	t.Logf("Input: %q", input)
	t.Logf("Expected: %q", expected)
	t.Logf("Got: %q", result)

	if result != expected {
		t.Errorf("MapCTypeToGo failed for executionBlocks type:\n  Input: %q\n  Expected: %q\n  Got: %q", input, expected, result)
	}
}

func TestMapCTypeToGo_ID_And_Class(t *testing.T) {
	tests := []struct {
		cType     string
		framework string
	}{
		{"id", "ObjectiveC"},
		{"Class", "ObjectiveC"},
		{"void (*)(id, void *)", "ObjectiveC"},
	}
	
	for _, tt := range tests {
		t.Run(tt.cType, func(t *testing.T) {
			got := MapCTypeToGo(tt.cType, tt.framework)
			t.Logf("MapCTypeToGo(%q, %q) = %q", tt.cType, tt.framework, got)
		})
	}
}

func TestParseBlockType_CFunctionPointers(t *testing.T) {
	tests := []struct {
		cType     string
		framework string
	}{
		{"void (*)(id, void *)", "ObjectiveC"},
		{"int (*)(Class, id)", "ObjectiveC"},
		{"id (*)(id)", "ObjectiveC"},
	}
	
	for _, tt := range tests {
		t.Run(tt.cType, func(t *testing.T) {
			got := parseBlockType(tt.cType, tt.framework)
			t.Logf("parseBlockType(%q, %q) = %q", tt.cType, tt.framework, got)
		})
	}
}
