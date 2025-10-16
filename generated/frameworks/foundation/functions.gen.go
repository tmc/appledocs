// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Foundation Functions
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

// Discovered functions (22 total):

// NSClassFromString(aClassName _, :  String) ->  AnyClass?) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// cgAffineTransform(string for, :  String) ->  CGAffineTransform) class   func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// cgPoint(string for, :  String) ->  CGPoint) class   func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// cgRect(string for, :  String) ->  CGRect) class   func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// cgSize(string for, :  String) ->  CGSize) class   func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// cgVector(string for, :  String) ->  CGVector) class   func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// string(size for, :  CGSize) ->  String) class   func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// string(vector for, :  CGVector) ->  String) class   func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// string(rect for, :  CGRect) ->  String) class   func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// string(point for, :  CGPoint) ->  String) class   func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// string(transform for, :  CGAffineTransform) ->  String) class   func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// NSCountFrames() extern   NSUInteger
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// NSFrameAddress(frame NSUInteger, );) extern   void  *
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// NSGetUncaughtExceptionHandler() func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// NSIsFreedObject(anObject id, );) extern   BOOL
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// NSLog(format NSString *, , ...);) extern   void
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// NSLogv(format _, args :  String,  _, :  CVaListPointer) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// NSRecordAllocationEvent(eventType int, object ,  id, );) extern   void
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// NSReturnAddress(frame NSUInteger, );) extern   void  *
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// NSSetUncaughtExceptionHandler(_: (( NSException) ->  Void)?)) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// NSTemporaryDirectory() func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// NXReadNSObjectFromCoder(decoder NSCoder *, );) extern   NSObject  *
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.5)
//
// Deprecated: This function is deprecated.
