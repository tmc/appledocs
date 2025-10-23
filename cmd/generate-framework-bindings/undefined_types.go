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
	BaseType   string // Inferred base type (e.g., "uint" for NS*Options, "int" for enums)
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
				mappedType := mapObjCTypeToGo(method.ReturnType, g.Framework)
				// If mapping failed, collect the original type
				if mappedType == "unsafe.Pointer" && !strings.Contains(method.ReturnType, "*") && !strings.Contains(method.ReturnType, "^") && !strings.Contains(method.ReturnType, "Block") {
					typeToCollect := stripObjCPrefix(method.ReturnType)
					collectTypeReferences(typeToCollect, undefined, g.Framework, typedefNames)
				} else {
					collectTypeReferences(mappedType, undefined, g.Framework, typedefNames)
					// Also collect the stripped ObjC type name for objc.Send[T] usage
					// When mappedType is objc.IObject or cross-framework, we still need the unqualified type
					if strings.HasPrefix(mappedType, "objc.") || strings.Contains(mappedType, ".") {
						strippedObjcType := stripObjCPrefix(method.ReturnType)
						collectTypeReferences(strippedObjcType, undefined, g.Framework, typedefNames)
					}
				}
			}

			// Check parameter types - use MAPPED types
			for _, param := range method.Parameters {
				if param.Type != "" {
					mappedType := mapObjCTypeToGo(param.Type, g.Framework)
					// If mapping failed, collect the original type
					if mappedType == "unsafe.Pointer" && !strings.Contains(param.Type, "*") && !strings.Contains(param.Type, "^") && !strings.Contains(param.Type, "Block") {
						typeToCollect := stripObjCPrefix(param.Type)
						collectTypeReferences(typeToCollect, undefined, g.Framework, typedefNames)
					} else {
						collectTypeReferences(mappedType, undefined, g.Framework, typedefNames)
					}
				}
			}
		}

		// Check properties - properties are always generated
		for _, prop := range cls.Properties {
			if prop.Type != "" {
				mappedType := mapObjCTypeToGo(prop.Type, g.Framework)
				// If mapping failed (returned unsafe.Pointer for non-pointer types),
				// collect the original ObjC type instead so it can be added as fallback
				if mappedType == "unsafe.Pointer" && !strings.Contains(prop.Type, "*") && !strings.Contains(prop.Type, "^") && !strings.Contains(prop.Type, "Block") {
					// Strip NS/CG/CA prefix from the ObjC type before collecting
					typeToCollect := prop.Type
					if strings.HasPrefix(typeToCollect, "NS") && len(typeToCollect) > 2 && typeToCollect[2] >= 'A' && typeToCollect[2] <= 'Z' {
						typeToCollect = typeToCollect[2:]
					} else if strings.HasPrefix(typeToCollect, "CG") && len(typeToCollect) > 2 && typeToCollect[2] >= 'A' && typeToCollect[2] <= 'Z' {
						typeToCollect = typeToCollect[2:]
					} else if strings.HasPrefix(typeToCollect, "CA") && len(typeToCollect) > 2 && typeToCollect[2] >= 'A' && typeToCollect[2] <= 'Z' {
						typeToCollect = typeToCollect[2:]
					}
					collectTypeReferences(typeToCollect, undefined, g.Framework, typedefNames)
				} else {
					collectTypeReferences(mappedType, undefined, g.Framework, typedefNames)
				}
			}
		}
	}

	// Collect from functions
	for _, fn := range g.Functions {
		if fn.ReturnType != "" && fn.ReturnType != "void" {
			mappedType := mapObjCTypeToGo(fn.ReturnType, g.Framework)
			collectTypeReferences(mappedType, undefined, g.Framework, typedefNames)
		}
		for _, param := range fn.Parameters {
			if param.Type != "" {
				mappedType := mapObjCTypeToGo(param.Type, g.Framework)
				collectTypeReferences(mappedType, undefined, g.Framework, typedefNames)
			}
		}
	}

	// Filter out types that are defined
	defined := g.getDefinedTypes()

	for name := range undefined {
		Debug.Undefined("filtering undefined type", name, g.Framework,
			"framework", g.Framework,
			"name", name,
			"defined", defined[name])
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

	// Get ALL methods - both instance and class methods
	// This ensures we collect type information from all generated methods
	for _, method := range cls.Methods {
		result = append(result, method)
	}

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

	Debug.Undefined("collectTypeReferences entry", typeStr, framework,
		"typeStr", typeStr,
		"framework", framework)
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

	// Skip if it's unsafe.Pointer
	if typeStr == "unsafe.Pointer" {
		return
	}

	// Handle slice and pointer types by recursively processing the element type
	if strings.HasPrefix(typeStr, "[]") {
		// Extract the element type from the slice
		elementType := strings.TrimPrefix(typeStr, "[]")
		collectTypeReferences(elementType, undefined, framework, typedefNames)
		return
	}
	if strings.HasPrefix(typeStr, "*") {
		// Extract the element type from the pointer
		elementType := strings.TrimPrefix(typeStr, "*")
		collectTypeReferences(elementType, undefined, framework, typedefNames)
		return
	}

	// Extract all CamelCase identifiers
	// Pattern matches: Word followed by optional CamelCase parts
	pattern := regexp.MustCompile(`\b([A-Z][a-zA-Z0-9]*)\b`)
	matches := pattern.FindAllString(typeStr, -1)

	for _, match := range matches {
		// Strip ObjC prefixes (NS, CG, CA) to match what TypeToInterfaceType will output
		// This ensures undefined types use the same names as the generated code
		// Example: NSDecodingFailurePolicy -> DecodingFailurePolicy
		strippedMatch := match
		if strings.HasPrefix(match, "NS") && len(match) > 2 && match[2] >= 'A' && match[2] <= 'Z' {
			strippedMatch = match[2:]
		} else if strings.HasPrefix(match, "CG") && len(match) > 2 && match[2] >= 'A' && match[2] <= 'Z' {
			strippedMatch = match[2:]
		} else if strings.HasPrefix(match, "CA") && len(match) > 2 && match[2] >= 'A' && match[2] <= 'Z' {
			strippedMatch = match[2:]
		}

		Debug.Undefined("match found", match, typeStr,
			"match", match,
			"strippedMatch", strippedMatch,
			"isBuiltin", isBuiltinType(strippedMatch),
			"isTypedef", typedefNames != nil && typedefNames[strippedMatch],
			"framework", framework)
		// Skip common built-ins and primitive types (check stripped version)
		if isBuiltinType(strippedMatch) {
			continue
		}

		// Skip if this match is a typedef (check both original and stripped)
		if typedefNames != nil && (typedefNames[match] || typedefNames[strippedMatch]) {
			continue
		}

		// Skip if already tracking (use stripped name as key)
		if _, exists := undefined[strippedMatch]; exists {
			undefined[strippedMatch].References++
			Debug.Undefined("incrementing reference count", strippedMatch, framework,
				"match", match,
				"strippedMatch", strippedMatch,
				"references", undefined[strippedMatch].References)
		} else {
			undefined[strippedMatch] = &UndefinedType{
				Name:       strippedMatch, // Use stripped name
				Framework:  framework,
				References: 1,
				BaseType:   inferBaseType(strippedMatch),
			}
			Debug.Undefined("adding new undefined type", strippedMatch, framework,
				"match", match,
				"strippedMatch", strippedMatch,
				"framework", framework)
		}
	}
}

