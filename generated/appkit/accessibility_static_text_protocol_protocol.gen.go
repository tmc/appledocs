// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"
)

// PAccessibilityStaticText is the NSAccessibilityStaticText protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as static text.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityStaticText
type PAccessibilityStaticText interface {
	// Required methods
	AccessibilityValue() foundation.String/* debug [protocol_interface/required_method]: AccessibilityValue */
	// Optional methods
	AccessibilityAttributedStringForRange(range_ corefoundation.Range) foundation.AttributedString
	HasAccessibilityAttributedStringForRange() bool
	AccessibilityVisibleCharacterRange() corefoundation.Range
	HasAccessibilityVisibleCharacterRange() bool
}
