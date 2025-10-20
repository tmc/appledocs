package main

import (
	"strings"
)

// TypeMapping represents a mapping from an Objective-C type to a Go type.
type TypeMapping struct {
	// ObjCType is the Objective-C type to match (e.g., "NSRect", "NSWindowStyleMask")
	ObjCType string
	// GoType is the corresponding Go type (e.g., "foundation.Rect", "WindowStyleMask")
	GoType string
	// Framework is the framework this type belongs to (e.g., "AppKit", "Foundation", "CoreGraphics")
	Framework string
	// RequiresImport is the import path if this type needs to be imported (e.g., "github.com/progrium/darwinkit/macos/foundation")
	RequiresImport string
}

// typeRegistry contains all known Objective-C to Go type mappings
var typeRegistry = []TypeMapping{
	// ==== Geometry types - Foundation framework ====
	// Foundation has its own geometry types that should be properly typed
	{ObjCType: "NSRect", GoType: "Rect", Framework: "Foundation"},
	{ObjCType: "NSSize", GoType: "Size", Framework: "Foundation"},
	{ObjCType: "NSPoint", GoType: "Point", Framework: "Foundation"},
	{ObjCType: "NSRange", GoType: "Range", Framework: "Foundation"},
	// CG geometry types in Foundation - cross-reference to CoreGraphics
	{ObjCType: "CGRect", GoType: "coregraphics.CGRect", Framework: "Foundation"},
	{ObjCType: "CGSize", GoType: "coregraphics.CGSize", Framework: "Foundation"},
	{ObjCType: "CGPoint", GoType: "coregraphics.CGPoint", Framework: "Foundation"},
	{ObjCType: "CGVector", GoType: "coregraphics.CGVector", Framework: "Foundation"},
	{ObjCType: "CGAffineTransform", GoType: "coregraphics.CGAffineTransform", Framework: "Foundation"},

	// ==== CoreGraphics types in CoreGraphics framework (unqualified) ====
	// When generating CoreGraphics itself, CG types should not be qualified
	// Struct types
	{ObjCType: "CGRect", GoType: "CGRect", Framework: "CoreGraphics"},
	{ObjCType: "CGSize", GoType: "CGSize", Framework: "CoreGraphics"},
	{ObjCType: "CGPoint", GoType: "CGPoint", Framework: "CoreGraphics"},
	{ObjCType: "CGAffineTransform", GoType: "CGAffineTransform", Framework: "CoreGraphics"},
	{ObjCType: "CGVector", GoType: "CGVector", Framework: "CoreGraphics"},
	{ObjCType: "CGFloat", GoType: "float64", Framework: "CoreGraphics"},
	// Opaque ref types
	{ObjCType: "CGColorRef", GoType: "CGColorRef", Framework: "CoreGraphics"},
	{ObjCType: "CGColorSpaceRef", GoType: "CGColorSpaceRef", Framework: "CoreGraphics"},
	{ObjCType: "CGContextRef", GoType: "CGContextRef", Framework: "CoreGraphics"},
	{ObjCType: "CGImageRef", GoType: "CGImageRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPathRef", GoType: "CGPathRef", Framework: "CoreGraphics"},
	{ObjCType: "CGLayerRef", GoType: "CGLayerRef", Framework: "CoreGraphics"},
	{ObjCType: "CGFontRef", GoType: "CGFontRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDataProviderRef", GoType: "CGDataProviderRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDataConsumerRef", GoType: "CGDataConsumerRef", Framework: "CoreGraphics"},
	{ObjCType: "CGFunctionRef", GoType: "CGFunctionRef", Framework: "CoreGraphics"},
	{ObjCType: "CGShadingRef", GoType: "CGShadingRef", Framework: "CoreGraphics"},
	{ObjCType: "CGGradientRef", GoType: "CGGradientRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPatternRef", GoType: "CGPatternRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFDocumentRef", GoType: "CGPDFDocumentRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFPageRef", GoType: "CGPDFPageRef", Framework: "CoreGraphics"},
	{ObjCType: "CGColorConversionInfoRef", GoType: "CGColorConversionInfoRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayConfigRef", GoType: "CGDisplayConfigRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayModeRef", GoType: "CGDisplayModeRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayStreamRef", GoType: "CGDisplayStreamRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayStreamUpdateRef", GoType: "CGDisplayStreamUpdateRef", Framework: "CoreGraphics"},
	{ObjCType: "CGEventSourceRef", GoType: "CGEventSourceRef", Framework: "CoreGraphics"},
	{ObjCType: "CGMutablePathRef", GoType: "CGMutablePathRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFArrayRef", GoType: "CGPDFArrayRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFContentStreamRef", GoType: "CGPDFContentStreamRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFDictionaryRef", GoType: "CGPDFDictionaryRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFObjectRef", GoType: "CGPDFObjectRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFOperatorTableRef", GoType: "CGPDFOperatorTableRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFScannerRef", GoType: "CGPDFScannerRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFStreamRef", GoType: "CGPDFStreamRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFStringRef", GoType: "CGPDFStringRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPSConverterRef", GoType: "CGPSConverterRef", Framework: "CoreGraphics"},
	{ObjCType: "CGRenderingBufferProviderRef", GoType: "CGRenderingBufferProviderRef", Framework: "CoreGraphics"},
	{ObjCType: "CGEventRef", GoType: "CGEventRef", Framework: "CoreGraphics"},

	// ==== Geometry types - AppKit framework ====
	// AppKit uses CoreGraphics types - import from CoreGraphics package
	{ObjCType: "NSRect", GoType: "coregraphics.CGRect", Framework: "AppKit"},
	{ObjCType: "CGRect", GoType: "coregraphics.CGRect", Framework: "AppKit"},
	{ObjCType: "NSSize", GoType: "coregraphics.CGSize", Framework: "AppKit"},
	{ObjCType: "CGSize", GoType: "coregraphics.CGSize", Framework: "AppKit"},
	{ObjCType: "NSPoint", GoType: "coregraphics.CGPoint", Framework: "AppKit"},
	{ObjCType: "CGPoint", GoType: "coregraphics.CGPoint", Framework: "AppKit"},
	{ObjCType: "NSRange", GoType: "foundation.Range", Framework: "AppKit"},
	{ObjCType: "CGAffineTransform", GoType: "coregraphics.CGAffineTransform", Framework: "AppKit"},
	// CoreGraphics opaque ref types in AppKit
	{ObjCType: "CGColorSpaceRef", GoType: "coregraphics.CGColorSpaceRef", Framework: "AppKit"},
	{ObjCType: "CGEventRef", GoType: "coregraphics.CGEventRef", Framework: "AppKit"},
	{ObjCType: "CGContextRef", GoType: "coregraphics.CGContextRef", Framework: "AppKit"},
	{ObjCType: "CGImageRef", GoType: "coregraphics.CGImageRef", Framework: "AppKit"},

	// ==== Geometry types - CoreImage framework ====
	// CoreImage uses CoreGraphics geometry types
	{ObjCType: "CGRect", GoType: "coregraphics.CGRect", Framework: "CoreImage"},
	{ObjCType: "CGSize", GoType: "coregraphics.CGSize", Framework: "CoreImage"},
	{ObjCType: "CGPoint", GoType: "coregraphics.CGPoint", Framework: "CoreImage"},
	{ObjCType: "CGAffineTransform", GoType: "coregraphics.CGAffineTransform", Framework: "CoreImage"},
	// CoreImage also uses NS-prefixed geometry types that map to CG types
	{ObjCType: "NSRect", GoType: "coregraphics.CGRect", Framework: "CoreImage"},
	{ObjCType: "NSSize", GoType: "coregraphics.CGSize", Framework: "CoreImage"},
	{ObjCType: "NSPoint", GoType: "coregraphics.CGPoint", Framework: "CoreImage"},
	{ObjCType: "Rect", GoType: "coregraphics.CGRect", Framework: "CoreImage"},
	{ObjCType: "Size", GoType: "coregraphics.CGSize", Framework: "CoreImage"},
	{ObjCType: "Point", GoType: "coregraphics.CGPoint", Framework: "CoreImage"},

	// ==== Geometry types - ScreenCaptureKit framework ====
	{ObjCType: "CGRect", GoType: "coregraphics.CGRect", Framework: "ScreenCaptureKit"},
	{ObjCType: "CGSize", GoType: "coregraphics.CGSize", Framework: "ScreenCaptureKit"},
	{ObjCType: "CGPoint", GoType: "coregraphics.CGPoint", Framework: "ScreenCaptureKit"},

	// ==== Geometry types - ScreenSaver framework ====
	// ScreenSaver uses CoreGraphics geometry types
	{ObjCType: "NSRect", GoType: "coregraphics.CGRect", Framework: "ScreenSaver"},
	{ObjCType: "CGRect", GoType: "coregraphics.CGRect", Framework: "ScreenSaver"},
	{ObjCType: "Rect", GoType: "coregraphics.CGRect", Framework: "ScreenSaver"},
	{ObjCType: "NSSize", GoType: "coregraphics.CGSize", Framework: "ScreenSaver"},
	{ObjCType: "CGSize", GoType: "coregraphics.CGSize", Framework: "ScreenSaver"},
	{ObjCType: "Size", GoType: "coregraphics.CGSize", Framework: "ScreenSaver"},
	{ObjCType: "NSPoint", GoType: "coregraphics.CGPoint", Framework: "ScreenSaver"},
	{ObjCType: "CGPoint", GoType: "coregraphics.CGPoint", Framework: "ScreenSaver"},
	{ObjCType: "Point", GoType: "coregraphics.CGPoint", Framework: "ScreenSaver"},

	// ==== AppKit types used in ScreenSaver framework ====
	// ScreenSaver uses AppKit's BackingStoreType enum
	{ObjCType: "NSBackingStoreType", GoType: "appkit.BackingStoreType", Framework: "ScreenSaver"},
	{ObjCType: "BackingStoreType", GoType: "appkit.BackingStoreType", Framework: "ScreenSaver"},

	// ==== AVFoundation types ====
	{ObjCType: "CGImageRef", GoType: "coregraphics.CGImageRef", Framework: "AVFoundation"},

	// ==== CoreML types ====
	{ObjCType: "CGImageRef", GoType: "coregraphics.CGImageRef", Framework: "CoreML"},
	{ObjCType: "Range", GoType: "foundation.Range", Framework: "CoreML"},
	{ObjCType: "NSRange", GoType: "foundation.Range", Framework: "CoreML"},

	// ==== MetalKit types ====
	{ObjCType: "CGImageRef", GoType: "coregraphics.CGImageRef", Framework: "MetalKit"},
	{ObjCType: "CGColorSpaceRef", GoType: "coregraphics.CGColorSpaceRef", Framework: "MetalKit"},

	// ==== CoreImage types ====
	{ObjCType: "CGImageRef", GoType: "coregraphics.CGImageRef", Framework: "CoreImage"},
	{ObjCType: "CGColorSpaceRef", GoType: "coregraphics.CGColorSpaceRef", Framework: "CoreImage"},
	{ObjCType: "CGLayerRef", GoType: "coregraphics.CGLayerRef", Framework: "CoreImage"},
	{ObjCType: "CGColorRef", GoType: "coregraphics.CGColorRef", Framework: "CoreImage"},
	{ObjCType: "CGContextRef", GoType: "coregraphics.CGContextRef", Framework: "CoreImage"},

	// AppKit window and view types (enums)
	{ObjCType: "NSWindowStyleMask", GoType: "WindowStyleMask", Framework: "AppKit"},
	{ObjCType: "NSBackingStoreType", GoType: "BackingStoreType", Framework: "AppKit"},
	{ObjCType: "NSWindowOrderingMode", GoType: "WindowOrderingMode", Framework: "AppKit"},
	{ObjCType: "NSWindowLevel", GoType: "WindowLevel", Framework: "AppKit"},

	// AppKit string types
	{ObjCType: "NSString *", GoType: "string", Framework: "AppKit"},

	// ==== UniformTypeIdentifiers types used in AppKit ====
	{ObjCType: "UTType", GoType: "uniformtypeidentifiers.UTType", Framework: "AppKit"},

	// Foundation date/time types - unqualified within Foundation
	{ObjCType: "NSTimeInterval", GoType: "TimeInterval", Framework: "Foundation"},
	// Foundation date/time types for AppKit - as float64 (no darwinkit imports)
	{ObjCType: "NSTimeInterval", GoType: "float64", Framework: "AppKit"},

	// ==== CoreGraphics types ====
	{ObjCType: "CGFloat", GoType: "float64", Framework: "CoreGraphics"},

	// ==== CoreText types ====
	// CoreText uses CoreGraphics types
	{ObjCType: "CGFloat", GoType: "float64", Framework: "CoreText"},
	{ObjCType: "CGRect", GoType: "coregraphics.CGRect", Framework: "CoreText"},
	{ObjCType: "CGSize", GoType: "coregraphics.CGSize", Framework: "CoreText"},
	{ObjCType: "CGPoint", GoType: "coregraphics.CGPoint", Framework: "CoreText"},
	{ObjCType: "CGAffineTransform", GoType: "coregraphics.CGAffineTransform", Framework: "CoreText"},

	// ==== CoreVideo types ====
	// CoreVideo uses CoreGraphics types
	{ObjCType: "CGFloat", GoType: "float64", Framework: "CoreVideo"},
	{ObjCType: "CGFloat", GoType: "float64", Framework: "QuartzCore"},
	{ObjCType: "CGRect", GoType: "coregraphics.CGRect", Framework: "CoreVideo"},
	{ObjCType: "CGSize", GoType: "coregraphics.CGSize", Framework: "CoreVideo"},
	{ObjCType: "CGPoint", GoType: "coregraphics.CGPoint", Framework: "CoreVideo"},
	{ObjCType: "CGAffineTransform", GoType: "coregraphics.CGAffineTransform", Framework: "CoreVideo"},
	{ObjCType: "CGAffineTransform", GoType: "coregraphics.CGAffineTransform", Framework: "QuartzCore"},

	// ==== CoreGraphics opaque ref types in QuartzCore ====
	// QuartzCore (Core Animation) frequently uses CG types
	{ObjCType: "CGColorRef", GoType: "coregraphics.CGColorRef", Framework: "QuartzCore"},
	{ObjCType: "CGColorSpaceRef", GoType: "coregraphics.CGColorSpaceRef", Framework: "QuartzCore"},
	{ObjCType: "CGContextRef", GoType: "coregraphics.CGContextRef", Framework: "QuartzCore"},
	{ObjCType: "CGImageRef", GoType: "coregraphics.CGImageRef", Framework: "QuartzCore"},
	{ObjCType: "CGPathRef", GoType: "coregraphics.CGPathRef", Framework: "QuartzCore"},
	{ObjCType: "CGMutablePathRef", GoType: "coregraphics.CGMutablePathRef", Framework: "QuartzCore"},
	{ObjCType: "CGLayerRef", GoType: "coregraphics.CGLayerRef", Framework: "QuartzCore"},
	{ObjCType: "CGFontRef", GoType: "coregraphics.CGFontRef", Framework: "QuartzCore"},

	// ==== Block/Closure types ====
	// Completion handlers and callbacks - map to proper function types
	// Generic completion handler: void (^)(NSError *)
	{ObjCType: "void (^)(NSError *)", GoType: "func(error objc.ID)", Framework: ""},
	{ObjCType: "void (^)(NSError * _Nullable)", GoType: "func(error objc.ID)", Framework: ""},
	// Generic completion handler with no parameters
	{ObjCType: "void (^)(void)", GoType: "func()", Framework: ""},
	// BOOL completion handler: void (^)(BOOL)
	{ObjCType: "void (^)(BOOL)", GoType: "func(success bool)", Framework: ""},

	// ==== Generic collection element types ====
	// NSArray element type - use objc.ID for element access
	{ObjCType: "id", GoType: "objc.ID", Framework: ""},
	{ObjCType: "id _Nullable", GoType: "objc.ID", Framework: ""},
	{ObjCType: "NSArray *", GoType: "objc.ID", Framework: ""},      // Will be wrapped with typed accessors
	{ObjCType: "NSDictionary *", GoType: "objc.ID", Framework: ""}, // Will be wrapped with typed accessors
	{ObjCType: "NSSet *", GoType: "objc.ID", Framework: ""},

	// Foundation edge enum - unqualified within Foundation
	{ObjCType: "NSRectEdge", GoType: "RectEdge", Framework: "Foundation"},
	// Foundation edge enum for AppKit - as int (no darwinkit imports)
	{ObjCType: "NSRectEdge", GoType: "int", Framework: "AppKit"},

	// ==== Foundation geometry types for other frameworks ====
	// FileProvider uses Foundation's Range type
	{ObjCType: "Range", GoType: "foundation.Range", Framework: "FileProvider"},
	{ObjCType: "NSRange", GoType: "foundation.Range", Framework: "FileProvider"},
	// MediaPlayer uses Foundation's Range type
	{ObjCType: "Range", GoType: "foundation.Range", Framework: "MediaPlayer"},
	{ObjCType: "NSRange", GoType: "foundation.Range", Framework: "MediaPlayer"},
	// ContactsUI uses Foundation geometry types
	{ObjCType: "Rect", GoType: "foundation.Rect", Framework: "ContactsUI"},
	{ObjCType: "NSRect", GoType: "foundation.Rect", Framework: "ContactsUI"},
	{ObjCType: "RectEdge", GoType: "foundation.RectEdge", Framework: "ContactsUI"},
	{ObjCType: "NSRectEdge", GoType: "foundation.RectEdge", Framework: "ContactsUI"},
	// AVKit uses Foundation's Rect type
	{ObjCType: "Rect", GoType: "foundation.Rect", Framework: "AVKit"},
	{ObjCType: "NSRect", GoType: "foundation.Rect", Framework: "AVKit"},
	// Quartz uses Foundation geometry types
	{ObjCType: "Rect", GoType: "foundation.Rect", Framework: "Quartz"},
	{ObjCType: "NSRect", GoType: "foundation.Rect", Framework: "Quartz"},
	{ObjCType: "Size", GoType: "foundation.Size", Framework: "Quartz"},
	{ObjCType: "NSSize", GoType: "foundation.Size", Framework: "Quartz"},
	// WebKit uses Foundation's Point type
	{ObjCType: "Point", GoType: "foundation.Point", Framework: "WebKit"},
	{ObjCType: "NSPoint", GoType: "foundation.Point", Framework: "WebKit"},

	// AppKit event types
	{ObjCType: "NSEventType", GoType: "EventType", Framework: "AppKit"},
	{ObjCType: "NSEventModifierFlags", GoType: "EventModifierFlags", Framework: "AppKit"},

	// ==== UserNotifications types ====
	// UserNotifications types used in Foundation - use objc.ID to avoid import cycles
	// Foundation imports UserNotifications, and UserNotifications imports Foundation,
	// creating a cycle if we use typed references
	{ObjCType: "UNNotificationAction", GoType: "objc.ID", Framework: "Foundation"},
}

