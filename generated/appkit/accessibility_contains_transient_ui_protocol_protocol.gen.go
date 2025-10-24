// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PAccessibilityContainsTransientUI is the NSAccessibilityContainsTransientUI protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to support dynamic UI changes.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityContainsTransientUI
type PAccessibilityContainsTransientUI interface {
	// Required methods
	AccessibilityPerformShowAlternateUI() bool/* debug [protocol_interface/required_method]: AccessibilityPerformShowAlternateUI */
	AccessibilityPerformShowDefaultUI() bool/* debug [protocol_interface/required_method]: AccessibilityPerformShowDefaultUI */
	IsAccessibilityAlternateUIVisible() bool/* debug [protocol_interface/required_method]: IsAccessibilityAlternateUIVisible */
}
