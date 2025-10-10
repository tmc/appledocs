package occ2go

import (
	"strings"
)

// MapCTypeToGo maps C types to Go types for a given framework.
func MapCTypeToGo(cType, framework string) string {
	cType = strings.TrimSpace(cType)

	// Framework-specific types
	if framework == "CoreGraphics" {
		switch {
		case strings.HasPrefix(cType, "CG") && strings.HasSuffix(cType, "Ref"):
			return cType // Already a Go type
		case cType == "CGFloat":
			return "CGFloat"
		case cType == "CGPoint":
			return "CGPoint"
		case cType == "CGSize":
			return "CGSize"
		case cType == "CGRect":
			return "CGRect"
		case cType == "CGAffineTransform":
			return "CGAffineTransform"
		}
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