// lookupTypeMapping finds a type mapping for the given Objective-C type.
// Returns the Go type and whether a mapping was found.
func lookupTypeMapping(objcType, framework string) (string, bool) {
	objcType = strings.TrimSpace(objcType)

	// Direct lookup - try framework-specific first
	for _, mapping := range typeRegistry {
		if mapping.ObjCType == objcType && mapping.Framework != "" && mapping.Framework == framework {
			return mapping.GoType, true
		}
	}

	// Then try framework-agnostic types
	for _, mapping := range typeRegistry {
		if mapping.ObjCType == objcType && mapping.Framework == "" {
			return mapping.GoType, true
		}
	}

	// Then try any matching type regardless of framework
	// (geometry types from Foundation are used in AppKit, etc.)
	for _, mapping := range typeRegistry {
		if mapping.ObjCType == objcType {
			return mapping.GoType, true
		}
	}

	// Try without pointer suffix for object types
	objcTypeNoPtr := strings.TrimSuffix(objcType, " *")
	if objcTypeNoPtr != objcType {
		// Direct lookup for no-pointer version
		for _, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr && mapping.Framework != "" && mapping.Framework == framework {
				return mapping.GoType, true
			}
		}

		// Then try framework-agnostic types
		for _, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr && mapping.Framework == "" {
				return mapping.GoType, true
			}
		}

		// Then try any matching type
		for _, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr {
				return mapping.GoType, true
			}
		}
	}

	return "", false
}

