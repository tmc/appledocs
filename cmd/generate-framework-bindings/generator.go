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

// GeneratorConfig contains all configuration parameters for the generator
type GeneratorConfig struct {
	Framework        string
	PackageName      string
	InputDir         string
	OutputModule     string
	Variant          string
	WithRefMethods   bool
	GenerateTests    bool
	GenerateExamples bool
	DebugEnabled     bool
}

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

	DebugEnabled bool

	Functions []*occ2go.ParsedFunction
	Classes   []*occ2go.ParsedClass
	Protocols []*occ2go.ParsedProtocol
	Enums     []*occ2go.ParsedEnum
	Typedefs  []*occ2go.ParsedTypedef
	Constants []*occ2go.ParsedConstant
	Structs   []*occ2go.ParsedStruct

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

	// Undefined types (cached from CollectUndefinedTypes())
	undefinedTypes map[string]*UndefinedType

	// Error collection
	Errors []error
}

// NewGenerator creates a new Generator instance from a config
func NewGenerator(config GeneratorConfig) *Generator {
	return &Generator{
		Framework:        config.Framework,
		PackageName:      config.PackageName,
		InputDir:         config.InputDir,
		OutputModule:     config.OutputModule,
		Variant:          config.Variant,
		WithRefMethods:   config.WithRefMethods,
		GenerateTests:    config.GenerateTests,
		GenerateExamples: config.GenerateExamples,
		DebugEnabled:     config.DebugEnabled,
		Errors:           make([]error, 0),
	}
}

// SetParsedData sets the parsed symbol data on the generator
func (g *Generator) SetParsedData(
	functions []*occ2go.ParsedFunction,
	classes []*occ2go.ParsedClass,
	protocols []*occ2go.ParsedProtocol,
	enums []*occ2go.ParsedEnum,
	typedefs []*occ2go.ParsedTypedef,
	constants []*occ2go.ParsedConstant,
	structs []*occ2go.ParsedStruct,
) {
	g.Functions = functions
	g.Classes = classes
	g.Protocols = protocols
	g.Enums = enums
	g.Typedefs = typedefs
	g.Constants = constants
	g.Structs = structs
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
	// Try exact match first
	_, ok := g.typedefIndex[typeName]
	if ok {
		Debug.TimeInterval("IsTypedefType check (exact)", typeName, "",
			"typeName", typeName,
			"found", true,
			"indexSize", len(g.typedefIndex))
		return true
	}

	// Also try lowercase version (for title-cased typedef names like "Unichar")
	// The index contains original names from Apple docs (e.g., "unichar")
	// but mapObjCTypeToGo may return title-cased names (e.g., "Unichar")
	lowercaseTypeName := ""
	if len(typeName) > 0 {
		lowercaseTypeName = strings.ToLower(typeName[:1]) + typeName[1:]
	}
	if lowercaseTypeName != "" && lowercaseTypeName != typeName {
		_, ok = g.typedefIndex[lowercaseTypeName]
		if ok {
			Debug.TimeInterval("IsTypedefType check (lowercase)", typeName, "",
				"typeName", typeName,
				"lowercaseTypeName", lowercaseTypeName,
				"found", true,
				"indexSize", len(g.typedefIndex))
			return true
		}
	}

	Debug.TimeInterval("IsTypedefType check (not found)", typeName, "",
		"typeName", typeName,
		"tried", lowercaseTypeName,
		"found", false,
		"indexSize", len(g.typedefIndex))
	return false
}

// GetTypedefGoName returns the Go name for a typedef given its ObjC name.
// For example: "unichar" -> "Unichar", "NSInteger" -> "Integer"
// Returns empty string if the typedef is not found.
func (g *Generator) GetTypedefGoName(objcName string) string {
	if objcName == "" {
		return ""
	}

	// Look up the typedef in the index
	var typedef *occ2go.ParsedTypedef
	if g.typedefIndex != nil {
		typedef = g.typedefIndex[objcName]
	} else {
		// Fallback to linear search if index not built
		for _, td := range g.Typedefs {
			if td.Name == objcName {
				typedef = td
				break
			}
		}
	}

	if typedef == nil {
		return ""
	}

	// Apply the same naming logic as the typedef template
	// See templates.txtar line 1481
	if strings.ToLower(g.Framework) == "objectivec" {
		return typedef.Name
	}
	return titleString(stripObjCPrefix(typedef.Name))
}

