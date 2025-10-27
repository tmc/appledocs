// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
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
	RotorResultForSearchParameters(rotor IAccessibilityCustomRotor, searchParameters IAccessibilityCustomRotorSearchParameters) IAccessibilityCustomRotorItemResult
}

// AccessibilityCustomRotorItemSearchDelegate is a delegate implementation builder for the PAccessibilityCustomRotorItemSearchDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AccessibilityCustomRotorItemSearchDelegate struct {
	_RotorResultForSearchParameters func(rotor IAccessibilityCustomRotor, searchParameters IAccessibilityCustomRotorSearchParameters) IAccessibilityCustomRotorItemResult
}

// SetRotorResultForSearchParameters sets the handler for the RotorResultForSearchParameters delegate method.
//
// Performs a search with the specified search parameters and returns the item result.
func (d *AccessibilityCustomRotorItemSearchDelegate) SetRotorResultForSearchParameters(f func(rotor IAccessibilityCustomRotor, searchParameters IAccessibilityCustomRotorSearchParameters) IAccessibilityCustomRotorItemResult) {
	d._RotorResultForSearchParameters = f
}

// RotorResultForSearchParameters implements the PAccessibilityCustomRotorItemSearchDelegate interface.
func (d *AccessibilityCustomRotorItemSearchDelegate) RotorResultForSearchParameters(rotor IAccessibilityCustomRotor, searchParameters IAccessibilityCustomRotorSearchParameters) IAccessibilityCustomRotorItemResult {
	if d._RotorResultForSearchParameters != nil {
		return d._RotorResultForSearchParameters(rotor, searchParameters)
	}
	var zero IAccessibilityCustomRotorItemResult
	return zero
}

// HasRotorResultForSearchParameters returns true if a handler for RotorResultForSearchParameters has been set.
func (d *AccessibilityCustomRotorItemSearchDelegate) HasRotorResultForSearchParameters() bool {
	return d._RotorResultForSearchParameters != nil
}

// AccessibilityCustomRotorItemSearchDelegateObject wraps an existing Objective-C object that conforms to the PAccessibilityCustomRotorItemSearchDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type AccessibilityCustomRotorItemSearchDelegateObject struct {
	objectivec.Object
}

// NewAccessibilityCustomRotorItemSearchDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSAccessibilityCustomRotorItemSearchDelegate protocol.
func NewAccessibilityCustomRotorItemSearchDelegateObject(obj objectivec.Object) *AccessibilityCustomRotorItemSearchDelegateObject {
	return &AccessibilityCustomRotorItemSearchDelegateObject{obj}
}

// Make sure AccessibilityCustomRotorItemSearchDelegateObject implements PAccessibilityCustomRotorItemSearchDelegate.
var _ PAccessibilityCustomRotorItemSearchDelegate = (*AccessibilityCustomRotorItemSearchDelegateObject)(nil)

// RotorResultForSearchParameters implements the PAccessibilityCustomRotorItemSearchDelegate interface.
// This required method is always available on objects conforming to RotorResultForSearchParameters.
func (o *AccessibilityCustomRotorItemSearchDelegateObject) RotorResultForSearchParameters(rotor IAccessibilityCustomRotor, searchParameters IAccessibilityCustomRotorSearchParameters) IAccessibilityCustomRotorItemResult {
	return objc.Send[IAccessibilityCustomRotorItemResult](o.ID, objc.Sel("rotor:resultForSearchParameters:"), rotor, searchParameters)
}