// getTypeImportPath returns the import path needed for a given Go type, if any.
func getTypeImportPath(objcType, framework string) string {
	objcType = strings.TrimSpace(objcType)

	for _, mapping := range typeRegistry {
		if mapping.ObjCType == objcType || (strings.TrimSuffix(objcType, " *") == mapping.ObjCType) {
			if mapping.Framework == "" || mapping.Framework == framework {
				return mapping.RequiresImport
			}
		}
	}

	return ""
}

// getAllAppKitEnumTypes returns all AppKit enum type names that we can generate.
// This helps in determining which types are available for generation.
func getAllAppKitEnumTypes() []string {
	var result []string
	seen := make(map[string]bool)

	for _, mapping := range typeRegistry {
		if mapping.Framework == "AppKit" && !strings.Contains(mapping.GoType, ".") {
			if !seen[mapping.GoType] {
				result = append(result, mapping.GoType)
				seen[mapping.GoType] = true
			}
		}
	}

	return result
}

// debugLogTypeMapping logs the actual type strings being looked up (for debugging)
// This helper is used during generation to understand what metadata types arrive
func debugLogTypeMapping(objcType, framework string, result string) {
	// This would normally log to stderr or a debug file
	// Enable with environment variable DEBUG_TYPE_MAPPING=1
}

