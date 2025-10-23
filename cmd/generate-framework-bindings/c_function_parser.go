package main

import (
	"regexp"
	"strings"
)

// CallbackTypeInfo holds information about a parsed callback type
type CallbackTypeInfo struct {
	GoType     string
	IsCallback bool
}

// typedefsNeedUnsafe checks if any typedef in the list needs the unsafe package
func typedefsNeedUnsafe(gen *Generator) bool {
	for _, typedef := range gen.Typedefs {
		info := parseCFunctionPointer(typedef.BaseType)
		if info.IsCallback && strings.Contains(info.GoType, "unsafe.Pointer") {
			return true
		}
	}
	return false
}

// parseCFunctionPointer parses a C function pointer type and converts it to a Go function type.
// Examples:
//   - "void *(*)(long, unsigned long, void *)" → "func(int, uint, unsafe.Pointer) unsafe.Pointer"
//   - "void (*)(void *, void *)" → "func(unsafe.Pointer, unsafe.Pointer)"
//   - "const struct __CFString *(*)(const void *)" → "func(unsafe.Pointer) StringRef"
func parseCFunctionPointer(cType string) CallbackTypeInfo {
	// Check if this is a function pointer: contains "(*)" pattern
	if !strings.Contains(cType, "(*)") {
		return CallbackTypeInfo{GoType: "", IsCallback: false}
	}

	// Pattern: returnType (*)(param1, param2, ...)
	// Example: void *(*)(long, unsigned long, void *)
	pattern := regexp.MustCompile(`^(.+?)\s*\(\*\)\s*\((.+?)\)$`)
	matches := pattern.FindStringSubmatch(cType)
	if matches == nil {
		// Try empty parameter list: returnType (*)(void) or returnType (*)()
		pattern = regexp.MustCompile(`^(.+?)\s*\(\*\)\s*\((void)?\)$`)
		matches = pattern.FindStringSubmatch(cType)
		if matches == nil {
			return CallbackTypeInfo{GoType: "", IsCallback: false}
		}
	}

	returnType := strings.TrimSpace(matches[1])
	paramString := ""
	if len(matches) > 2 {
		paramString = strings.TrimSpace(matches[2])
	}

	// Parse return type
	goReturnType := cTypeToGoType(returnType)

	// Parse parameters
	var goParams []string
	if paramString != "" && paramString != "void" {
		// Split by comma, handling nested parentheses
		params := splitCParameters(paramString)
		for _, param := range params {
			param = strings.TrimSpace(param)
			if param != "" && param != "void" {
				goParams = append(goParams, cTypeToGoType(param))
			}
		}
	}

	// Build Go function signature
	if len(goParams) == 0 {
		if goReturnType == "" {
			return CallbackTypeInfo{GoType: "func()", IsCallback: true}
		}
		return CallbackTypeInfo{GoType: "func() " + goReturnType, IsCallback: true}
	}

	paramsStr := strings.Join(goParams, ", ")
	if goReturnType == "" {
		return CallbackTypeInfo{GoType: "func(" + paramsStr + ")", IsCallback: true}
	}
	return CallbackTypeInfo{GoType: "func(" + paramsStr + ") " + goReturnType, IsCallback: true}
}

// splitCParameters splits C function parameters by comma, handling nested parentheses
func splitCParameters(params string) []string {
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
				result = append(result, current.String())
				current.Reset()
			} else {
				current.WriteRune(ch)
			}
		default:
			current.WriteRune(ch)
		}
	}

	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}

// cTypeToGoType converts a C type to a Go type for use in function signatures
func cTypeToGoType(cType string) string {
	cType = strings.TrimSpace(cType)

	// Remove const qualifier
	cType = strings.TrimPrefix(cType, "const ")
	cType = strings.TrimSpace(cType)

	// Remove enum keyword (e.g., "enum CFComparisonResult" → "CFComparisonResult")
	cType = strings.TrimPrefix(cType, "enum ")
	cType = strings.TrimSpace(cType)

	// Handle void (no return type)
	if cType == "void" {
		return ""
	}

	// Handle pointers
	isPointer := strings.HasSuffix(cType, "*")
	baseType := strings.TrimSuffix(cType, "*")
	baseType = strings.TrimSpace(baseType)

	// Map common C types to Go types
	switch baseType {
	case "void":
		return "unsafe.Pointer"
	case "char":
		if isPointer {
			return "string" // char* → string
		}
		return "byte"
	case "unsigned char":
		return "uint8"
	case "short":
		return "int16"
	case "unsigned short":
		return "uint16"
	case "int":
		return "int32"
	case "unsigned int", "unsigned":
		return "uint32"
	case "long":
		return "int"
	case "unsigned long":
		return "uint"
	case "long long":
		return "int64"
	case "unsigned long long":
		return "uint64"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "struct __CFString", "CFStringRef":
		return "StringRef"
	case "struct __CFAllocator", "CFAllocatorRef":
		return "AllocatorRef"
	case "struct __CFArray", "CFArrayRef":
		return "ArrayRef"
	case "struct __CFDictionary", "CFDictionaryRef":
		return "DictionaryRef"
	case "struct __CFData", "CFDataRef":
		return "DataRef"
	case "struct __CFURL", "CFURLRef":
		return "URLRef"
	case "struct __CFNumber", "CFNumberRef":
		return "NumberRef"
	case "struct __CFBoolean", "CFBooleanRef":
		return "BooleanRef"
	}

	// Handle other struct pointers
	if strings.HasPrefix(baseType, "struct __CF") && isPointer {
		// struct __CFSomething * → SomethingRef
		name := strings.TrimPrefix(baseType, "struct __CF")
		return name + "Ref"
	}

	// Strip CF prefix from types (CFComparisonResult → ComparisonResult)
	if strings.HasPrefix(baseType, "CF") && len(baseType) > 2 {
		nextChar := baseType[2]
		if nextChar >= 'A' && nextChar <= 'Z' {
			// Has CF prefix followed by uppercase - strip it
			baseType = baseType[2:]
		}
	}

	// If we have a pointer to an unknown type, use unsafe.Pointer
	if isPointer {
		return "unsafe.Pointer"
	}

	// Default: return the base type (possibly with CF prefix stripped)
	return baseType
}
