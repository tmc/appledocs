// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit


import (
	"unsafe"

	"github.com/ebitengine/purego"
	objc "github.com/ebitengine/purego/objc"
	corefoundation "github.com/tmc/appledocs/generated/corefoundation"
)


// AppKit Functions (68 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSAccessibilityActionDescription func(AccessibilityActionName) unsafe.Pointer
	_NSAccessibilityRoleDescriptionForUIElement func(objc.ID) unsafe.Pointer
	_NSAccessibilityRoleDescription func(AccessibilityRole, AccessibilitySubrole) unsafe.Pointer
	_NSAccessibilityPostNotification func(objc.ID, AccessibilityNotificationName)
	_NSAccessibilityPostNotificationWithUserInfo func(objc.ID, AccessibilityNotificationName, unsafe.Pointer)
	_NSAccessibilityRaiseBadArgumentException func(objc.ID, AccessibilityAttributeName, objc.ID)
	_NSAccessibilityPointInView func(unsafe.Pointer, corefoundation.CGPoint) corefoundation.CGPoint
	_NSAccessibilityFrameInView func(unsafe.Pointer, corefoundation.CGRect) corefoundation.CGRect
	_NSAccessibilitySetMayContainProtectedContent func(bool) bool
	_NSAccessibilityUnignoredAncestor func(objc.ID) objc.ID
	_NSAccessibilityUnignoredChildren func(unsafe.Pointer) unsafe.Pointer
	_NSAccessibilityUnignoredChildrenForOnlyChild func(objc.ID) unsafe.Pointer
	_NSAccessibilityUnignoredDescendant func(objc.ID) objc.ID
	_NSApplicationLoad func() bool
	_NSApplicationMain func(int, unsafe.Pointer) int
	_NSAvailableWindowDepths func() unsafe.Pointer
	_NSBeep func()
	_NSBeginAlertSheet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, objc.ID, objc.SEL, objc.SEL, unsafe.Pointer, unsafe.Pointer)
	_NSBeginCriticalAlertSheet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, objc.ID, objc.SEL, objc.SEL, unsafe.Pointer, unsafe.Pointer)
	_NSBeginInformationalAlertSheet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, objc.ID, objc.SEL, objc.SEL, unsafe.Pointer, unsafe.Pointer)
	_NSBestDepth func(ColorSpaceName, int, int, bool, unsafe.Pointer) WindowDepth
	_NSNumberOfColorComponents func(ColorSpaceName) int
	_NSConvertGlyphsToPackedGlyphs func(unsafe.Pointer, int, MultibyteGlyphPacking, unsafe.Pointer) int
	_NSCopyBits func(int, corefoundation.CGRect, corefoundation.CGPoint)
	_NSCountWindows func(unsafe.Pointer)
	_NSCountWindowsForContext func(int, unsafe.Pointer)
	_NSDisableScreenUpdates func()
	_NSDottedFrameRect func(corefoundation.CGRect)
	_NSDrawBitmap func(corefoundation.CGRect, int, int, int, int, int, int, bool, bool, ColorSpaceName, unsafe.Pointer, unsafe.Pointer)
	_NSDrawColorTiledRects func(corefoundation.CGRect, corefoundation.CGRect, unsafe.Pointer, unsafe.Pointer, int) corefoundation.CGRect
	_NSDrawTiledRects func(corefoundation.CGRect, corefoundation.CGRect, unsafe.Pointer, []float64, int) corefoundation.CGRect
	_NSEnableScreenUpdates func()
	_NSSetFocusRingStyle func(FocusRingPlacement)
	_NSGetAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) objc.ID
	_NSGetCriticalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) objc.ID
	_NSGetInformationalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) objc.ID
	_NSGetWindowServerMemory func(int, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_NSHighlightRect func(corefoundation.CGRect)
	_NSInterfaceStyleForKey func(unsafe.Pointer, unsafe.Pointer) InterfaceStyle
	_NSIsControllerMarker func(objc.ID) bool
	_NSOpenGLGetOption func(OpenGLGlobalOption, unsafe.Pointer)
	_NSOpenGLGetVersion func(unsafe.Pointer, unsafe.Pointer)
	_NSOpenGLSetOption func(OpenGLGlobalOption, unsafe.Pointer)
	_NSCreateFileContentsPboardType func(unsafe.Pointer) PasteboardType
	_NSCreateFilenamePboardType func(unsafe.Pointer) PasteboardType
	_NSGetFileType func(PasteboardType) unsafe.Pointer
	_NSGetFileTypes func([]unsafe.Pointer) []unsafe.Pointer
	_NSReadPixel func(corefoundation.CGPoint) unsafe.Pointer
	_NSReleaseAlertPanel func(objc.ID)
	_NSRunAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_NSRunAlertPanelRelativeToWindow func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_NSRunCriticalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_NSRunCriticalAlertPanelRelativeToWindow func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_NSRunInformationalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_NSRunInformationalAlertPanelRelativeToWindow func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_NSSetShowsServicesMenuItem func(unsafe.Pointer, bool) int
	_NSShowAnimationEffect func(AnimationEffect, corefoundation.CGPoint, corefoundation.CGSize, objc.ID, objc.SEL, unsafe.Pointer)
	_NSShowsServicesMenuItem func(unsafe.Pointer) bool
	_NSUpdateDynamicServices func()
	_NSBitsPerPixelFromDepth func(WindowDepth) int
	_NSBitsPerSampleFromDepth func(WindowDepth) int
	_NSColorSpaceFromDepth func(WindowDepth) ColorSpaceName
	_NSPlanarFromDepth func(WindowDepth) bool
	_NSWindowList func(int, int)
	_NSWindowListForContext func(int, int, int)
	_NSPerformService func(unsafe.Pointer, unsafe.Pointer) bool
	_NSRegisterServicesProvider func(objc.ID, ServiceProviderName)
	_NSUnregisterServicesProvider func(ServiceProviderName)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSAccessibilityActionDescription, lib, "NSAccessibilityActionDescription")
	tryRegister(&_NSAccessibilityRoleDescriptionForUIElement, lib, "NSAccessibilityRoleDescriptionForUIElement")
	tryRegister(&_NSAccessibilityRoleDescription, lib, "NSAccessibilityRoleDescription")
	tryRegister(&_NSAccessibilityPostNotification, lib, "NSAccessibilityPostNotification")
	tryRegister(&_NSAccessibilityPostNotificationWithUserInfo, lib, "NSAccessibilityPostNotificationWithUserInfo")
	tryRegister(&_NSAccessibilityRaiseBadArgumentException, lib, "NSAccessibilityRaiseBadArgumentException")
	tryRegister(&_NSAccessibilityPointInView, lib, "NSAccessibilityPointInView")
	tryRegister(&_NSAccessibilityFrameInView, lib, "NSAccessibilityFrameInView")
	tryRegister(&_NSAccessibilitySetMayContainProtectedContent, lib, "NSAccessibilitySetMayContainProtectedContent")
	tryRegister(&_NSAccessibilityUnignoredAncestor, lib, "NSAccessibilityUnignoredAncestor")
	tryRegister(&_NSAccessibilityUnignoredChildren, lib, "NSAccessibilityUnignoredChildren")
	tryRegister(&_NSAccessibilityUnignoredChildrenForOnlyChild, lib, "NSAccessibilityUnignoredChildrenForOnlyChild")
	tryRegister(&_NSAccessibilityUnignoredDescendant, lib, "NSAccessibilityUnignoredDescendant")
	tryRegister(&_NSApplicationLoad, lib, "NSApplicationLoad")
	tryRegister(&_NSApplicationMain, lib, "NSApplicationMain")
	tryRegister(&_NSAvailableWindowDepths, lib, "NSAvailableWindowDepths")
	tryRegister(&_NSBeep, lib, "NSBeep")
	tryRegister(&_NSBeginAlertSheet, lib, "NSBeginAlertSheet")
	tryRegister(&_NSBeginCriticalAlertSheet, lib, "NSBeginCriticalAlertSheet")
	tryRegister(&_NSBeginInformationalAlertSheet, lib, "NSBeginInformationalAlertSheet")
	tryRegister(&_NSBestDepth, lib, "NSBestDepth")
	tryRegister(&_NSNumberOfColorComponents, lib, "NSNumberOfColorComponents")
	tryRegister(&_NSConvertGlyphsToPackedGlyphs, lib, "NSConvertGlyphsToPackedGlyphs")
	tryRegister(&_NSCopyBits, lib, "NSCopyBits")
	tryRegister(&_NSCountWindows, lib, "NSCountWindows")
	tryRegister(&_NSCountWindowsForContext, lib, "NSCountWindowsForContext")
	tryRegister(&_NSDisableScreenUpdates, lib, "NSDisableScreenUpdates")
	tryRegister(&_NSDottedFrameRect, lib, "NSDottedFrameRect")
	tryRegister(&_NSDrawBitmap, lib, "NSDrawBitmap")
	tryRegister(&_NSDrawColorTiledRects, lib, "NSDrawColorTiledRects")
	tryRegister(&_NSDrawTiledRects, lib, "NSDrawTiledRects")
	tryRegister(&_NSEnableScreenUpdates, lib, "NSEnableScreenUpdates")
	tryRegister(&_NSSetFocusRingStyle, lib, "NSSetFocusRingStyle")
	tryRegister(&_NSGetAlertPanel, lib, "NSGetAlertPanel")
	tryRegister(&_NSGetCriticalAlertPanel, lib, "NSGetCriticalAlertPanel")
	tryRegister(&_NSGetInformationalAlertPanel, lib, "NSGetInformationalAlertPanel")
	tryRegister(&_NSGetWindowServerMemory, lib, "NSGetWindowServerMemory")
	tryRegister(&_NSHighlightRect, lib, "NSHighlightRect")
	tryRegister(&_NSInterfaceStyleForKey, lib, "NSInterfaceStyleForKey")
	tryRegister(&_NSIsControllerMarker, lib, "NSIsControllerMarker")
	tryRegister(&_NSOpenGLGetOption, lib, "NSOpenGLGetOption")
	tryRegister(&_NSOpenGLGetVersion, lib, "NSOpenGLGetVersion")
	tryRegister(&_NSOpenGLSetOption, lib, "NSOpenGLSetOption")
	tryRegister(&_NSCreateFileContentsPboardType, lib, "NSCreateFileContentsPboardType")
	tryRegister(&_NSCreateFilenamePboardType, lib, "NSCreateFilenamePboardType")
	tryRegister(&_NSGetFileType, lib, "NSGetFileType")
	tryRegister(&_NSGetFileTypes, lib, "NSGetFileTypes")
	tryRegister(&_NSReadPixel, lib, "NSReadPixel")
	tryRegister(&_NSReleaseAlertPanel, lib, "NSReleaseAlertPanel")
	tryRegister(&_NSRunAlertPanel, lib, "NSRunAlertPanel")
	tryRegister(&_NSRunAlertPanelRelativeToWindow, lib, "NSRunAlertPanelRelativeToWindow")
	tryRegister(&_NSRunCriticalAlertPanel, lib, "NSRunCriticalAlertPanel")
	tryRegister(&_NSRunCriticalAlertPanelRelativeToWindow, lib, "NSRunCriticalAlertPanelRelativeToWindow")
	tryRegister(&_NSRunInformationalAlertPanel, lib, "NSRunInformationalAlertPanel")
	tryRegister(&_NSRunInformationalAlertPanelRelativeToWindow, lib, "NSRunInformationalAlertPanelRelativeToWindow")
	tryRegister(&_NSSetShowsServicesMenuItem, lib, "NSSetShowsServicesMenuItem")
	tryRegister(&_NSShowAnimationEffect, lib, "NSShowAnimationEffect")
	tryRegister(&_NSShowsServicesMenuItem, lib, "NSShowsServicesMenuItem")
	tryRegister(&_NSUpdateDynamicServices, lib, "NSUpdateDynamicServices")
	tryRegister(&_NSBitsPerPixelFromDepth, lib, "NSBitsPerPixelFromDepth")
	tryRegister(&_NSBitsPerSampleFromDepth, lib, "NSBitsPerSampleFromDepth")
	tryRegister(&_NSColorSpaceFromDepth, lib, "NSColorSpaceFromDepth")
	tryRegister(&_NSPlanarFromDepth, lib, "NSPlanarFromDepth")
	tryRegister(&_NSWindowList, lib, "NSWindowList")
	tryRegister(&_NSWindowListForContext, lib, "NSWindowListForContext")
	tryRegister(&_NSPerformService, lib, "NSPerformService")
	tryRegister(&_NSRegisterServicesProvider, lib, "NSRegisterServicesProvider")
	tryRegister(&_NSUnregisterServicesProvider, lib, "NSUnregisterServicesProvider")
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
//
// Added in macOS .
// Returns a standard description for an action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/description
func NSAccessibilityActionDescription(action AccessibilityActionName) unsafe.Pointer {
	return _NSAccessibilityActionDescription(action)
}

