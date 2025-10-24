// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PAccessibilityLayoutItem is the NSAccessibilityLayoutItem protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a layout item.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityLayoutItem
type PAccessibilityLayoutItem interface {
	// Optional methods
	SetAccessibilityFrame(frame Rect /* not a class type */)
	HasSetAccessibilityFrame() bool
}
