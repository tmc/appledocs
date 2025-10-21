package main

import (
	"regexp"
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// UndefinedType represents a type that's referenced but not defined
type UndefinedType struct {
	Name       string // e.g., "MKLocalSearchCompletion"
	Framework  string // e.g., "MapKit"
	References int    // how many times it's referenced
}

// CollectUndefinedTypes scans all methods and properties for types that are referenced
// but not defined in the generated code.
// IMPORTANT: This only scans methods/properties that will actually be generated,
// using the same filtering logic as the templates.
func (g *Generator) CollectUndefinedTypes() map[string]*UndefinedType {
	undefined := make(map[string]*UndefinedType)

	// Create a set of typedef names to exclude from undefined types
	typedefNames := make(map[string]bool)
	for _, typedef := range g.Typedefs {
		if typedef.Name != "" {
			typedefNames[typedef.Name] = true
		}
	}

	// Collect from class methods (only those that will be generated)
	for _, cls := range g.Classes {
		// Get the methods that will actually be generated
		// This mirrors the template logic which uses filterPropertyMethods, prepareClassMethods, etc.
		generatedMethods := g.getGeneratedMethods(cls)

		for _, method := range generatedMethods {
			// Check return type - use MAPPED type, not raw ObjC type
			if method.ReturnType != "" && method.ReturnType != "void" {
				mappedType := resolveType(g.Framework, method.ReturnType)
				collectTypeReferences(mappedType, undefined, g.Framework, typedefNames)
			}

			// Check parameter types - use MAPPED types
			for _, param := range method.Parameters {
				if param.Type != "" {
					mappedType := resolveType(g.Framework, param.Type)
					collectTypeReferences(mappedType, undefined, g.Framework, typedefNames)
				}
			}
		}

		// Check properties - properties are always generated
		for _, prop := range cls.Properties {
			if prop.Type != "" {
				mappedType := resolveType(g.Framework, prop.Type)
				collectTypeReferences(mappedType, undefined, g.Framework, typedefNames)
			}
		}
	}

	// Collect from functions
	for _, fn := range g.Functions {
		if fn.ReturnType != "" && fn.ReturnType != "void" {
			mappedType := resolveType(g.Framework, fn.ReturnType)
			collectTypeReferences(mappedType, undefined, g.Framework, typedefNames)
		}
		for _, param := range fn.Parameters {
			if param.Type != "" {
				mappedType := resolveType(g.Framework, param.Type)
				collectTypeReferences(mappedType, undefined, g.Framework, typedefNames)
			}
		}
	}

	// Filter out types that are defined
	defined := g.getDefinedTypes()
	for name := range undefined {
		if defined[name] {
			delete(undefined, name)
		}
	}

	return undefined
}

// getGeneratedMethods returns only the methods that will actually be generated
// for a class, using the same filtering logic as the templates.
func (g *Generator) getGeneratedMethods(cls *occ2go.ParsedClass) []*occ2go.ParsedMethod {
	if cls == nil {
		return nil
	}

	result := make([]*occ2go.ParsedMethod, 0)

	// Get property methods (filtered to exclude property accessors that are auto-generated)
	propertyMethods := filterPropertyMethods(cls)
	result = append(result, propertyMethods...)

	// Get init methods (constructor-style methods)
	initMethods := prepareInitMethodsWithClassName(cls.Name, cls.Methods)
	result = append(result, initMethods...)

	// Get class methods
	classMethods := prepareClassMethods(cls.Methods)
	result = append(result, classMethods...)

	// Deduplicate in case there's overlap
	seen := make(map[string]bool)
	deduplicated := make([]*occ2go.ParsedMethod, 0, len(result))
	for _, m := range result {
		key := m.Selector + ":" + string(rune(boolToInt(m.IsClassMethod)))
		if !seen[key] {
			seen[key] = true
			deduplicated = append(deduplicated, m)
		}
	}

	return deduplicated
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// collectTypeReferences extracts type names from a type string and adds them to undefined.
// IMPORTANT: This should only be called on MAPPED Go types (after resolveType), not raw ObjC types.
func collectTypeReferences(typeStr string, undefined map[string]*UndefinedType, framework string, typedefNames map[string]bool) {
	if typeStr == "" || typeStr == "void" {
		return
	}

	// Skip types that are defined as typedefs
	if typedefNames != nil && typedefNames[typeStr] {
		return
	}

	// Skip types that are package-qualified (e.g., "coregraphics.CGColorRef")
	// These are imports from other frameworks and don't need undefined type declarations
	if strings.Contains(typeStr, ".") {
		return
	}

	// Skip primitive/builtin types
	if isBuiltinType(typeStr) {
		return
	}

	// Skip if it's unsafe.Pointer or other special cases
	if typeStr == "unsafe.Pointer" || strings.HasPrefix(typeStr, "[]") || strings.HasPrefix(typeStr, "*") {
		return
	}

	// Extract all CamelCase identifiers
	// Pattern matches: Word followed by optional CamelCase parts
	pattern := regexp.MustCompile(`\b([A-Z][a-zA-Z0-9]*)\b`)
	matches := pattern.FindAllString(typeStr, -1)

	for _, match := range matches {
		// Skip common built-ins and primitive types
		if isBuiltinType(match) {
			continue
		}

		// Skip if this match is a typedef
		if typedefNames != nil && typedefNames[match] {
			continue
		}

		// Skip if already tracking
		if _, exists := undefined[match]; exists {
			undefined[match].References++
		} else {
			undefined[match] = &UndefinedType{
				Name:       match,
				Framework:  framework,
				References: 1,
			}
		}
	}
}

// getDefinedTypes returns a map of types that are defined in generated code
func (g *Generator) getDefinedTypes() map[string]bool {
	defined := make(map[string]bool)

	// Add built-in types
	for _, name := range []string{
		"void", "bool", "int", "uint", "int8", "uint8", "int16", "uint16",
		"int32", "uint32", "int64", "uint64", "float32", "float64",
		"string", "byte", "rune", "error", "interface{}",
		"unsafe.Pointer", "unsafe",
	} {
		defined[name] = true
	}

	// Add Objective-C types that map to Go built-ins
	// These should never be generated as fallback types
	for _, name := range []string{
		"BOOL",         // maps to bool
		"NSInteger",    // maps to int
		"NSUInteger",   // maps to uint
		"CGFloat",      // maps to float64
		"TimeInterval", // maps to float64
	} {
		defined[name] = true
	}

	// Add types from current framework's classes and protocols
	for _, cls := range g.Classes {
		defined[cls.Name] = true
		// Also add Go struct name
		defined[classToStructName(cls.Name)] = true
	}

	for _, proto := range g.Protocols {
		defined[proto.Name] = true
		// Add interface name
		defined["I"+proto.Name] = true
	}

	// Add enum types
	for _, enum := range g.Enums {
		defined[enum.Name] = true
		// Also add the stripped version (e.g., "SMAppServiceStatus" becomes "AppServiceStatus")
		defined[stripObjCPrefix(enum.Name)] = true
	}

	// Add ref types
	for _, ref := range g.refTypes {
		defined[ref] = true
	}

	// Add types from cross-framework registry that belong to OTHER frameworks
	// This prevents generating fallback types for classes that exist in other frameworks
	// e.g., UniformTypeIdentifiers shouldn't define NSArray - it exists in Foundation
	// But we only mark them as defined if they're from a DIFFERENT framework
	currentFrameworkPkg := strings.ToLower(g.Framework)
	for typeName, pkgName := range crossFrameworkTypeRegistry {
		if pkgName != currentFrameworkPkg {
			defined[typeName] = true
			// Also add the NS-prefixed version (e.g., if registry has "Error", also mark "NSError" as defined)
			// This handles cases where docs reference "NSError" but Foundation generates "Error"
			objcName := "NS" + typeName
			defined[objcName] = true
		}
	}

	// Add framework-specific types that are defined in templates or as classes
	frameworkSpecificTypes := map[string][]string{
		"CoreGraphics": {"CGFloat", "CGPoint", "CGSize", "CGRect", "CGAffineTransform", "CGVector", "Range", "Size", "Point", "Rect"},
		"Foundation":   {"TimeInterval", "Point", "Size", "Rect", "Range", "RectEdge"},
		"AppKit":       {"WindowStyleMask", "BackingStoreType", "WindowOrderingMode", "WindowLevel", "EventType", "EventModifierFlags"},
		"QuartzCore":   {"CGFloat"},
		"ObjectiveC":   {"Protocol"}, // Protocol is a class, not a fallback type
	}

	if types, ok := frameworkSpecificTypes[g.Framework]; ok {
		for _, t := range types {
			defined[t] = true
		}
	}

	return defined
}

// isBuiltinType checks if a type is a Go built-in
func isBuiltinType(name string) bool {
	builtins := map[string]bool{
		"void": true, "bool": true, "byte": true, "rune": true,
		"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
		"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true,
		"uintptr": true, "float32": true, "float64": true, "complex64": true, "complex128": true,
		"string": true, "error": true, "any": true,
		// Common patterns
		"Pointer": true, "unsafe": true,
	}
	return builtins[name]
}

// GetUndefinedTypesForTemplate returns undefined types formatted for template use
func (g *Generator) GetUndefinedTypesForTemplate() []UndefinedType {
	undefined := g.CollectUndefinedTypes()
	result := make([]UndefinedType, 0, len(undefined))

	// Convert map to sorted slice
	for _, ut := range undefined {
		result = append(result, *ut)
	}

	// Sort by name for consistent output
	// Simple bubble sort for stability
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i].Name > result[j].Name {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}

// RequiredImports analyzes types to determine which framework imports are needed
func (g *Generator) RequiredImports() map[string]bool {
	imports := make(map[string]bool)

	// Check if we need CoreGraphics
	if g.hasCoreGraphicsTypes() {
		if g.Framework != "CoreGraphics" {
			imports["coregraphics"] = true
		}
	}

	// Check if we need Foundation
	if g.hasFoundationTypes() {
		if g.Framework != "Foundation" {
			imports["foundation"] = true
		}
	}

	return imports
}

// hasCoreGraphicsTypes checks if any method uses CoreGraphics types
func (g *Generator) hasCoreGraphicsTypes() bool {
	cgTypes := map[string]bool{
		"CGColorRef": true, "CGContextRef": true, "CGPDFPageRef": true,
		"CGPoint": true, "CGSize": true, "CGRect": true, "CGFloat": true,
	}

	for _, cls := range g.Classes {
		for _, method := range cls.Methods {
			if method.ReturnType != "" && cgTypes[method.ReturnType] {
				return true
			}
			for _, param := range method.Parameters {
				if cgTypes[param.Type] {
					return true
				}
			}
		}
	}

	for _, fn := range g.Functions {
		if fn.ReturnType != "" && cgTypes[fn.ReturnType] {
			return true
		}
		for _, param := range fn.Parameters {
			if cgTypes[param.Type] {
				return true
			}
		}
	}

	return false
}

// hasFoundationTypes checks if any method uses Foundation types
func (g *Generator) hasFoundationTypes() bool {
	foundationTypes := map[string]bool{
		"NSString": true, "NSArray": true, "NSDictionary": true,
		"NSNumber": true, "NSDate": true, "NSURL": true,
	}

	for _, cls := range g.Classes {
		for _, method := range cls.Methods {
			if method.ReturnType != "" && foundationTypes[method.ReturnType] {
				return true
			}
			for _, param := range method.Parameters {
				if foundationTypes[param.Type] {
					return true
				}
			}
		}
	}

	return false
}