// Returns a standard role description for a user interface element.
//
// Added in macOS .
// Returns a standard role description for a user interface element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/description(for:)
func NSAccessibilityRoleDescriptionForUIElement(element objc.ID) unsafe.Pointer {
	return _NSAccessibilityRoleDescriptionForUIElement(element)
}

// Returns a standard description for a role and subrole.
//
// Added in macOS .
// Returns a standard description for a role and subrole.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/description(with:)
func NSAccessibilityRoleDescription(role AccessibilityRole, subrole AccessibilitySubrole) unsafe.Pointer {
	return _NSAccessibilityRoleDescription(role, subrole)
}

// Sends a notification to any observing assistive apps.
//
// Added in macOS .
// Sends a notification to any observing assistive apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/post(element:notification:)
func NSAccessibilityPostNotification(element objc.ID, notification AccessibilityNotificationName) {
	_NSAccessibilityPostNotification(element, notification)
}

// Sends a notification and an optional user info dictionary to any observing assistive apps.
//
// Added in macOS 10.7.
// Sends a notification and an optional user info dictionary to any observing assistive apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/post(element:notification:userInfo:)
func NSAccessibilityPostNotificationWithUserInfo(element objc.ID, notification AccessibilityNotificationName, userInfo unsafe.Pointer) {
	_NSAccessibilityPostNotificationWithUserInfo(element, notification, userInfo)
}

