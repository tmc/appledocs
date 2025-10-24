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

// extractBlockComponents extracts return type and parameter types from a block type string.
// This handles nested blocks by tracking parenthesis depth.
// Supports both named and unnamed blocks:
//   - Unnamed: "void (^)(void)"
//   - Named: "void (^blockName)(void)"
// Examples:
//   "void (^)(void)" -> ("void", "void", true)
//   "NSProgress * (^)(void (^)(NSData *, NSError *))" -> ("NSProgress *", "void (^)(NSData *, NSError *)", true)
//   "void (^completionHandler)(NSData *, NSError *)" -> ("void", "NSData *, NSError *", true)
func extractBlockComponents(blockType string) (returnType string, paramTypes string, ok bool) {
	// Find the (^ marker - could be (^) or (^name)
	caretIdx := strings.Index(blockType, "(^")
	if caretIdx == -1 {
		return "", "", false
	}

	// Return type is everything before (^
	returnType = strings.TrimSpace(blockType[:caretIdx])

	// Find the closing paren of (^) or (^name) - we need to skip past the block name if present
	// Start after "(^"
	nameEndIdx := caretIdx + 2
	for nameEndIdx < len(blockType) && blockType[nameEndIdx] != ')' {
		nameEndIdx++
	}
	if nameEndIdx >= len(blockType) {
		return "", "", false
	}
	// Now nameEndIdx points to the ')' after (^ or (^name

	// Find the opening paren after (^) or (^name)
	startIdx := nameEndIdx + 1 // Start after the closing paren
	for startIdx < len(blockType) && blockType[startIdx] == ' ' {
		startIdx++
	}
	if startIdx >= len(blockType) || blockType[startIdx] != '(' {
		return "", "", false
	}

	// Track parenthesis depth to find matching closing paren
	depth := 0
	endIdx := -1
	for i := startIdx; i < len(blockType); i++ {
		if blockType[i] == '(' {
			depth++
		} else if blockType[i] == ')' {
			depth--
			if depth == 0 {
				endIdx = i
				break
			}
		}
	}

	if endIdx == -1 {
		return "", "", false
	}

	// Parameter types are between the parentheses (excluding the parens themselves)
	paramTypes = blockType[startIdx+1 : endIdx]

	return returnType, paramTypes, true
}

// parseBlockType converts Objective-C block syntax to Go function type.
// Examples:
//   "void (^)(void)"          -> "func()"
//   "id (^)(NSString *)"      -> "func(objc.ID) objc.ID"
//   "BOOL (^)(id, NSError *)" -> "func(objc.ID, objc.ID) bool"
//   "NSProgress * (^)(void (^)(NSData *, NSError *))" -> "func(func(unsafe.Pointer, unsafe.Pointer)) unsafe.Pointer"
func parseBlockType(blockType, framework string) string {
	blockType = strings.TrimSpace(blockType)

	// Extract block components manually to handle nested blocks
	returnType, paramTypes, ok := extractBlockComponents(blockType)
	if !ok {
		// Not a block type
		return ""
	}

	returnType = strings.TrimSpace(returnType)
	paramTypes = strings.TrimSpace(paramTypes)

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
	case cType == "NSTimeInterval", cType == "TimeInterval":
		return "float64" // NSTimeInterval is typedef for double
	case cType == "size_t":
		return "uintptr"
	case cType == "int8_t":
		return "int8"
	case cType == "int16_t":
		return "int16"
	case cType == "int32_t":
		return "int32"
	case cType == "int64_t":
		return "int64"
	case cType == "uint8_t":
		return "uint8"
	case cType == "uint16_t":
		return "uint16"
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
	case cType == "Class":
		return "objc.Class" // Objective-C class type
	case cType == "SEL":
		return "objc.SEL" // Objective-C selector type
	case cType == "Method":
		return "Method" // Objective-C method type - defined as typedef
	case cType == "Ivar":
		return "Ivar" // Objective-C instance variable type - defined as typedef
	case cType == "Category":
		return "Category" // Objective-C category type - defined as typedef
	case cType == "objc_property_t":
		return "Objc_property_t" // Objective-C property type - defined as typedef
	case cType == "objc_exception_handler":
		return "Objc_exception_handler"
	case cType == "objc_exception_matcher":
		return "Objc_exception_matcher"
	case cType == "objc_exception_preprocessor":
		return "Objc_exception_preprocessor"
	case cType == "objc_func_loadImage":
		return "Objc_func_loadImage"
	case cType == "objc_hook_getClass":
		return "Objc_hook_getClass"
	case cType == "objc_hook_getImageName":
		return "Objc_hook_getImageName"
	case cType == "objc_hook_lazyClassNamer":
		return "Objc_hook_lazyClassNamer"
	case cType == "objc_uncaught_exception_handler":
		return "Objc_uncaught_exception_handler"
	case cType == "objc_objectptr_t":
		return "Objc_objectptr_t"
	case cType == "objc_zone_t":
		return "Objc_zone_t"
	case cType == "IMP":
		return "IMP" // Objective-C method implementation pointer - defined as typedef
	case cType == "id":
		return "objc.ID" // Objective-C object reference
	case strings.Contains(cType, "*"):
		// Handle pointer types
		// Remove const and * to get base type
		baseType := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(cType, "const", ""), "*", ""))

		// Special case: char * is commonly used for C strings
		// Map to *byte for better type safety (can convert to string when needed)
		if baseType == "char" || baseType == "CChar" {
			return "unsafe.Pointer" // Keep as unsafe.Pointer for compatibility with purego
		}

		// Handle pointer-to-primitive types as slices for array parameters
		// e.g., "const CGFloat *" -> []float64, "const float *" -> []float32
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
