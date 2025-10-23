package occ2go

import (
	"regexp"
	"strings"
)

var (
	// blockTypeRegex matches Objective-C block syntax: ReturnType (^)(ParamTypes)
	// Examples:
	//   void (^)(void)           -> matches
	//   id (^)(NSString *)       -> matches
	//   BOOL (^)(id, NSError *)  -> matches
	blockTypeRegex = regexp.MustCompile(`^(.+?)\s*\(\^\)\s*\(([^)]*)\)$`)
)

// parseBlockType converts Objective-C block syntax to Go function type.
// Examples:
//   "void (^)(void)"          -> "func()"
//   "id (^)(NSString *)"      -> "func(objc.ID) objc.ID"
//   "BOOL (^)(id, NSError *)" -> "func(objc.ID, objc.ID) bool"
func parseBlockType(blockType, framework string) string {
	blockType = strings.TrimSpace(blockType)

	matches := blockTypeRegex.FindStringSubmatch(blockType)
	if matches == nil {
		// Not a block type
		return ""
	}

	returnType := strings.TrimSpace(matches[1])
	paramTypes := strings.TrimSpace(matches[2])

	// Map return type
	goReturnType := MapCTypeToGo(returnType, framework)

	// Map parameter types
	var goParams []string
	if paramTypes != "" && paramTypes != "void" {
		// Split by comma, accounting for pointer types
		params := splitParameters(paramTypes)
		for _, param := range params {
			param = strings.TrimSpace(param)
			if param != "" {
				goParam := MapCTypeToGo(param, framework)
				goParams = append(goParams, goParam)
			}
		}
	}

	// Build function signature
	result := "func("
	if len(goParams) > 0 {
		result += strings.Join(goParams, ", ")
	}
	result += ")"

	if goReturnType != "" {
		result += " " + goReturnType
	}

	return result
}

// splitParameters splits parameter list by comma, handling pointer types correctly.
// Examples:
//   "id, NSError *"     -> ["id", "NSError *"]
//   "NSString *, BOOL"  -> ["NSString *", "BOOL"]
func splitParameters(params string) []string {
	var result []string
	var current strings.Builder
	depth := 0

	for _, ch := range params {
		switch ch {
		case '(':
			depth++
			current.WriteRune(ch)
		case ')':
			depth--
			current.WriteRune(ch)
		case ',':
			if depth == 0 {
				result = append(result, strings.TrimSpace(current.String()))
				current.Reset()
			} else {
				current.WriteRune(ch)
			}
		default:
			current.WriteRune(ch)
		}
	}

	if current.Len() > 0 {
		result = append(result, strings.TrimSpace(current.String()))
	}

	return result
}

// parseGenericType extracts type name and generic parameter from types like "NSArray<T>"
func parseGenericType(typeName string) (baseType string, genericParam string, isGeneric bool) {
	// Check for generic syntax: Type<GenericParam>
	openBracket := strings.Index(typeName, "<")
	closeBracket := strings.LastIndex(typeName, ">")

	if openBracket != -1 && closeBracket != -1 && closeBracket > openBracket {
		baseType = strings.TrimSpace(typeName[:openBracket])
		genericParam = strings.TrimSpace(typeName[openBracket+1 : closeBracket])
		return baseType, genericParam, true
	}

	return typeName, "", false
}

// MapCTypeToGo maps C types to Go types for a given framework.
func MapCTypeToGo(cType, framework string) string {
	cType = strings.TrimSpace(cType)

	// Check if this is a generic type (e.g., NSArray<void (^)(void)>)
	if baseType, genericParam, isGeneric := parseGenericType(cType); isGeneric {
		// Recursively map the generic parameter
		mappedParam := MapCTypeToGo(genericParam, framework)

		// For NSArray with mapped generic param, return slice type
		if strings.HasPrefix(baseType, "NSArray") {
			return "[]" + mappedParam
		}

		// For other generic types, fall through to normal handling
		// (Could add more special cases here if needed)
	}

	// Check if this is a block type
	if blockFunc := parseBlockType(cType, framework); blockFunc != "" {
		return blockFunc
	}

	// CoreGraphics types (framework-agnostic - these types are used across frameworks)
	switch {
	// Special case: CGImageSourceRef and CGImageDestinationRef are from ImageIO framework, not CoreGraphics
	// Map to unsafe.Pointer until ImageIO framework is generated
	case cType == "CGImageSourceRef", cType == "CGImageDestinationRef":
		return "unsafe.Pointer"
	case strings.HasPrefix(cType, "CG") && strings.HasSuffix(cType, "Ref"):
		// CG*Ref types are opaque pointers - return as-is for qualification by resolveType
		return cType
	case cType == "CGFloat":
		return "CGFloat"
	case cType == "CGPoint":
		return "CGPoint"
	case cType == "CGSize":
		return "CGSize"
	case cType == "CGRect":
		return "CGRect"
	case cType == "CGVector":
		return "CGVector"
	case cType == "CGAffineTransform":
		return "CGAffineTransform"
	// NS geometry types that are typedef'd to CG types
	case cType == "NSPoint":
		return "CGPoint"
	case cType == "NSSize":
		return "CGSize"
	case cType == "NSRect":
		return "CGRect"
	// Bare geometry type names (sometimes used in Swift/modern APIs)
	case cType == "Point":
		return "CGPoint"
	case cType == "Size":
		return "CGSize"
	case cType == "Rect":
		return "CGRect"
	case cType == "Vector":
		return "CGVector"
	}

	// Common C types
	switch {
	case cType == "void":
		return ""
	case cType == "int":
		return "int"
	case cType == "NSInteger":
		return "int64"
	case cType == "NSUInteger":
		return "uint64"
	case cType == "size_t":
		return "uintptr"
	case cType == "uint32_t":
		return "uint32"
	case cType == "uint64_t":
		return "uint64"
	case cType == "float":
		return "float32"
	case cType == "double":
		return "float64"
	case cType == "bool", cType == "BOOL":
		return "bool"
	case strings.Contains(cType, "*"):
		// Handle pointer-to-primitive types as slices
		// e.g., "const CGFloat *" -> []float64, "const float *" -> []float32
		baseType := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(cType, "const", ""), "*", ""))
		switch baseType {
		case "CGFloat":
			return "[]float64"
		case "float":
			return "[]float32"
		case "double":
			return "[]float64"
		case "int":
			return "[]int"
		case "uint32_t":
			return "[]uint32"
		case "uint64_t":
			return "[]uint64"
		default:
			return "unsafe.Pointer"
		}
	default:
		// Default to unsafe.Pointer for unknown types
		return "unsafe.Pointer"
	}
}
