// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
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
	AccessibilityElementWithToken(token AccessibilityLoadingToken) unsafe.Pointer
	// Optional methods
	AccessibilityRangeInTargetElementWithToken(token AccessibilityLoadingToken) foundation.Range
	HasAccessibilityRangeInTargetElementWithToken() bool
}
