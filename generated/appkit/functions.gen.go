// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// AppKit Functions (85 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSAccessibilityActionDescription func(unsafe.Pointer) unsafe.Pointer
	_NSAccessibilityRoleDescriptionForUIElement func(unsafe.Pointer) unsafe.Pointer
	_NSAccessibilityPostNotification func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSAccessibilityPostNotificationWithUserInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSAccessibilityRaiseBadArgumentException func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSAccessibilityPointInView func(unsafe.Pointer, coregraphics.CGPoint) coregraphics.CGPoint
	_NSAccessibilityFrameInView func(unsafe.Pointer, coregraphics.CGRect) coregraphics.CGRect
	_NSAccessibilityUnignoredAncestor func(unsafe.Pointer) unsafe.Pointer
	_NSAccessibilityUnignoredChildren func(unsafe.Pointer) unsafe.Pointer
	_NSAccessibilityUnignoredChildrenForOnlyChild func(unsafe.Pointer) unsafe.Pointer
	_NSApplicationMain func(int, unsafe.Pointer) int
	_NSAvailableWindowDepths func() unsafe.Pointer
	_NSBeep func() unsafe.Pointer
	_NSBeginAlertSheet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSBeginCriticalAlertSheet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSBeginInformationalAlertSheet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSBestDepth func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool, unsafe.Pointer) unsafe.Pointer
	_NSNumberOfColorComponents func(unsafe.Pointer) unsafe.Pointer
	_NSConvertGlyphsToPackedGlyphs func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCopyBits func(unsafe.Pointer, coregraphics.CGRect, coregraphics.CGPoint) unsafe.Pointer
	_NSCountWindows func(unsafe.Pointer) unsafe.Pointer
	_NSCountWindowsForContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDisableScreenUpdates func() unsafe.Pointer
	_NSDottedFrameRect func(coregraphics.CGRect) unsafe.Pointer
	_NSDrawButton func(coregraphics.CGRect, coregraphics.CGRect) unsafe.Pointer
	_NSDrawColorTiledRects func(coregraphics.CGRect, coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) coregraphics.CGRect
	_NSDrawDarkBezel func(coregraphics.CGRect, coregraphics.CGRect) unsafe.Pointer
	_NSDrawGrayBezel func(coregraphics.CGRect, coregraphics.CGRect) unsafe.Pointer
	_NSDrawGroove func(coregraphics.CGRect, coregraphics.CGRect) unsafe.Pointer
	_NSDrawLightBezel func(coregraphics.CGRect, coregraphics.CGRect) unsafe.Pointer
	_NSDrawNinePartImage func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, float64, bool) unsafe.Pointer
	_NSDrawThreePartImage func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool, unsafe.Pointer, float64, bool) unsafe.Pointer
	_NSDrawTiledRects func(coregraphics.CGRect, coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) coregraphics.CGRect
	_NSDrawWhiteBezel func(coregraphics.CGRect, coregraphics.CGRect) unsafe.Pointer
	_NSDrawWindowBackground func(coregraphics.CGRect) unsafe.Pointer
	_NSEnableScreenUpdates func() unsafe.Pointer
	_NSEraseRect func(coregraphics.CGRect) unsafe.Pointer
	_NSSetFocusRingStyle func(unsafe.Pointer) unsafe.Pointer
	_NSFrameRect func(coregraphics.CGRect) unsafe.Pointer
	_NSFrameRectWithWidth func(coregraphics.CGRect, float64) unsafe.Pointer
	_NSFrameRectWithWidthUsingOperation func(coregraphics.CGRect, float64, unsafe.Pointer) unsafe.Pointer
	_NSGetAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSGetCriticalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSGetInformationalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSGetWindowServerMemory func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSHighlightRect func(coregraphics.CGRect) unsafe.Pointer
	_NSInterfaceStyleForKey func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSIsControllerMarker func(unsafe.Pointer) bool
	_NSOpenGLGetOption func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSOpenGLGetVersion func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSOpenGLSetOption func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCreateFileContentsPboardType func(unsafe.Pointer) unsafe.Pointer
	_NSCreateFilenamePboardType func(unsafe.Pointer) unsafe.Pointer
	_NSGetFileType func(unsafe.Pointer) unsafe.Pointer
	_NSGetFileTypes func(unsafe.Pointer) unsafe.Pointer
	_NSPerformService func(unsafe.Pointer, unsafe.Pointer) bool
	_NSReadPixel func(coregraphics.CGPoint) unsafe.Pointer
	_NSRectClip func(coregraphics.CGRect) unsafe.Pointer
	_NSRectClipList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRectFill func(coregraphics.CGRect) unsafe.Pointer
	_NSRectFillList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRectFillListUsingOperation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRectFillListWithColors func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRectFillListWithColorsUsingOperation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRectFillListWithGrays func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRectFillUsingOperation func(coregraphics.CGRect, unsafe.Pointer) unsafe.Pointer
	_NSRegisterServicesProvider func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSReleaseAlertPanel func(unsafe.Pointer) unsafe.Pointer
	_NSRunAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRunAlertPanelRelativeToWindow func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRunCriticalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRunCriticalAlertPanelRelativeToWindow func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRunInformationalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSRunInformationalAlertPanelRelativeToWindow func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSSetShowsServicesMenuItem func(unsafe.Pointer, bool) unsafe.Pointer
	_NSShowAnimationEffect func(unsafe.Pointer, coregraphics.CGPoint, coregraphics.CGSize, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSShowsServicesMenuItem func(unsafe.Pointer) bool
	_NSUnregisterServicesProvider func(unsafe.Pointer) unsafe.Pointer
	_NSUpdateDynamicServices func() unsafe.Pointer
	_NSBitsPerPixelFromDepth func(unsafe.Pointer) unsafe.Pointer
	_NSBitsPerSampleFromDepth func(unsafe.Pointer) unsafe.Pointer
	_NSColorSpaceFromDepth func(unsafe.Pointer) unsafe.Pointer
	_NSPlanarFromDepth func(unsafe.Pointer) bool
	_NSWindowList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSWindowListForContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSAccessibilityActionDescription, lib, "NSAccessibilityActionDescription")
	tryRegister(&_NSAccessibilityRoleDescriptionForUIElement, lib, "NSAccessibilityRoleDescriptionForUIElement")
	tryRegister(&_NSAccessibilityPostNotification, lib, "NSAccessibilityPostNotification")
	tryRegister(&_NSAccessibilityPostNotificationWithUserInfo, lib, "NSAccessibilityPostNotificationWithUserInfo")
	tryRegister(&_NSAccessibilityRaiseBadArgumentException, lib, "NSAccessibilityRaiseBadArgumentException")
	tryRegister(&_NSAccessibilityPointInView, lib, "NSAccessibilityPointInView")
	tryRegister(&_NSAccessibilityFrameInView, lib, "NSAccessibilityFrameInView")
	tryRegister(&_NSAccessibilityUnignoredAncestor, lib, "NSAccessibilityUnignoredAncestor")
	tryRegister(&_NSAccessibilityUnignoredChildren, lib, "NSAccessibilityUnignoredChildren")
	tryRegister(&_NSAccessibilityUnignoredChildrenForOnlyChild, lib, "NSAccessibilityUnignoredChildrenForOnlyChild")
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
	tryRegister(&_NSDrawButton, lib, "NSDrawButton")
	tryRegister(&_NSDrawColorTiledRects, lib, "NSDrawColorTiledRects")
	tryRegister(&_NSDrawDarkBezel, lib, "NSDrawDarkBezel")
	tryRegister(&_NSDrawGrayBezel, lib, "NSDrawGrayBezel")
	tryRegister(&_NSDrawGroove, lib, "NSDrawGroove")
	tryRegister(&_NSDrawLightBezel, lib, "NSDrawLightBezel")
	tryRegister(&_NSDrawNinePartImage, lib, "NSDrawNinePartImage")
	tryRegister(&_NSDrawThreePartImage, lib, "NSDrawThreePartImage")
	tryRegister(&_NSDrawTiledRects, lib, "NSDrawTiledRects")
	tryRegister(&_NSDrawWhiteBezel, lib, "NSDrawWhiteBezel")
	tryRegister(&_NSDrawWindowBackground, lib, "NSDrawWindowBackground")
	tryRegister(&_NSEnableScreenUpdates, lib, "NSEnableScreenUpdates")
	tryRegister(&_NSEraseRect, lib, "NSEraseRect")
	tryRegister(&_NSSetFocusRingStyle, lib, "NSSetFocusRingStyle")
	tryRegister(&_NSFrameRect, lib, "NSFrameRect")
	tryRegister(&_NSFrameRectWithWidth, lib, "NSFrameRectWithWidth")
	tryRegister(&_NSFrameRectWithWidthUsingOperation, lib, "NSFrameRectWithWidthUsingOperation")
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
	tryRegister(&_NSPerformService, lib, "NSPerformService")
	tryRegister(&_NSReadPixel, lib, "NSReadPixel")
	tryRegister(&_NSRectClip, lib, "NSRectClip")
	tryRegister(&_NSRectClipList, lib, "NSRectClipList")
	tryRegister(&_NSRectFill, lib, "NSRectFill")
	tryRegister(&_NSRectFillList, lib, "NSRectFillList")
	tryRegister(&_NSRectFillListUsingOperation, lib, "NSRectFillListUsingOperation")
	tryRegister(&_NSRectFillListWithColors, lib, "NSRectFillListWithColors")
	tryRegister(&_NSRectFillListWithColorsUsingOperation, lib, "NSRectFillListWithColorsUsingOperation")
	tryRegister(&_NSRectFillListWithGrays, lib, "NSRectFillListWithGrays")
	tryRegister(&_NSRectFillUsingOperation, lib, "NSRectFillUsingOperation")
	tryRegister(&_NSRegisterServicesProvider, lib, "NSRegisterServicesProvider")
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
	tryRegister(&_NSUnregisterServicesProvider, lib, "NSUnregisterServicesProvider")
	tryRegister(&_NSUpdateDynamicServices, lib, "NSUpdateDynamicServices")
	tryRegister(&_NSBitsPerPixelFromDepth, lib, "NSBitsPerPixelFromDepth")
	tryRegister(&_NSBitsPerSampleFromDepth, lib, "NSBitsPerSampleFromDepth")
	tryRegister(&_NSColorSpaceFromDepth, lib, "NSColorSpaceFromDepth")
	tryRegister(&_NSPlanarFromDepth, lib, "NSPlanarFromDepth")
	tryRegister(&_NSWindowList, lib, "NSWindowList")
	tryRegister(&_NSWindowListForContext, lib, "NSWindowListForContext")
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


