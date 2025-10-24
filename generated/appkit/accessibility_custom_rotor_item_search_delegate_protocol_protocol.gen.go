// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PAccessibilityCustomRotorItemSearchDelegate is the NSAccessibilityCustomRotorItemSearchDelegate protocol interface.
//
// A delegate for a custom rotor that finds the next item result after performing a search with the specified search parameters.
//
// Availability:
//   - macOS 10.13+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityCustomRotorItemSearchDelegate
type PAccessibilityCustomRotorItemSearchDelegate interface {
	// Required methods
	RotorResultForSearchParameters(rotor IAccessibilityCustomRotor, searchParameters IAccessibilityCustomRotorSearchParameters) AccessibilityCustomRotorItemResult
}

// AccessibilityCustomRotorItemSearchDelegate is a delegate implementation builder for the PAccessibilityCustomRotorItemSearchDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AccessibilityCustomRotorItemSearchDelegate struct {
	_RotorResultForSearchParameters func(rotor IAccessibilityCustomRotor, searchParameters IAccessibilityCustomRotorSearchParameters) AccessibilityCustomRotorItemResult
}

// SetRotorResultForSearchParameters sets the handler for the RotorResultForSearchParameters delegate method.
//
// Performs a search with the specified search parameters and returns the item result.
func (d *AccessibilityCustomRotorItemSearchDelegate) SetRotorResultForSearchParameters(f func(rotor IAccessibilityCustomRotor, searchParameters IAccessibilityCustomRotorSearchParameters) AccessibilityCustomRotorItemResult) {
	d._RotorResultForSearchParameters = f
}

// RotorResultForSearchParameters implements the PAccessibilityCustomRotorItemSearchDelegate interface.
func (d *AccessibilityCustomRotorItemSearchDelegate) RotorResultForSearchParameters(rotor IAccessibilityCustomRotor, searchParameters IAccessibilityCustomRotorSearchParameters) AccessibilityCustomRotorItemResult {
	if d._RotorResultForSearchParameters != nil {
		return d._RotorResultForSearchParameters(rotor, searchParameters)
	}
	var zero AccessibilityCustomRotorItemResult
	return zero
}

// HasRotorResultForSearchParameters returns true if a handler for RotorResultForSearchParameters has been set.
func (d *AccessibilityCustomRotorItemSearchDelegate) HasRotorResultForSearchParameters() bool {
	return d._RotorResultForSearchParameters != nil
}
