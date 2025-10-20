package occ2go

import (
	"strings"
)

// MapCTypeToGo maps C types to Go types for a given framework.
func MapCTypeToGo(cType, framework string) string {
	cType = strings.TrimSpace(cType)

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
	}

	// Common C types
	switch {
	case cType == "void":
		return ""
	case cType == "int":
		return "int"
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
		return "unsafe.Pointer"
	default:
		// Default to unsafe.Pointer for unknown types
		return "unsafe.Pointer"
	}
}
