// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"
)

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
	AccessibilityFrameForRange(range_ foundation.Range) corefoundation.CGRect
	AccessibilityLineForIndex(index int) int
	AccessibilityRangeForLine(lineNumber int) foundation.Range
	AccessibilityStringForRange(range_ foundation.Range) foundation.String
}
