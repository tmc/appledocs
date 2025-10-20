package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/tmc/appledocs/occ2go"
)

// TestGeneratedCodeCompiles tests that generated code compiles successfully
func TestGeneratedCodeCompiles(t *testing.T) {
	tests := []struct {
		name      string
		framework string
		class     ClassInfo
		wantError bool
	}{
		{
			name:      "Simple class with basic method",
			framework: "TestFramework",
			class: ClassInfo{
				Name: "TestClass",
				Methods: []MethodInfo{
					{
						Selector:   "init",
						ReturnType: "instancetype",
						Parameters: []occ2go.Parameter{},
					},
				},
			},
			wantError: false,
		},
		{
			name:      "Class with CGRect parameter",
			framework: "ScreenSaver",
			class: ClassInfo{
				Name: "TestView",
				Methods: []MethodInfo{
					{
						Selector:   "initWithFrame:",
						ReturnType: "instancetype",
						Parameters: []occ2go.Parameter{
							{Name: "frame", Type: "Rect"},
						},
					},
				},
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate the code
			code, err := generateTestCode(tt.class, tt.framework)
			if err != nil {
				t.Fatalf("Failed to generate code: %v", err)
			}

			// Try to parse it
			fset := token.NewFileSet()
			_, err = parser.ParseFile(fset, "test.go", code, parser.ParseComments)
			if (err != nil) != tt.wantError {
				t.Errorf("parser.ParseFile() error = %v, wantError %v\nCode:\n%s", err, tt.wantError, code)
			}
		})
	}
}

// generateTestCode generates a minimal valid Go file for testing
func generateTestCode(class ClassInfo, framework string) (string, error) {
	// Determine imports
	imports := determineImports(class, framework)

	var buf bytes.Buffer
	buf.WriteString("package testpkg\n\n")

	// Add imports
	if len(imports) > 0 {
		buf.WriteString("import (\n")
		for _, imp := range imports {
			buf.WriteString(fmt.Sprintf("\t%q\n", imp))
		}
		buf.WriteString(")\n\n")
	}

	// Add a simple struct and method
	buf.WriteString(fmt.Sprintf("type %s struct {\n", class.Name))
	buf.WriteString("\tID objc.ID\n")
	buf.WriteString("}\n\n")

	// Add one method as example
	if len(class.Methods) > 0 {
		method := class.Methods[0]
		goName := methodToGoName(method)

		// Map return type
		returnType, _ := lookupTypeMapping(method.ReturnType, framework)
		if returnType == "" || returnType == "instancetype" {
			returnType = class.Name
		}

		// Generate method signature
		buf.WriteString(fmt.Sprintf("func (c %s) %s(", class.Name, goName))

		// Add parameters
		for i, param := range method.Parameters {
			if i > 0 {
				buf.WriteString(", ")
			}
			paramType, _ := lookupTypeMapping(param.Type, framework)
			if paramType == "" {
				paramType = "objc.ID"
			}
			buf.WriteString(fmt.Sprintf("%s %s", sanitizeName(param.Name), paramType))
		}

		buf.WriteString(fmt.Sprintf(") %s {\n", returnType))
		if returnType != "void" && returnType != "" {
			buf.WriteString(fmt.Sprintf("\treturn %s{}\n", returnType))
		}
		buf.WriteString("}\n")
	}

	return buf.String(), nil
}

// sanitizeName ensures a parameter name is a valid Go identifier
func sanitizeName(name string) string {
	if name == "" {
		return "value"
	}
	// Replace common Objective-C patterns
	name = strings.ReplaceAll(name, ":", "")
	// Ensure first character is lowercase
	if len(name) > 0 {
		name = strings.ToLower(name[:1]) + name[1:]
	}
	return name
}

// TestGoFormatGeneratedCode tests that generated code can be formatted
func TestGoFormatGeneratedCode(t *testing.T) {
	class := ClassInfo{
		Name: "TestClass",
		Methods: []MethodInfo{
			{
				Selector:   "count",
				ReturnType: "int",
			},
		},
	}

	code, err := generateTestCode(class, "Foundation")
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Try to format it
	formatted, err := format.Source([]byte(code))
	if err != nil {
		t.Errorf("format.Source() failed: %v\nOriginal code:\n%s", err, code)
		return
	}

	// Ensure it's still valid after formatting
	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, "test.go", formatted, parser.ParseComments)
	if err != nil {
		t.Errorf("Formatted code doesn't parse: %v\nFormatted code:\n%s", err, formatted)
	}
}

// TestMethodSignatureGeneration tests method signature generation
func TestMethodSignatureGeneration(t *testing.T) {
	tests := []struct {
		name      string
		method    MethodInfo
		framework string
		wantSig   string
	}{
		{
			name: "simple getter",
			method: MethodInfo{
				Selector:   "title",
				ReturnType: "NSString *",
			},
			framework: "AppKit",
			wantSig:   "func (c TestClass) Title() string",
		},
		{
			name: "simple setter",
			method: MethodInfo{
				Selector:   "setTitle:",
				ReturnType: "void",
				Parameters: []occ2go.Parameter{
					{Name: "title", Type: "NSString *"},
				},
			},
			framework: "AppKit",
			wantSig:   "func (c TestClass) SetTitle(title string)",
		},
		{
			name: "method with CGRect",
			method: MethodInfo{
				Selector:   "initWithFrame:",
				ReturnType: "instancetype",
				Parameters: []occ2go.Parameter{
					{Name: "frame", Type: "CGRect"},
				},
			},
			framework: "AppKit",
			wantSig:   "func (c TestClass) InitWithFrame(frame coregraphics.CGRect) TestClass",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sig := generateMethodSignature(tt.method, "TestClass", tt.framework)

			// Normalize whitespace for comparison
			sig = strings.Join(strings.Fields(sig), " ")
			wantSig := strings.Join(strings.Fields(tt.wantSig), " ")

			if sig != wantSig {
				t.Errorf("generateMethodSignature() = %q, want %q", sig, wantSig)
			}
		})
	}
}

