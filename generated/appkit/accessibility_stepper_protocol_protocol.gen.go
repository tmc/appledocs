// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PAccessibilityStepper is the NSAccessibilityStepper protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a stepper.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityStepper
type PAccessibilityStepper interface {
	// Required methods
	AccessibilityLabel() foundation.String
	AccessibilityPerformDecrement() bool
	AccessibilityPerformIncrement() bool
	// Optional methods
	AccessibilityValue() objc.ID
	HasAccessibilityValue() bool
}
