// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PAccessibilityElement is the NSAccessibilityElement protocol interface.
//
// A role-based protocol that declares the minimum interface necessary to interact with an assistive app.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityElementProtocol
type PAccessibilityElement interface {
	// Required methods
	AccessibilityFrame() Rect/* debug [protocol_interface/required_method]: AccessibilityFrame */
	AccessibilityParent() objc.ID/* debug [protocol_interface/required_method]: AccessibilityParent */
	// Optional methods
	AccessibilityIdentifier() foundation.String
	HasAccessibilityIdentifier() bool
	IsAccessibilityFocused() bool
	HasIsAccessibilityFocused() bool
}