// Raises an error if the parameter is the wrong type or has an illegal value
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.1.
// Raises an error if the parameter is the wrong type or has an illegal value
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/raiseBadArgumentException(_:_:_:)
func NSAccessibilityRaiseBadArgumentException(element objc.ID, attribute AccessibilityAttributeName, value objc.ID) {
	_NSAccessibilityRaiseBadArgumentException(element, attribute, value)
}

// Returns the point in screen coordinates.
//
// Added in macOS 10.10.
// Returns the point in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/screenPoint(fromView:point:)
func NSAccessibilityPointInView(parentView unsafe.Pointer, point corefoundation.CGPoint) corefoundation.CGPoint {
	return _NSAccessibilityPointInView(parentView, point)
}

// Returns the frame in screen coordinates.
//
// Added in macOS 10.10.
// Returns the frame in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/screenRect(fromView:rect:)
func NSAccessibilityFrameInView(parentView unsafe.Pointer, frame corefoundation.CGRect) corefoundation.CGRect {
	return _NSAccessibilityFrameInView(parentView, frame)
}

// Sets whether the app may have protected content.
//
// Added in macOS .
// Sets whether the app may have protected content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/setMayContainProtectedContent(_:)
func NSAccessibilitySetMayContainProtectedContent(flag bool) bool {
	return _NSAccessibilitySetMayContainProtectedContent(flag)
}