// IsStringBasedTypedef checks if a typedef is based on NSString *.
// String-based typedefs (like NSCalendarIdentifier, NSErrorDomain) need special
// handling in method returns because objc.Send returns a String struct but the
// typedef is actually a string type alias.
func (g *Generator) IsStringBasedTypedef(typeName string) bool {
	if typeName == "" {
		return false
	}

	// Look up the typedef
	var typedef *occ2go.ParsedTypedef
	if g.typedefIndex != nil {
		typedef = g.typedefIndex[typeName]
	} else {
		// Fallback to linear search if index not built
		for _, td := range g.Typedefs {
			if td.Name == typeName {
				typedef = td
				break
			}
		}
	}

	if typedef == nil {
		return false
	}

	// Check if base type is NSString * or NSString*
	baseType := strings.TrimSpace(typedef.BaseType)
	return baseType == "NSString *" || baseType == "NSString*"
}

// TypeToInterfaceType converts a struct type name to its interface type name using
// data-driven type checking. For example: "Data" becomes "IData", "Window" becomes "IWindow".
// For qualified types: "foundation.Coder" becomes "foundation.ICoder".
// Types that don't have interfaces (primitives, slices, enums, typedefs, structs) are returned unchanged.
// GeneratorFuncs wraps a Generator and provides template functions with access to
// the Generator's data-driven type checking (enum/typedef/class indices).
// This allows templates to use accurate type resolution instead of heuristics.
func (g *Generator) TypeToInterfaceType(goType string) string {
	// Foundation types that should map to primitives
	// TimeInterval is NSTimeInterval which is typedef for double
	if goType == "TimeInterval" || goType == "NSTimeInterval" {
		return "float64"
	}

	// Don't convert function types (blocks)
	if strings.HasPrefix(goType, "func(") {
		return goType
	}

	// Handle slices FIRST - before checking for qualified types
	// This ensures []foundation.Number is processed correctly (element type contains ".")
	// NOTE: For slices, we keep the CONCRETE type ([]Connection not []IConnection) because
	// in Go you cannot convert between []ConcreteType and []InterfaceType,even if ConcreteType implements InterfaceType.
	// We process the element type but strip any "I" interface prefix from the result.
	if strings.HasPrefix(goType, "[]") {
		elemType := goType[2:]
		convertedElemType := g.TypeToInterfaceType(elemType)

		// If the element type was converted to IObject (hierarchy violation fallback), use []objc.ID
		if convertedElemType == "IObject" || convertedElemType == "objectivec.IObject" {
			return "[]objc.ID"
		}

		// Check for hierarchy violations in qualified element types
		if strings.Contains(convertedElemType, ".") {
			parts := strings.SplitN(convertedElemType, ".", 2)
			if len(parts) == 2 {
				pkg := parts[0]
				typeName := parts[1]

				// Check framework hierarchy
				currentLevel := getFrameworkLevel(strings.ToLower(g.Framework))
				targetLevel := getFrameworkLevel(pkg)
				if currentLevel >= 0 && targetLevel > currentLevel {
					// Hierarchy violation - use []objc.ID instead
					return "[]objc.ID"
				}

				// Strip I prefix from qualified type: "foundation.IData" -> "foundation.Data"
				// But DO NOT strip from objc.ID, objc.IObject, objc.IClass, etc. (objc runtime types)
				if pkg != "objc" && len(typeName) > 1 && typeName[0] == 'I' && typeName[1] >= 'A' && typeName[1] <= 'Z' {
					convertedElemType = pkg + "." + typeName[1:]
				}
			}
		} else if len(convertedElemType) > 1 && convertedElemType[0] == 'I' && convertedElemType[1] >= 'A' && convertedElemType[1] <= 'Z' {
			// Strip I prefix: "IConnection" -> "Connection"
			convertedElemType = convertedElemType[1:]
		}
		return "[]" + convertedElemType
	}

	// Handle qualified types (e.g., "foundation.Coder" -> "foundation.ICoder")
	if strings.Contains(goType, ".") {
		parts := strings.SplitN(goType, ".", 2)
		if len(parts) == 2 {
			pkg := parts[0]
			typeName := parts[1]

			// Special case: In ObjectiveC framework, objc.ID should become IObject
			if pkg == "objc" && typeName == "ID" && strings.ToLower(g.Framework) == "objectivec" {
				return "IObject"
			}

			// Don't convert runtime types (objc.ID, unsafe.Pointer, etc.)
			if pkg == "objc" || pkg == "unsafe" {
				return goType
			}

			// If the package matches the current framework, strip the qualification
			// e.g., "objectivec.IObject" in ObjectiveC framework becomes "IObject"
			if pkg == strings.ToLower(g.Framework) {
				return typeName
			}

			// For objectivec package types in other frameworks, keep them
			if pkg == "objectivec" {
				return goType
			}

			// Check for framework hierarchy violations (lower-level importing higher-level)
			// This must come BEFORE struct check to ensure hierarchy violations fall back to IObject
			// For ObjectiveC framework (level 0), use IObject (unqualified) for ANY other framework
			// For other frameworks, use objc.IObject (the alias to github.com/tmc/appledocs/generated/objc)
			currentLevel := getFrameworkLevel(strings.ToLower(g.Framework))
			targetLevel := getFrameworkLevel(pkg)
			fallbackType := "objc.IObject"
			if strings.ToLower(g.Framework) == "objectivec" {
				fallbackType = "IObject"
			}
			if targetLevel > currentLevel {
				// Higher-level framework (or unknown framework, which defaults to level 999)
				Debug.TypeMap("framework layering violation in TypeToInterfaceType", typeName, fallbackType,
					"framework", g.Framework,
					"currentLevel", currentLevel,
					"targetPkg", pkg,
					"targetLevel", targetLevel,
					"objcType", goType)
				return fallbackType
			}

			// After hierarchy violation checks pass, check if this is a struct type
			// Structs should remain as structs, not become interfaces (e.g., coregraphics.AffineTransform)
			// Use the struct registry for data-driven decision making
			if crossFrameworkStructRegistry[goType] {
				// This is a registered struct type
				// Don't convert to interface
				Debug.TypeMap("TypeToInterfaceType: struct type, not converting", goType, goType,
					"goType", goType,
					"reason", "registered as struct")
				return goType
			}

			// Recursively convert the type part
			interfaceType := g.TypeToInterfaceType(typeName)

			// IMPORTANT: If the recursive call returned a fallback type (objc.IObject or objectivec.IObject),
			// don't re-qualify it with the original package. Fallback types should remain as-is.
			// This fixes appledocs-496 where uniformtypeidentifiers.UTType → objc.IObject
			// was being incorrectly transformed to uniformtypeidentifiers.objc.IObject
			if strings.HasPrefix(interfaceType, "objc.IObject") || strings.HasPrefix(interfaceType, "objectivec.IObject") {
				return interfaceType
			}

			return pkg + "." + interfaceType
		}
		return goType + " /* malformed qualified type */"
	}

	// Handle pointer to typedef (e.g., "unichar *" -> "*Unichar")
	// Check for space-separated pointer syntax from ObjC
	if strings.HasSuffix(goType, " *") {
		baseType := strings.TrimSuffix(goType, " *")
		if g.IsTypedefType(baseType) {
			if goName := g.GetTypedefGoName(baseType); goName != "" {
				return "*" + goName
			}
		}
		// Not a typedef, fall through to general pointer handling
	}

	// Don't convert primitives, pointers, maps
	if strings.HasPrefix(goType, "*") ||
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

	// If it already starts with I and next char is uppercase, it's already an interface
	if strings.HasPrefix(goType, "I") && len(goType) > 1 && goType[1] >= 'A' && goType[1] <= 'Z' {
		return goType
	}

	// DATA-DRIVEN: Check if this is an enum or typedef - don't convert those
	// For enums, return the goType as-is (which is already the stripped/Go name)
	// The goType parameter has already been mapped by mapObjCTypeToGo
	if g.IsEnumType(goType) {
		Debug.EnumType("ToInterfaceType enum check", goType, "",
			"goType", goType,
			"isEnum", true,
			"returning", goType)
		return goType // Return the mapped Go type name, not the original ObjC name
	}
	if g.IsTypedefType(goType) {
		// Try to get the Go typedef name
		if goName := g.GetTypedefGoName(goType); goName != "" {
			Debug.TypeMap("TypeToInterfaceType typedef resolved", goType, goName,
				"objcType", goType,
				"goName", goName)
			return goName
		}
		// Fallback if typedef lookup fails (shouldn't happen but defensive)
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
	Debug.TypeMap("TypeToInterfaceType class check", goType, baseType,
		"goType", goType,
		"baseType", baseType,
		"isClassGoType", g.IsClassType(goType),
		"isClassBaseType", g.IsClassType(baseType))
	if !g.IsClassType(goType) && !g.IsClassType(baseType) {
		// Check if this type is known in the cross-framework registry
		// If it is, it's a class from another framework
		// First, try the type as-is, then try the unqualified name if it's qualified
		var lookupType string = goType
		var frameworkPkg string
		var found bool

		// Try direct lookup first
		if frameworkPkg, found = crossFrameworkTypeRegistry[goType]; !found {
			// If goType is already qualified (e.g., "foundation.NSString"),
			// extract the unqualified name and try that
			if strings.Contains(goType, ".") {
				parts := strings.SplitN(goType, ".", 2)
				if len(parts) == 2 {
					lookupType = parts[1]
					frameworkPkg, found = crossFrameworkTypeRegistry[lookupType]
				}
			}
		}

		if found {
			// Check if this would violate the framework hierarchy
			currentLevel := getFrameworkLevel(strings.ToLower(g.Framework))
			targetLevel := getFrameworkLevel(frameworkPkg)

			if currentLevel >= 0 && targetLevel > currentLevel {
				// Hierarchy violation - use objc.IObject
				Debug.TypeMap("cross-framework type violates hierarchy", goType, "objc.IObject",
					"goType", goType,
					"currentFramework", g.Framework,
					"currentLevel", currentLevel,
					"targetFramework", frameworkPkg,
					"targetLevel", targetLevel,
					"returning", "objc.IObject")
				if debugTypeAnnotations {
					return fmt.Sprintf("objc.IObject /* type-resolution: cross-framework=%s.%s, objc-type=%s, reason=hierarchy-violation */", frameworkPkg, goType, goType)
				}
				return "objc.IObject /* cross-framework: " + goType + " */"
			}

			// No hierarchy violation - use the properly qualified type
			// If goType is already qualified (contains "."), use it as-is with "I" prefix
			// Otherwise, qualify it with the framework package
			var qualifiedType string
			if strings.Contains(goType, ".") {
				// Already qualified, just add the "I" prefix
				parts := strings.SplitN(goType, ".", 2)
				qualifiedType = parts[0] + ".I" + parts[1]
			} else {
				// Not qualified yet, add framework package
				qualifiedType = frameworkPkg + ".I" + goType
			}
			Debug.TypeMap("cross-framework type allowed by hierarchy", goType, qualifiedType,
				"goType", goType,
				"currentFramework", g.Framework,
				"currentLevel", currentLevel,
				"targetFramework", frameworkPkg,
				"targetLevel", targetLevel,
				"returning", qualifiedType)
			return qualifiedType
		}
		if baseType != goType {
			// Similar lookup for baseType: try direct, then try unqualified
			var baseFrameworkPkg string
			var baseFound bool

			// Try direct lookup first
			if baseFrameworkPkg, baseFound = crossFrameworkTypeRegistry[baseType]; !baseFound {
				// If baseType is already qualified, extract the unqualified name and try that
				if strings.Contains(baseType, ".") {
					parts := strings.SplitN(baseType, ".", 2)
					if len(parts) == 2 {
						baseFrameworkPkg, baseFound = crossFrameworkTypeRegistry[parts[1]]
					}
				}
			}

			if baseFound {
				// Check if this would violate the framework hierarchy
				currentLevel := getFrameworkLevel(strings.ToLower(g.Framework))
				targetLevel := getFrameworkLevel(baseFrameworkPkg)

				if currentLevel >= 0 && targetLevel > currentLevel {
					// Hierarchy violation - use objc.IObject
					Debug.TypeMap("cross-framework type violates hierarchy (stripped)", baseType, "objc.IObject",
						"baseType", baseType,
						"currentFramework", g.Framework,
						"currentLevel", currentLevel,
						"targetFramework", baseFrameworkPkg,
						"targetLevel", targetLevel,
						"returning", "objc.IObject")
					if debugTypeAnnotations {
						return fmt.Sprintf("objc.IObject /* type-resolution: cross-framework=%s.%s, objc-type=%s, stripped-from=%s, reason=hierarchy-violation */", baseFrameworkPkg, baseType, baseType, goType)
					}
					return "objc.IObject /* cross-framework: " + baseType + " */"
				}

				// No hierarchy violation - use the properly qualified type
				// If baseType is already qualified (contains "."), use it as-is with "I" prefix
				// Otherwise, qualify it with the framework package
				var qualifiedType string
				if strings.Contains(baseType, ".") {
					// Already qualified, just add the "I" prefix
					parts := strings.SplitN(baseType, ".", 2)
					qualifiedType = parts[0] + ".I" + parts[1]
				} else {
					// Not qualified yet, add framework package
					qualifiedType = baseFrameworkPkg + ".I" + baseType
				}
				Debug.TypeMap("cross-framework type allowed by hierarchy (stripped)", baseType, qualifiedType,
					"baseType", baseType,
					"currentFramework", g.Framework,
					"currentLevel", currentLevel,
					"targetFramework", baseFrameworkPkg,
					"targetLevel", targetLevel,
					"returning", qualifiedType)
				return qualifiedType
			}
		}

		Debug.TypeMap("NOT a class, returning unchanged", goType, "",
			"goType", goType,
			"returning", goType)
		if debugTypeAnnotations {
			return goType + " /* type-resolution: not-a-class, returned-as-is */"
		}
		return goType + " /* not a class type */"
	}

	// Convert to interface type: "Data" -> "IData"
	// Use the stripped base type for the interface name
	Debug.TypeMap("IS a class, converting to interface", goType, baseType,
		"goType", goType,
		"baseType", baseType,
		"returning", "I"+baseType)
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
			Debug.EnumDedup("merging duplicate enum", enum.Name, "",
				"enumName", enum.Name,
				"newCases", len(enum.Cases),
				"existingCases", len(existing.Cases))
			for _, newCase := range enum.Cases {
				// Check if this case already exists
				isDuplicate := false
				for _, existingCase := range existing.Cases {
					if existingCase.Name == newCase.Name {
						isDuplicate = true
						Debug.EnumDedup("skipping duplicate case", newCase.Name, enum.Name,
							"caseName", newCase.Name,
							"enumName", enum.Name)
						break
					}
				}
				if !isDuplicate {
					Debug.EnumDedup("adding new case", newCase.Name, enum.Name,
						"caseName", newCase.Name,
						"enumName", enum.Name)
					existing.Cases = append(existing.Cases, newCase)
				}
			}
		} else {
			// First time seeing this enum - but also deduplicate its cases
			Debug.EnumDedup("first occurrence of enum", enum.Name, "",
				"enumName", enum.Name,
				"caseCount", len(enum.Cases))

			// Deduplicate cases within this enum
			casesSeen := make(map[string]bool)
			deduplicatedCases := make([]*occ2go.ParsedEnumCase, 0, len(enum.Cases))
			for _, enumCase := range enum.Cases {
				if !casesSeen[enumCase.Name] {
					casesSeen[enumCase.Name] = true
					deduplicatedCases = append(deduplicatedCases, enumCase)
				} else {
					Debug.EnumDedup("removing duplicate case from first", enumCase.Name, enum.Name,
						"caseName", enumCase.Name,
						"enumName", enum.Name)
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

	// Deduplicate typedefs by STRIPPED name - keep only the first occurrence
	// This prevents collisions when both NSFoo and CFCFoo strip to "Foo"
	typedefsSeen := make(map[string]*occ2go.ParsedTypedef)
	deduplicatedTypedefs := make([]*occ2go.ParsedTypedef, 0, len(g.Typedefs))
	for _, typedef := range g.Typedefs {
		strippedName := stripObjCPrefix(typedef.Name)
		if existing, seen := typedefsSeen[strippedName]; !seen {
			typedefsSeen[strippedName] = typedef
			deduplicatedTypedefs = append(deduplicatedTypedefs, typedef)
		} else {
			// Log which typedef is being skipped due to name collision
			Debug.TypeMap("typedef name collision - skipping duplicate", typedef.Name, strippedName,
				"skippedName", typedef.Name,
				"skippedBaseType", typedef.BaseType,
				"keptName", existing.Name,
				"keptBaseType", existing.BaseType,
				"strippedName", strippedName,
				"framework", g.Framework)
		}
	}
	g.Typedefs = deduplicatedTypedefs

	// Add synthetic typedefs from config.yaml
	// These are types that should be defined but aren't documented by Apple
	if config != nil && config.SyntheticTypedefs != nil {
		if syntheticTypedefs, ok := config.SyntheticTypedefs[g.Framework]; ok {
			for _, st := range syntheticTypedefs {
				// If this has enum_values, create an enum instead of a typedef
				if len(st.EnumValues) > 0 {
					// Check if enum already exists
					enumExists := false
					for _, e := range g.Enums {
						if e.Name == st.Name {
							enumExists = true
							break
						}
					}
					if !enumExists {
						// Create enum with cases
						cases := make([]*occ2go.ParsedEnumCase, len(st.EnumValues))
						for i, ev := range st.EnumValues {
							cases[i] = &occ2go.ParsedEnumCase{
								Name:     ev.Name,
								Value:    fmt.Sprintf("%d", ev.Value),
								IntValue: ev.Value,
								Abstract: ev.Abstract,
							}
						}
						g.Enums = append(g.Enums, &occ2go.ParsedEnum{
							Name:     st.Name,
							BaseType: st.BaseType,
							Cases:    cases,
							Abstract: st.Abstract,
						})
					}
				} else {
					// NOTE: Parsed typedefs keep their original names with prefixes (e.g., CGDisplayReservationInterval)
					// The template strips prefixes when generating code, but the typedef.Name field is unchanged
					// So we must search using the ORIGINAL name from config (with prefix)

					// Check if typedef already exists and update it, or add new one
					typedefExists := false
					for i, td := range g.Typedefs {
						if td.Name == st.Name {
							// Replace existing typedef with synthetic one
							g.Typedefs[i] = &occ2go.ParsedTypedef{
								Name:     st.Name, // Keep original name with prefix
								BaseType: st.BaseType,
								Abstract: st.Abstract,
							}
							typedefExists = true
							break
						}
					}
					if !typedefExists {
						g.Typedefs = append(g.Typedefs, &occ2go.ParsedTypedef{
							Name:     st.Name, // Keep original name with prefix
							BaseType: st.BaseType,
							Abstract: st.Abstract,
						})
					}
				}
			}
		}
	}

	// Deduplicate structs by name and fields within each struct
	structsSeen := make(map[string]*occ2go.ParsedStruct)
	deduplicatedStructs := make([]*occ2go.ParsedStruct, 0, len(g.Structs))
	for _, structDef := range g.Structs {
		if existing, exists := structsSeen[structDef.Name]; exists {
			// Struct already exists, deduplicate fields
			fieldsSeen := make(map[string]bool)
			for _, field := range existing.Fields {
				fieldsSeen[field.Name] = true
			}
			// Add any new fields from this occurrence
			for _, field := range structDef.Fields {
				if !fieldsSeen[field.Name] {
					existing.Fields = append(existing.Fields, field)
					fieldsSeen[field.Name] = true
				}
			}
		} else {
			// First occurrence of this struct, deduplicate its fields
			fieldsSeen := make(map[string]bool)
			deduplicatedFields := make([]*occ2go.ParsedStructField, 0, len(structDef.Fields))
			for _, field := range structDef.Fields {
				if !fieldsSeen[field.Name] {
					fieldsSeen[field.Name] = true
					deduplicatedFields = append(deduplicatedFields, field)
				}
			}
			structDef.Fields = deduplicatedFields
			structsSeen[structDef.Name] = structDef
			deduplicatedStructs = append(deduplicatedStructs, structDef)
		}
	}
	g.Structs = deduplicatedStructs

	// Strip prefixes from struct names based on framework conventions
	// - CoreFoundation: Strip CF prefix (CFRange → Range) but keep CG prefix for geometry types
	// - Other frameworks: Keep original names
	for _, structDef := range g.Structs {
		if strings.EqualFold(g.Framework, "CoreFoundation") {
			// In CoreFoundation, strip CF prefix but preserve CG prefix for geometry types
			if strings.HasPrefix(structDef.Name, "CF") && !strings.HasPrefix(structDef.Name, "CG") {
				structDef.Name = structDef.Name[2:] // Strip "CF" prefix (CFRange → Range)
			}
		}
	}

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
		// Also index by stripped name for easier lookup
		stripped := stripObjCPrefix(typedef.Name)
		if stripped != typedef.Name {
			g.typedefIndex[stripped] = typedef
		}
		Debug.TimeInterval("building typedef index", typedef.Name, "",
			"typedefName", typedef.Name,
			"strippedName", stripped,
			"baseType", typedef.BaseType)
	}
	Debug.TimeInterval("built typedef index", "", "",
		"entryCount", len(g.typedefIndex),
		"hasTimeInterval", g.typedefIndex["TimeInterval"] != nil,
		"hasNSTimeInterval", g.typedefIndex["NSTimeInterval"] != nil)
	if g.typedefIndex["TimeInterval"] != nil {
		Debug.TimeInterval("TimeInterval found in index", "TimeInterval", "",
			"baseType", g.typedefIndex["TimeInterval"].BaseType)
	}
	if g.typedefIndex["NSTimeInterval"] != nil {
		Debug.TimeInterval("NSTimeInterval found in index", "NSTimeInterval", "",
			"baseType", g.typedefIndex["NSTimeInterval"].BaseType)
	}

	if g.Framework == "ObjectiveC" && Debug != nil {
		Debug.Log(DebugUndefined, "About to CollectUndefinedTypes", nil, "typedef_count", len(g.Typedefs))
		for i, td := range g.Typedefs {
			if i < 15 {
				Debug.Log(DebugUndefined, "Typedef entry", nil, "index", i, "name", td.Name)
			}
		}
	}

	// Cache undefined types for test generation
	g.undefinedTypes = g.CollectUndefinedTypes()
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

				// APPLEDOCS-450: Don't generate stubs for cross-framework parents
				// If the parent class resolves to another framework (has "."), it belongs to that framework.
				// We should NOT create a stub copy in this framework as that creates import cycles and redundancy.
				// The parent framework will define the class properly.
				//
				// Example: CKOperation inherits from NSOperation
				//   - resolveType("CloudKit", "Operation") returns "foundation.Operation"
				//   - Has ".", so it's from Foundation framework
				//   - DON'T create a stub in CloudKit - use Foundation's definition instead

				needsStub := false
				if !strings.Contains(resolvedType, ".") {
					// No framework qualifier means it's local/undefined - create a stub
					// This handles cases where parent classes are truly missing from documentation
					needsStub = true
				}
				// If contains ".", it's from another framework - don't create stub
				// If it's empty string (not found), it will also be treated as needsStub = false
				// and will use objectivec.Object as fallback during class generation

				if needsStub {
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

	// Populate package-level type registries before template execution
	// This ensures TypeToInterfaceType and mapCTypeToGoWithFramework can check
	// if a type is defined in the current framework
	currentFrameworkClasses = make(map[string]bool)
	for _, class := range g.Classes {
		currentFrameworkClasses[stripObjCPrefix(class.Name)] = true
		currentFrameworkClasses[class.Name] = true
	}

	currentFrameworkTypedefs = make(map[string]bool)
	for _, typedef := range g.Typedefs {
		currentFrameworkTypedefs[stripObjCPrefix(typedef.Name)] = true
		currentFrameworkTypedefs[typedef.Name] = true
	}

	currentFrameworkStructs = make(map[string]bool)
	for _, structDef := range g.Structs {
		currentFrameworkStructs[stripObjCPrefix(structDef.Name)] = true
		currentFrameworkStructs[structDef.Name] = true
	}
	// Add manual types from manualFrameworkTypes (e.g., geometry types from rect_types.go)
	if manualTypes, exists := manualFrameworkTypes[strings.ToLower(g.Framework)]; exists {
		for _, typeName := range manualTypes {
			currentFrameworkStructs[typeName] = true
		}
	}

	currentFrameworkEnums = make(map[string]bool)
	for _, enum := range g.Enums {
		currentFrameworkEnums[stripObjCPrefix(enum.Name)] = true
		currentFrameworkEnums[enum.Name] = true
	}

	// Create master template and parse all sub-templates
	// Chain Funcs calls: first core "dumb" functions, then GeneratorFuncs methods
	gf := GeneratorFuncs{Generator: g}
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
