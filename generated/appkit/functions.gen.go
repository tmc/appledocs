// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// AppKit Functions (21 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSAccessibilityActionDescription func(unsafe.Pointer) unsafe.Pointer
	_NSAccessibilityRoleDescriptionForUIElement func(unsafe.Pointer) unsafe.Pointer
	_NSAccessibilityPostNotificationWithUserInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_NSApplicationLoad func() bool
	_NSAvailableWindowDepths func() unsafe.Pointer
	_NSBeep func()
	_NSBestDepth func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool, unsafe.Pointer) unsafe.Pointer
	_NSNumberOfColorComponents func(unsafe.Pointer) unsafe.Pointer
	_NSDrawNinePartImage func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, float64, bool)
	_NSDrawThreePartImage func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool, unsafe.Pointer, float64, bool)
	_NSIsControllerMarker func(unsafe.Pointer) bool
	_NSCreateFileContentsPboardType func(unsafe.Pointer) unsafe.Pointer
	_NSCreateFilenamePboardType func(unsafe.Pointer) unsafe.Pointer
	_NSGetFileType func(unsafe.Pointer) unsafe.Pointer
	_NSGetFileTypes func(unsafe.Pointer) unsafe.Pointer
	_NSUpdateDynamicServices func()
	_NSBitsPerPixelFromDepth func(unsafe.Pointer) unsafe.Pointer
	_NSBitsPerSampleFromDepth func(unsafe.Pointer) unsafe.Pointer
	_NSColorSpaceFromDepth func(unsafe.Pointer) unsafe.Pointer
	_NSPlanarFromDepth func(unsafe.Pointer) bool
	_NSWindowList func(unsafe.Pointer, unsafe.Pointer)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSAccessibilityActionDescription, lib, "NSAccessibilityActionDescription")
	tryRegister(&_NSAccessibilityRoleDescriptionForUIElement, lib, "NSAccessibilityRoleDescriptionForUIElement")
	tryRegister(&_NSAccessibilityPostNotificationWithUserInfo, lib, "NSAccessibilityPostNotificationWithUserInfo")
	tryRegister(&_NSApplicationLoad, lib, "NSApplicationLoad")
	tryRegister(&_NSAvailableWindowDepths, lib, "NSAvailableWindowDepths")
	tryRegister(&_NSBeep, lib, "NSBeep")
	tryRegister(&_NSBestDepth, lib, "NSBestDepth")
	tryRegister(&_NSNumberOfColorComponents, lib, "NSNumberOfColorComponents")
	tryRegister(&_NSDrawNinePartImage, lib, "NSDrawNinePartImage")
	tryRegister(&_NSDrawThreePartImage, lib, "NSDrawThreePartImage")
	tryRegister(&_NSIsControllerMarker, lib, "NSIsControllerMarker")
	tryRegister(&_NSCreateFileContentsPboardType, lib, "NSCreateFileContentsPboardType")
	tryRegister(&_NSCreateFilenamePboardType, lib, "NSCreateFilenamePboardType")
	tryRegister(&_NSGetFileType, lib, "NSGetFileType")
	tryRegister(&_NSGetFileTypes, lib, "NSGetFileTypes")
	tryRegister(&_NSUpdateDynamicServices, lib, "NSUpdateDynamicServices")
	tryRegister(&_NSBitsPerPixelFromDepth, lib, "NSBitsPerPixelFromDepth")
	tryRegister(&_NSBitsPerSampleFromDepth, lib, "NSBitsPerSampleFromDepth")
	tryRegister(&_NSColorSpaceFromDepth, lib, "NSColorSpaceFromDepth")
	tryRegister(&_NSPlanarFromDepth, lib, "NSPlanarFromDepth")
	tryRegister(&_NSWindowList, lib, "NSWindowList")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Returns a standard description for an action.

// Returns a standard description for an action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/description
func NSAccessibilityActionDescription(action unsafe.Pointer) unsafe.Pointer {
	return _NSAccessibilityActionDescription(action)
}

// Returns a standard role description for a user interface element.

// Returns a standard role description for a user interface element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/description(for:)
func NSAccessibilityRoleDescriptionForUIElement(element unsafe.Pointer) unsafe.Pointer {
	return _NSAccessibilityRoleDescriptionForUIElement(element)
}

// Sends a notification and an optional user info dictionary to any observing assistive apps.
//
// Added in macOS 10.7.
// Sends a notification and an optional user info dictionary to any observing assistive apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/post(element:notification:userInfo:)
func NSAccessibilityPostNotificationWithUserInfo(element unsafe.Pointer, notification unsafe.Pointer, userInfo unsafe.Pointer) {
	_NSAccessibilityPostNotificationWithUserInfo(element, notification, userInfo)
}

// Startup function to call when running Cocoa code from a Carbon application.

// Startup function to call when running Cocoa code from a Carbon application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplicationLoad
func NSApplicationLoad() bool {
	return _NSApplicationLoad()
}

// Returns the available window depth values.

// Returns the available window depth values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAvailableWindowDepths
func NSAvailableWindowDepths() unsafe.Pointer {
	return _NSAvailableWindowDepths()
}