// Returns an unignored accessibility object, ascending the hierarchy, if necessary.
//
// Added in macOS .
// Returns an unignored accessibility object, ascending the hierarchy, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredAncestor(of:)
func NSAccessibilityUnignoredAncestor(element objc.ID) objc.ID {
	return _NSAccessibilityUnignoredAncestor(element)
}

// Returns a list of unignored accessibility objects, descending the hierarchy, if necessary.
//
// Added in macOS .
// Returns a list of unignored accessibility objects, descending the hierarchy, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredChildren(from:)
func NSAccessibilityUnignoredChildren(originalChildren unsafe.Pointer) unsafe.Pointer {
	return _NSAccessibilityUnignoredChildren(originalChildren)
}

// Returns a list of unignored accessibility objects, descending the hierarchy, if necessary.
//
// Added in macOS .
// Returns a list of unignored accessibility objects, descending the hierarchy, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredChildrenForOnlyChild(from:)
func NSAccessibilityUnignoredChildrenForOnlyChild(originalChild objc.ID) unsafe.Pointer {
	return _NSAccessibilityUnignoredChildrenForOnlyChild(originalChild)
}

// Returns an unignored accessibility object, descending the hierarchy, if necessary.
//
// Added in macOS .
// Returns an unignored accessibility object, descending the hierarchy, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredDescendant(of:)
func NSAccessibilityUnignoredDescendant(element objc.ID) objc.ID {
	return _NSAccessibilityUnignoredDescendant(element)
}

