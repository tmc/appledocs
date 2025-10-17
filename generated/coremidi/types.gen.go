// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import "unsafe"

// CFDataRef is a CoreGraphics opaque type.
type CFDataRef unsafe.Pointer

// CFRunLoopRef is a CoreGraphics opaque type.
type CFRunLoopRef unsafe.Pointer

// CFStringRef is a CoreGraphics opaque type.
type CFStringRef unsafe.Pointer

// MIDIClientRef is a CoreGraphics opaque type.
type MIDIClientRef unsafe.Pointer

// MIDIDeviceListRef is a CoreGraphics opaque type.
type MIDIDeviceListRef unsafe.Pointer

// MIDIDeviceRef is a CoreGraphics opaque type.
type MIDIDeviceRef unsafe.Pointer

// MIDIDriverRef is a CoreGraphics opaque type.
type MIDIDriverRef unsafe.Pointer

// MIDIEndpointRef is a CoreGraphics opaque type.
type MIDIEndpointRef unsafe.Pointer

// MIDIEntityRef is a CoreGraphics opaque type.
type MIDIEntityRef unsafe.Pointer

// MIDIPortRef is a CoreGraphics opaque type.
type MIDIPortRef unsafe.Pointer

// MIDISetupRef is a CoreGraphics opaque type.
type MIDISetupRef unsafe.Pointer

// MIDIThruConnectionRef is a CoreGraphics opaque type.
type MIDIThruConnectionRef unsafe.Pointer


// Common CoreGraphics struct types
type CGFloat = float64

type CGPoint struct {
	X CGFloat
	Y CGFloat
}

type CGSize struct {
	Width  CGFloat
	Height CGFloat
}

type CGRect struct {
	Origin CGPoint
	Size   CGSize
}

type CGAffineTransform struct {
	A  CGFloat
	B  CGFloat
	C  CGFloat
	D  CGFloat
	Tx CGFloat
	Ty CGFloat
}

// Common type aliases
type Range = CGPoint  // NSRange
type Size = CGSize    // NSSize
type Point = CGPoint  // NSPoint
type Rect = CGRect    // NSRect
type TimeInterval = float64  // NSTimeInterval


