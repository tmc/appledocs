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

// Discovered functions (21 total):

// cgAffineTransform(string for, :  String) ->  CGAffineTransform) class   func
// cgPoint(string for, :  String) ->  CGPoint) class   func
// cgRect(string for, :  String) ->  CGRect) class   func
// cgSize(string for, :  String) ->  CGSize) class   func
// cgVector(string for, :  String) ->  CGVector) class   func

// string(size for, :  CGSize) ->  String) class   func
// string(vector for, :  CGVector) ->  String) class   func
// string(rect for, :  CGRect) ->  String) class   func
// string(point for, :  CGPoint) ->  String) class   func
// string(transform for, :  CGAffineTransform) ->  String) class   func

// NSCountFrames() extern   NSUInteger
// NSFrameAddress(frame NSUInteger, );) extern   void  *
// NSGetUncaughtExceptionHandler() func
// NSIsFreedObject(anObject id, );) extern   BOOL
// NSLog(format NSString *, , ...);) extern   void

// NSLogv(format _, args :  String,  _, :  CVaListPointer) func
// NSRecordAllocationEvent(eventType int, object ,  id, );) extern   void
// NSReturnAddress(frame NSUInteger, );) extern   void  *
// NSSetUncaughtExceptionHandler(_: (( NSException) ->  Void)?)) func
// NSTemporaryDirectory() func

// NXReadNSObjectFromCoder(decoder NSCoder *, );) extern   NSObject  *
