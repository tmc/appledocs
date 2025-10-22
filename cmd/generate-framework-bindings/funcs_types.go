package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// mapCTypeToGoWithFramework wraps occ2go.MapCTypeToGo and applies framework-specific type mappings.
// This ensures C types like CGAffineTransform are properly qualified with their framework package.
func mapCTypeToGoWithFramework(cType, framework string) string {
	// First apply occ2go's basic C type mapping
	goType := occ2go.MapCTypeToGo(cType, framework)

	// Then apply our framework-specific mapping to add package qualifiers
	// For example, CGAffineTransform -> coregraphics.CGAffineTransform
	mapped := mapObjCTypeToGo(goType, framework)

	return mapped
}

// mapObjCTypeToGo maps Objective-C types to Go types for darwinkit style.
// Examples:
//
//	NSString * -> string
//	id -> objc.Object
//	NSButton * -> Button (interface type in parameters)
//	NSRect -> foundation.Rect
//	NSWindowStyleMask -> WindowStyleMask
func mapObjCTypeToGo(objcType, framework string) string {
	objcType = strings.TrimSpace(objcType)

	if os.Getenv("DEBUG_TYPEMAP") == "1" {
		if strings.Contains(objcType, "Hotspot") || (framework == "Foundation" && strings.Contains(objcType, "NE")) {
			fmt.Fprintf(os.Stderr, "DEBUG mapObjCTypeToGo ENTRY: objcType=%q framework=%s\n", objcType, framework)
		}
	}

	// Strip self-package qualifications from Swift documentation
	// Swift docs often use module.Type format (e.g., uniformtypeidentifiers.UTType)
	// When generating the same framework, we should use unqualified names
	if framework != "" {
		// Build package prefix (e.g., "uniformtypeidentifiers." from "UniformTypeIdentifiers")
		packagePrefix := strings.ToLower(framework) + "."
		// Check for exact package.Type pattern
		if strings.HasPrefix(strings.ToLower(objcType), packagePrefix) {
			// Extract the type name after the dot
			// E.g., "uniformtypeidentifiers.UTType" -> "UTType"
			parts := strings.SplitN(objcType, ".", 2)
			if len(parts) == 2 {
				objcType = parts[1]
			}
		}
	}

	// Strip type qualifiers (__kindof, const, etc.) using occ2go utility
	objcType = occ2go.StripTypeQualifiers(objcType)

	// Handle id<Protocol> pattern (e.g., "id<NSFetchRequestResult>" -> "objc.ID")
	// This is Objective-C's protocol conformance syntax
	if occ2go.IsProtocolType(objcType) {
		return "objc.ID"
	}

	// Handle []id<Protocol> pattern (e.g., "[]id<NSFetchRequestResult>" -> "[]objc.ID")
	if occ2go.IsArrayType(objcType) {
		elementType := occ2go.GetArrayElementType(objcType)
		if occ2go.IsProtocolType(elementType) {
			return "[]objc.ID"
		}
	}

	// Handle array types that are already converted by occ2go (e.g., "[]void (^)(void)" -> "[]unsafe.Pointer")
	// This handles cases where occ2go has already converted NSArray<T> to []T
	// We need to recursively map the element type
	if occ2go.IsArrayType(objcType) {
		elementType := occ2go.GetArrayElementType(objcType)
		goElementType := mapObjCTypeToGo(elementType, framework)
		return "[]" + goElementType
	}

	// Handle Objective-C generic types (e.g., NSArray<NSString *>, NSArray<SCDisplay *>)
	if occ2go.IsGenericType(objcType) {
		// Extract NSArray element type: NSArray<ElementType *> -> []ElementType
		// Handle both "NSArray<T>" and "NSArray<T> *" patterns
		elementType := occ2go.ExtractGenericElementType(objcType)
		if elementType != "" {
			// Successfully extracted NSArray element type

			// Special case: NSString -> string
			if elementType == "NSString" {
				return "[]string"
			}

			// Strip common Apple prefixes from element types
			// This ensures NSArray<SCDisplay *> -> []Display, NSArray<NSButton *> -> []Button
			strippedType := stripObjCPrefix(elementType)
			if strippedType != elementType {
				// Prefix was stripped - check if this type exists in current framework
				if currentFrameworkClasses[strippedType] {
					// Type is defined in current framework, safe to use
					return "[]" + strippedType
				}
				// Cross-framework reference - check type mapping registry for explicit mapping
				// Foundation defines URL and Number, not NSURL and NSNumber
				// Try both the stripped type name and the original element type
				if goType, found := lookupTypeMapping(strippedType, framework); found {
					return "[]" + goType
				}
				if goType, found := lookupTypeMapping(elementType, framework); found {
					return "[]" + goType
				}
				// Last resort - fall back to unsafe.Pointer for array elements
				return "[]unsafe.Pointer"
			}

			// For other types, try to map them
			goElementType := mapObjCTypeToGo(elementType, framework)
			if goElementType == "unsafe.Pointer" {
				// If mapping failed, use the element type directly
				return "[]" + elementType
			}
			return "[]" + goElementType
		}
		// For other generic types (NSDictionary, etc.), fall back to unsafe.Pointer
		return "unsafe.Pointer"
	}

	// Special built-in types (before checking pointers)
	switch objcType {
	case "id":
		return "objc.ID"
	case "Class":
		return "objc.Class"
	case "SEL":
		return "objc.SEL"
	case "BOOL":
		return "bool"
	case "NSInteger", "Int":
		return "int"
	case "NSUInteger", "UInt":
		return "uint"
	case "unsigned long long", "UInt64", "uint64_t":
		return "uint64"
	case "CGFloat", "Double":
		return "float64"
	case "Float":
		return "float32"
	case "void":
		return ""
	case "String", "String?":
		// Swift string types map to Go string
		return "string"
	}

	// Check the type mapping registry first (includes both with and without pointers)
	// This must come before the block check so that mapped block types (e.g., void (^)(void) -> func())
	// are handled correctly
	if goType, found := lookupTypeMapping(objcType, framework); found {
		if os.Getenv("DEBUG_TYPEMAP") == "1" {
			if strings.Contains(objcType, "NSObject") || strings.Contains(objcType, "CellAttribute") || strings.Contains(objcType, "Coder") {
				fmt.Fprintf(os.Stderr, "DEBUG mapObjCTypeToGo: found in registry objcType=%s goType=%s framework=%s\n", objcType, goType, framework)
			}
		}
		return goType
	}

	// Debug: check if NSObject wasn't found
	if os.Getenv("DEBUG_TYPEMAP") == "1" && strings.Contains(objcType, "NSObject") {
		fmt.Fprintf(os.Stderr, "DEBUG mapObjCTypeToGo: NSObject NOT found in registry, objcType=%s framework=%s\n", objcType, framework)
	}

	// Handle Objective-C blocks (e.g., void (^)(NSModalResponse))
	// Blocks are closures that cannot be easily represented in Go, so map to unsafe.Pointer
	// This comes after type registry check so explicitly mapped blocks can use proper Go types
	if occ2go.IsBlockType(objcType) {
		return "unsafe.Pointer"
	}

	// Handle pointers for types not in the registry
	isPointer := occ2go.IsPointerType(objcType)
	objcTypeNoPtr := occ2go.StripPointer(objcType)

	// Special case: NSString * -> string (most common string parameter type)
	if isPointer && objcTypeNoPtr == "NSString" {
		return "string"
	}

	// Also handle NSString without pointer (from property types in docs)
	if objcType == "NSString" {
		return "string"
	}

	// Check registry again for type without pointer
	if isPointer && objcTypeNoPtr != objcType {
		if goType, found := lookupTypeMapping(objcTypeNoPtr, framework); found {
			return goType
		}
	}

	//  For pointer types to ObjC classes, try stripping prefix BEFORE falling back to MapCTypeToGo
	// This allows cross-framework type resolution to work properly
	goType := ""
	if isPointer && objcTypeNoPtr != "" {
		strippedType := stripObjCPrefix(objcTypeNoPtr)
		// Debug ALL pointer types when framework is Foundation
		if os.Getenv("DEBUG_TYPEMAP") == "1" && framework == "Foundation" {
			fmt.Fprintf(os.Stderr, "DEBUG mapObjCTypeToGo pointer: objcType=%s objcTypeNoPtr=%s strippedType=%s\n",
				objcType, objcTypeNoPtr, strippedType)
		}
		if strippedType != objcTypeNoPtr {
			// Successfully stripped a prefix - this is likely an ObjC class type
			// Use the stripped type and let resolveType find the right framework
			goType = strippedType
			if os.Getenv("DEBUG_TYPEMAP") == "1" && framework == "Foundation" {
				fmt.Fprintf(os.Stderr, "DEBUG mapObjCTypeToGo: set goType=%s from %s\n", goType, objcType)
			}
		}
	}

	// Fall back to occ2go mapping if we haven't resolved it yet
	if goType == "" {
		goType = occ2go.MapCTypeToGo(objcType, framework)
		if os.Getenv("DEBUG_TYPEMAP") == "1" && (strings.Contains(objcType, "Hotspot") || strings.Contains(objcType, "CellAttribute")) {
			fmt.Fprintf(os.Stderr, "DEBUG mapObjCTypeToGo: after occ2go.MapCTypeToGo goType=%q objcType=%s\n", goType, objcType)
		}
	}

	// Never return empty string for a type - default to unsafe.Pointer
	if goType == "" {
		return "unsafe.Pointer"
	}

	// Resolve cross-framework types (e.g., CGAffineTransform -> coregraphics.CGAffineTransform)
	resolvedType := resolveType(framework, goType)
	if os.Getenv("DEBUG_TYPEMAP") == "1" && (strings.Contains(objcType, "AttributedString") || strings.Contains(objcType, "NSApplication") || objcType == "NSString *" || goType == "string") {
		fmt.Fprintf(os.Stderr, "DEBUG mapObjCTypeToGo: objcType=%s framework=%s before resolve goType=%s, after resolve=%s\n", objcType, framework, goType, resolvedType)
	}
	goType = resolvedType

	// Check for framework hierarchy violations - if resolved type references a higher-level framework,
	// map to generic objectivec.IObject instead to avoid import cycles
	if strings.Contains(goType, ".") && framework != "" {
		parts := strings.Split(goType, ".")
		if len(parts) >= 2 {
			targetFramework := parts[0]
			currentLevel := getFrameworkLevel(strings.ToLower(framework))
			targetLevel := getFrameworkLevel(targetFramework)

			if currentLevel >= 0 && targetLevel > currentLevel {
				// Hierarchy violation - map to objectivec.IObject
				if os.Getenv("DEBUG_TYPEMAP") == "1" {
					fmt.Fprintf(os.Stderr, "DEBUG mapObjCTypeToGo: hierarchy violation %s (level %d) -> %s (level %d), mapping to objectivec.IObject\n",
						framework, currentLevel, targetFramework, targetLevel)
				}
				return "objectivec.IObject"
			}
		}
	}

	if os.Getenv("DEBUG_TYPEMAP") == "1" && (strings.Contains(objcType, "NSApplication") || goType == "appkit.string") {
		fmt.Fprintf(os.Stderr, "DEBUG mapObjCTypeToGo EXIT: objcType=%s framework=%s returning=%s\n", objcType, framework, goType)
	}

	return goType
}

