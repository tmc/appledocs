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

// Discovered functions (43 total):

// raiseBadArgumentException(element _, attribute :  Any!,  _, value :  NSAccessibility. Attribute!,  _, :  Any!)) static   func
//
// Availability:
//   - macOS 10.1+ (Deprecated in 10.11)
//
// Deprecated: This function is deprecated.

// screenPoint(parentView fromView, :  NSView,  point:  NSPoint) ->  NSPoint) static   func
//
// Availability:
//   - macOS 10.10+

// screenRect(parentView fromView, frame :  NSView,  rect, :  NSRect) ->  NSRect) static   func
//
// Availability:
//   - macOS 10.10+


// NSApplicationMain(argc int, argv ,  const  char *, []);) extern   int

// NSBeep() extern   void

// NSBeginAlertSheet(title NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, modalDelegate ,  id, didEndSelector ,  SEL, didDismissSelector ,  SEL, contextInfo ,  void *, msgFormat ,  NSString *, , ...);) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.


// NSBeginCriticalAlertSheet(title NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, modalDelegate ,  id, didEndSelector ,  SEL, didDismissSelector ,  SEL, contextInfo ,  void *, msgFormat ,  NSString *, , ...);) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSBeginInformationalAlertSheet(title NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, modalDelegate ,  id, didEndSelector ,  SEL, didDismissSelector ,  SEL, contextInfo ,  void *, msgFormat ,  NSString *, , ...);) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSConvertGlyphsToPackedGlyphs(glBuf _, count :  UnsafeMutablePointer< NSGlyph>,  _, packing :  Int,  _, packedGlyphs :  NSMultibyteGlyphPacking,  _, :  UnsafeMutablePointer< CChar>) ->  Int) func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.13)
//
// Deprecated: This function is deprecated.


// NSCopyBits(srcGState _, srcRect :  Int,  _, destPoint :  NSRect,  _, :  NSPoint) func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSCountWindows(count NSInteger *, );) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// NSCountWindowsForContext(context NSInteger, count ,  NSInteger *, );) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.


// NSDisableScreenUpdates() func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.11)
//
// Deprecated: This function is deprecated.

// NSDottedFrameRect(rect _, :  NSRect) func

// NSDrawColorTiledRects(boundsRect _, clipRect :  NSRect,  _, sides :  NSRect,  _, colors :  UnsafePointer< NSRectEdge>,  _, count :  AutoreleasingUnsafeMutablePointer< NSColor>,  _, :  Int) ->  NSRect) func


// NSEnableScreenUpdates() func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.11)
//
// Deprecated: This function is deprecated.

// NSGetAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   id
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.


// NSGetCriticalAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   id
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSGetInformationalAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   id
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSGetWindowServerMemory(context _, virtualMemory :  Int,  _, windowBackingMemory :  UnsafeMutablePointer< Int>,  _, windowDumpString :  UnsafeMutablePointer< Int>,  _, :  AutoreleasingUnsafeMutablePointer< NSString>) ->  Int) func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.


// NSInterfaceStyleForKey(key NSString *, responder ,  NSResponder *, );) extern   NSInterfaceStyle
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// NSIsControllerMarker(object _, :  Any?) ->  Bool) func

// NSOpenGLGetOption(pname NSOpenGLGlobalOption, param ,  GLint *, );) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.


// NSOpenGLGetVersion(major GLint *, minor ,  GLint *, );) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// NSOpenGLSetOption(pname NSOpenGLGlobalOption, param ,  GLint, );) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// fileContentsType(fileType forPathExtension, :  String) ->  NSPasteboard. PasteboardType!) static   func


// fileNameType(fileType forPathExtension, :  String) ->  NSPasteboard. PasteboardType!) static   func

// representedPathExtension() var

// representedPathExtensions(pboardTypes from, : [ NSPasteboard. PasteboardType]) -> [ String]?) static   func


// NSReadPixel(passedPoint _, :  NSPoint) ->  NSColor?) func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.14)
//
// Deprecated: This function is deprecated.

// NSReleaseAlertPanel(panel _, :  Any!)) func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSRunAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.


// NSRunAlertPanelRelativeToWindow(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, , ...);) extern   NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.0)
//
// Deprecated: This function is deprecated.

// NSRunCriticalAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSRunCriticalAlertPanelRelativeToWindow(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, , ...);) extern   NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.0)
//
// Deprecated: This function is deprecated.


// NSRunInformationalAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.10)
//
// Deprecated: This function is deprecated.

// NSRunInformationalAlertPanelRelativeToWindow(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, , ...);) extern   NSInteger
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.0)
//
// Deprecated: This function is deprecated.

// NSSetShowsServicesMenuItem(itemName _, enabled :  String,  _, :  Bool) ->  Int) func


// NSShowAnimationEffect(animationEffect NSAnimationEffect, centerLocation ,  NSPoint, size ,  NSSize, animationDelegate ,  id, didEndSelector ,  SEL, contextInfo ,  void *, );) extern   void
//
// Availability:
//   - macOS 10.3+ (Deprecated in 14.0)
//
// Deprecated: This function is deprecated.

// NSShowsServicesMenuItem(itemName _, :  String) ->  Bool) func

// NSWindowList(size NSInteger, list ,  NSInteger, []);) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.


// NSWindowListForContext(context NSInteger, size ,  NSInteger, list ,  NSInteger, []);) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