// Sends a notification to any observing assistive apps.

// Sends a notification to any observing assistive apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/post(element:notification:)

func NSAccessibilityPostNotification(element unsafe.Pointer, notification unsafe.Pointer) {
	_NSAccessibilityPostNotification(element, notification)
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


// Raises an error if the parameter is the wrong type or has an illegal value
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.1.

// Raises an error if the parameter is the wrong type or has an illegal value
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/raiseBadArgumentException(_:_:_:)

func NSAccessibilityRaiseBadArgumentException(element unsafe.Pointer, attribute unsafe.Pointer, value unsafe.Pointer) {
	_NSAccessibilityRaiseBadArgumentException(element, attribute, value)
	}


// Returns the point in screen coordinates.
//
// Added in macOS 10.10.

// Returns the point in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/screenPoint(fromView:point:)

func NSAccessibilityPointInView(parentView unsafe.Pointer, point coregraphics.CGPoint) coregraphics.CGPoint {
	return _NSAccessibilityPointInView(parentView, point)
	}


// Returns the frame in screen coordinates.
//
// Added in macOS 10.10.

// Returns the frame in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/screenRect(fromView:rect:)

func NSAccessibilityFrameInView(parentView unsafe.Pointer, frame coregraphics.CGRect) coregraphics.CGRect {
	return _NSAccessibilityFrameInView(parentView, frame)
	}


// Returns an unignored accessibility object, ascending the hierarchy, if necessary.

// Returns an unignored accessibility object, ascending the hierarchy, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredAncestor(of:)

func NSAccessibilityUnignoredAncestor(element unsafe.Pointer) unsafe.Pointer {
	return _NSAccessibilityUnignoredAncestor(element)
	}


// Returns a list of unignored accessibility objects, descending the hierarchy, if necessary.

// Returns a list of unignored accessibility objects, descending the hierarchy, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredChildren(from:)

func NSAccessibilityUnignoredChildren(originalChildren unsafe.Pointer) unsafe.Pointer {
	return _NSAccessibilityUnignoredChildren(originalChildren)
	}


// Returns a list of unignored accessibility objects, descending the hierarchy, if necessary.

// Returns a list of unignored accessibility objects, descending the hierarchy, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredChildrenForOnlyChild(from:)

func NSAccessibilityUnignoredChildrenForOnlyChild(originalChild unsafe.Pointer) unsafe.Pointer {
	return _NSAccessibilityUnignoredChildrenForOnlyChild(originalChild)
	}


// Called by the main function to create and run the application.

// Called by the main function to create and run the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplicationMain

func NSApplicationMain(argc int, argv unsafe.Pointer) int {
	return _NSApplicationMain(argc, argv)
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


// Creates and runs an alert sheet.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.

// Creates and runs an alert sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBeginAlertSheet

func NSBeginAlertSheet(title unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer, modalDelegate unsafe.Pointer, didEndSelector unsafe.Pointer, didDismissSelector unsafe.Pointer, contextInfo unsafe.Pointer, msgFormat unsafe.Pointer) {
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

func NSBeginCriticalAlertSheet(title unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer, modalDelegate unsafe.Pointer, didEndSelector unsafe.Pointer, didDismissSelector unsafe.Pointer, contextInfo unsafe.Pointer, msgFormat unsafe.Pointer) {
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

func NSBeginInformationalAlertSheet(title unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer, modalDelegate unsafe.Pointer, didEndSelector unsafe.Pointer, didDismissSelector unsafe.Pointer, contextInfo unsafe.Pointer, msgFormat unsafe.Pointer) {
	_NSBeginInformationalAlertSheet(title, defaultButton, alternateButton, otherButton, docWindow, modalDelegate, didEndSelector, didDismissSelector, contextInfo, msgFormat)
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


// Prepares a set of glyphs for processing by character-based routines.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.0.

// Prepares a set of glyphs for processing by character-based routines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSConvertGlyphsToPackedGlyphs(_:_:_:_:)

func NSConvertGlyphsToPackedGlyphs(glBuf unsafe.Pointer, count unsafe.Pointer, packing unsafe.Pointer, packedGlyphs unsafe.Pointer) unsafe.Pointer {
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

func NSCopyBits(srcGState unsafe.Pointer, srcRect coregraphics.CGRect, destPoint coregraphics.CGPoint) {
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

func NSCountWindowsForContext(context unsafe.Pointer, count unsafe.Pointer) {
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

// Draws a bordered rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDottedFrameRect(_:)

func NSDottedFrameRect(rect coregraphics.CGRect) {
	_NSDottedFrameRect(rect)
	}


// Draws a gray-filled rectangle representing a user-interface button.

// Draws a gray-filled rectangle representing a user-interface button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawButton(_:_:)

func NSDrawButton(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawButton(rect, clipRect)
	}


// Draws a single-color, bordered rectangle.

// Draws a single-color, bordered rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawColorTiledRects(_:_:_:_:_:)

func NSDrawColorTiledRects(boundsRect coregraphics.CGRect, clipRect coregraphics.CGRect, sides unsafe.Pointer, colors unsafe.Pointer, count unsafe.Pointer) coregraphics.CGRect {
	return _NSDrawColorTiledRects(boundsRect, clipRect, sides, colors, count)
	}


// Draws a dark gray-filled rectangle with a bezel border.

// Draws a dark gray-filled rectangle with a bezel border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawDarkBezel(_:_:)

func NSDrawDarkBezel(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawDarkBezel(rect, clipRect)
	}


// Draws a gray-filled rectangle with a bezel border.

// Draws a gray-filled rectangle with a bezel border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawGrayBezel(_:_:)

func NSDrawGrayBezel(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawGrayBezel(rect, clipRect)
	}


// Draws a gray-filled rectangle with a groove border.

// Draws a gray-filled rectangle with a groove border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawGroove(_:_:)

func NSDrawGroove(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawGroove(rect, clipRect)
	}


// Draws a white-filled rectangle with a bezel border.

// Draws a white-filled rectangle with a bezel border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawLightBezel(_:_:)

func NSDrawLightBezel(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawLightBezel(rect, clipRect)
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


// Draws rectangles with borders.

// Draws rectangles with borders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawTiledRects(_:_:_:_:_:)

func NSDrawTiledRects(boundsRect coregraphics.CGRect, clipRect coregraphics.CGRect, sides unsafe.Pointer, grays unsafe.Pointer, count unsafe.Pointer) coregraphics.CGRect {
	return _NSDrawTiledRects(boundsRect, clipRect, sides, grays, count)
	}


// Draws a white-filled rectangle with a bezel border.

// Draws a white-filled rectangle with a bezel border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawWhiteBezel(_:_:)

func NSDrawWhiteBezel(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawWhiteBezel(rect, clipRect)
	}


// Draws the window’s default background pattern into the specified rectangle of the currently focused view.

// Draws the window’s default background pattern into the specified rectangle of the currently focused view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawWindowBackground(_:)

func NSDrawWindowBackground(rect coregraphics.CGRect) {
	_NSDrawWindowBackground(rect)
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


// Erases the specified rect by filling it with white.

// Erases the specified rect by filling it with white.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEraseRect(_:)

func NSEraseRect(rect coregraphics.CGRect) {
	_NSEraseRect(rect)
	}


// Specifies how the system draws the focus ring.

// Specifies how the system draws the focus ring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/set()

func NSSetFocusRingStyle(placement unsafe.Pointer) {
	_NSSetFocusRingStyle(placement)
	}


// Draws a bordered rectangle.

// Draws a bordered rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFrameRect

func NSFrameRect(rect coregraphics.CGRect) {
	_NSFrameRect(rect)
	}


// Draws a bordered rectangle.

// Draws a bordered rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFrameRectWithWidth

func NSFrameRectWithWidth(rect coregraphics.CGRect, frameWidth float64) {
	_NSFrameRectWithWidth(rect, frameWidth)
	}


// Draws a bordered rectangle using the specified compositing operation.

// Draws a bordered rectangle using the specified compositing operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFrameRectWithWidthUsingOperation

func NSFrameRectWithWidthUsingOperation(rect coregraphics.CGRect, frameWidth float64, op unsafe.Pointer) {
	_NSFrameRectWithWidthUsingOperation(rect, frameWidth, op)
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

func NSGetAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) unsafe.Pointer {
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

func NSGetCriticalAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) unsafe.Pointer {
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

func NSGetInformationalAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) unsafe.Pointer {
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

func NSGetWindowServerMemory(context unsafe.Pointer, virtualMemory unsafe.Pointer, windowBackingMemory unsafe.Pointer, windowDumpString unsafe.Pointer) unsafe.Pointer {
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

func NSHighlightRect(rect coregraphics.CGRect) {
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

func NSInterfaceStyleForKey(key unsafe.Pointer, responder unsafe.Pointer) unsafe.Pointer {
	return _NSInterfaceStyleForKey(key, responder)
	}


// Tests whether a given object is special marker object used for indicating the state of a selection in relation to a key.

// Tests whether a given object is special marker object used for indicating the state of a selection in relation to a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSIsControllerMarker(_:)

func NSIsControllerMarker(object unsafe.Pointer) bool {
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

func NSOpenGLGetOption(pname unsafe.Pointer, param unsafe.Pointer) {
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

func NSOpenGLSetOption(pname unsafe.Pointer, param unsafe.Pointer) {
	_NSOpenGLSetOption(pname, param)
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


// Programmatically invokes a Services menu service.

// Programmatically invokes a Services menu service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPerformService(_:_:)

func NSPerformService(itemName unsafe.Pointer, pboard unsafe.Pointer) bool {
	return _NSPerformService(itemName, pboard)
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

func NSReadPixel(passedPoint coregraphics.CGPoint) unsafe.Pointer {
	return _NSReadPixel(passedPoint)
	}


// Modifies the current clipping path by intersecting it with the passed rect.

// Modifies the current clipping path by intersecting it with the passed rect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectClip

func NSRectClip(rect coregraphics.CGRect) {
	_NSRectClip(rect)
	}


// Modifies the current clipping path by intersecting it with the passed rect.

// Modifies the current clipping path by intersecting it with the passed rect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectClipList

func NSRectClipList(rects unsafe.Pointer, count unsafe.Pointer) {
	_NSRectClipList(rects, count)
	}


// Fills the passed rectangle with the current color.

// Fills the passed rectangle with the current color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFill

func NSRectFill(rect coregraphics.CGRect) {
	_NSRectFill(rect)
	}


// Fills the rectangles in the passed list with the current fill color.

// Fills the rectangles in the passed list with the current fill color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillList

func NSRectFillList(rects unsafe.Pointer, count unsafe.Pointer) {
	_NSRectFillList(rects, count)
	}


// Fills the rectangles in a list using the current fill color and specified compositing operation.

// Fills the rectangles in a list using the current fill color and specified compositing operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillListUsingOperation

func NSRectFillListUsingOperation(rects unsafe.Pointer, count unsafe.Pointer, op unsafe.Pointer) {
	_NSRectFillListUsingOperation(rects, count, op)
	}


// Fills the rectangles in the passed list with the passed list of colors.

// Fills the rectangles in the passed list with the passed list of colors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillListWithColors

func NSRectFillListWithColors(rects unsafe.Pointer, colors unsafe.Pointer, num unsafe.Pointer) {
	_NSRectFillListWithColors(rects, colors, num)
	}


// Fills the rectangles in a list using the specified colors and compositing operation.

// Fills the rectangles in a list using the specified colors and compositing operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillListWithColorsUsingOperation

func NSRectFillListWithColorsUsingOperation(rects unsafe.Pointer, colors unsafe.Pointer, num unsafe.Pointer, op unsafe.Pointer) {
	_NSRectFillListWithColorsUsingOperation(rects, colors, num, op)
	}


// Fills the rectangles in the passed list with the passed list of grays.

// Fills the rectangles in the passed list with the passed list of grays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillListWithGrays

func NSRectFillListWithGrays(rects unsafe.Pointer, grays unsafe.Pointer, num unsafe.Pointer) {
	_NSRectFillListWithGrays(rects, grays, num)
	}


// Fills a rectangle using the current fill color and the specified compositing operation.

// Fills a rectangle using the current fill color and the specified compositing operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillUsingOperation

func NSRectFillUsingOperation(rect coregraphics.CGRect, op unsafe.Pointer) {
	_NSRectFillUsingOperation(rect, op)
	}


// Registers a service provider.

// Registers a service provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRegisterServicesProvider(_:_:)

func NSRegisterServicesProvider(provider unsafe.Pointer, name unsafe.Pointer) {
	_NSRegisterServicesProvider(provider, name)
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

func NSReleaseAlertPanel(panel unsafe.Pointer) {
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

func NSRunAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) unsafe.Pointer {
	return _NSRunAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
	}


// NSRunAlertPanelRelativeToWindow is a AppKit function.
//
// Deprecated: This function was deprecated in macOS 10.0.
//
// Added in macOS 10.0.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunAlertPanelRelativeToWindow

func NSRunAlertPanelRelativeToWindow(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer) unsafe.Pointer {
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

func NSRunCriticalAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) unsafe.Pointer {
	return _NSRunCriticalAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
	}


// NSRunCriticalAlertPanelRelativeToWindow is a AppKit function.
//
// Deprecated: This function was deprecated in macOS 10.0.
//
// Added in macOS 10.0.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunCriticalAlertPanelRelativeToWindow

func NSRunCriticalAlertPanelRelativeToWindow(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer) unsafe.Pointer {
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

func NSRunInformationalAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) unsafe.Pointer {
	return _NSRunInformationalAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
	}


// NSRunInformationalAlertPanelRelativeToWindow is a AppKit function.
//
// Deprecated: This function was deprecated in macOS 10.0.
//
// Added in macOS 10.0.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunInformationalAlertPanelRelativeToWindow

func NSRunInformationalAlertPanelRelativeToWindow(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer) unsafe.Pointer {
	return _NSRunInformationalAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow)
	}


// Specifies whether an item should be included in Services menus.

// Specifies whether an item should be included in Services menus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSetShowsServicesMenuItem(_:_:)

func NSSetShowsServicesMenuItem(itemName unsafe.Pointer, enabled bool) unsafe.Pointer {
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

func NSShowAnimationEffect(animationEffect unsafe.Pointer, centerLocation coregraphics.CGPoint, size coregraphics.CGSize, animationDelegate unsafe.Pointer, didEndSelector unsafe.Pointer, contextInfo unsafe.Pointer) {
	_NSShowAnimationEffect(animationEffect, centerLocation, size, animationDelegate, didEndSelector, contextInfo)
	}


// Specifies whether a Services menu item is currently enabled.

// Specifies whether a Services menu item is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShowsServicesMenuItem(_:)

func NSShowsServicesMenuItem(itemName unsafe.Pointer) bool {
	return _NSShowsServicesMenuItem(itemName)
	}


// Unregisters a service provider.

// Unregisters a service provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnregisterServicesProvider(_:)

func NSUnregisterServicesProvider(name unsafe.Pointer) {
	_NSUnregisterServicesProvider(name)
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


// Gets information about an application’s onscreen windows.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.

// Gets information about an application’s onscreen windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowListForContext

func NSWindowListForContext(context unsafe.Pointer, size unsafe.Pointer, list unsafe.Pointer) {
	_NSWindowListForContext(context, size, list)
	}