// Startup function to call when running Cocoa code from a Carbon application.
//
// Added in macOS .
// Startup function to call when running Cocoa code from a Carbon application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplicationLoad
func NSApplicationLoad() bool {
	return _NSApplicationLoad()
}

// Called by the main function to create and run the application.
//
// Added in macOS .
// Called by the main function to create and run the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplicationMain
func NSApplicationMain(argc int, argv unsafe.Pointer) int {
	return _NSApplicationMain(argc, argv)
}

// Returns the available window depth values.
//
// Added in macOS .
// Returns the available window depth values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAvailableWindowDepths
func NSAvailableWindowDepths() unsafe.Pointer {
	return _NSAvailableWindowDepths()
}

// Plays the system beep.
//
// Added in macOS .
// Plays the system beep.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBeep
func NSBeep() {
	_NSBeep()
}

// Creates and runs an alert sheet.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Creates and runs an alert sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBeginAlertSheet
func NSBeginAlertSheet(title unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer, modalDelegate objc.ID, didEndSelector objc.SEL, didDismissSelector objc.SEL, contextInfo unsafe.Pointer, msgFormat unsafe.Pointer) {
	_NSBeginAlertSheet(title, defaultButton, alternateButton, otherButton, docWindow, modalDelegate, didEndSelector, didDismissSelector, contextInfo, msgFormat)
}

// Creates and runs a critical alert sheet.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Creates and runs a critical alert sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBeginCriticalAlertSheet
func NSBeginCriticalAlertSheet(title unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer, modalDelegate objc.ID, didEndSelector objc.SEL, didDismissSelector objc.SEL, contextInfo unsafe.Pointer, msgFormat unsafe.Pointer) {
	_NSBeginCriticalAlertSheet(title, defaultButton, alternateButton, otherButton, docWindow, modalDelegate, didEndSelector, didDismissSelector, contextInfo, msgFormat)
}

// Creates and runs an informational alert sheet.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Creates and runs an informational alert sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBeginInformationalAlertSheet
func NSBeginInformationalAlertSheet(title unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer, modalDelegate objc.ID, didEndSelector objc.SEL, didDismissSelector objc.SEL, contextInfo unsafe.Pointer, msgFormat unsafe.Pointer) {
	_NSBeginInformationalAlertSheet(title, defaultButton, alternateButton, otherButton, docWindow, modalDelegate, didEndSelector, didDismissSelector, contextInfo, msgFormat)
}

// Attempts to return a window depth adequate for the specified parameters.
//
// Added in macOS .
// Attempts to return a window depth adequate for the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBestDepth
func NSBestDepth(colorSpace ColorSpaceName, bps int, bpp int, planar bool, exactMatch unsafe.Pointer) WindowDepth {
	return _NSBestDepth(colorSpace, bps, bpp, planar, exactMatch)
}

// Returns the number of color components in the specified color space.
//
// Added in macOS .
// Returns the number of color components in the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/numberOfColorComponents
func NSNumberOfColorComponents(colorSpaceName ColorSpaceName) int {
	return _NSNumberOfColorComponents(colorSpaceName)
}

