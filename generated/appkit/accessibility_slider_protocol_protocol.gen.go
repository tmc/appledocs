// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
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
	AccessibilityLabel() foundation.String/* debug [protocol_interface/required_method]: AccessibilityLabel */
	AccessibilityPerformDecrement() bool/* debug [protocol_interface/required_method]: AccessibilityPerformDecrement */
	AccessibilityPerformIncrement() bool/* debug [protocol_interface/required_method]: AccessibilityPerformIncrement */
	AccessibilityValue() objc.ID/* debug [protocol_interface/required_method]: AccessibilityValue */
}