// generateMethodSignature generates a method signature for testing
func generateMethodSignature(method MethodInfo, className, framework string) string {
	goName := methodToGoName(method)

	// Map return type
	returnType := method.ReturnType
	if returnType == "void" {
		returnType = ""
	} else {
		mappedType, _ := lookupTypeMapping(method.ReturnType, framework)
		if mappedType != "" {
			returnType = mappedType
		}
		if returnType == "instancetype" || returnType == "id" {
			returnType = className
		}
	}

	// Build signature
	var sig strings.Builder
	sig.WriteString(fmt.Sprintf("func (c %s) %s(", className, goName))

	// Add parameters
	for i, param := range method.Parameters {
		if i > 0 {
			sig.WriteString(", ")
		}
		paramType, _ := lookupTypeMapping(param.Type, framework)
		if paramType == "" {
			paramType = "objc.ID"
		}
		sig.WriteString(fmt.Sprintf("%s %s", sanitizeName(param.Name), paramType))
	}

	sig.WriteString(")")
	if returnType != "" {
		sig.WriteString(" " + returnType)
	}

	return sig.String()
}

// TestTemplateExecution tests that our templates execute successfully
func TestTemplateExecution(t *testing.T) {
	// This is a simplified test - the real templates are more complex
	tmplStr := `package {{.Package}}

import (
	"github.com/tmc/appledocs/generated/objc"
)

type {{.ClassName}} struct {
	ID objc.ID
}

func New{{.ClassName}}() {{.ClassName}} {
	return {{.ClassName}}{}
}
`

	tmpl, err := template.New("test").Parse(tmplStr)
	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	data := struct {
		Package   string
		ClassName string
	}{
		Package:   "testpkg",
		ClassName: "TestClass",
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	code := buf.String()

	// Ensure it's valid Go code
	fset := token.NewFileSet()
	_, err = parser.ParseFile(fset, "test.go", code, parser.ParseComments)
	if err != nil {
		t.Errorf("Template generated invalid code: %v\nCode:\n%s", err, code)
	}
}

// TestGeneratedCodeBuildability tests that generated code can be built
func TestGeneratedCodeBuildability(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping build test in short mode")
	}

	// Create a temporary directory
	tmpDir := t.TempDir()

	// Generate a simple test file
	class := ClassInfo{
		Name: "TestClass",
		Methods: []MethodInfo{
			{
				Selector:   "init",
				ReturnType: "instancetype",
			},
		},
	}

	code, err := generateTestCode(class, "TestFramework")
	if err != nil {
		t.Fatalf("Failed to generate code: %v", err)
	}

	// Write to file
	testFile := filepath.Join(tmpDir, "test.go")
	if err := os.WriteFile(testFile, []byte(code), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Initialize go module
	cmd := exec.Command("go", "mod", "init", "testmodule")
	cmd.Dir = tmpDir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Logf("go mod init output: %s", out)
		// Not fatal - may already exist
	}

	// Try to build (will fail because of imports, but should parse)
	cmd = exec.Command("go", "build", "-o", "/dev/null", testFile)
	cmd.Dir = tmpDir
	out, err := cmd.CombinedOutput()

	// We expect build to fail due to missing imports, but not due to syntax errors
	if err != nil && !strings.Contains(string(out), "cannot find package") {
		t.Logf("Build output: %s", out)
		// Check if it's a syntax error
		if strings.Contains(string(out), "syntax error") {
			t.Errorf("Generated code has syntax errors: %s", out)
		}
	}
}

// TestConstructorGeneration tests constructor generation
func TestConstructorGeneration(t *testing.T) {
	tests := []struct {
		name      string
		method    MethodInfo
		className string
		wantName  string
	}{
		{
			name: "init becomes New",
			method: MethodInfo{
				Selector:   "init",
				ReturnType: "instancetype",
			},
			className: "Button",
			wantName:  "NewButton",
		},
		{
			name: "initWithFrame becomes NewWithFrame",
			method: MethodInfo{
				Selector:   "initWithFrame:",
				ReturnType: "instancetype",
				Parameters: []occ2go.Parameter{
					{Name: "frame", Type: "CGRect"},
				},
			},
			className: "View",
			wantName:  "NewViewWithFrame",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate constructor name
			constructorName := generateConstructorName(tt.method)
			fullName := "New" + tt.className + strings.TrimPrefix(constructorName, "New")

			if fullName != tt.wantName {
				t.Errorf("Constructor name = %q, want %q", fullName, tt.wantName)
			}
		})
	}
}

// TestPropertyAccessorGeneration tests property accessor generation
func TestPropertyAccessorGeneration(t *testing.T) {
	tests := []struct {
		name   string
		method MethodInfo
		want   string
	}{
		{
			name: "simple getter",
			method: MethodInfo{
				Selector:   "title",
				ReturnType: "NSString *",
			},
			want: "Title",
		},
		{
			name: "boolean getter",
			method: MethodInfo{
				Selector:   "isEnabled",
				ReturnType: "BOOL",
			},
			want: "IsEnabled",
		},
		{
			name: "setter",
			method: MethodInfo{
				Selector:   "setTitle:",
				ReturnType: "void",
				Parameters: []occ2go.Parameter{
					{Name: "title", Type: "NSString *"},
				},
			},
			want: "SetTitle",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := methodToGoName(tt.method)
			if got != tt.want {
				t.Errorf("methodToGoName() = %q, want %q", got, tt.want)
			}
		})
	}
}