// Prepares a set of glyphs for processing by character-based routines.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.0.
// Prepares a set of glyphs for processing by character-based routines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSConvertGlyphsToPackedGlyphs(_:_:_:_:)
func NSConvertGlyphsToPackedGlyphs(glBuf unsafe.Pointer, count int, packing MultibyteGlyphPacking, packedGlyphs unsafe.Pointer) int {
	return _NSConvertGlyphsToPackedGlyphs(glBuf, count, packing, packedGlyphs)
}

// Copies a bitmap image to the location specified by a destination point.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Copies a bitmap image to the location specified by a destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCopyBits(_:_:_:)
func NSCopyBits(srcGState int, srcRect corefoundation.CGRect, destPoint corefoundation.CGPoint) {
	_NSCopyBits(srcGState, srcRect, destPoint)
}

// Counts the number of onscreen windows.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Counts the number of onscreen windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCountWindows
func NSCountWindows(count unsafe.Pointer) {
	_NSCountWindows(count)
}

// Counts the number of onscreen windows belonging to a particular application.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Counts the number of onscreen windows belonging to a particular application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCountWindowsForContext
func NSCountWindowsForContext(context int, count unsafe.Pointer) {
	_NSCountWindowsForContext(context, count)
}

// Disables screen updates.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.0.
// Disables screen updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDisableScreenUpdates()
func NSDisableScreenUpdates() {
	_NSDisableScreenUpdates()
}

// Draws a bordered rectangle.
//
// Added in macOS .
// Draws a bordered rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDottedFrameRect(_:)
func NSDottedFrameRect(rect corefoundation.CGRect) {
	_NSDottedFrameRect(rect)
}

// Draws a bitmap image.
//
// Added in macOS .
// Draws a bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawBitmap(_:_:_:_:_:_:_:_:_:_:_:)
func NSDrawBitmap(rect corefoundation.CGRect, width int, height int, bps int, spp int, bpp int, bpr int, isPlanar bool, hasAlpha bool, colorSpaceName ColorSpaceName, data unsafe.Pointer, p11 unsafe.Pointer) {
	_NSDrawBitmap(rect, width, height, bps, spp, bpp, bpr, isPlanar, hasAlpha, colorSpaceName, data, p11)
}

// Draws a single-color, bordered rectangle.
//
// Added in macOS .
// Draws a single-color, bordered rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawColorTiledRects(_:_:_:_:_:)
func NSDrawColorTiledRects(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides unsafe.Pointer, colors unsafe.Pointer, count int) corefoundation.CGRect {
	return _NSDrawColorTiledRects(boundsRect, clipRect, sides, colors, count)
}

// Draws rectangles with borders.
//
// Added in macOS .
// Draws rectangles with borders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawTiledRects(_:_:_:_:_:)
func NSDrawTiledRects(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides unsafe.Pointer, grays []float64, count int) corefoundation.CGRect {
	return _NSDrawTiledRects(boundsRect, clipRect, sides, grays, count)
}

// Enables screen updates.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.0.
// Enables screen updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEnableScreenUpdates()
func NSEnableScreenUpdates() {
	_NSEnableScreenUpdates()
}

// Specifies how the system draws the focus ring.
//
// Added in macOS .
// Specifies how the system draws the focus ring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/set()
func NSSetFocusRingStyle(placement FocusRingPlacement) {
	_NSSetFocusRingStyle(placement)
}

// Returns an alert panel.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Returns an alert panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGetAlertPanel
func NSGetAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) objc.ID {
	return _NSGetAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
}

// Returns an alert panel to display a critical message.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Returns an alert panel to display a critical message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGetCriticalAlertPanel
func NSGetCriticalAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) objc.ID {
	return _NSGetCriticalAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
}

// Returns an alert panel to display an informational message.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Returns an alert panel to display an informational message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGetInformationalAlertPanel
func NSGetInformationalAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) objc.ID {
	return _NSGetInformationalAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
}

// Returns the amount of memory being used by a context.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.0.
// Returns the amount of memory being used by a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGetWindowServerMemory(_:_:_:_:)
func NSGetWindowServerMemory(context int, virtualMemory unsafe.Pointer, windowBackingMemory unsafe.Pointer, windowDumpString unsafe.Pointer) int {
	return _NSGetWindowServerMemory(context, virtualMemory, windowBackingMemory, windowDumpString)
}

