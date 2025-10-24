// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PAccessibilitySwitch is the NSAccessibilitySwitch protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a switch.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilitySwitch
type PAccessibilitySwitch interface {
	// Required methods
	AccessibilityValue() foundation.String/* debug [protocol_interface/required_method]: AccessibilityValue */
	// Optional methods
	AccessibilityPerformDecrement() bool
	HasAccessibilityPerformDecrement() bool
	AccessibilityPerformIncrement() bool
	HasAccessibilityPerformIncrement() bool
}
