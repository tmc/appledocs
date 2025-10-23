package foundation

import "github.com/tmc/appledocs/generated/corefoundation"

// Geometry type aliases for Foundation functions that use CoreFoundation types.
// These allow Foundation functions like NSPointFromString to work correctly
// by aliasing to the actual CoreFoundation definitions.
//
// Note: CGPoint, CGSize, and CGRect are defined in CoreFoundation, not CoreGraphics.
// See: https://developer.apple.com/documentation/corefoundation/cgpoint

// Point is a type alias for CoreFoundation CGPoint (NSPoint in Foundation).
type Point = corefoundation.CGPoint

// Size is a type alias for CoreFoundation CGSize (NSSize in Foundation).
type Size = corefoundation.CGSize

// Rect is a type alias for CoreFoundation CGRect (NSRect in Foundation).
type Rect = corefoundation.CGRect
