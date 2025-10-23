package main

import (
	"fmt"
	"os"
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
	// ==== Foundation framework types ====
	// Foundation geometry types
	{ObjCType: "NSRect", GoType: "Rect", Framework: "Foundation"},
	{ObjCType: "NSSize", GoType: "Size", Framework: "Foundation"},
	{ObjCType: "NSPoint", GoType: "Point", Framework: "Foundation"},
	{ObjCType: "NSRange", GoType: "Range", Framework: "Foundation"},
	// Foundation enum/typedef types that conflict with other frameworks
	{ObjCType: "NSFormattingContext", GoType: "int", Framework: "Foundation"}, // Enum - use int to avoid coreimage.Context conflict
	{ObjCType: "FormattingContext", GoType: "int", Framework: "Foundation"},   // Stripped version
	{ObjCType: "Formatter.Context", GoType: "int", Framework: "Foundation"},   // Nested type (from Swift docs)
	{ObjCType: "Context", GoType: "int", Framework: "Foundation"},             // Bare Context in Foundation = FormattingContext
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
	// Opaque ref types - strip CG prefix to match generated typedefs
	{ObjCType: "CGColorRef", GoType: "ColorRef", Framework: "CoreGraphics"},
	{ObjCType: "CGColorSpaceRef", GoType: "ColorSpaceRef", Framework: "CoreGraphics"},
	{ObjCType: "CGContextRef", GoType: "ContextRef", Framework: "CoreGraphics"},
	{ObjCType: "CGImageRef", GoType: "ImageRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPathRef", GoType: "PathRef", Framework: "CoreGraphics"},
	{ObjCType: "CGLayerRef", GoType: "LayerRef", Framework: "CoreGraphics"},
	{ObjCType: "CGFontRef", GoType: "FontRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDataProviderRef", GoType: "DataProviderRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDataConsumerRef", GoType: "DataConsumerRef", Framework: "CoreGraphics"},
	{ObjCType: "CGFunctionRef", GoType: "FunctionRef", Framework: "CoreGraphics"},
	{ObjCType: "CGShadingRef", GoType: "ShadingRef", Framework: "CoreGraphics"},
	{ObjCType: "CGGradientRef", GoType: "GradientRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPatternRef", GoType: "PatternRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFDocumentRef", GoType: "PDFDocumentRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFPageRef", GoType: "PDFPageRef", Framework: "CoreGraphics"},
	{ObjCType: "CGColorConversionInfoRef", GoType: "ColorConversionInfoRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayConfigRef", GoType: "DisplayConfigRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayModeRef", GoType: "DisplayModeRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayStreamRef", GoType: "DisplayStreamRef", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayStreamUpdateRef", GoType: "DisplayStreamUpdateRef", Framework: "CoreGraphics"},
	{ObjCType: "CGEventSourceRef", GoType: "EventSourceRef", Framework: "CoreGraphics"},
	{ObjCType: "CGMutablePathRef", GoType: "MutablePathRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFArrayRef", GoType: "PDFArrayRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFContentStreamRef", GoType: "PDFContentStreamRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFDictionaryRef", GoType: "PDFDictionaryRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFObjectRef", GoType: "PDFObjectRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFOperatorTableRef", GoType: "PDFOperatorTableRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFScannerRef", GoType: "PDFScannerRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFStreamRef", GoType: "PDFStreamRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPDFStringRef", GoType: "PDFStringRef", Framework: "CoreGraphics"},
	{ObjCType: "CGPSConverterRef", GoType: "PSConverterRef", Framework: "CoreGraphics"},
	{ObjCType: "CGRenderingBufferProviderRef", GoType: "RenderingBufferProviderRef", Framework: "CoreGraphics"},
	{ObjCType: "CGEventRef", GoType: "EventRef", Framework: "CoreGraphics"},
	// Common C typedef types
	{ObjCType: "CGError", GoType: "Error", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayReservationInterval", GoType: "DisplayReservationInterval", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayFadeReservationToken", GoType: "DisplayFadeReservationToken", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayCount", GoType: "DisplayCount", Framework: "CoreGraphics"},
	{ObjCType: "CGDisplayErr", GoType: "DisplayErr", Framework: "CoreGraphics"},

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

	// ==== AppKit types ====
	// AppKit class types
	{ObjCType: "NSImage", GoType: "Image", Framework: "AppKit"},
	{ObjCType: "NSImage *", GoType: "Image", Framework: "AppKit"},
	{ObjCType: "NSImageSymbolConfiguration", GoType: "ImageSymbolConfiguration", Framework: "AppKit"},
	{ObjCType: "NSImageSymbolConfiguration *", GoType: "ImageSymbolConfiguration", Framework: "AppKit"},

	// AppKit enum types
	{ObjCType: "NSCellImagePosition", GoType: "CellImagePosition", Framework: "AppKit"},
	{ObjCType: "NSImageScaling", GoType: "ImageScaling", Framework: "AppKit"},
	{ObjCType: "NSWindowStyleMask", GoType: "WindowStyleMask", Framework: "AppKit"},
	{ObjCType: "WindowStyleMask", GoType: "WindowStyleMask", Framework: "AppKit"}, // Go name
	{ObjCType: "NSBackingStoreType", GoType: "BackingStoreType", Framework: "AppKit"},
	{ObjCType: "BackingStoreType", GoType: "BackingStoreType", Framework: "AppKit"}, // Go name
	{ObjCType: "NSWindowOrderingMode", GoType: "WindowOrderingMode", Framework: "AppKit"},
	{ObjCType: "NSWindowLevel", GoType: "WindowLevel", Framework: "AppKit"},

	// NSString is a universal type that maps to Go string across all frameworks
	// Framework: "" makes it framework-agnostic so it never gets qualified
	{ObjCType: "NSString *", GoType: "string", Framework: ""},

	// ==== UniformTypeIdentifiers types ====
	// Removed - let automatic type resolution handle UTType

	// Foundation date/time types - unqualified within Foundation
	{ObjCType: "NSTimeInterval", GoType: "TimeInterval", Framework: "Foundation"},

	// Foundation enum/options types - preserve NS prefix for enums defined in docs
	{ObjCType: "NSISO8601DateFormatOptions", GoType: "NSISO8601DateFormatOptions", Framework: "Foundation"},
	{ObjCType: "NSTimeInterval", GoType: "TimeInterval", Framework: "ObjectiveC"},
	{ObjCType: "NSObject", GoType: "objectivec.IObject", Framework: "*"}, // NSObject should always map to objectivec.IObject
	{ObjCType: "Object", GoType: "objectivec.IObject", Framework: "*"},   // Swift Object type (stripped NSObject) maps to objectivec.IObject
	// Foundation date/time types for AppKit - as float64 (no darwinkit imports)
	{ObjCType: "NSTimeInterval", GoType: "float64", Framework: "AppKit"},
	// Default for all other frameworks - use qualified foundation.TimeInterval
	{ObjCType: "NSTimeInterval", GoType: "foundation.TimeInterval", Framework: ""},

	// Foundation class types - unqualified within Foundation
	{ObjCType: "NSURL", GoType: "URL", Framework: "Foundation"},
	{ObjCType: "URL", GoType: "URL", Framework: "Foundation"},
	{ObjCType: "NSNumber", GoType: "Number", Framework: "Foundation"},
	{ObjCType: "Number", GoType: "Number", Framework: "Foundation"},
	// Foundation class types for AppKit - use qualified foundation types
	{ObjCType: "NSURL", GoType: "foundation.URL", Framework: "AppKit"},
	{ObjCType: "URL", GoType: "foundation.URL", Framework: "AppKit"},
	{ObjCType: "NSNumber", GoType: "foundation.Number", Framework: "AppKit"},
	{ObjCType: "Number", GoType: "foundation.Number", Framework: "AppKit"},
	// Default for all other frameworks - use qualified foundation types
	{ObjCType: "NSURL", GoType: "foundation.URL", Framework: ""},
	{ObjCType: "URL", GoType: "foundation.URL", Framework: ""},
	{ObjCType: "NSNumber", GoType: "foundation.Number", Framework: ""},
	{ObjCType: "Number", GoType: "foundation.Number", Framework: ""},

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

	// ==== CoreGraphics types in ImageCaptureCore ====
	{ObjCType: "CGImageRef", GoType: "coregraphics.CGImageRef", Framework: "ImageCaptureCore"},

	// ==== CoreGraphics types in ScreenCaptureKit ====
	{ObjCType: "CGImageRef", GoType: "coregraphics.CGImageRef", Framework: "ScreenCaptureKit"},

	// ==== CloudKit types in CoreData ====
	{ObjCType: "CKShare", GoType: "cloudkit.CKShare", Framework: "CoreData"},

	// ==== Foundation geometry types in ParavirtualizedGraphics ====
	{ObjCType: "NSSize", GoType: "foundation.Size", Framework: "ParavirtualizedGraphics"},

	// ==== Foundation geometry types in Virtualization ====
	{ObjCType: "NSSize", GoType: "foundation.Size", Framework: "Virtualization"},

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
	// Quartz uses CoreGraphics types
	{ObjCType: "CGColorSpaceRef", GoType: "coregraphics.CGColorSpaceRef", Framework: "Quartz"},
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
	{ObjCType: "UNNotificationAction", GoType: "objectivec.IObject", Framework: "Foundation"},

	// Foundation ↔ UniformTypeIdentifiers import cycle:
	// Foundation methods reference UTType, but UniformTypeIdentifiers imports Foundation (for URL, Number, etc.)
	// Break the cycle by mapping UTType to objectivec.IObject in Foundation
	{ObjCType: "UTType", GoType: "objectivec.IObject", Framework: "Foundation"},
	{ObjCType: "UTType *", GoType: "objectivec.IObject", Framework: "Foundation"},
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
			// If the type is from a different framework, qualify it
			if mapping.Framework != "" && mapping.Framework != framework {
				// Don't qualify types that are already qualified (unsafe.Pointer, objc.ID, etc.)
				if strings.Contains(mapping.GoType, ".") {
					return mapping.GoType, true
				}
				return strings.ToLower(mapping.Framework) + "." + mapping.GoType, true
			}
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
				// If the type is from a different framework, qualify it
				if mapping.Framework != "" && mapping.Framework != framework {
					// Don't qualify types that are already qualified (unsafe.Pointer, objc.ID, etc.)
					if strings.Contains(mapping.GoType, ".") {
						return mapping.GoType, true
					}
					return strings.ToLower(mapping.Framework) + "." + mapping.GoType, true
				}
				return mapping.GoType, true
			}
		}
	}

	// Check cross-framework type registry (auto-discovered types)
	// Try with both ObjC names (NSColor, NSImageScaling) and Go names (Color, ImageScaling)
	// First try the stripped name (preferred) to avoid returning NSCellAttribute when we want CellAttribute
	strippedType := stripObjCPrefix(objcType)
	if os.Getenv("DEBUG_TYPEMAP") == "1" && (strings.Contains(objcType, "Coder") || strings.Contains(objcType, "Error")) {
		fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: objcType=%s strippedType=%s (stripped=%v)\n",
			objcType, strippedType, strippedType != objcType)
	}

	// IMPORTANT: Check if the stripped type exists in the current framework BEFORE checking cross-framework registry
	// This prevents "Cursor" in CloudKit from resolving to appkit.Cursor instead of CKQueryCursor
	if strippedType != objcType {
		// Check if this stripped type is a class in the current framework
		if currentFrameworkClasses[strippedType] {
			if os.Getenv("DEBUG_IMPORTS") == "1" || os.Getenv("DEBUG_TYPEMAP") == "1" {
				fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: stripped type %s found in current framework %s classes, using unqualified name\n",
					strippedType, framework)
			}
			return strippedType, true
		}

		// Check if it's an enum in the current framework
		// Enums are generated with stripped names (e.g., ComparisonResult not NSComparisonResult)
		// so we must return the stripped name to match the generated enum type
		if os.Getenv("DEBUG_TYPEMAP") == "1" && strings.Contains(objcType, "Comparison") {
			fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: checking if %s is in currentFrameworkEnums (size=%d)\n",
				strippedType, len(currentFrameworkEnums))
		}
		if currentFrameworkEnums[strippedType] {
			if os.Getenv("DEBUG_IMPORTS") == "1" || os.Getenv("DEBUG_TYPEMAP") == "1" {
				fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: stripped type %s found in current framework %s enums, returning STRIPPED name %s\n",
					strippedType, framework, strippedType)
			}
			return strippedType, true // Return STRIPPED name to match generated enum types
		}

		// HEURISTIC: If currentFrameworkEnums is empty (not populated yet), use a heuristic:
		// Types with NS/CG/CA prefix that are NOT pointer types are likely enums
		// Classes are always used as pointers (*), but enums are value types
		// Only apply this if objcType does NOT contain " *" (not a pointer type)
		// IMPORTANT: Return STRIPPED name since enums are generated without prefixes
		if len(currentFrameworkEnums) == 0 && !strings.Contains(objcType, " *") {
			// Type has a prefix and is not a pointer type - likely an enum
			if os.Getenv("DEBUG_TYPEMAP") == "1" {
				fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: HEURISTIC - %s is non-pointer with prefix, likely enum, returning STRIPPED name %s\n",
					strippedType, strippedType)
			}
			return strippedType, true // Return STRIPPED name to match generated enum types
		}

		// Check if it's a typedef in the current framework
		if currentFrameworkTypedefs[strippedType] {
			if os.Getenv("DEBUG_IMPORTS") == "1" || os.Getenv("DEBUG_TYPEMAP") == "1" {
				fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: stripped type %s found in current framework %s typedefs, using unqualified name\n",
					strippedType, framework)
			}
			return strippedType, true
		}

		// Not in current framework, check cross-framework registry
		if frameworkPkg, found := crossFrameworkTypeRegistry[strippedType]; found {
			if os.Getenv("DEBUG_TYPEMAP") == "1" && strings.Contains(objcType, "Coder") {
				fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: stripped objcType=%s strippedType=%s frameworkPkg=%s currentFramework=%s\n",
					objcType, strippedType, frameworkPkg, framework)
			}
			// Don't qualify types with their own framework name
			if strings.ToLower(framework) == frameworkPkg {
				return strippedType, true
			}
			// NEVER qualify Go built-in primitives, even if they appear in cross-framework registry
			if isGoPrimitive(strippedType) {
				return strippedType, true
			}
			result := frameworkPkg + "." + strippedType
			if os.Getenv("DEBUG_TYPEMAP") == "1" && strings.Contains(objcType, "Coder") {
				fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: returning qualified type: %s\n", result)
			}
			return result, true
		}
		if os.Getenv("DEBUG_TYPEMAP") == "1" && (strings.Contains(objcType, "Coder") || strings.Contains(objcType, "Error")) {
			fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: stripped type %s NOT FOUND in registry\n", strippedType)
		}
	}

	// Then try the original name as fallback
	if frameworkPkg, found := crossFrameworkTypeRegistry[objcType]; found {
		if os.Getenv("DEBUG_TYPEMAP") == "1" && (strings.Contains(objcType, "Coder") || strings.Contains(objcType, "Error") || strings.Contains(objcType, "Operation")) {
			fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: original objcType=%s frameworkPkg=%s currentFramework=%s\n",
				objcType, frameworkPkg, framework)
		}
		// Don't qualify types with their own framework name
		if strings.ToLower(framework) == frameworkPkg {
			return objcType, true
		}
		// NEVER qualify Go built-in primitives, even if they appear in cross-framework registry
		if isGoPrimitive(objcType) {
			return objcType, true
		}
		result := frameworkPkg + "." + objcType
		if os.Getenv("DEBUG_TYPEMAP") == "1" && (strings.Contains(objcType, "Coder") || strings.Contains(objcType, "Error") || strings.Contains(objcType, "Operation")) {
			fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: returning qualified: %s\n", result)
		}
		return result, true
	}
	if os.Getenv("DEBUG_TYPEMAP") == "1" && (strings.Contains(objcType, "Coder") || strings.Contains(objcType, "Error") || strings.Contains(objcType, "Operation")) {
		fmt.Fprintf(os.Stderr, "DEBUG lookupTypeMapping: NOT FOUND in registry: objcType=%s\n", objcType)
	}

	// Try without pointer suffix in cross-framework registry
	if objcTypeNoPtr != objcType {
		if frameworkPkg, found := crossFrameworkTypeRegistry[objcTypeNoPtr]; found {
			// Don't qualify types with their own framework name
			if strings.ToLower(framework) == frameworkPkg {
				return objcTypeNoPtr, true
			}
			// NEVER qualify Go built-in primitives, even if they appear in cross-framework registry
			if isGoPrimitive(objcTypeNoPtr) {
				return objcTypeNoPtr, true
			}
			return frameworkPkg + "." + objcTypeNoPtr, true
		}
	}

	return "", false
}

// isGoPrimitive checks if a type name is a Go built-in primitive type.
// These types should NEVER be qualified with a package name.
func isGoPrimitive(typeName string) bool {
	goPrimitives := map[string]bool{
		"string":         true,
		"int":            true,
		"int8":           true,
		"int16":          true,
		"int32":          true,
		"int64":          true,
		"uint":           true,
		"uint8":          true,
		"uint16":         true,
		"uint32":         true,
		"uint64":         true,
		"float32":        true,
		"float64":        true,
		"bool":           true,
		"byte":           true,
		"rune":           true,
		"uintptr":        true,
		"unsafe.Pointer": true,
		"objc.ID":        true,
		"objc.Class":     true,
		"objc.SEL":       true,
	}
	return goPrimitives[typeName]
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
// UNUSED: Commented out as unreachable code
/*
func debugLogTypeMapping(objcType, framework string, result string) {
	// This would normally log to stderr or a debug file
	// Enable with environment variable DEBUG_TYPE_MAPPING=1
}
*/

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