// getFrameworkLevel returns the hierarchy level for a framework (0-4), or -1 if unknown
func getFrameworkLevel(framework string) int {
	// Access the frameworkLevels map from framework_hierarchy.go
	// We need to import this or duplicate the levels here
	// For now, duplicate the essential levels
	levels := map[string]int{
		"objc":             0,
		"objectivec":       0,
		"coregraphics":     1,
		"corefoundation":   1,
		"foundation":       1,
		"coretext":         1,
		"iosurface":        1,
		"coreimage":        2,
		"quartzcore":       2,
		"coreaudio":        2,
		"coremidi":         2,
		"imageio":          2,
		"coredata":         2,
		"corelocation":     2,
		"corespotlight":    2,
		"network":          2,
		"security":         2,
		"corebluetooth":    2,
		"corevideo":        2,
		"coreml":           2,
		"vision":           2,
		"naturallanguage":  2,
		"appkit":           3,
		"uikit":            3,
		"webkit":           3,
		"pdfkit":           3,
		"networkextension": 3,
		"avfoundation":     4,
		"avfaudio":         4,
		"avkit":            4,
		"avrouting":        4,
		"audiotoolbox":     4,
		"cloudkit":         4,
		"contacts":         4,
		"contactsui":       4,
		"gameplaykit":      4,
		"intents":          4,
		"intentsui":        4,
		"metal":            4,
		"metalkit":         4,
		"eventkit":         4,
		"healthkit":        4,
		"homekit":          4,
		"mapkit":           4,
		"messages":         4,
		"storekit":         4,
		"usernotifications": 4,
	}
	if level, ok := levels[strings.ToLower(framework)]; ok {
		return level
	}
	return -1
}

