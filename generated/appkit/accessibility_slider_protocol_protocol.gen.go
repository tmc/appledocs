// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PAccessibilitySlider is the NSAccessibilitySlider protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a slider.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilitySlider
type PAccessibilitySlider interface {
	// Required methods
	AccessibilityLabel() foundation.String
	AccessibilityPerformDecrement() bool
	AccessibilityPerformIncrement() bool
	AccessibilityValue() objc.ID
}
