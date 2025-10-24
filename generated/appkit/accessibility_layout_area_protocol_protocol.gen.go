// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PAccessibilityLayoutArea is the NSAccessibilityLayoutArea protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a layout area.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityLayoutArea
type PAccessibilityLayoutArea interface {
	// Required methods
	AccessibilityChildren() foundation.Array/* debug [protocol_interface/required_method]: AccessibilityChildren */
	AccessibilityLabel() foundation.String/* debug [protocol_interface/required_method]: AccessibilityLabel */
	AccessibilitySelectedChildren() foundation.Array/* debug [protocol_interface/required_method]: AccessibilitySelectedChildren */
}
