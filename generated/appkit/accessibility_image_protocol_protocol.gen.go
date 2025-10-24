// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PAccessibilityImage is the NSAccessibilityImage protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as an image.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityImage
type PAccessibilityImage interface {
	// Required methods
	AccessibilityLabel() foundation.String
}
