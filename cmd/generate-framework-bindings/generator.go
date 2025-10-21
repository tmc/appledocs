package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/tmc/appledocs/occ2go"
)

// Generator encapsulates the state and methods for generating bindings
type Generator struct {
	Framework        string
	PackageName      string
	InputDir         string
	OutputModule     string
	Variant          string
	WithRefMethods   bool
	GenerateTests    bool
	GenerateExamples bool

	Functions []*occ2go.ParsedFunction
	Classes   []*occ2go.ParsedClass
	Protocols []*occ2go.ParsedProtocol
	Enums     []*occ2go.ParsedEnum
	Typedefs  []*occ2go.ParsedTypedef
	Constants []*occ2go.ParsedConstant

	// Computed/cached data
	frameworkAbstract string
	frameworkURL      string
	refTypes          []string
	typeMethods       map[string][]*occ2go.ParsedFunction
	typeToRef         map[string]string

	// Error collection
	Errors []error
}

// NewGenerator creates a new Generator instance
func NewGenerator(framework, packageName, inputDir, outputModule, variant string, withRefMethods, generateTests, generateExamples bool) *Generator {
	return &Generator{
		Framework:        framework,
		PackageName:      packageName,
		InputDir:         inputDir,
		OutputModule:     outputModule,
		Variant:          variant,
		WithRefMethods:   withRefMethods,
		GenerateTests:    generateTests,
		GenerateExamples: generateExamples,
		Errors:           make([]error, 0),
	}
}

// AddError adds an error to the error collection
func (g *Generator) AddError(err error) {
	if err != nil {
		g.Errors = append(g.Errors, err)
	}
}

// prepare computes cached data needed for generation
func (g *Generator) prepare() {
	g.frameworkAbstract, g.frameworkURL, _ = loadFrameworkMetadata(g.InputDir, g.Framework)
	// Build typedef names map to exclude from refTypes
	typedefNames := make(map[string]bool)
	for _, typedef := range g.Typedefs {
		if typedef.Name != "" {
			typedefNames[typedef.Name] = true
		}
	}
	g.refTypes = extractRefTypes(g.Functions, getFrameworkPrefix(g.Framework), typedefNames)
	g.typeMethods = groupFunctionsByType(g.Functions, g.Framework)
	g.typeToRef = make(map[string]string)
	for _, refType := range g.refTypes {
		prefix := getFrameworkPrefix(g.Framework)
		if strings.HasPrefix(refType, prefix) && strings.HasSuffix(refType, "Ref") {
			typeName := strings.TrimSuffix(strings.TrimPrefix(refType, prefix), "Ref")
			g.typeToRef[typeName] = refType
		}
	}
}

// SortClassesByDependency sorts classes topologically so parent classes come before children.
// This ensures that when a class extends another class in the same framework, the parent
// is generated first.
func (g *Generator) SortClassesByDependency() {
	if len(g.Classes) == 0 {
		return
	}

	// Build a map of class names for quick lookup
	classMap := make(map[string]*occ2go.ParsedClass)
	for _, cls := range g.Classes {
		classMap[cls.Name] = cls
	}

	// Track visit state: 0 = unvisited, 1 = visiting, 2 = visited
	visited := make(map[string]int)
	sorted := make([]*occ2go.ParsedClass, 0, len(g.Classes))

	// Depth-first search for topological sort
	var visit func(className string) bool
	visit = func(className string) bool {
		if visited[className] == 2 {
			return true // Already processed
		}
		if visited[className] == 1 {
			// Cycle detected - shouldn't happen with proper inheritance, but handle gracefully
			return false
		}

		cls, exists := classMap[className]
		if !exists {
			// Class not in this framework (e.g., NSObject, or from another framework)
			return true
		}

		visited[className] = 1 // Mark as visiting

		// Visit parent first if it exists in this framework
		if cls.SuperClass != "" && cls.SuperClass != "NSObject" {
			if !visit(cls.SuperClass) {
				return false // Cycle detected
			}
		}

		visited[className] = 2 // Mark as visited
		sorted = append(sorted, cls)
		return true
	}

	// Visit all classes
	for _, cls := range g.Classes {
		if visited[cls.Name] == 0 {
			visit(cls.Name)
		}
	}

	// Update the classes with the sorted order
	g.Classes = sorted
}

