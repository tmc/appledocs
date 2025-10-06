package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTypeGenerator(t *testing.T) {
	// Create temp directory for test
	tmpDir := t.TempDir()
	outputPath := filepath.Join(tmpDir, "test_types.go")

	// Create type generator
	gen := NewTypeGenerator("testpkg", outputPath)

	// Add some test schemas
	gen.schemas["TestStruct"] = &Schema{
		Name: "TestStruct",
		Fields: map[string]*FieldInfo{
			"Name": {
				JSONType:   "name",
				GoType:     "string",
				IsOptional: false,
			},
			"Count": {
				JSONType:   "count",
				GoType:     "int",
				IsOptional: true,
			},
			"Tags": {
				JSONType:   "tags",
				GoType:     "string",
				IsArray:    true,
				IsOptional: true,
			},
		},
	}

	// Generate types
	if err := gen.GenerateTypes(); err != nil {
		t.Fatalf("GenerateTypes failed: %v", err)
	}

	// Read generated file
	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	content := string(data)

	// Check package declaration
	if !strings.Contains(content, "package testpkg") {
		t.Error("Package declaration not found")
	}

	// Check type declaration
	if !strings.Contains(content, "type TestStruct struct") {
		t.Error("Type declaration not found")
	}

	// Check fields
	if !strings.Contains(content, "Name string") {
		t.Error("Name field not found or has wrong type")
	}

	if !strings.Contains(content, "Count int") {
		t.Error("Count field not found or has wrong type")
	}

	if !strings.Contains(content, "Tags []string") {
		t.Error("Tags field not found or has wrong array type")
	}

	// Check JSON tags
	if !strings.Contains(content, "`json:\"name\"`") {
		t.Error("JSON tag for Name field incorrect")
	}

	if !strings.Contains(content, "`json:\"count,omitempty\"`") {
		t.Error("JSON tag for Count field incorrect")
	}

	if !strings.Contains(content, "`json:\"tags,omitempty\"`") {
		t.Error("JSON tag for Tags field incorrect")
	}
}

func TestToGoTypeName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello_world", "HelloWorld"},
		{"api_url", "APIURL"},
		{"test-case", "TestCase"},
		{"json_data", "JSONData"},
		{"id", "ID"},
		{"user_id", "UserID"},
		{"http_request", "HTTPRequest"},
		{"simple", "Simple"},
		{"", ""},
	}

	for _, tt := range tests {
		result := toGoTypeName(tt.input)
		if result != tt.expected {
			t.Errorf("toGoTypeName(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestInferGoType(t *testing.T) {
	tests := []struct {
		value      interface{}
		goType     string
		isArray    bool
	}{
		{nil, "interface{}", false},
		{true, "bool", false},
		{false, "bool", false},
		{42.0, "int", false},
		{42.5, "float64", false},
		{"hello", "string", false},
		{[]interface{}{}, "interface{}", true},
		{[]interface{}{"a", "b"}, "string", true},
		{[]interface{}{1.0, 2.0}, "int", true},
		{map[string]interface{}{}, "map[string]interface{}", false},
	}

	for _, tt := range tests {
		goType, isArray := inferGoType(tt.value)
		if goType != tt.goType {
			t.Errorf("inferGoType(%v) type = %q, want %q", tt.value, goType, tt.goType)
		}
		if isArray != tt.isArray {
			t.Errorf("inferGoType(%v) isArray = %v, want %v", tt.value, isArray, tt.isArray)
		}
	}
}

func TestAnalyzeObject(t *testing.T) {
	gen := NewTypeGenerator("test", "/tmp/test.go")

	// Test object
	obj := map[string]interface{}{
		"name":  "test",
		"count": 42.0,
		"active": true,
		"tags":  []interface{}{"a", "b", "c"},
		"metadata": map[string]interface{}{
			"created": "2025-01-01",
		},
	}

	gen.analyzeObject("TestObject", obj, 0)

	// Check that schema was created
	schema, exists := gen.schemas["TestObject"]
	if !exists {
		t.Fatal("TestObject schema not created")
	}

	// Check fields
	if len(schema.Fields) != 5 {
		t.Errorf("Expected 5 fields, got %d", len(schema.Fields))
	}

	// Check specific fields
	nameField := schema.Fields["Name"]
	if nameField == nil || nameField.GoType != "string" {
		t.Error("Name field not correctly analyzed")
	}

	countField := schema.Fields["Count"]
	if countField == nil || countField.GoType != "int" {
		t.Error("Count field not correctly analyzed")
	}

	tagsField := schema.Fields["Tags"]
	if tagsField == nil || tagsField.GoType != "string" || !tagsField.IsArray {
		t.Error("Tags field not correctly analyzed as array")
	}

	// Check nested object schema was created
	if _, exists := gen.schemas["TestObjectMetadata"]; !exists {
		t.Error("Nested Metadata schema not created")
	}
}

func TestTruncateExample(t *testing.T) {
	tests := []struct {
		input    interface{}
		maxLen   int
	}{
		{"short", 50},
		{"this is a very long string that should be truncated to fit the maximum length allowed for examples", 50},
		{42, 50},
		{[]string{"a", "b", "c"}, 50},
	}

	for _, tt := range tests {
		result := truncateExample(tt.input)
		if len(result) > tt.maxLen {
			t.Errorf("truncateExample(%v) = %q (len %d), want len <= %d", tt.input, result, len(result), tt.maxLen)
		}
	}
}
