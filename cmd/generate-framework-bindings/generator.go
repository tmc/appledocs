package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
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

	// Indexes for O(1) lookups (built in prepare())
	classIndex   map[string]*occ2go.ParsedClass
	enumIndex    map[string]*occ2go.ParsedEnum
	typedefIndex map[string]*occ2go.ParsedTypedef

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

// IsClassType checks if a given type name (without package prefix) is an ObjC class
// by looking it up in the classIndex. This is used to determine if a type should
// be converted to an interface type (IClassName) or left as-is.
func (g *Generator) IsClassType(typeName string) bool {
	if typeName == "" {
		return false
	}

	// If index not built yet (shouldn't happen after prepare()), use old method
	if g.classIndex == nil {
		for _, class := range g.Classes {
			structName := classToStructName(class.Name)
			if structName == typeName || class.Name == typeName {
				return true
			}
		}
		return false
	}

	// O(1) lookup in index
	if _, ok := g.classIndex[typeName]; ok {
		return true
	}

	// Strip I prefix and try again
	if strings.HasPrefix(typeName, "I") && len(typeName) > 1 && typeName[1] >= 'A' && typeName[1] <= 'Z' {
		lookupName := typeName[1:]
		if _, ok := g.classIndex[lookupName]; ok {
			return true
		}
	}

	// Try with NS prefix
	nsName := "NS" + typeName
	if _, ok := g.classIndex[nsName]; ok {
		return true
	}

	return false
}

// IsEnumType checks if a given type name is an enum by looking it up in the enumIndex.
func (g *Generator) IsEnumType(typeName string) bool {
	if typeName == "" {
		return false
	}

	// If index not built yet, use old method
	if g.enumIndex == nil {
		for _, enum := range g.Enums {
			if enum.Name == typeName {
				return true
			}
			// Also check with NS prefix added (for backward compat with stripped names)
			if "NS"+typeName == enum.Name {
				return true
			}
		}
		return false
	}

	// O(1) lookup in index
	// The index includes both full names (NSDataCompressionAlgorithm)
	// and stripped names (DataCompressionAlgorithm) for backward compatibility
	_, ok := g.enumIndex[typeName]
	return ok
}

// IsTypedefType checks if a given type name is a typedef by looking it up in the typedefIndex.
func (g *Generator) IsTypedefType(typeName string) bool {
	if typeName == "" {
		return false
	}

	// If index not built yet, use old method
	if g.typedefIndex == nil {
		for _, typedef := range g.Typedefs {
			if typedef.Name == typeName {
				return true
			}
		}
		return false
	}

	// O(1) lookup in index
	_, ok := g.typedefIndex[typeName]
	if os.Getenv("DEBUG_TIMEINTERVAL") == "1" && typeName == "TimeInterval" {
		fmt.Fprintf(os.Stderr, "DEBUG: IsTypedefType(TimeInterval) = %v, index size = %d\n", ok, len(g.typedefIndex))
	}
	return ok
}

