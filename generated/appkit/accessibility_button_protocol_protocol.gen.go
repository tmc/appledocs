// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PAccessibilityButton is the NSAccessibilityButton protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a button.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityButton
type PAccessibilityButton interface {
	// Required methods
	AccessibilityLabel() foundation.String/* debug [protocol_interface/required_method]: AccessibilityLabel */
	AccessibilityPerformPress() bool/* debug [protocol_interface/required_method]: AccessibilityPerformPress */
}