// Plays the system beep.

// Plays the system beep.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBeep
func NSBeep() {
	_NSBeep()
}

// Attempts to return a window depth adequate for the specified parameters.

// Attempts to return a window depth adequate for the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBestDepth
func NSBestDepth(colorSpace unsafe.Pointer, bps unsafe.Pointer, bpp unsafe.Pointer, planar bool, exactMatch unsafe.Pointer) unsafe.Pointer {
	return _NSBestDepth(colorSpace, bps, bpp, planar, exactMatch)
}

// Returns the number of color components in the specified color space.

// Returns the number of color components in the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/numberOfColorComponents
func NSNumberOfColorComponents(colorSpaceName unsafe.Pointer) unsafe.Pointer {
	return _NSNumberOfColorComponents(colorSpaceName)
}

// Draws a nine-part tiled image.
//
// Added in macOS 10.5.
// Draws a nine-part tiled image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawNinePartImage(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func NSDrawNinePartImage(frame coregraphics.CGRect, topLeftCorner unsafe.Pointer, topEdgeFill unsafe.Pointer, topRightCorner unsafe.Pointer, leftEdgeFill unsafe.Pointer, centerFill unsafe.Pointer, rightEdgeFill unsafe.Pointer, bottomLeftCorner unsafe.Pointer, bottomEdgeFill unsafe.Pointer, bottomRightCorner unsafe.Pointer, op unsafe.Pointer, alphaFraction float64, flipped bool) {
	_NSDrawNinePartImage(frame, topLeftCorner, topEdgeFill, topRightCorner, leftEdgeFill, centerFill, rightEdgeFill, bottomLeftCorner, bottomEdgeFill, bottomRightCorner, op, alphaFraction, flipped)
}

// Draws a three-part tiled image.
//
// Added in macOS 10.5.
// Draws a three-part tiled image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawThreePartImage(_:_:_:_:_:_:_:_:)
func NSDrawThreePartImage(frame coregraphics.CGRect, startCap unsafe.Pointer, centerFill unsafe.Pointer, endCap unsafe.Pointer, vertical bool, op unsafe.Pointer, alphaFraction float64, flipped bool) {
	_NSDrawThreePartImage(frame, startCap, centerFill, endCap, vertical, op, alphaFraction, flipped)
}

// Tests whether a given object is special marker object used for indicating the state of a selection in relation to a key.

// Tests whether a given object is special marker object used for indicating the state of a selection in relation to a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSIsControllerMarker(_:)
func NSIsControllerMarker(object unsafe.Pointer) bool {
	return _NSIsControllerMarker(object)
}

// Returns a pasteboard type based on the passed file type.

// Returns a pasteboard type based on the passed file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileContentsType(forPathExtension:)
func NSCreateFileContentsPboardType(fileType unsafe.Pointer) unsafe.Pointer {
	return _NSCreateFileContentsPboardType(fileType)
}

// Returns a pasteboard type based on the passed file type.

// Returns a pasteboard type based on the passed file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileNameType(forPathExtension:)
func NSCreateFilenamePboardType(fileType unsafe.Pointer) unsafe.Pointer {
	return _NSCreateFilenamePboardType(fileType)
}

// A file type based on the passed pasteboard type.

// A file type based on the passed pasteboard type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/representedPathExtension
func NSGetFileType(pboardType unsafe.Pointer) unsafe.Pointer {
	return _NSGetFileType(pboardType)
}

// Returns an array of file types based on the passed pasteboard types.

// Returns an array of file types based on the passed pasteboard types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/representedPathExtensions(from:)
func NSGetFileTypes(pboardTypes unsafe.Pointer) unsafe.Pointer {
	return _NSGetFileTypes(pboardTypes)
}

// Causes the services information for the system to be updated.

// Causes the services information for the system to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUpdateDynamicServices()
func NSUpdateDynamicServices() {
	_NSUpdateDynamicServices()
}

// Returns the bits per pixel for the specified window depth.

// Returns the bits per pixel for the specified window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/bitsPerPixel
func NSBitsPerPixelFromDepth(depth unsafe.Pointer) unsafe.Pointer {
	return _NSBitsPerPixelFromDepth(depth)
}

// Returns the bits per sample for the specified window depth.

// Returns the bits per sample for the specified window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/bitsPerSample
func NSBitsPerSampleFromDepth(depth unsafe.Pointer) unsafe.Pointer {
	return _NSBitsPerSampleFromDepth(depth)
}

// Returns the name of the color space corresponding to the passed window depth.

// Returns the name of the color space corresponding to the passed window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/colorSpaceName
func NSColorSpaceFromDepth(depth unsafe.Pointer) unsafe.Pointer {
	return _NSColorSpaceFromDepth(depth)
}

// Returns whether the specified window depth is planar.

// Returns whether the specified window depth is planar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/isPlanar
func NSPlanarFromDepth(depth unsafe.Pointer) bool {
	return _NSPlanarFromDepth(depth)
}

// Gets information about onscreen windows.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Gets information about onscreen windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowList
func NSWindowList(size unsafe.Pointer, list unsafe.Pointer) {
	_NSWindowList(size, list)
}