// getAllMappedTypes returns all ObjC types in the registry for debugging
func getAllMappedTypes() []TypeMapping {
	return typeRegistry
}

// lookupTypeMappingDetails finds the full TypeMapping for a given Objective-C type.
func lookupTypeMappingDetails(objcType, framework string) *TypeMapping {
	objcType = strings.TrimSpace(objcType)

	// Direct lookup - try framework-specific first
	for i, mapping := range typeRegistry {
		if mapping.ObjCType == objcType && mapping.Framework != "" && mapping.Framework == framework {
			return &typeRegistry[i]
		}
	}

	// Then try framework-agnostic types
	for i, mapping := range typeRegistry {
		if mapping.ObjCType == objcType && mapping.Framework == "" {
			return &typeRegistry[i]
		}
	}

	// Then try any matching type
	for i, mapping := range typeRegistry {
		if mapping.ObjCType == objcType {
			return &typeRegistry[i]
		}
	}

	// Try without pointer suffix
	objcTypeNoPtr := strings.TrimSuffix(objcType, " *")
	if objcTypeNoPtr != objcType {
		for i, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr && mapping.Framework != "" && mapping.Framework == framework {
				return &typeRegistry[i]
			}
		}

		for i, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr && mapping.Framework == "" {
				return &typeRegistry[i]
			}
		}

		for i, mapping := range typeRegistry {
			if mapping.ObjCType == objcTypeNoPtr {
				return &typeRegistry[i]
			}
		}
	}

	return nil
}

// isFrameworkLocalType checks if an import path is for the same framework being generated.
// For example: if framework is "Foundation" and import is "github.com/progrium/darwinkit/macos/foundation",
// this returns true because both refer to Foundation.
func isFrameworkLocalType(importPath, framework string) bool {
	framework = strings.ToLower(framework)
	importPath = strings.ToLower(importPath)

	// Check if the framework name appears in the import path
	// This is a simplified check - in a real scenario we might want to be more precise
	return strings.Contains(importPath, framework)
}
