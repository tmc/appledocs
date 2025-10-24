// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PAccessibilityElementLoading is the NSAccessibilityElementLoading protocol interface.
//
// A role-based protocol that declares the minimum interface necessary for an accessibility element to support loading.
//
// Availability:
//   - macOS 10.13+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityElementLoading
type PAccessibilityElementLoading interface {
	// Required methods
	AccessibilityElementWithToken(token objc.IObject /* cross-framework: AccessibilityLoadingToken */) objc.ID
	// Optional methods
	AccessibilityRangeInTargetElementWithToken(token objc.IObject /* cross-framework: AccessibilityLoadingToken */) corefoundation.Range
	HasAccessibilityRangeInTargetElementWithToken() bool
}
