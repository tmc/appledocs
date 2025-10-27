// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PAccessibilityCheckBox is the NSAccessibilityCheckBox protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a checkbox.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityCheckBox
type PAccessibilityCheckBox interface {
	// Required methods
	AccessibilityValue() foundation.Number
}
