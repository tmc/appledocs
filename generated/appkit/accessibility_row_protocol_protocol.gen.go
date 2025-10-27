// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PAccessibilityRow is the NSAccessibilityRow protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a row for a table, list, or outline view.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityRow
type PAccessibilityRow interface {
	// Required methods
	AccessibilityIndex() int
	// Optional methods
	AccessibilityDisclosureLevel() int
	HasAccessibilityDisclosureLevel() bool
}
