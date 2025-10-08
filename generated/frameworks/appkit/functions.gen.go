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
// screenPoint(parentView fromView, :  NSView,  point:  NSPoint) ->  NSPoint) static   func
// screenRect(parentView fromView, frame :  NSView,  rect, :  NSRect) ->  NSRect) static   func
// NSApplicationMain(argc int, argv ,  const  char *, []);) extern   int
// NSBeep() extern   void

// NSBeginAlertSheet(title NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, modalDelegate ,  id, didEndSelector ,  SEL, didDismissSelector ,  SEL, contextInfo ,  void *, msgFormat ,  NSString *, , ...);) extern   void
// NSBeginCriticalAlertSheet(title NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, modalDelegate ,  id, didEndSelector ,  SEL, didDismissSelector ,  SEL, contextInfo ,  void *, msgFormat ,  NSString *, , ...);) extern   void
// NSBeginInformationalAlertSheet(title NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, modalDelegate ,  id, didEndSelector ,  SEL, didDismissSelector ,  SEL, contextInfo ,  void *, msgFormat ,  NSString *, , ...);) extern   void
// NSConvertGlyphsToPackedGlyphs(glBuf _, count :  UnsafeMutablePointer< NSGlyph>,  _, packing :  Int,  _, packedGlyphs :  NSMultibyteGlyphPacking,  _, :  UnsafeMutablePointer< CChar>) ->  Int) func
// NSCopyBits(srcGState _, srcRect :  Int,  _, destPoint :  NSRect,  _, :  NSPoint) func

// NSCountWindows(count NSInteger *, );) extern   void
// NSCountWindowsForContext(context NSInteger, count ,  NSInteger *, );) extern   void
// NSDisableScreenUpdates() func
// NSDottedFrameRect(rect _, :  NSRect) func
// NSDrawColorTiledRects(boundsRect _, clipRect :  NSRect,  _, sides :  NSRect,  _, colors :  UnsafePointer< NSRectEdge>,  _, count :  AutoreleasingUnsafeMutablePointer< NSColor>,  _, :  Int) ->  NSRect) func

// NSEnableScreenUpdates() func
// NSGetAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   id
// NSGetCriticalAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   id
// NSGetInformationalAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   id

// NSGetWindowServerMemory(context _, virtualMemory :  Int,  _, windowBackingMemory :  UnsafeMutablePointer< Int>,  _, windowDumpString :  UnsafeMutablePointer< Int>,  _, :  AutoreleasingUnsafeMutablePointer< NSString>) ->  Int) func
// NSInterfaceStyleForKey(key NSString *, responder ,  NSResponder *, );) extern   NSInterfaceStyle
// NSIsControllerMarker(object _, :  Any?) ->  Bool) func
// NSOpenGLGetOption(pname NSOpenGLGlobalOption, param ,  GLint *, );) extern   void
// NSOpenGLGetVersion(major GLint *, minor ,  GLint *, );) extern   void

// NSOpenGLSetOption(pname NSOpenGLGlobalOption, param ,  GLint, );) extern   void
// fileContentsType(fileType forPathExtension, :  String) ->  NSPasteboard. PasteboardType!) static   func
// fileNameType(fileType forPathExtension, :  String) ->  NSPasteboard. PasteboardType!) static   func
// representedPathExtension() var
// representedPathExtensions(pboardTypes from, : [ NSPasteboard. PasteboardType]) -> [ String]?) static   func

// NSReadPixel(passedPoint _, :  NSPoint) ->  NSColor?) func
// NSReleaseAlertPanel(panel _, :  Any!)) func
// NSRunAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   NSInteger
// NSRunAlertPanelRelativeToWindow(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, , ...);) extern   NSInteger
// NSRunCriticalAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   NSInteger

// NSRunCriticalAlertPanelRelativeToWindow(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, , ...);) extern   NSInteger
// NSRunInformationalAlertPanel(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, , ...);) extern   NSInteger
// NSRunInformationalAlertPanelRelativeToWindow(title NSString *, msgFormat ,  NSString *, defaultButton ,  NSString *, alternateButton ,  NSString *, otherButton ,  NSString *, docWindow ,  NSWindow *, , ...);) extern   NSInteger
// NSSetShowsServicesMenuItem(itemName _, enabled :  String,  _, :  Bool) ->  Int) func
// NSShowAnimationEffect(animationEffect NSAnimationEffect, centerLocation ,  NSPoint, size ,  NSSize, animationDelegate ,  id, didEndSelector ,  SEL, contextInfo ,  void *, );) extern   void

// NSShowsServicesMenuItem(itemName _, :  String) ->  Bool) func
// NSWindowList(size NSInteger, list ,  NSInteger, []);) extern   void
// NSWindowListForContext(context NSInteger, size ,  NSInteger, list ,  NSInteger, []);) extern   void