// TypeToInterfaceType converts a struct type name to its interface type name using
// data-driven type checking. For example: "Data" becomes "IData", "Window" becomes "IWindow".
// For qualified types: "foundation.Coder" becomes "foundation.ICoder".
// Types that don't have interfaces (primitives, slices, enums, typedefs, structs) are returned unchanged.
// GeneratorFuncs wraps a Generator and provides template functions with access to
// the Generator's data-driven type checking (enum/typedef/class indices).
// This allows templates to use accurate type resolution instead of heuristics.
func (g *Generator) TypeToInterfaceType(goType string) string {
	// Handle qualified types (e.g., "foundation.Coder" -> "foundation.ICoder")
	if strings.Contains(goType, ".") {
		parts := strings.SplitN(goType, ".", 2)
		if len(parts) == 2 {
			pkg := parts[0]
			typeName := parts[1]

			// Don't convert runtime types (objc.ID, unsafe.Pointer, etc.)
			if pkg == "objc" || pkg == "unsafe" || pkg == "objectivec" {
				return goType
			}

			// Don't convert CoreGraphics types (structs and refs)
			if strings.HasPrefix(typeName, "CG") {
				return goType
			}

			// Recursively convert the type part
			interfaceType := g.TypeToInterfaceType(typeName)
			return pkg + "." + interfaceType
		}
		return goType
	}

	// Don't convert primitives, slices, pointers
	if strings.HasPrefix(goType, "[]") ||
		strings.HasPrefix(goType, "*") ||
		strings.HasPrefix(goType, "map[") ||
		goType == "string" ||
		goType == "int" ||
		goType == "int64" ||
		goType == "uint" ||
		goType == "uint64" ||
		goType == "float32" ||
		goType == "float64" ||
		goType == "bool" ||
		goType == "unsafe.Pointer" {
		return goType
	}

	// Don't convert CoreGraphics types (structs and refs like CGPoint, CGContextRef)
	if strings.HasPrefix(goType, "CG") {
		return goType
	}

	// Don't convert NSInteger/NSUInteger - these are typedefs, not classes
	if strings.HasPrefix(goType, "NS") && (strings.HasSuffix(goType, "Integer") || strings.HasSuffix(goType, "UInteger")) {
		return goType
	}

	// REMOVED: Hardcoded geometry/alias checks - now fully data-driven
	// TimeInterval and other template-defined types are added to Typedefs synthetically
	// during preparation, so IsTypedefType check above catches them.

	// If it already starts with I and next char is uppercase, it's already an interface
	if strings.HasPrefix(goType, "I") && len(goType) > 1 && goType[1] >= 'A' && goType[1] <= 'Z' {
		return goType
	}

	// DATA-DRIVEN: Check if this is an enum or typedef - don't convert those
	// For enums, return the full name (with NS/CG prefix) if available
	if g.IsEnumType(goType) {
		if enum, ok := g.enumIndex[goType]; ok {
			if os.Getenv("DEBUG_ENUM_TYPE") == "1" && strings.Contains(enum.Name, "Compression") {
				fmt.Fprintf(os.Stderr, "DEBUG: ToInterfaceType: %s -> %s (enum)\n", goType, enum.Name)
			}
			return enum.Name // Return the full enum name (e.g., NSDataCompressionAlgorithm)
		}
		return goType
	}
	if g.IsTypedefType(goType) {
		return goType
	}

	// Strip ObjC prefixes first before checking if it's a class
	// This handles cases where mapObjCTypeToGo returns "NSData" instead of "Data"
	baseType := goType
	if strings.HasPrefix(goType, "NS") && len(goType) > 2 && goType[2] >= 'A' && goType[2] <= 'Z' {
		baseType = goType[2:]
	} else if strings.HasPrefix(goType, "CG") && len(goType) > 2 && goType[2] >= 'A' && goType[2] <= 'Z' {
		baseType = goType[2:]
	} else if strings.HasPrefix(goType, "CA") && len(goType) > 2 && goType[2] >= 'A' && goType[2] <= 'Z' {
		baseType = goType[2:]
	}

	// DATA-DRIVEN: Check if this is actually a class type by looking it up
	// Check both the original type and the stripped version
	// If it's not a class, it's likely a struct - don't convert
	if os.Getenv("DEBUG_TYPEMAP") == "1" && strings.Contains(goType, "Character") {
		fmt.Fprintf(os.Stderr, "DEBUG TypeToInterfaceType: goType=%s baseType=%s IsClass(goType)=%v IsClass(baseType)=%v\n",
			goType, baseType, g.IsClassType(goType), g.IsClassType(baseType))
	}
	if !g.IsClassType(goType) && !g.IsClassType(baseType) {
		if os.Getenv("DEBUG_TYPEMAP") == "1" && strings.Contains(goType, "Character") {
			fmt.Fprintf(os.Stderr, "DEBUG TypeToInterfaceType: NOT a class, returning goType=%s unchanged\n", goType)
		}
		return goType
	}

	// Convert to interface type: "Data" -> "IData"
	// Use the stripped base type for the interface name
	if os.Getenv("DEBUG_TYPEMAP") == "1" && strings.Contains(goType, "Character") {
		fmt.Fprintf(os.Stderr, "DEBUG TypeToInterfaceType: IS a class, returning I%s\n", baseType)
	}
	return "I" + baseType
}

