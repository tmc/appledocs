// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
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
	AccessibilityElementWithToken(token AccessibilityLoadingToken /* typedef */) unsafe.Pointer/* debug [protocol_interface/required_method]: AccessibilityElementWithToken */
	// Optional methods
	AccessibilityRangeInTargetElementWithToken(token AccessibilityLoadingToken /* typedef */) corefoundation.Range
	HasAccessibilityRangeInTargetElementWithToken() bool
}
