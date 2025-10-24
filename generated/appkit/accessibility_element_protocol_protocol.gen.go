// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
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
	AccessibilityFrame() corefoundation.Rect
	AccessibilityParent() objc.ID
	// Optional methods
	AccessibilityIdentifier() foundation.String
	HasAccessibilityIdentifier() bool
	IsAccessibilityFocused() bool
	HasIsAccessibilityFocused() bool
}
