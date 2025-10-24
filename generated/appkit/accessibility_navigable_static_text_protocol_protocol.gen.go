// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PAccessibilityNavigableStaticText is the NSAccessibilityNavigableStaticText protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as navigable static text.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityNavigableStaticText
type PAccessibilityNavigableStaticText interface {
	// Required methods
	AccessibilityFrameForRange(range_ corefoundation.Range) corefoundation.Rect
	AccessibilityLineForIndex(index int) int
	AccessibilityRangeForLine(lineNumber int) corefoundation.Range
	AccessibilityStringForRange(range_ corefoundation.Range) foundation.String
}
