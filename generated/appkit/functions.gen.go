// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
	objc "github.com/ebitengine/purego/objc"
	corefoundation "github.com/tmc/appledocs/generated/corefoundation"
)


// AppKit Functions (38 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSAccessibilityRaiseBadArgumentException func(objc.ID, unsafe.Pointer, objc.ID)
	_NSBeep func()
	_NSBeginAlertSheet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, objc.ID, objc.SEL, objc.SEL, unsafe.Pointer, unsafe.Pointer)
	_NSBeginCriticalAlertSheet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, objc.ID, objc.SEL, objc.SEL, unsafe.Pointer, unsafe.Pointer)
	_NSBeginInformationalAlertSheet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, objc.ID, objc.SEL, objc.SEL, unsafe.Pointer, unsafe.Pointer)
	_NSConvertGlyphsToPackedGlyphs func(unsafe.Pointer, int64, unsafe.Pointer, unsafe.Pointer) int64
	_NSCopyBits func(int64, corefoundation.Rect, corefoundation.Point)
	_NSCountWindows func(unsafe.Pointer)
	_NSCountWindowsForContext func(int64, unsafe.Pointer)
	_NSDisableScreenUpdates func()
	_NSDottedFrameRect func(corefoundation.Rect)
	_NSDrawColorTiledRects func(corefoundation.Rect, corefoundation.Rect, unsafe.Pointer, unsafe.Pointer, int64) corefoundation.Rect
	_NSEnableScreenUpdates func()
	_NSGetAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) objc.ID
	_NSGetCriticalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) objc.ID
	_NSGetInformationalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) objc.ID
	_NSGetWindowServerMemory func(int64, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int64
	_NSInterfaceStyleForKey func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSOpenGLGetOption func(unsafe.Pointer, unsafe.Pointer)
	_NSOpenGLGetVersion func(unsafe.Pointer, unsafe.Pointer)
	_NSOpenGLSetOption func(unsafe.Pointer, unsafe.Pointer)
	_NSCreateFileContentsPboardType func(unsafe.Pointer) unsafe.Pointer
	_NSCreateFilenamePboardType func(unsafe.Pointer) unsafe.Pointer
	_NSGetFileType func(unsafe.Pointer) unsafe.Pointer
	_NSGetFileTypes func([]unsafe.Pointer) []unsafe.Pointer
	_NSReadPixel func(corefoundation.Point) unsafe.Pointer
	_NSReleaseAlertPanel func(objc.ID)
	_NSRunAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int64
	_NSRunAlertPanelRelativeToWindow func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int64
	_NSRunCriticalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int64
	_NSRunCriticalAlertPanelRelativeToWindow func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int64
	_NSRunInformationalAlertPanel func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int64
	_NSRunInformationalAlertPanelRelativeToWindow func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int64
	_NSSetShowsServicesMenuItem func(unsafe.Pointer, bool) int64
	_NSShowAnimationEffect func(unsafe.Pointer, corefoundation.Point, corefoundation.Size, objc.ID, objc.SEL, unsafe.Pointer)
	_NSShowsServicesMenuItem func(unsafe.Pointer) bool
	_NSWindowList func(int64, int64)
	_NSWindowListForContext func(int64, int64, int64)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSAccessibilityRaiseBadArgumentException, lib, "NSAccessibilityRaiseBadArgumentException")
	tryRegister(&_NSBeep, lib, "NSBeep")
	tryRegister(&_NSBeginAlertSheet, lib, "NSBeginAlertSheet")
	tryRegister(&_NSBeginCriticalAlertSheet, lib, "NSBeginCriticalAlertSheet")
	tryRegister(&_NSBeginInformationalAlertSheet, lib, "NSBeginInformationalAlertSheet")
	tryRegister(&_NSConvertGlyphsToPackedGlyphs, lib, "NSConvertGlyphsToPackedGlyphs")
	tryRegister(&_NSCopyBits, lib, "NSCopyBits")
	tryRegister(&_NSCountWindows, lib, "NSCountWindows")
	tryRegister(&_NSCountWindowsForContext, lib, "NSCountWindowsForContext")
	tryRegister(&_NSDisableScreenUpdates, lib, "NSDisableScreenUpdates")
	tryRegister(&_NSDottedFrameRect, lib, "NSDottedFrameRect")
	tryRegister(&_NSDrawColorTiledRects, lib, "NSDrawColorTiledRects")
	tryRegister(&_NSEnableScreenUpdates, lib, "NSEnableScreenUpdates")
	tryRegister(&_NSGetAlertPanel, lib, "NSGetAlertPanel")
	tryRegister(&_NSGetCriticalAlertPanel, lib, "NSGetCriticalAlertPanel")
	tryRegister(&_NSGetInformationalAlertPanel, lib, "NSGetInformationalAlertPanel")
	tryRegister(&_NSGetWindowServerMemory, lib, "NSGetWindowServerMemory")
	tryRegister(&_NSInterfaceStyleForKey, lib, "NSInterfaceStyleForKey")
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



