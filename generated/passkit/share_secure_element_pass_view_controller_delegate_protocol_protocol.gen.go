// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PShareSecureElementPassViewControllerDelegate is the PKShareSecureElementPassViewControllerDelegate protocol interface.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.passkit/documentation/PassKit/PKShareSecureElementPassViewControllerDelegate
type PShareSecureElementPassViewControllerDelegate interface {
	// Required methods
	ShareSecureElementPassViewControllerDidFinishWithResult(controller ShareSecureElementPassViewController /* not a class type */, result ShareSecureElementPassResult)/* debug [protocol_interface/required_method]: ShareSecureElementPassViewControllerDidFinishWithResult */
	// Optional methods
	ShareSecureElementPassViewControllerDidCreateShareURLActivationCode(controller ShareSecureElementPassViewController /* not a class type */, universalShareURL objc.IObject /* cross-framework: NSURL */, activationCode objc.IObject /* cross-framework: NSString */)
	HasShareSecureElementPassViewControllerDidCreateShareURLActivationCode() bool
}

// ShareSecureElementPassViewControllerDelegate is a delegate implementation builder for the PShareSecureElementPassViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ShareSecureElementPassViewControllerDelegate struct {
	_ShareSecureElementPassViewControllerDidCreateShareURLActivationCode func(controller ShareSecureElementPassViewController /* not a class type */, universalShareURL objc.IObject /* cross-framework: NSURL */, activationCode objc.IObject /* cross-framework: NSString */)
	_ShareSecureElementPassViewControllerDidFinishWithResult func(controller ShareSecureElementPassViewController /* not a class type */, result ShareSecureElementPassResult)
}

// SetShareSecureElementPassViewControllerDidCreateShareURLActivationCode sets the handler for the ShareSecureElementPassViewControllerDidCreateShareURLActivationCode delegate method.
func (d *ShareSecureElementPassViewControllerDelegate) SetShareSecureElementPassViewControllerDidCreateShareURLActivationCode(f func(controller ShareSecureElementPassViewController /* not a class type */, universalShareURL objc.IObject /* cross-framework: NSURL */, activationCode objc.IObject /* cross-framework: NSString */)) {
	d._ShareSecureElementPassViewControllerDidCreateShareURLActivationCode = f
}

// SetShareSecureElementPassViewControllerDidFinishWithResult sets the handler for the ShareSecureElementPassViewControllerDidFinishWithResult delegate method.
func (d *ShareSecureElementPassViewControllerDelegate) SetShareSecureElementPassViewControllerDidFinishWithResult(f func(controller ShareSecureElementPassViewController /* not a class type */, result ShareSecureElementPassResult)) {
	d._ShareSecureElementPassViewControllerDidFinishWithResult = f
}

// ShareSecureElementPassViewControllerDidCreateShareURLActivationCode implements the PShareSecureElementPassViewControllerDelegate interface.
func (d *ShareSecureElementPassViewControllerDelegate) ShareSecureElementPassViewControllerDidCreateShareURLActivationCode(controller ShareSecureElementPassViewController /* not a class type */, universalShareURL objc.IObject /* cross-framework: NSURL */, activationCode objc.IObject /* cross-framework: NSString */) {
	if d._ShareSecureElementPassViewControllerDidCreateShareURLActivationCode != nil {
		d._ShareSecureElementPassViewControllerDidCreateShareURLActivationCode(controller, universalShareURL, activationCode)
	}
}

// HasShareSecureElementPassViewControllerDidCreateShareURLActivationCode returns true if a handler for ShareSecureElementPassViewControllerDidCreateShareURLActivationCode has been set.
func (d *ShareSecureElementPassViewControllerDelegate) HasShareSecureElementPassViewControllerDidCreateShareURLActivationCode() bool {
	return d._ShareSecureElementPassViewControllerDidCreateShareURLActivationCode != nil
}

// ShareSecureElementPassViewControllerDidFinishWithResult implements the PShareSecureElementPassViewControllerDelegate interface.
func (d *ShareSecureElementPassViewControllerDelegate) ShareSecureElementPassViewControllerDidFinishWithResult(controller ShareSecureElementPassViewController /* not a class type */, result ShareSecureElementPassResult) {
	if d._ShareSecureElementPassViewControllerDidFinishWithResult != nil {
		d._ShareSecureElementPassViewControllerDidFinishWithResult(controller, result)
	}
}

// HasShareSecureElementPassViewControllerDidFinishWithResult returns true if a handler for ShareSecureElementPassViewControllerDidFinishWithResult has been set.
func (d *ShareSecureElementPassViewControllerDelegate) HasShareSecureElementPassViewControllerDidFinishWithResult() bool {
	return d._ShareSecureElementPassViewControllerDidFinishWithResult != nil
}