// GenerateMissingParentStubs creates stub class definitions for parent classes
// that are referenced but not defined in the current framework.
// This handles cases where documentation doesn't include abstract base classes.
func (g *Generator) GenerateMissingParentStubs() []*occ2go.ParsedClass {
	if len(g.Classes) == 0 {
		return nil
	}

	// Build a map of existing classes
	existing := make(map[string]bool)
	for _, cls := range g.Classes {
		existing[cls.Name] = true
	}

	// Find all referenced parent classes that don't exist
	missing := make(map[string]bool)
	for _, cls := range g.Classes {
		if cls.SuperClass != "" && cls.SuperClass != "NSObject" {
			if !existing[cls.SuperClass] {
				// Check if it's a cross-framework type by checking the type registry
				// IMPORTANT: We need to temporarily clear currentFrameworkClasses to prevent
				// false matches. For example, in MetalKit, MTKView strips to "View", but
				// NSView (the parent class) also strips to "View". We don't want resolveType
				// to think NSView is a local type just because MTKView exists.
				savedClasses := currentFrameworkClasses
				currentFrameworkClasses = make(map[string]bool)

				structName := classToStructName(cls.SuperClass)
				resolvedType := resolveType(g.Framework, structName)

				// Restore the original map
				currentFrameworkClasses = savedClasses

				// Only create stub if it's not from another framework
				if !strings.Contains(resolvedType, ".") {
					missing[cls.SuperClass] = true
				}
			}
		}
	}

	// Generate stub classes
	stubs := make([]*occ2go.ParsedClass, 0, len(missing))
	for className := range missing {
		stub := &occ2go.ParsedClass{
			Name:       className,
			SuperClass: "NSObject",
			Methods:    []*occ2go.ParsedMethod{},
			Properties: []*occ2go.ParsedProperty{},
			Comment:    fmt.Sprintf("Auto-generated stub for missing parent class %s", className),
			Abstract:   fmt.Sprintf("A parent class referenced by other %s classes.", g.Framework),
		}
		stubs = append(stubs, stub)
	}

	return stubs
}

// Helper methods for templates

// FunctionCount returns the number of functions
func (g *Generator) FunctionCount() int {
	return len(g.Functions)
}

// ClassCount returns the number of classes
func (g *Generator) ClassCount() int {
	return len(g.Classes)
}

// ProtocolCount returns the number of protocols
func (g *Generator) ProtocolCount() int {
	return len(g.Protocols)
}

// EnumCount returns the number of enums
func (g *Generator) EnumCount() int {
	return len(g.Enums)
}

// TypedefCount returns the number of typedefs
func (g *Generator) TypedefCount() int {
	return len(g.Typedefs)
}

// ConstantCount returns the number of constants
func (g *Generator) ConstantCount() int {
	return len(g.Constants)
}

// MinVersion returns the minimum macOS version
func (g *Generator) MinVersion() string {
	return findMinimumMacOSVersion(g.Functions)
}

// Abstract returns the framework abstract
func (g *Generator) Abstract() string {
	return g.frameworkAbstract
}

// DocURL returns the framework documentation URL
func (g *Generator) DocURL() string {
	return g.frameworkURL
}

// RefTypes returns the ref types
func (g *Generator) RefTypes() []string {
	return g.refTypes
}

// TypeMethods returns the type methods map
func (g *Generator) TypeMethods() map[string][]*occ2go.ParsedFunction {
	return g.typeMethods
}

// TypeToRef returns the type to ref map
func (g *Generator) TypeToRef() map[string]string {
	return g.typeToRef
}

// Count returns function count (for backward compatibility)
func (g *Generator) Count() int {
	return len(g.Functions)
}

// GenerateTxtarFromModule generates the entire txtar output using the module template
func (g *Generator) GenerateTxtarFromModule(w io.Writer) error {
	g.prepare()

	// Load all templates as named templates
	moduleContent, err := getTemplateVariant("module", g.Variant)
	if err != nil {
		return err
	}

	// Create master template and parse all sub-templates
	tmpl := template.New("module").Funcs(templateFuncs)

	// Dynamically discover all templates from the archive (except "module")
	templateFiles := make(map[string]bool)

	// Collect from base archive
	for _, file := range templateArchive.Files {
		if file.Name != "module" && file.Name != "" {
			templateFiles[file.Name] = true
		}
	}

	// Collect from variant archives
	if g.Variant != "" {
		variants := strings.Split(g.Variant, ",")
		for _, v := range variants {
			v = strings.TrimSpace(v)
			if archive, ok := variantArchives[v]; ok {
				for _, file := range archive.Files {
					if file.Name != "module" && file.Name != "" {
						templateFiles[file.Name] = true
					}
				}
			}
		}
	}

	// Parse all discovered templates as associated templates
	for filename := range templateFiles {
		content, err := getTemplateVariant(filename, g.Variant)
		if err != nil {
			g.AddError(fmt.Errorf("warning: skipping template %s: %w", filename, err))
			continue // Skip if template doesn't exist
		}
		_, err = tmpl.New(filename).Parse(content)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", filename, err)
		}
	}

	// Parse the module template last
	tmpl, err = tmpl.Parse(moduleContent)
	if err != nil {
		return fmt.Errorf("failed to parse module template: %w", err)
	}

	// Execute template - pass Generator directly as context
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, g); err != nil {
		return err
	}

	// Replace #-- with -- to convert to proper txtar format
	output := strings.ReplaceAll(buf.String(), "#-- ", "-- ")
	_, err = w.Write([]byte(output))
	return err
}
