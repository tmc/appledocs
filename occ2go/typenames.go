package occ2go

import "strings"

// typenames.go contains utilities for parsing and normalizing Objective-C type names.
// These functions handle the syntactic aspects of ObjC types without framework-specific knowledge.

// StripTypeQualifiers removes Objective-C type qualifiers like __kindof, const, etc.
// Examples:
//
//	__kindof NSView * -> NSView *
//	const char * -> char *
func StripTypeQualifiers(typeName string) string {
	typeName = strings.TrimSpace(typeName)

	// Remove __kindof qualifier
	typeName = strings.TrimPrefix(typeName, "__kindof ")

	// Remove const qualifier
	typeName = strings.TrimPrefix(typeName, "const ")

	// Handle in array types
	if strings.HasPrefix(typeName, "[]__kindof ") {
		typeName = "[]" + strings.TrimPrefix(typeName, "[]__kindof ")
	}

	return typeName
}

// IsPointerType checks if a type is an Objective-C pointer type.
// Examples:
//
//	NSString * -> true
//	int -> false
//	id -> false (special case, not a pointer syntactically)
func IsPointerType(typeName string) bool {
	typeName = strings.TrimSpace(typeName)
	return strings.HasSuffix(typeName, "*")
}

// StripPointer removes the pointer suffix from a type name.
// Examples:
//
//	NSString * -> NSString
//	int -> int
func StripPointer(typeName string) string {
	typeName = strings.TrimSpace(typeName)
	return strings.TrimSpace(strings.TrimSuffix(typeName, "*"))
}

// IsBlockType checks if a type is an Objective-C block type.
// Blocks use the ^ syntax (e.g., void (^)(void))
func IsBlockType(typeName string) bool {
	return strings.Contains(typeName, "^")
}

// IsProtocolType checks if a type is an id<Protocol> pattern.
// Examples:
//
//	id<NSCopying> -> true
//	id -> false
//	NSString -> false
func IsProtocolType(typeName string) bool {
	typeName = strings.TrimSpace(typeName)
	return strings.HasPrefix(typeName, "id<") && strings.Contains(typeName, ">")
}

// IsArrayType checks if a type is an array type (already converted by parser).
// Examples:
//
//	[]int -> true
//	NSArray<NSString *> -> false (not yet converted)
//	int -> false
func IsArrayType(typeName string) bool {
	return strings.HasPrefix(strings.TrimSpace(typeName), "[]")
}

// GetArrayElementType extracts the element type from an array type.
// Examples:
//
//	[]int -> int
//	[]NSString -> NSString
//	int -> "" (not an array)
func GetArrayElementType(typeName string) string {
	typeName = strings.TrimSpace(typeName)
	if !strings.HasPrefix(typeName, "[]") {
		return ""
	}
	return strings.TrimPrefix(typeName, "[]")
}

// IsGenericType checks if a type uses Objective-C generics syntax.
// Examples:
//
//	NSArray<NSString *> -> true
//	NSDictionary<NSString *, id> -> true
//	NSString -> false
func IsGenericType(typeName string) bool {
	return strings.Contains(typeName, "<") && strings.Contains(typeName, ">")
}

// ExtractGenericElementType extracts the element type from NSArray<T> syntax.
// Returns empty string if not an NSArray generic or if parsing fails.
// Examples:
//
//	NSArray<NSString *> -> NSString
//	NSArray<NSString *> * -> NSString
//	NSDictionary<K, V> -> "" (not supported yet)
func ExtractGenericElementType(typeName string) string {
	typeName = strings.TrimSpace(typeName)

	// Only handle NSArray<T> pattern
	if !strings.HasPrefix(typeName, "NSArray<") {
		return ""
	}

	// Extract element type between < and >
	start := strings.Index(typeName, "<")
	end := strings.LastIndex(typeName, ">")

	if start < 0 || end <= start {
		return ""
	}

	elementType := strings.TrimSpace(typeName[start+1 : end])

	// Strip __kindof qualifier from element type
	elementType = strings.TrimPrefix(elementType, "__kindof ")

	// Remove trailing * from pointer types
	elementType = strings.TrimSpace(strings.TrimSuffix(elementType, "*"))

	// Strip protocol conformance syntax: NSView<Protocol> -> NSView
	if protocolStart := strings.Index(elementType, "<"); protocolStart > 0 {
		if strings.HasSuffix(elementType, ">") {
			elementType = strings.TrimSpace(elementType[:protocolStart])
		}
	}

	return elementType
}

