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
	AccessibilityFrameForRange(range_ corefoundation.Range) Rect/* debug [protocol_interface/required_method]: AccessibilityFrameForRange */
	AccessibilityLineForIndex(index int) int/* debug [protocol_interface/required_method]: AccessibilityLineForIndex */
	AccessibilityRangeForLine(lineNumber int) corefoundation.Range/* debug [protocol_interface/required_method]: AccessibilityRangeForLine */
	AccessibilityStringForRange(range_ corefoundation.Range) foundation.String/* debug [protocol_interface/required_method]: AccessibilityStringForRange */
}