// Highlights the specified rect by filling it with white.
//
// Deprecated: This function was deprecated in macOS 10.0.
//
// Added in macOS 10.0.
// Highlights the specified rect by filling it with white.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHighlightRect
func NSHighlightRect(rect corefoundation.CGRect) {
	_NSHighlightRect(rect)
}

// Returns an interface style value for the specified key and responder.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns an interface style value for the specified key and responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInterfaceStyleForKey
func NSInterfaceStyleForKey(key unsafe.Pointer, responder unsafe.Pointer) InterfaceStyle {
	return _NSInterfaceStyleForKey(key, responder)
}

// Tests whether a given object is special marker object used for indicating the state of a selection in relation to a key.
//
// Added in macOS .
// Tests whether a given object is special marker object used for indicating the state of a selection in relation to a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSIsControllerMarker(_:)
func NSIsControllerMarker(object objc.ID) bool {
	return _NSIsControllerMarker(object)
}

// Returns global OpenGL options.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.0.
// Returns global OpenGL options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGetOption
func NSOpenGLGetOption(pname OpenGLGlobalOption, param unsafe.Pointer) {
	_NSOpenGLGetOption(pname, param)
}

// Returns the NSOpenGL version numbers.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.0.
// Returns the NSOpenGL version numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGetVersion
func NSOpenGLGetVersion(major unsafe.Pointer, minor unsafe.Pointer) {
	_NSOpenGLGetVersion(major, minor)
}

// Sets global OpenGL options.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.0.
// Sets global OpenGL options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLSetOption
func NSOpenGLSetOption(pname OpenGLGlobalOption, param unsafe.Pointer) {
	_NSOpenGLSetOption(pname, param)
}

// Returns a pasteboard type based on the passed file type.
//
// Added in macOS .
// Returns a pasteboard type based on the passed file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileContentsType(forPathExtension:)
func NSCreateFileContentsPboardType(fileType unsafe.Pointer) PasteboardType {
	return _NSCreateFileContentsPboardType(fileType)
}

// Returns a pasteboard type based on the passed file type.
//
// Added in macOS .
// Returns a pasteboard type based on the passed file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileNameType(forPathExtension:)
func NSCreateFilenamePboardType(fileType unsafe.Pointer) PasteboardType {
	return _NSCreateFilenamePboardType(fileType)
}

// A file type based on the passed pasteboard type.
//
// Added in macOS .
// A file type based on the passed pasteboard type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/representedPathExtension
func NSGetFileType(pboardType PasteboardType) unsafe.Pointer {
	return _NSGetFileType(pboardType)
}

// Returns an array of file types based on the passed pasteboard types.
//
// Added in macOS .
// Returns an array of file types based on the passed pasteboard types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/representedPathExtensions(from:)
func NSGetFileTypes(pboardTypes []unsafe.Pointer) []unsafe.Pointer {
	return _NSGetFileTypes(pboardTypes)
}

// Reads the color of the pixel at the specified location.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.0.
// Reads the color of the pixel at the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSReadPixel(_:)
func NSReadPixel(passedPoint corefoundation.CGPoint) unsafe.Pointer {
	return _NSReadPixel(passedPoint)
}

// Disposes of an alert panel.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Disposes of an alert panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSReleaseAlertPanel(_:)
func NSReleaseAlertPanel(panel objc.ID) {
	_NSReleaseAlertPanel(panel)
}

// Creates an alert panel.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Creates an alert panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunAlertPanel
func NSRunAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) int {
	return _NSRunAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
}

// NSRunAlertPanelRelativeToWindow is a AppKit function.
//
// Deprecated: This function was deprecated in macOS 10.0.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunAlertPanelRelativeToWindow
func NSRunAlertPanelRelativeToWindow(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer) int {
	return _NSRunAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow)
}

// Creates and runs a critical alert panel.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Creates and runs a critical alert panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunCriticalAlertPanel
func NSRunCriticalAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) int {
	return _NSRunCriticalAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
}