// getDefinedTypes returns a map of types that are defined in generated code
func (g *Generator) getDefinedTypes() map[string]bool {
	defined := make(map[string]bool)

	Debug.Undefined("getDefinedTypes entry", g.Framework, "",
		"framework", g.Framework,
		"classCount", len(g.Classes))
	for i, cls := range g.Classes {
		if i < 10 {
			Debug.Undefined("class sample", cls.Name, g.Framework,
				"index", i,
				"name", cls.Name,
				"framework", g.Framework)
		}
	}

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
		structName := classToStructName(cls.Name)
		defined[structName] = true
		Debug.Undefined("class type defined", cls.Name, structName,
			"framework", g.Framework,
			"className", cls.Name,
			"structName", structName)
	}

	for _, proto := range g.Protocols {
		defined[proto.Name] = true
		// Add interface name
		defined["I"+proto.Name] = true
	}

	// Add enum types
	for _, enum := range g.Enums {
		Debug.Undefined("marking enum as defined", enum.Name, g.Framework,
			"framework", g.Framework,
			"enumName", enum.Name)
		defined[enum.Name] = true
		// Also add the stripped version (e.g., "SMAppServiceStatus" becomes "AppServiceStatus")
		stripped := stripObjCPrefix(enum.Name)
		defined[stripped] = true
		Debug.Undefined("marking stripped enum as defined", stripped, g.Framework,
			"framework", g.Framework,
			"stripped", stripped,
			"original", enum.Name)
	}

	// Add struct types
	for _, strct := range g.Structs {
		defined[strct.Name] = true
		// Also add the stripped version (e.g., "CFRange" becomes "Range")
		stripped := stripObjCPrefix(strct.Name)
		defined[stripped] = true
		Debug.Undefined("marking struct as defined", strct.Name, g.Framework,
			"framework", g.Framework,
			"structName", strct.Name,
			"stripped", stripped)
	}

	// Add ref types
	for _, ref := range g.refTypes {
		defined[ref] = true
	}

	// Add types from cross-framework registry
	// ONLY mark types that belong to THIS framework as defined
	// Types from other frameworks should NOT be marked as defined, so they can be
	// generated as undefined type aliases (needed for objc.Send[T] calls)
	currentFrameworkPkg := strings.ToLower(g.Framework)
	for typeName, pkgName := range crossFrameworkTypeRegistry {
		// Only mark if it belongs to THIS framework
		if pkgName == currentFrameworkPkg {
			defined[typeName] = true
			// Also mark the NS-prefixed version
			objcName := "NS" + typeName
			Debug.Undefined("cross-framework registry (current)", typeName, objcName,
				"framework", g.Framework,
				"typeName", typeName,
				"objcName", objcName,
				"pkgName", pkgName)
			defined[objcName] = true
		}
		// Do NOT mark types from other frameworks as defined
		// They need to be generated as undefined type aliases
	}

	// Add framework-specific types that are defined manually in non-generated files
	// These types are provided in manually-written files like rect_types.go
	frameworkSpecificTypes := map[string][]string{
		"CoreGraphics": {"Float", "Point", "Size", "Rect", "AffineTransform", "AffineTransformComponents"},
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
	Debug.Undefined("GetUndefinedTypesForTemplate called", g.Framework, "",
		"classCount", len(g.Classes),
		"framework", g.Framework)
	undefined := g.CollectUndefinedTypes()
	for name := range undefined {
		Debug.Undefined("found undefined type", name, g.Framework,
			"framework", g.Framework,
			"name", name)
	}
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

// inferBaseType infers the appropriate base type for an undefined type
func inferBaseType(typeName string) string {
	// Default to int for undefined enum types
	// These are types that exist in docs but weren't extracted
	return "int"
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
