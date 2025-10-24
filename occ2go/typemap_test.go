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
			want:      "func() objc.ID", // id maps to objc.ID type
		},
		{
			name:      "block with single param",
			blockType: "void (^)(id)",
			framework: "Foundation",
			want:      "func(objc.ID)", // id maps to objc.ID type
		},
		{
			name:      "block with multiple params",
			blockType: "BOOL (^)(id, NSError *)",
			framework: "Foundation",
			want:      "func(objc.ID, unsafe.Pointer) bool", // id maps to objc.ID type
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
		{
			name:      "nested block - NSItemProvider loadHandler",
			blockType: "NSProgress * (^)(void (^)(NSData *, NSError *))",
			framework: "Foundation",
			want:      "func(func(unsafe.Pointer, unsafe.Pointer)) unsafe.Pointer",
		},
		{
			name:      "nested block with parameter name",
			blockType: "NSProgress * (^)(void (^completionHandler)(NSData * data, NSError * error))",
			framework: "Foundation",
			want:      "func(func(unsafe.Pointer, unsafe.Pointer)) unsafe.Pointer",
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

// TestParseBlockType_ErrorCases tests error conditions and edge cases in block type parsing
func TestParseBlockType_ErrorCases(t *testing.T) {
	tests := []struct {
		name      string
		blockType string
		framework string
		want      string // Empty string means not a block type
	}{
		{
			name:      "not a block type - simple type",
			blockType: "NSString *",
			framework: "Foundation",
			want:      "",
		},
		{
			name:      "not a block type - missing caret",
			blockType: "void ()(void)",
			framework: "Foundation",
			want:      "",
		},
		{
			name:      "malformed block - missing opening paren",
			blockType: "void (^)void)",
			framework: "Foundation",
			want:      "",
		},
		{
			name:      "malformed block - missing closing paren",
			blockType: "void (^)(void",
			framework: "Foundation",
			want:      "",
		},
		{
			name:      "malformed block - unbalanced nested parens",
			blockType: "NSProgress * (^)(void (^)(NSData *, NSError *",
			framework: "Foundation",
			want:      "",
		},
		{
			name:      "malformed block - extra closing paren (still parses correctly)",
			blockType: "void (^)(void))",
			framework: "Foundation",
			want:      "func()", // Extra paren is ignored - block is still valid
		},
		{
			name:      "empty string",
			blockType: "",
			framework: "Foundation",
			want:      "",
		},
		{
			name:      "only whitespace",
			blockType: "   ",
			framework: "Foundation",
			want:      "",
		},
		{
			name:      "block with only caret marker",
			blockType: "(^)",
			framework: "Foundation",
			want:      "",
		},
		{
			name:      "deeply nested blocks - triple nesting",
			blockType: "void (^)(void (^)(void (^)(void)))",
			framework: "Foundation",
			want:      "func(func(func()))",
		},
		{
			name:      "block with complex parameter types (includes generic array)",
			blockType: "id (^)(NSString *, NSArray<NSString *> *, NSError **)",
			framework: "Foundation",
			want:      "func(unsafe.Pointer, []unsafe.Pointer, unsafe.Pointer) objc.ID", // NSArray<T> maps to []T, id maps to objc.ID
		},
		{
			name:      "block with pointer return type",
			blockType: "NSString * (^)(void)",
			framework: "Foundation",
			want:      "func() unsafe.Pointer",
		},
		{
			name:      "block with no parameters but explicit void - id return",
			blockType: "id (^)(void)",
			framework: "Foundation",
			want:      "func() objc.ID", // id maps to objc.ID, not unsafe.Pointer
		},
		{
			name:      "block with named parameters in nested block",
			blockType: "void (^)(void (^handler)(NSString *message, NSInteger code))",
			framework: "Foundation",
			want:      "func(func(unsafe.Pointer, unsafe.Pointer))", // NSInteger without framework context maps to unsafe.Pointer
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
			want:      "func(objc.ID, unsafe.Pointer) bool", // id maps to objc.ID type
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
			want:      "[]func(objc.ID, unsafe.Pointer) bool", // id maps to objc.ID type
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