// NSRunCriticalAlertPanelRelativeToWindow is a AppKit function.
//
// Deprecated: This function was deprecated in macOS 10.0.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunCriticalAlertPanelRelativeToWindow
func NSRunCriticalAlertPanelRelativeToWindow(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer) int {
	return _NSRunCriticalAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow)
}

// Creates and runs an informational alert panel.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Creates and runs an informational alert panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunInformationalAlertPanel
func NSRunInformationalAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) int {
	return _NSRunInformationalAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
}

// NSRunInformationalAlertPanelRelativeToWindow is a AppKit function.
//
// Deprecated: This function was deprecated in macOS 10.0.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunInformationalAlertPanelRelativeToWindow
func NSRunInformationalAlertPanelRelativeToWindow(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer) int {
	return _NSRunInformationalAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow)
}

// Specifies whether an item should be included in Services menus.
//
// Added in macOS .
// Specifies whether an item should be included in Services menus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSetShowsServicesMenuItem(_:_:)
func NSSetShowsServicesMenuItem(itemName unsafe.Pointer, enabled bool) int {
	return _NSSetShowsServicesMenuItem(itemName, enabled)
}

// Runs a system animation effect.
//
// Deprecated: This function was deprecated in macOS 14.0.
//
// Added in macOS 10.3.
// Runs a system animation effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShowAnimationEffect
func NSShowAnimationEffect(animationEffect AnimationEffect, centerLocation corefoundation.CGPoint, size corefoundation.CGSize, animationDelegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer) {
	_NSShowAnimationEffect(animationEffect, centerLocation, size, animationDelegate, didEndSelector, contextInfo)
}

// Specifies whether a Services menu item is currently enabled.
//
// Added in macOS .
// Specifies whether a Services menu item is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShowsServicesMenuItem(_:)
func NSShowsServicesMenuItem(itemName unsafe.Pointer) bool {
	return _NSShowsServicesMenuItem(itemName)
}

// Causes the services information for the system to be updated.
//
// Added in macOS .
// Causes the services information for the system to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUpdateDynamicServices()
func NSUpdateDynamicServices() {
	_NSUpdateDynamicServices()
}

// Returns the bits per pixel for the specified window depth.
//
// Added in macOS .
// Returns the bits per pixel for the specified window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/bitsPerPixel
func NSBitsPerPixelFromDepth(depth WindowDepth) int {
	return _NSBitsPerPixelFromDepth(depth)
}

// Returns the bits per sample for the specified window depth.
//
// Added in macOS .
// Returns the bits per sample for the specified window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/bitsPerSample
func NSBitsPerSampleFromDepth(depth WindowDepth) int {
	return _NSBitsPerSampleFromDepth(depth)
}

// Returns the name of the color space corresponding to the passed window depth.
//
// Added in macOS .
// Returns the name of the color space corresponding to the passed window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/colorSpaceName
func NSColorSpaceFromDepth(depth WindowDepth) ColorSpaceName {
	return _NSColorSpaceFromDepth(depth)
}

// Returns whether the specified window depth is planar.
//
// Added in macOS .
// Returns whether the specified window depth is planar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/isPlanar
func NSPlanarFromDepth(depth WindowDepth) bool {
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
func NSWindowList(size int, list int) {
	_NSWindowList(size, list)
}

// Gets information about an application’s onscreen windows.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Gets information about an application’s onscreen windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowListForContext
func NSWindowListForContext(context int, size int, list int) {
	_NSWindowListForContext(context, size, list)
}

// Programmatically invokes a Services menu service.
//
// Added in macOS .
// Programmatically invokes a Services menu service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPerformService(_:_:)
func NSPerformService(itemName unsafe.Pointer, pboard unsafe.Pointer) bool {
	return _NSPerformService(itemName, pboard)
}

// Registers a service provider.
//
// Added in macOS .
// Registers a service provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRegisterServicesProvider(_:_:)
func NSRegisterServicesProvider(provider objc.ID, name ServiceProviderName) {
	_NSRegisterServicesProvider(provider, name)
}

// Unregisters a service provider.
//
// Added in macOS .
// Unregisters a service provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnregisterServicesProvider(_:)
func NSUnregisterServicesProvider(name ServiceProviderName) {
	_NSUnregisterServicesProvider(name)
}