// prepare computes cached data needed for generation
func (g *Generator) prepare() {
	g.frameworkAbstract, g.frameworkURL, _ = loadFrameworkMetadata(g.InputDir, g.Framework)

	// Deduplicate all parsed data structures to prevent duplicate declarations
	// This handles cases where documentation includes the same symbol multiple times

	// Deduplicate classes by name - keep only the first occurrence
	classesSeen := make(map[string]bool)
	deduplicatedClasses := make([]*occ2go.ParsedClass, 0, len(g.Classes))
	for _, class := range g.Classes {
		if !classesSeen[class.Name] {
			classesSeen[class.Name] = true
			deduplicatedClasses = append(deduplicatedClasses, class)
		}
	}
	g.Classes = deduplicatedClasses

	// Deduplicate protocols by name - keep only the first occurrence
	protocolsSeen := make(map[string]bool)
	deduplicatedProtocols := make([]*occ2go.ParsedProtocol, 0, len(g.Protocols))
	for _, protocol := range g.Protocols {
		if !protocolsSeen[protocol.Name] {
			protocolsSeen[protocol.Name] = true
			deduplicatedProtocols = append(deduplicatedProtocols, protocol)
		}
	}
	g.Protocols = deduplicatedProtocols

	// Deduplicate functions by name - keep only the first occurrence
	functionsSeen := make(map[string]bool)
	deduplicatedFunctions := make([]*occ2go.ParsedFunction, 0, len(g.Functions))
	for _, function := range g.Functions {
		if !functionsSeen[function.Name] {
			functionsSeen[function.Name] = true
			deduplicatedFunctions = append(deduplicatedFunctions, function)
		}
	}
	g.Functions = deduplicatedFunctions

	// Deduplicate constants by name - keep only the first occurrence
	constantsSeen := make(map[string]bool)
	deduplicatedConstants := make([]*occ2go.ParsedConstant, 0, len(g.Constants))
	for _, constant := range g.Constants {
		if !constantsSeen[constant.Name] {
			constantsSeen[constant.Name] = true
			deduplicatedConstants = append(deduplicatedConstants, constant)
		}
	}
	g.Constants = deduplicatedConstants

	// Deduplicate enums by name - merge cases from duplicate enums
	enumsMap := make(map[string]*occ2go.ParsedEnum)
	enumsOrder := make([]string, 0, len(g.Enums))

	for _, enum := range g.Enums {
		if existing, exists := enumsMap[enum.Name]; exists {
			// Merge cases from duplicate enum into existing enum
			if os.Getenv("DEBUG_ENUM_DEDUP") == "1" {
				fmt.Fprintf(os.Stderr, "DEBUG: Merging duplicate enum %s with %d new cases into existing %d cases\n",
					enum.Name, len(enum.Cases), len(existing.Cases))
			}
			for _, newCase := range enum.Cases {
				// Check if this case already exists
				isDuplicate := false
				for _, existingCase := range existing.Cases {
					if existingCase.Name == newCase.Name {
						isDuplicate = true
						if os.Getenv("DEBUG_ENUM_DEDUP") == "1" {
							fmt.Fprintf(os.Stderr, "DEBUG:   Skipping duplicate case %s\n", newCase.Name)
						}
						break
					}
				}
				if !isDuplicate {
					if os.Getenv("DEBUG_ENUM_DEDUP") == "1" {
						fmt.Fprintf(os.Stderr, "DEBUG:   Adding new case %s\n", newCase.Name)
					}
					existing.Cases = append(existing.Cases, newCase)
				}
			}
		} else {
			// First time seeing this enum - but also deduplicate its cases
			if os.Getenv("DEBUG_ENUM_DEDUP") == "1" {
				fmt.Fprintf(os.Stderr, "DEBUG: First occurrence of enum %s with %d cases\n", enum.Name, len(enum.Cases))
			}

			// Deduplicate cases within this enum
			casesSeen := make(map[string]bool)
			deduplicatedCases := make([]*occ2go.ParsedEnumCase, 0, len(enum.Cases))
			for _, enumCase := range enum.Cases {
				if !casesSeen[enumCase.Name] {
					casesSeen[enumCase.Name] = true
					deduplicatedCases = append(deduplicatedCases, enumCase)
				} else if os.Getenv("DEBUG_ENUM_DEDUP") == "1" {
					fmt.Fprintf(os.Stderr, "DEBUG:   Removing duplicate case %s from first occurrence\n", enumCase.Name)
				}
			}
			enum.Cases = deduplicatedCases

			enumsMap[enum.Name] = enum
			enumsOrder = append(enumsOrder, enum.Name)
		}
	}

	// Build final deduplicated list in original order
	deduplicatedEnums := make([]*occ2go.ParsedEnum, 0, len(enumsMap))
	for _, name := range enumsOrder {
		deduplicatedEnums = append(deduplicatedEnums, enumsMap[name])
	}
	g.Enums = deduplicatedEnums

	// Deduplicate typedefs by name - keep only the first occurrence
	typedefsSeen := make(map[string]bool)
	deduplicatedTypedefs := make([]*occ2go.ParsedTypedef, 0, len(g.Typedefs))
	for _, typedef := range g.Typedefs {
		if !typedefsSeen[typedef.Name] {
			typedefsSeen[typedef.Name] = true
			deduplicatedTypedefs = append(deduplicatedTypedefs, typedef)
		}
	}
	g.Typedefs = deduplicatedTypedefs

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

	// Extract missing enums automatically from undefined types
	if os.Getenv("EXTRACT_MISSING_ENUMS") == "1" {
		if os.Getenv("VERBOSE") == "1" {
			fmt.Fprintf(os.Stderr, "EXTRACT_MISSING_ENUMS=1, calling ExtractMissingEnums() for %s\n", g.Framework)
		}
		extractedEnums, err := g.ExtractMissingEnums()
		if err != nil && os.Getenv("VERBOSE") == "1" {
			fmt.Fprintf(os.Stderr, "ExtractMissingEnums error: %v\n", err)
		}
		if err == nil && len(extractedEnums) > 0 {
			if os.Getenv("VERBOSE") == "1" {
				fmt.Fprintf(os.Stderr, "Auto-extracted %d missing enums for %s\n", len(extractedEnums), g.Framework)
			}
			// Add extracted enums to existing enums
			g.Enums = append(g.Enums, extractedEnums...)
		} else if os.Getenv("VERBOSE") == "1" {
			fmt.Fprintf(os.Stderr, "No enums extracted for %s\n", g.Framework)
		}
	}

	// Build indexes for O(1) lookups
	// This significantly improves performance for IsClassType, IsEnumType, IsTypedefType
	g.classIndex = make(map[string]*occ2go.ParsedClass, len(g.Classes))
	for _, class := range g.Classes {
		g.classIndex[class.Name] = class
		// Also index by stripped name for easier lookup
		stripped := classToStructName(class.Name)
		if stripped != class.Name {
			g.classIndex[stripped] = class
		}
	}

	g.enumIndex = make(map[string]*occ2go.ParsedEnum, len(g.Enums))
	for _, enum := range g.Enums {
		g.enumIndex[enum.Name] = enum
		// Also index by stripped name
		stripped := stripObjCPrefix(enum.Name)
		if stripped != enum.Name {
			g.enumIndex[stripped] = enum
		}
	}

	g.typedefIndex = make(map[string]*occ2go.ParsedTypedef, len(g.Typedefs))
	for _, typedef := range g.Typedefs {
		g.typedefIndex[typedef.Name] = typedef
		if os.Getenv("DEBUG_TIMEINTERVAL") == "1" && typedef.Name == "TimeInterval" {
			fmt.Fprintf(os.Stderr, "DEBUG: Building typedef index - found TimeInterval\n")
		}
	}
	if os.Getenv("DEBUG_TIMEINTERVAL") == "1" {
		fmt.Fprintf(os.Stderr, "DEBUG: Built typedef index with %d entries, TimeInterval present = %v\n", len(g.typedefIndex), g.typedefIndex["TimeInterval"] != nil)
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
// Note: prepare() should be called by the caller before calling this method
func (g *Generator) GenerateTxtarFromModule(w io.Writer) error {
	// Load all templates as named templates
	moduleContent, err := getTemplateVariant("module", g.Variant)
	if err != nil {
		return err
	}

	// Create master template and parse all sub-templates
	// Chain Funcs calls: first core "dumb" functions, then GeneratorFuncs methods
	gf := GeneratorFuncs{g}
	tmpl := template.New("module").Funcs(templateFuncs).Funcs(gf.Funcs())

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