// StripProtocolConformance removes protocol conformance syntax from a type.
// Objective-C uses Type<Protocol> to indicate protocol conformance.
// Examples:
//
//	NSView<NSCollectionViewElement> -> NSView
//	NSView<Protocol> * -> NSView *
//	id<NSCopying> -> id
//	NSString -> NSString (no change)
func StripProtocolConformance(typeName string) string {
	typeName = strings.TrimSpace(typeName)

	// Handle id<Protocol> specially - return as-is
	if strings.HasPrefix(typeName, "id<") {
		return typeName
	}

	// Find protocol syntax
	if protocolStart := strings.Index(typeName, "<"); protocolStart > 0 {
		// Find the closing >
		protocolEnd := strings.Index(typeName, ">")
		if protocolEnd > protocolStart {
			// Extract the base type and any suffix after >
			baseType := strings.TrimSpace(typeName[:protocolStart])
			suffix := strings.TrimSpace(typeName[protocolEnd+1:])
			if suffix != "" {
				return baseType + " " + suffix
			}
			return baseType
		}
	}

	return typeName
}

// StripPackageQualification removes package qualification from a type name.
// This handles Swift-style module.Type syntax.
// Examples:
//
//	foundation.NSString -> NSString
//	uniformtypeidentifiers.UTType -> UTType
//	NSString -> NSString (no change)
func StripPackageQualification(typeName string) string {
	typeName = strings.TrimSpace(typeName)

	// Check for package.Type pattern (lowercase package prefix)
	parts := strings.SplitN(typeName, ".", 2)
	if len(parts) == 2 {
		// Only strip if first part is all lowercase (package name)
		packageName := parts[0]
		if packageName == strings.ToLower(packageName) {
			return parts[1]
		}
	}

	return typeName
}

// NormalizeTypeName applies all common normalizations to a type name.
// This is a convenience function that combines multiple normalization steps.
// Examples:
//
//	__kindof NSView * -> NSView
//	foundation.NSString * -> NSString
//	const char * -> char
func NormalizeTypeName(typeName string) string {
	typeName = strings.TrimSpace(typeName)
	typeName = StripTypeQualifiers(typeName)
	typeName = StripPackageQualification(typeName)
	typeName = StripProtocolConformance(typeName)
	return typeName
}

// IsPrimitiveType checks if a type is a C/ObjC primitive type.
// Examples:
//
//	int -> true
//	BOOL -> true
//	NSInteger -> true
//	NSString -> false
func IsPrimitiveType(typeName string) bool {
	typeName = strings.TrimSpace(typeName)

	primitives := map[string]bool{
		"void":               true,
		"bool":               true,
		"BOOL":               true,
		"int":                true,
		"unsigned int":       true,
		"long":               true,
		"unsigned long":      true,
		"long long":          true,
		"unsigned long long": true,
		"short":              true,
		"unsigned short":     true,
		"char":               true,
		"unsigned char":      true,
		"float":              true,
		"double":             true,
		"int8_t":             true,
		"uint8_t":            true,
		"int16_t":            true,
		"uint16_t":           true,
		"int32_t":            true,
		"uint32_t":           true,
		"int64_t":            true,
		"uint64_t":           true,
		"size_t":             true,
		"ssize_t":            true,
		"NSInteger":          true,
		"NSUInteger":         true,
		"CGFloat":            true,
	}

	return primitives[typeName]
}

// IsSpecialType checks if a type is a special Objective-C runtime type.
// Examples:
//
//	id -> true
//	Class -> true
//	SEL -> true
//	NSString -> false
func IsSpecialType(typeName string) bool {
	typeName = strings.TrimSpace(typeName)

	specialTypes := map[string]bool{
		"id":    true,
		"Class": true,
		"SEL":   true,
	}

	return specialTypes[typeName]
}