// resolveType resolves a type name to its fully qualified name, handling cross-framework dependencies.
// Takes the current framework context and a type name (e.g., "MutableAttributedString") and returns
// either the unqualified name (if it's in the same framework) or a qualified name (e.g., "foundation.MutableAttributedString").
// This helper is used in templates to properly reference types that may come from other frameworks.
//
// Examples:
//
//	resolveType("AppKit", "Button") -> "Button" (same framework)
//	resolveType("AppKit", "MutableAttributedString") -> "foundation.MutableAttributedString" (cross-framework)
//	resolveType("Foundation", "Array") -> "Array" (same framework)
func resolveType(framework, typeName string) string {
	if os.Getenv("DEBUG_TYPEMAP") == "1" && (typeName == "string" || strings.Contains(typeName, "string")) {
		fmt.Fprintf(os.Stderr, "DEBUG resolveType ENTRY: framework=%s typeName=%q\n", framework, typeName)
	}

	if typeName == "" {
		return ""
	}

	// Never qualify Go primitives - they should always be unqualified
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
	}
	if goPrimitives[typeName] {
		if os.Getenv("DEBUG_TYPEMAP") == "1" && typeName == "string" {
			fmt.Fprintf(os.Stderr, "DEBUG resolveType: returning primitive 'string' unqualified for framework %s\n", framework)
		}
		return typeName
	}

	// ALSO check for capital-S String which should map to lowercase string
	// This happens when NSString typedef resolves to "String" instead of "string"
	if typeName == "String" {
		if os.Getenv("DEBUG_TYPEMAP") == "1" {
			fmt.Fprintf(os.Stderr, "DEBUG resolveType: converting 'String' to 'string' for framework %s\n", framework)
		}
		return "string"
	}

	// Strip self-package qualifications (e.g., foundation.NSString in Foundation -> NSString)
	// This prevents incorrect qualification like foundation.NSOrderedCollectionChange in the foundation package
	packagePrefix := strings.ToLower(framework) + "."
	if strings.HasPrefix(typeName, packagePrefix) {
		return strings.TrimPrefix(typeName, packagePrefix)
	}

	// Common Foundation base classes that other frameworks inherit from
	foundationTypes := map[string]bool{
		"MutableAttributedString": true,
		"AttributedString":        true,
		"Array":                   true,
		"MutableArray":            true,
		"Dictionary":              true,
		"MutableDictionary":       true,
		"Set":                     true,
		"MutableSet":              true,
		"String":                  true,
		"MutableString":           true,
		"Data":                    true,
		"MutableData":             true,
		"Date":                    true,
		"URL":                     true,
		"NSURL":                   true, // Include NS-prefixed version
		"URLRequest":              true,
		"MutableURLRequest":       true,
		"Value":                   true,
		"Number":                  true,
		"NSNumber":                true, // Include NS-prefixed version
		"URLSession":              true,
		"URLSessionTask":          true,
		"URLSessionDataTask":      true,
		"URLSessionUploadTask":    true,
		"URLSessionDownloadTask":  true,
		"URLSessionStreamTask":    true,
		"Enumerator":              true, // For OSLog.OSLogEnumerator
		"Operation":               true,
		"OperationQueue":          true,
		"Expression":              true, // NSExpression - used by CoreData
		"ExtensionContext":        true, // NSExtensionContext - used by AuthenticationServices
		"Coder":                   true, // NSCoder - base class for archiving
		"KeyedArchiver":           true, // NSKeyedArchiver
		"KeyedUnarchiver":         true, // NSKeyedUnarchiver - used by MetalPerformanceShaders
		"TimeInterval":            true, // NSTimeInterval - type alias for float64 used by many frameworks
		"Error":                   true, // NSError - error handling across all frameworks
		"NSError":                 true, // Include NS-prefixed version
	}

	// QuartzCore types used by other frameworks
	quartzCoreTypes := map[string]bool{
		"Layer":          true, // CALayer
		"Animation":      true, // CAAnimation
		"MediaTiming":    true, // CAMediaTiming protocol
		"Transaction":    true, // CATransaction
		"TransformLayer": true, // CATransformLayer
		"OpenGLLayer":    true, // CAOpenGLLayer - used by NSOpenGLLayer in AppKit
	}

	// AppKit types used by other frameworks (common base classes)
	appKitTypes := map[string]bool{
		"Responder":            true, // NSResponder
		"View":                 true, // NSView
		"Control":              true, // NSControl
		"Window":               true, // NSWindow
		"ViewController":       true, // NSViewController
		"NavigationController": true, // NSNavigationController (though less common on macOS)
		"Panel":                true, // NSPanel
		"Application":          true, // NSApplication
		"Document":             true, // NSDocument
		"WindowController":     true, // NSWindowController
		"Menu":                 true, // NSMenu
		"MenuItem":             true, // NSMenuItem
	}

	// CoreGraphics types used by other frameworks
	coreGraphicsTypes := map[string]bool{
		// Struct types
		"CGAffineTransform": true,
		"CGPoint":           true,
		"CGSize":            true,
		"CGRect":            true,
		"CGVector":          true,
		"CGFloat":           true,
		// Opaque ref types
		"CGColorRef":            true,
		"CGColorSpaceRef":       true,
		"CGContextRef":          true,
		"CGImageRef":            true,
		"CGImageSourceRef":      true,
		"CGImageDestinationRef": true,
		"CGPathRef":             true,
		"CGLayerRef":            true,
		"CGFontRef":             true,
		"CGDataProviderRef":     true,
		"CGDataConsumerRef":     true,
		"CGFunctionRef":         true,
		"CGShadingRef":          true,
		"CGGradientRef":         true,
		"CGPatternRef":          true,
		"CGPDFDocumentRef":      true,
		"CGPDFPageRef":          true,
	}

	// Check if the type exists in current framework FIRST before adding qualifications
	// This prevents self-imports (e.g., coregraphics.CGAffineTransform in CoreGraphics)
	// Check classes, enums, and typedefs - all stored with stripped ObjC prefixes
	strippedTypeName := stripObjCPrefix(typeName)
	if currentFrameworkClasses[strippedTypeName] || currentFrameworkEnums[strippedTypeName] || currentFrameworkTypedefs[strippedTypeName] {
		// DEBUG: Uncomment to debug same-framework type resolution
		// fmt.Fprintf(os.Stderr, "DEBUG resolveType: Found '%s' (stripped: '%s') in current framework '%s', returning as-is\n", typeName, strippedTypeName, framework)
		// It's in the current framework, return as-is
		return typeName
	}

	// If we're in CoreGraphics framework, all types are local
	if framework == "CoreGraphics" && coreGraphicsTypes[typeName] {
		return typeName
	}

	// If this is a known CoreGraphics type and we're not in CoreGraphics, qualify it
	if coreGraphicsTypes[typeName] {
		// Make sure the type has the CG prefix for proper type reference
		if !strings.HasPrefix(typeName, "CG") {
			return "coregraphics.CG" + typeName
		}
		return "coregraphics." + typeName
	}

	// If we're in QuartzCore framework, all types are local
	if framework == "QuartzCore" && quartzCoreTypes[typeName] {
		return typeName
	}

	// If this is a known QuartzCore type and we're not in QuartzCore, qualify it
	if quartzCoreTypes[typeName] {
		return "quartzcore." + typeName
	}

	// If we're in Foundation framework, all types are local
	if framework == "Foundation" && foundationTypes[typeName] {
		return typeName
	}

	// If we're in AppKit (or other frameworks that embed NSObject), Foundation types are also local
	// since NSObject/Foundation is embedded in the object hierarchy
	// This includes most UI/system frameworks that depend on Foundation
	if (framework == "AppKit" || framework == "QuartzCore" || framework == "CoreData" ||
		framework == "Accessibility" || framework == "Accounts" || framework == "AddressBook" ||
		framework == "AdServices" || framework == "AdSupport" || framework == "Automator" ||
		framework == "CallKit" || framework == "ClassKit" || framework == "CloudKit" ||
		framework == "Collaboration" || framework == "Contacts" || framework == "ContactsUI" ||
		framework == "CoreLocationUI" || framework == "CryptoKit" || framework == "Darwin" ||
		framework == "DeviceCheck" || framework == "DocumentPickerUI" || framework == "EventKit" ||
		framework == "EventKitUI" || framework == "ExtensionKit" || framework == "FileProvider" ||
		framework == "FileProviderUI" || framework == "GameController" || framework == "GameKit" ||
		framework == "GLKit" || framework == "HealthKit" || framework == "HealthKitUI" ||
		framework == "HomeKit" || framework == "IOSurface" || framework == "LocalAuthentication" ||
		framework == "MapKit" || framework == "MediaAccessibility" || framework == "MediaKit" ||
		framework == "MessageUI" || framework == "Messages" || framework == "Metal" ||
		framework == "MetalKit" || framework == "MetalPerformanceShaders" || framework == "ModelIO" ||
		framework == "MultipeerConnectivity" || framework == "NaturalLanguage" || framework == "Network" ||
		framework == "NotificationCenter" || framework == "PDFKit" || framework == "PencilKit" ||
		framework == "Photos" || framework == "PhotosUI" || framework == "PlaygroundSupport" ||
		framework == "PushKit" || framework == "QuickLook" || framework == "RealityKit" ||
		framework == "SafariServices" || framework == "SceneKit" || framework == "ScreenTime" ||
		framework == "Security" || framework == "SensorKit" || framework == "ServiceManagement" ||
		framework == "SharedWithYou" || framework == "SharedWithYouCore" || framework == "ShazamKit" ||
		framework == "SiriKit" || framework == "Social" || framework == "SoundAnalysis" ||
		framework == "Speech" || framework == "SpriteKit" || framework == "StoreKit" ||
		framework == "SwiftUI" || framework == "SystemConfiguration" || framework == "ThreadNetwork" ||
		framework == "UserNotifications" || framework == "UserNotificationsUI" || framework == "VideoSubscriberAccount" ||
		framework == "VideoToolbox" || framework == "Vision" || framework == "VisionKit" ||
		framework == "WatchConnectivity" || framework == "WatchKit" || framework == "WebKit" ||
		framework == "WidgetKit") && foundationTypes[typeName] {
		return typeName
	}

	// If this is a known Foundation type and we're not in Foundation/AppKit, qualify it
	if foundationTypes[typeName] {
		return "foundation." + typeName
	}

	// If we're in AppKit framework, all types are local
	if framework == "AppKit" && appKitTypes[typeName] {
		return typeName
	}

	// If this is a known AppKit type and we're not in AppKit, qualify it
	if appKitTypes[typeName] {
		return "appkit." + typeName
	}

	// Check if we know about this type from the cross-framework registry
	if frameworkPkg, found := crossFrameworkTypeRegistry[typeName]; found {
		// Don't qualify types with their own framework name (e.g., foundation.NSString in foundation package)
		if strings.ToLower(framework) == frameworkPkg {
			return typeName
		}
		// NEVER qualify Go built-in primitives, even if they appear in cross-framework registry
		// This prevents errors like "appkit.string" when NSString resolves to "string"
		if isGoPrimitive(typeName) {
			return typeName
		}
		return frameworkPkg + "." + typeName
	}

	// Before falling back to unsafe.Pointer, check if this type belongs to the current framework
	// based on naming conventions. For example, in AppKit, types like NSView, NSButton, NSTextCheckingResult
	// should be returned as-is, not qualified with appkit.
	// This handles types that aren't classes (so not in currentFrameworkClasses) but are still
	// defined in the current framework's types.gen.go file.
	if framework != "" {
		// Check common framework prefixes
		frameworkPrefixes := map[string][]string{
			"AppKit":           {"NS", "AK"},
			"Foundation":       {"NS", "CF"},
			"CoreGraphics":     {"CG"},
			"QuartzCore":       {"CA"},
			"CoreImage":        {"CI"},
			"CoreData":         {"NS", "CD"},
			"AVFoundation":     {"AV"},
			"Metal":            {"MTL"},
			"MetalKit":         {"MTK"},
			"SpriteKit":        {"SK"},
			"SceneKit":         {"SCN"},
			"CoreML":           {"ML"},
			"Vision":           {"VN"},
			"CoreLocation":     {"CL"},
			"MapKit":           {"MK"},
			"PhotoKit":         {"PH"},
			"Photos":           {"PH"},
			"ScreenCaptureKit": {"SC"},
		}

		if prefixes, ok := frameworkPrefixes[framework]; ok {
			for _, prefix := range prefixes {
				if strings.HasPrefix(typeName, prefix) {
					// Type likely belongs to current framework, return as-is
					return typeName
				}
			}
		}
	}

	// Last resort: fall back to unsafe.Pointer for truly unknown types
	// This should be rare with a well-populated registry
	return "unsafe.Pointer"
}

// parameterToGoType converts an Objective-C parameter to a Go type.
// Uses the occ2go.MapCTypeToGo function with framework context.
func parameterToGoType(param occ2go.Parameter, framework string) string {
	paramType := strings.TrimSpace(param.Type)
	return occ2go.MapCTypeToGo(paramType, framework)
}

// wrapObjCReturn generates the return statement for converting objc.ID to Go types.
// It handles special cases like bool conversion and objc.Object mapping.
// Examples:
//
//	wrapObjCReturn("bool") -> "ret != 0"
//	wrapObjCReturn("objc.Object") -> "objc.ID(ret)"
//	wrapObjCReturn("int") -> "int(ret)"
func wrapObjCReturn(goType string) string {
	switch goType {
	case "bool":
		return "ret != 0"
	case "objc.ID":
		return "ret"
	case "unsafe.Pointer":
		return "unsafe.Pointer(ret)"
	default:
		// Default cast
		return fmt.Sprintf("%s(ret)", goType)
	}
}
