// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PAccessibilityRadioButton is the NSAccessibilityRadioButton protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a radio button.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityRadioButton
type PAccessibilityRadioButton interface {
	// Required methods
	AccessibilityValue() foundation.Number/* debug [protocol_interface/required_method]: AccessibilityValue */
}
