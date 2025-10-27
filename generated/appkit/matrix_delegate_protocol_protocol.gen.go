// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PMatrixDelegate is the NSMatrixDelegate protocol interface.
//
// The   protocol defines the optional methods implemented by delegates of   objects.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSMatrixDelegate
type PMatrixDelegate interface {
}

// MatrixDelegate is a delegate implementation builder for the PMatrixDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MatrixDelegate struct {
}

// MatrixDelegateObject wraps an existing Objective-C object that conforms to the PMatrixDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type MatrixDelegateObject struct {
	objectivec.Object
}

// NewMatrixDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSMatrixDelegate protocol.
func NewMatrixDelegateObject(obj objectivec.Object) *MatrixDelegateObject {
	return &MatrixDelegateObject{obj}
}

// Make sure MatrixDelegateObject implements PMatrixDelegate.
var _ PMatrixDelegate = (*MatrixDelegateObject)(nil)