// Raises an error if the parameter is the wrong type or has an illegal value
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.1.
// Raises an error if the parameter is the wrong type or has an illegal value
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/raiseBadArgumentException(_:_:_:)
func NSAccessibilityRaiseBadArgumentException(element objc.ID, attribute unsafe.Pointer, value objc.ID) {
	_NSAccessibilityRaiseBadArgumentException(element, attribute, value)
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

// Prepares a set of glyphs for processing by character-based routines.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.0.
// Prepares a set of glyphs for processing by character-based routines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSConvertGlyphsToPackedGlyphs(_:_:_:_:)
func NSConvertGlyphsToPackedGlyphs(glBuf unsafe.Pointer, count int64, packing unsafe.Pointer, packedGlyphs unsafe.Pointer) int64 {
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
func NSCopyBits(srcGState int64, srcRect corefoundation.Rect, destPoint corefoundation.Point) {
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
func NSCountWindowsForContext(context int64, count unsafe.Pointer) {
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
func NSDottedFrameRect(rect corefoundation.Rect) {
	_NSDottedFrameRect(rect)
}

// Draws a single-color, bordered rectangle.

// Draws a single-color, bordered rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawColorTiledRects(_:_:_:_:_:)
func NSDrawColorTiledRects(boundsRect corefoundation.Rect, clipRect corefoundation.Rect, sides unsafe.Pointer, colors unsafe.Pointer, count int64) corefoundation.Rect {
	return _NSDrawColorTiledRects(boundsRect, clipRect, sides, colors, count)
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
func NSGetWindowServerMemory(context int64, virtualMemory unsafe.Pointer, windowBackingMemory unsafe.Pointer, windowDumpString unsafe.Pointer) int64 {
	return _NSGetWindowServerMemory(context, virtualMemory, windowBackingMemory, windowDumpString)
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
func NSReadPixel(passedPoint corefoundation.Point) unsafe.Pointer {
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
func NSRunAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) int64 {
	return _NSRunAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
}

// NSRunAlertPanelRelativeToWindow is a AppKit function.
//
// Deprecated: This function was deprecated in macOS 10.0.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunAlertPanelRelativeToWindow
func NSRunAlertPanelRelativeToWindow(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer) int64 {
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
func NSRunCriticalAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) int64 {
	return _NSRunCriticalAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
}

// NSRunCriticalAlertPanelRelativeToWindow is a AppKit function.
//
// Deprecated: This function was deprecated in macOS 10.0.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunCriticalAlertPanelRelativeToWindow
func NSRunCriticalAlertPanelRelativeToWindow(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer) int64 {
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
func NSRunInformationalAlertPanel(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer) int64 {
	return _NSRunInformationalAlertPanel(title, msgFormat, defaultButton, alternateButton, otherButton)
}

// NSRunInformationalAlertPanelRelativeToWindow is a AppKit function.
//
// Deprecated: This function was deprecated in macOS 10.0.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRunInformationalAlertPanelRelativeToWindow
func NSRunInformationalAlertPanelRelativeToWindow(title unsafe.Pointer, msgFormat unsafe.Pointer, defaultButton unsafe.Pointer, alternateButton unsafe.Pointer, otherButton unsafe.Pointer, docWindow unsafe.Pointer) int64 {
	return _NSRunInformationalAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow)
}

// Specifies whether an item should be included in Services menus.

// Specifies whether an item should be included in Services menus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSetShowsServicesMenuItem(_:_:)
func NSSetShowsServicesMenuItem(itemName unsafe.Pointer, enabled bool) int64 {
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
func NSShowAnimationEffect(animationEffect unsafe.Pointer, centerLocation corefoundation.Point, size corefoundation.Size, animationDelegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer) {
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

// Gets information about onscreen windows.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.0.
// Gets information about onscreen windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowList
func NSWindowList(size int64, list int64) {
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
func NSWindowListForContext(context int64, size int64, list int64) {
	_NSWindowListForContext(context, size, list)
}



