// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration

// DiskArbitration Functions
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

// Discovered functions (5 total):

// DAApprovalSessionScheduleWithRunLoop(session DAApprovalSessionRef, runLoop ,  CFRunLoopRef, runLoopMode ,  CFStringRef, );) extern   void
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+

// DADiskCreateFromBSDName(allocator _, session :  CFAllocator?,  _, name :  DASession,  _, :  UnsafePointer< CChar>) ->  DADisk?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+

// DADiskGetBSDName(disk _, :  DADisk) ->  UnsafePointer< CChar>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+

// DADissenterGetStatus(dissenter _, :  DADissenter) ->  DAReturn) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+

// DASessionCreate(allocator _, :  CFAllocator?) ->  DASession?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.4+
