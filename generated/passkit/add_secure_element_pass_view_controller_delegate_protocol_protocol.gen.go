// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PAddSecureElementPassViewControllerDelegate is the PKAddSecureElementPassViewControllerDelegate protocol interface.
//
// The methods for responding to the life cycle events of a Secure Element pass.
//
// Availability:
//   - Mac Catalyst 13.4+
//   - iOS 13.4+
//   - iPadOS 13.4+
//   - visionOS 1.0+
//
// See: doc://com.apple.passkit/documentation/PassKit/PKAddSecureElementPassViewControllerDelegate
type PAddSecureElementPassViewControllerDelegate interface {
	// Required methods
	AddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError(controller AddSecureElementPassViewController /* not a class type */, passes []SecureElementPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: AddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError */
	// Optional methods
	AddSecureElementPassViewControllerDidFinishAddingSecureElementPassError(controller AddSecureElementPassViewController /* not a class type */, pass SecureElementPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */)
	HasAddSecureElementPassViewControllerDidFinishAddingSecureElementPassError() bool
}

// AddSecureElementPassViewControllerDelegate is a delegate implementation builder for the PAddSecureElementPassViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AddSecureElementPassViewControllerDelegate struct {
	_AddSecureElementPassViewControllerDidFinishAddingSecureElementPassError func(controller AddSecureElementPassViewController /* not a class type */, pass SecureElementPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */)
	_AddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError func(controller AddSecureElementPassViewController /* not a class type */, passes []SecureElementPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */)
}

// SetAddSecureElementPassViewControllerDidFinishAddingSecureElementPassError sets the handler for the AddSecureElementPassViewControllerDidFinishAddingSecureElementPassError delegate method.
//
// Tells the delegate when PassKit finishes adding a Secure Element pass.
func (d *AddSecureElementPassViewControllerDelegate) SetAddSecureElementPassViewControllerDidFinishAddingSecureElementPassError(f func(controller AddSecureElementPassViewController /* not a class type */, pass SecureElementPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */)) {
	d._AddSecureElementPassViewControllerDidFinishAddingSecureElementPassError = f
}

// SetAddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError sets the handler for the AddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError delegate method.
//
// Tells the delegate when PassKit finishes adding one or more Secure Element passes.
func (d *AddSecureElementPassViewControllerDelegate) SetAddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError(f func(controller AddSecureElementPassViewController /* not a class type */, passes []SecureElementPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */)) {
	d._AddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError = f
}

// AddSecureElementPassViewControllerDidFinishAddingSecureElementPassError implements the PAddSecureElementPassViewControllerDelegate interface.
func (d *AddSecureElementPassViewControllerDelegate) AddSecureElementPassViewControllerDidFinishAddingSecureElementPassError(controller AddSecureElementPassViewController /* not a class type */, pass SecureElementPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */) {
	if d._AddSecureElementPassViewControllerDidFinishAddingSecureElementPassError != nil {
		d._AddSecureElementPassViewControllerDidFinishAddingSecureElementPassError(controller, pass, error_)
	}
}

// HasAddSecureElementPassViewControllerDidFinishAddingSecureElementPassError returns true if a handler for AddSecureElementPassViewControllerDidFinishAddingSecureElementPassError has been set.
func (d *AddSecureElementPassViewControllerDelegate) HasAddSecureElementPassViewControllerDidFinishAddingSecureElementPassError() bool {
	return d._AddSecureElementPassViewControllerDidFinishAddingSecureElementPassError != nil
}

// AddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError implements the PAddSecureElementPassViewControllerDelegate interface.
func (d *AddSecureElementPassViewControllerDelegate) AddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError(controller AddSecureElementPassViewController /* not a class type */, passes []SecureElementPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */) {
	if d._AddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError != nil {
		d._AddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError(controller, passes, error_)
	}
}

// HasAddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError returns true if a handler for AddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError has been set.
func (d *AddSecureElementPassViewControllerDelegate) HasAddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError() bool {
	return d._AddSecureElementPassViewControllerDidFinishAddingSecureElementPassesError != nil
}
