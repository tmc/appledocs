// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// AppKit Functions
//
// This file contains function declarations discovered from Apple's documentation.
// To use these functions, you need to:
//   1. Map C types to Go types
//   2. Create function variables
//   3. Register them with purego.RegisterLibFunc
//
// Example:
//   var CGContextSetRGBFillColor func(c CGContextRef, red, green, blue, alpha CGFloat)
//   purego.RegisterLibFunc(&CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")

// Discovered functions (39 total):

// NSAccessibilityRaiseBadArgumentException(element id, attribute ,  NSAccessibilityAttributeName, value ,  id, )
//
// Availability:
//   - macOS 10.1+ (Deprecated in 10.11)
//
// Deprecated: This function is deprecated.

// NSAccessibilityPointInView(parentView NSView  *, point ,  NSPoint, ) NSPoint
//
// Availability:
//   - macOS 10.10+

// NSAccessibilityFrameInView(parentView NSView  *, frame ,  NSRect, ) NSRect
//
// Availability:
//   - macOS 10.10+


// NSApplicationMain(argc int, argv ,  const char  *, []) int

// NSBeginAlertSheet(title NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, docWindow ,  NSWindow  *, modalDelegate ,  id, didEndSelector ,  SEL, didDismissSelector ,  SEL, contextInfo ,  void  *, msgFormat ,  NSString  *, , ...)
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSBeginCriticalAlertSheet(title NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, docWindow ,  NSWindow  *, modalDelegate ,  id, didEndSelector ,  SEL, didDismissSelector ,  SEL, contextInfo ,  void  *, msgFormat ,  NSString  *, , ...)
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.


// NSBeginInformationalAlertSheet(title NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, docWindow ,  NSWindow  *, modalDelegate ,  id, didEndSelector ,  SEL, didDismissSelector ,  SEL, contextInfo ,  void  *, msgFormat ,  NSString  *, , ...)
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSConvertGlyphsToPackedGlyphs(glBuf NSGlyph  *, count ,  NSInteger, packing ,  NSMultibyteGlyphPacking, packedGlyphs ,  char  *, ) NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.13)
//
// Deprecated: This function is deprecated.

// NSCopyBits(srcGState NSInteger, srcRect ,  NSRect, destPoint ,  NSPoint, )
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.


// NSCountWindows(count NSInteger  *, )
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// NSCountWindowsForContext(context NSInteger, count ,  NSInteger  *, )
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// NSDottedFrameRect(rect NSRect, )


// NSDrawColorTiledRects(boundsRect NSRect, clipRect ,  NSRect, sides ,  const NSRectEdge  *, colors ,  NSColor  * *, count ,  NSInteger, ) NSRect

// NSGetAlertPanel(title NSString  *, msgFormat ,  NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, , ...) id
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSGetCriticalAlertPanel(title NSString  *, msgFormat ,  NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, , ...) id
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.


// NSGetInformationalAlertPanel(title NSString  *, msgFormat ,  NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, , ...) id
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSGetWindowServerMemory(context NSInteger, virtualMemory ,  NSInteger  *, windowBackingMemory ,  NSInteger  *, windowDumpString ,  NSString  * *, ) NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// NSInterfaceStyleForKey(key NSString  *, responder ,  NSResponder  *, ) NSInterfaceStyle
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// NSIsControllerMarker(object id, ) BOOL

// NSOpenGLGetOption(pname NSOpenGLGlobalOption, param ,  GLint  *, )
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// NSOpenGLGetVersion(major GLint  *, minor ,  GLint  *, )
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.


// NSOpenGLSetOption(pname NSOpenGLGlobalOption, param ,  GLint, )
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// NSCreateFileContentsPboardType(fileType NSString  *, ) NSPasteboardType

// NSCreateFilenamePboardType(fileType NSString  *, ) NSPasteboardType


// NSGetFileType(pboardType NSPasteboardType, ) NSString  *

// NSGetFileTypes(pboardTypes NSArray<NSString *>  *, ) NSArray<NSString *>  *

// NSReadPixel(passedPoint NSPoint, ) NSColor  *
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.


// NSReleaseAlertPanel(panel id, )
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSRunAlertPanel(title NSString  *, msgFormat ,  NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, , ...) NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSRunAlertPanelRelativeToWindow(title NSString  *, msgFormat ,  NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, docWindow ,  NSWindow  *, , ...) NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.0)
//
// Deprecated: This function is deprecated.


// NSRunCriticalAlertPanel(title NSString  *, msgFormat ,  NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, , ...) NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSRunCriticalAlertPanelRelativeToWindow(title NSString  *, msgFormat ,  NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, docWindow ,  NSWindow  *, , ...) NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.0)
//
// Deprecated: This function is deprecated.

// NSRunInformationalAlertPanel(title NSString  *, msgFormat ,  NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, , ...) NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.


// NSRunInformationalAlertPanelRelativeToWindow(title NSString  *, msgFormat ,  NSString  *, defaultButton ,  NSString  *, alternateButton ,  NSString  *, otherButton ,  NSString  *, docWindow ,  NSWindow  *, , ...) NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.0)
//
// Deprecated: This function is deprecated.

// NSSetShowsServicesMenuItem(itemName NSString  *, enabled ,  BOOL, ) NSInteger

// NSShowAnimationEffect(animationEffect NSAnimationEffect, centerLocation ,  NSPoint, size ,  NSSize, animationDelegate ,  id, didEndSelector ,  SEL, contextInfo ,  void  *, )
//
// Availability:
//   - macOS 10.3+ (Deprecated in 14.0)
//
// Deprecated: This function is deprecated.


// NSShowsServicesMenuItem(itemName NSString  *, ) BOOL

// NSWindowList(size NSInteger, list ,  NSInteger, [])
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// NSWindowListForContext(context NSInteger, size ,  NSInteger, list ,  NSInteger, [])
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.


